package xcfgstructpb

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	//配置文件所在路径
	// cfgPath = "../../apps/config/files"
	//生成pb文件存放路径
	// pbPath    = "../../pbfiles/"
	end_chat  = "\r\n"
	end_begin = "\r\n\t"
	pb_array  = "[]"
)

var (
	//已经存在的pb 检测配置重名
	had_pbs   map[string]bool
	had_yyact map[string]bool
	wfile     *os.File
	// 组织常规配置大结构
	bigmsg   *bytes.Buffer
	bigindex int
	// 组织活动配置大结构
	yyactmsg   *bytes.Buffer
	yyactindex int
	// 用于子结构 临时写入用
	filemsg   *bytes.Buffer
	fileIndex int
)

func endChat() {
	wfile.WriteString(end_chat)
}

func writeHead(issystem bool) {
	wfile.WriteString(`syntax ="proto3";
package pb;
option go_package = "../pb;pb";

//---------------说明---------------------------
//自动生成的pb 配置文件
//----------------------------------------------
	`)
	endChat()

	if issystem {
		bigmsg.WriteString("message MsgSysConfigs {")
	} else {
		bigmsg.WriteString("message MsgConfigs {")

		yyactmsg.WriteString("message MsgYYactConfigs {")
		yyactmsg.WriteString(end_begin)
		yyactmsg.WriteString("map<string,bytes> unhandle = 1;")
	}

	bigmsg.WriteString(end_begin)
	bigmsg.WriteString("map<string,bytes> unhandle = 1;")
	bigmsg.WriteString(end_begin)
}

func writeBig(path, pbname, key string) {
	if strings.HasPrefix(path, "YY") {
		writeFile(path, pbname, key)
		return
	}
	bigmsg.WriteString(end_begin)
	if key != "" {
		bigmsg.WriteString("map<" + filterKey(key) + ", " + pbname + ">")
	} else {
		bigmsg.WriteString(pbname)
	}
	bigmsg.WriteString(" ")
	bigmsg.WriteString(path)
	bigmsg.WriteString(" = ")
	bigmsg.WriteString(strconv.Itoa(bigindex))
	bigindex++
	bigmsg.WriteString(";")
}

func writeYYact(path, pbname string) {
	yyactmsg.WriteString(end_begin)
	yyactmsg.WriteString("map<uint32, " + pbname + ">")
	yyactmsg.WriteString(" ")
	yyactmsg.WriteString(path)
	yyactmsg.WriteString(" = ")
	yyactmsg.WriteString(strconv.Itoa(yyactindex))
	yyactindex++
	yyactmsg.WriteString(";")
}

func writeFile(path, pbname, key string) {
	filemsg.WriteString(end_begin)
	if key != "" {
		filemsg.WriteString("map<" + filterKey(key) + ", " + pbname + ">")
	} else {
		filemsg.WriteString(pbname)
	}
	filemsg.WriteString(" ")
	filemsg.WriteString(path)
	filemsg.WriteString(" = ")
	filemsg.WriteString(strconv.Itoa(fileIndex))
	fileIndex++
	filemsg.WriteString(";")
}

func filterKey(key string) string {
	//enum枚举类型不能做key
	if key[0] >= 'A' && key[0] <= 'Z' {
		return "uint32"
	}
	return key
}

func writeFin(issystem bool) {
	bigmsg.WriteString(end_chat)
	bigmsg.WriteString("}")
	wfile.WriteString(bigmsg.String())
	endChat()

	if issystem {
		return
	}
	yyactmsg.WriteString(end_chat)
	yyactmsg.WriteString("}")
	wfile.WriteString(yyactmsg.String())
	endChat()

	wfile.WriteString(`message MsgAllConfigs{
	MsgConfigs Configs = 1;
	MsgYYactConfigs Yyacts = 2;
}
	`)
}

func Run(cfgPath, pbPath string, issystem bool) {
	had_pbs = make(map[string]bool)
	var err error
	wfile, err = os.OpenFile(pbPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("Open write File %v error with : %v", pbPath, err)
	}
	defer wfile.Close()
	if issystem {
		had_yyact = make(map[string]bool)
	}

	bigmsg = bytes.NewBuffer([]byte{})
	bigindex = 2
	yyactmsg = bytes.NewBuffer([]byte{})
	yyactindex = 2

	writeHead(issystem)
	ReloadPath(cfgPath)
	writeFin(issystem)
}
