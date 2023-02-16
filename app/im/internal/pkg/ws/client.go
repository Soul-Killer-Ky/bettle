package ws

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"

	_ "github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc/encoding"
)

type ImClient struct {
	ID     uint
	Socket *websocket.Conn
	Send   chan []byte

	log *log.Helper
}

var (
	Text  = websocket.TextMessage  // 文本消息指令
	Clone = websocket.CloseMessage // 关闭指令
)

func NewImClient(id uint, conn *websocket.Conn, log *log.Helper) *ImClient {
	return &ImClient{
		ID:     id,
		Socket: conn,
		Send:   make(chan []byte),

		log: log,
	}
}

// Read 消息投递
func (c *ImClient) Read() {
	defer func() {
		ImManager.Unregister <- c
		c.Socket.Close()
	}()
	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			ImManager.Unregister <- c
			c.Socket.Close()
			break
		}
		c.PullMessageHandler(message)
	}
}

// Write 从客户端消费消息
func (c *ImClient) Write() {
	defer c.Socket.Close()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(Clone, []byte{})
				return
			}
			c.Socket.WriteMessage(Text, message)
		}
	}
}

// PullMessageHandler 消息处理方法
func (c *ImClient) PullMessageHandler(message []byte) {
	if len(message) < 0 {
		return
	}
	if string(message) == "HeartBeat" {
		LaunchTicklingAckMsg([]byte(`{"code":0,"data":"heartbeat ok"}`), c)
		return
	}

	msg := new(ImClientMessage)
	jsonCodec := encoding.GetCodec("json")
	err := jsonCodec.Unmarshal(message, &msg)
	if err != nil {
		c.log.Errorf("消息解析异常: %s", err)
		return
	}

	if msg.ChannelType == 1 {
		data := fmt.Sprintf(`{"code":200,"msg":"%s","from_id":%v,"to_id":%v,"status":"0","msg_type":%v,"channel_type":%v}`,
			msg.Msg, msg.FromId, msg.ToId, msg.MsgType, msg.ChannelType)
		LaunchTicklingAckMsg([]byte(data), c)
	}
	messageByte, _ := jsonCodec.Marshal(&ImServerMessage{Sender: c.ID, Mes: msg})
	ImManager.Broadcast <- messageByte

	return
}

func LaunchTicklingAckMsg(msg []byte, conn *ImClient) {
	conn.Socket.WriteMessage(websocket.TextMessage, msg)
	return
}
