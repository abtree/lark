package websocket

import (
	"encoding/json"
	"lark/com/pkgs/xlog"
	"lark/pb"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	User     *pb.UserInfo            //用户信息
	LoginTs  int64                   //上线时间
	conn     *websocket.Conn         //连接socket
	hub      *Hub                    //所在的hub
	sendChan chan *pb.WebSocketProto //发送消息队列
	exit     chan struct{}           //关闭通道
	closed   atomic.Bool             //退出标识
}

func newClient(h *Hub, c *websocket.Conn, user *pb.UserInfo) *Client {
	return &Client{
		User:     user,
		conn:     c,
		hub:      h,
		sendChan: make(chan *pb.WebSocketProto, h.cfg.ChanWriteSize),
		exit:     make(chan struct{}),
	}
}

func (t *Client) listen() {
	go t.readLoop()
	go t.writeLoop()
}

func (t *Client) readLoop() {
	defer func() {
		if r := recover(); r != nil {
			xlog.Error(r, debug.Stack())
		}
		t.offline()
	}()
	cfg := t.hub.cfg
	t.conn.SetReadLimit(int64(cfg.ReadMaxBufferSize))
	t.conn.SetReadDeadline(time.Now().Add(time.Duration(cfg.PongWait) * time.Second))
	t.conn.SetPongHandler(t.pongHandler)
	t.conn.SetCloseHandler(t.closeHandler)
	for {
		msgType, buf, err := t.conn.ReadMessage()
		if err != nil {
			// if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			// }
			break //报错直接退出
		}
		if t.closed.Load() {
			break //已经关闭了 退出
		}
		if msgType == websocket.CloseMessage {
			break //收到退出消息 退出
		}
		if msgType != websocket.TextMessage {
			continue //不是我们能处理的消息
		}
		//限流(丢弃消息)
		if t.hub.ReadChanFull() {
			continue
		}
		msg := &pb.WebSocketProto{}
		err = json.Unmarshal(buf, msg)
		if err != nil {
			xlog.Warn(err.Error())
			continue
		}
		t.hub.readChan <- &Wstask{
			Sender: t.User,
			Msg:    msg,
		}
	}
}

func (t *Client) writeLoop() {
	pingTicker := time.NewTicker(time.Duration(t.hub.cfg.PingWait) * time.Second)
	defer func() {
		if r := recover(); r != nil {
			xlog.Error(r, debug.Stack())
		}
		pingTicker.Stop()
		t.offline()
	}()
	for {
		select {
		case message, ok := <-t.sendChan:
			if !ok {
				return //发送消息异常
			}
			if t.closed.Load() {
				return
			}
			if err := t.conn.SetWriteDeadline(time.Now().Add(time.Duration(t.hub.cfg.WriteWait) * time.Second)); err != nil {
				xlog.Warn(err.Error())
				return
			}
			wc, err := t.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				xlog.Warn(err.Error())
				return
			}
			byts, err := json.Marshal(message)
			if err != nil {
				xlog.Warn(err.Error())
				return
			}
			_, err = wc.Write(byts)
			if err != nil {
				xlog.Warn(err.Error())
				return
			}
			err = wc.Close()
			if err != nil {
				xlog.Warn(err.Error())
				return
			}
		case _, ok := <-pingTicker.C:
			if !ok {
				return //心跳异常
			}
			if t.closed.Load() {
				return //已经停止
			}
			if err := t.conn.SetWriteDeadline(time.Now().Add(time.Duration(t.hub.cfg.WriteWait) * time.Second)); err != nil {
				xlog.Warn(err.Error())
				return
			}
			if err := t.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				xlog.Warn(err.Error())
				return
			}
		case <-t.exit:
			return
		}
	}
}

func (t *Client) offline() {
	if t.closed.Load() {
		return
	}
	t.closed.Store(true)
	t.hub.unregisterChan <- t
	t.exit <- struct{}{}
	nowAt := time.Now()
	nowAt = nowAt.Add(time.Duration(t.hub.cfg.RWDeadLine) * time.Millisecond)
	t.conn.SetWriteDeadline(nowAt)
	t.conn.SetReadDeadline(nowAt)
	//耗时操作
	t.conn.Close()
}

func (t *Client) Online() {
	t.LoginTs = time.Now().Unix()
}

func (t *Client) Send(msg *pb.WebSocketProto) {
	if t.closed.Load() {
		return
	}
	if len(t.sendChan) >= int(t.hub.cfg.ChanWriteSize) {
		return
	}
	t.sendChan <- msg
}

func (t *Client) pongHandler(dat string) error {
	cfg := t.hub.cfg
	err := t.conn.SetReadDeadline(time.Now().Add(time.Duration(cfg.PongWait) * time.Second))
	if err != nil {
		xlog.Warn(err.Error())
		t.offline()
	}
	return err
}

func (t *Client) closeHandler(code int, text string) error {
	t.offline()
	return nil
}
