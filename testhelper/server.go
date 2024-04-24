package testhelper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"lpstest/log"
	"net"
	"strconv"
	"sync/atomic"
)

var logger = log.DLogger()

// 表示服务器请求的结构体
type ServerReq struct {
	ID       int64
	Operands []int
	Operator string
}

// 表示 服务器响应的结构体
type ServerResp struct {
	ID      int64
	Formula string
	Result  int
	Err     error
}

func op(operands []int, operator string) int {
	var result int
	switch {
	case operator == "+":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result += v
			}
		}
	case operator == "-":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result -= v
			}
		}
	case operator == "*":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result *= v
			}
		}
	case operator == "/":
		for _, v := range operands {
			if result == 0 {
				result = v
			} else {
				result /= v
			}
		}
	}
	return result
}

// 根据参数生成字符串形式的公司
func genFormula(operands []int, operator string, result int, equal bool) string {
	var buff bytes.Buffer
	n := len(operands)
	for i := 0; i < n; i++ {
		if i > 0 {
			buff.WriteString(" ")
			buff.WriteString(operator)
			buff.WriteString(" ")
		}

		buff.WriteString(strconv.Itoa(operands[i]))
	}
	if equal {
		buff.WriteString(" = ")
	} else {
		buff.WriteString(" != ")
	}
	buff.WriteString(strconv.Itoa(result))
	return buff.String()
}

// 会把参数 sresp 代表的请求转换为数据并发送连接。
func reqHandler(conn net.Conn) {
	var errMsg string
	var sresp ServerResp
	req, err := read(conn, DELIM)
	if err != nil {
		errMsg = fmt.Sprintf("Server: Req Read Error: %s", err)
	} else {
		var sreq ServerReq
		err := json.Unmarshal(req, &sreq)
		if err != nil {
			errMsg = fmt.Sprintf("Server: Req Unmarshal Error: %s", err)
		} else {
			sresp.ID = sreq.ID
			sresp.Result = op(sreq.Operands, sreq.Operator)
			sresp.Formula = genFormula(sreq.Operands, sreq.Operator, sresp.Result, true)
		}
	}
	if errMsg != "" {
		sresp.Err = errors.New(errMsg)
	}
	bytes, err := json.Marshal(sresp)
	if err != nil {
		logger.Errorf("Server: Resp Marshal Error: %s", err)
	}
	logger.Infof("Server: Resp Marshal info: %s", sresp)
	_, err = write(conn, bytes, DELIM)
	if err != nil {
		logger.Errorf("Server: Resp Write error: %s", err)
	}
}

// 表示基于 TCP 协议的服务器
type TCPServer struct {
	listenner net.Listener
	active    uint32 // 0-未激活；1-已激活
}

// 新建一个基于 TCP 协议的服务器
func NewTCPServer() *TCPServer {
	return &TCPServer{}
}

func (server *TCPServer) init(addr string) error {
	if !atomic.CompareAndSwapUint32(&server.active, 0, 1) {
		return nil
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		atomic.StoreUint32(&server.active, 0)
		return err
	}
	server.listenner = ln
	return nil
}

func (server *TCPServer) Listen(addr string) error {
	err := server.init(addr)
	if err != nil {
		return err
	}
	go func() {
		for {
			if atomic.LoadUint32(&server.active) != 1 {
				break
			}
			conn, err := server.listenner.Accept()
			if err != nil {
				if atomic.LoadUint32(&server.active) == 1 {
					logger.Errorf("Server: Request Acception Error: %s\n", err)
				} else {
					logger.Warnf("Server: Broken acception because of closed network connection.")
				}
				continue
			}
			go reqHandler(conn)
		}
	}()
	return nil
}

func (server *TCPServer) Close() bool {
	if !atomic.CompareAndSwapUint32(&server.active, 1, 0) {
		return false
	}
	server.listenner.Close()
	return true
}
