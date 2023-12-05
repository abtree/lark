package websocket

import (
	"context"
	"lark/com/pkgs/xgin"
	"lark/com/pkgs/xjwt"
	"lark/com/pkgs/xlog"
	"lark/com/utils"
	"lark/pb"
	"net/http"
	"runtime/debug"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

type Wstask struct {
	Sender *pb.UserInfo
	Msg    *pb.WebSocketProto
}

type Hub struct {
	cfg            *pb.WebSocketCfg //配置
	upgrader       websocket.Upgrader
	registerChan   chan *Client       //注册client队列
	unregisterChan chan *Client       //退出client队列
	readChan       chan *Wstask       //消息队列
	callback       func(*Wstask)      //注册的消息处理函数
	clients        sync.Map           //注册的客户端
	cancle         context.CancelFunc //退出函数
	wg             sync.WaitGroup     //等待退出
}

func NewHub(cfg *pb.WebSocketCfg, call func(*Wstask)) *Hub {
	return &Hub{
		cfg: cfg,
		upgrader: websocket.Upgrader{
			ReadBufferSize:    int(cfg.ReadBufferSize),
			WriteBufferSize:   int(cfg.WriteBufferSize),
			EnableCompression: false, //不压缩消息
			//解决跨域问题
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		registerChan:   make(chan *Client, cfg.ChanRegisterSize),
		unregisterChan: make(chan *Client, cfg.ChanUnregisterSize),
		readChan:       make(chan *Wstask, cfg.ChanReadSize),
		callback:       call,
		clients:        sync.Map{},
	}
}

// 注册客户端
func (t *Hub) registerClient(cli *Client) {
	//首先检测是否已经注册了
	if c, ok := t.clients.Load(cli.User.Id); ok && c != nil {
		c.(*Client).offline() //老客户端下线
	}
	t.clients.Store(cli.User.Id, cli) //新客户端上线
	cli.Online()
}

// 提出客户端
func (t *Hub) unregisterClient(cli *Client) {
	cli.offline()
	t.clients.Delete(cli.User.Id)
}

// 运行hub
func (t *Hub) Run() {
	defer func() {
		if r := recover(); r != nil {
			xlog.Error(r, debug.Stack())
		}
	}()
	ctx, cancle := context.WithCancel(context.Background())
	t.cancle = cancle
	//开启注册线程
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case client := <-t.registerChan:
				t.registerClient(client)
			case client := <-t.unregisterChan:
				t.unregisterClient(client)
			case <-ctx.Done():
				return
			}
		}
	}()
	//开启读线程(多个)
	for i := uint32(0); i < t.cfg.RoutineRead; i++ {
		t.wg.Add(1)
		go func() {
			defer t.wg.Done()
			for {
				select {
				case msg := <-t.readChan:
					t.callback(msg)
				case <-ctx.Done():
					return
				}
			}
		}()
	}
}

// 退出
func (t *Hub) Exit() {
	t.cancle()
	t.wg.Wait()
}

// 发送消息
func (t *Hub) SendMsg(guid uint32, msg *pb.WebSocketProto) pb.ServerError {
	cli, ok := t.clients.Load(guid)
	if !ok || cli == nil {
		return pb.ServerError_WsClientOffline
	}
	cli.(*Client).Send(msg)
	return pb.ServerError_NilError
}

// 广播消息
func (t *Hub) Broadcast(msg *pb.WebSocketProto) {
	t.clients.Range(func(k, v any) bool {
		v.(*Client).Send(msg)
		return true
	})
}

// 获取在线人数
func (t *Hub) NumOfOnline() int {
	i := 0
	t.clients.Range(func(k, v any) bool {
		i++
		return true
	})
	return i
}

// 判断读通道是否已满
func (t *Hub) ReadChanFull() bool {
	return len(t.readChan) >= int(t.cfg.ChanReadSize)
}

// 消息处理函数(客户端登录)
func (t *Hub) Handler(c *gin.Context) {
	cs := t.NumOfOnline()
	if cs >= int(t.cfg.MaxConn) {
		xgin.Result(c, &xgin.Resp{
			Code: int32(pb.ServerError_WsHubMaxConn),
			Msg:  utils.Const_Ws_Hub_Max_Connetcions,
		})
		return
	}
	str := xgin.ParseFromHeader(c)
	if str == "" {
		xgin.Result(c, &xgin.Resp{
			Code: int32(pb.ServerError_WsVerifyError),
			Msg:  utils.Const_Ws_UserInfo_Error,
		})
		return
	}
	token, err := xjwt.Decode(str)
	if err != nil {
		xlog.Warn(err.Error())
		xgin.Result(c, &xgin.Resp{
			Code: int32(pb.ServerError_WsVerifyError),
			Msg:  err.Error(),
		})
		return
	}
	user := &pb.UserInfo{}
	err = proto.Unmarshal(token.UserData, user)
	if err != nil {
		xlog.Warn(err.Error())
		xgin.Result(c, &xgin.Resp{
			Code: int32(pb.ServerError_WsVerifyError),
			Msg:  err.Error(),
		})
		return
	}
	conn, err := t.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		xlog.Warn(err.Error())
		xgin.Result(c, &xgin.Resp{
			Code: int32(pb.ServerError_WsUpgraderFailed),
			Msg:  err.Error(),
		})
		return
	}
	client := newClient(t, conn, user)
	client.listen()
	t.registerChan <- client
}
