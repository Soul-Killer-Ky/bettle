package ws

import (
	"github.com/go-kratos/kratos/v2/encoding"
	"sync"

	_ "github.com/go-kratos/kratos/v2/encoding/json"
	"github.com/gorilla/websocket"
)

type ImClientManager struct {
	ActiveClients map[uint]*ImClient // 存放在线用户连接
	Broadcast     chan []byte        // 收集消息分发到客户端
	Register      chan *ImClient     // 新注册的长连接
	Unregister    chan *ImClient     // 已注销的长连接

	mu sync.RWMutex
}

// ImManager 客户端管理器
var ImManager = ImClientManager{
	ActiveClients: make(map[uint]*ImClient),
	Broadcast:     make(chan []byte),
	Register:      make(chan *ImClient),
	Unregister:    make(chan *ImClient),
}

func (manager *ImClientManager) SetClientInfo(c *ImClient) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	manager.ActiveClients[c.ID] = c
}

func (manager *ImClientManager) RemoveClientInfo(c *ImClient) {
	manager.mu.Lock()
	defer manager.mu.Unlock()
	delete(manager.ActiveClients, c.ID)
}

// LaunchOnlineMsg 上线消息通知
func (manager *ImClientManager) LaunchOnlineMsg(id uint) {
	jsonCodec := encoding.GetCodec("json")
	message, _ := jsonCodec.Marshal(&ImOnlineMessage{Code: connOk, Msg: "用户上线啦", ID: id, ChannelType: 3})
	for _, client := range manager.ActiveClients {
		err := client.Socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
	return
}

func (manager *ImClientManager) LaunchOfflineMsg(id uint) {
	jsonCodec := encoding.GetCodec("json")
	message, _ := jsonCodec.Marshal(&ImOnlineMessage{Code: connOk, Msg: "用户离线了", ID: id, ChannelType: 3})
	for _, client := range manager.ActiveClients {
		err := client.Socket.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
	return
}

func (manager *ImClientManager) Start() {
	for {
		select {
		case c := <-ImManager.Register:
			manager.SetClientInfo(c)      // 设置客户端信息
			manager.LaunchOnlineMsg(c.ID) // 用户在线下发通知
			//node.SetUserServiceNode(c.ID) // 设置用户节点
			//err := pool.AntsPool.Submit(func() {
			//	MqPersonalConsumption(c, c.ID) //离线消息同步
			//})
			//if err != nil {
			//	zaplog.Error("MqPersonalConsumption submit fail: %s", err.Error())
			//}
		case c := <-ImManager.Unregister:
			manager.RemoveClientInfo(c)
			manager.LaunchOfflineMsg(c.ID)
			//PushUserOfflineNotification(manager, conn) // 设置用户离线状态
			//node.DelUserServiceNode(conn.ID)           // 删除用户节点
		case _ = <-ImManager.Broadcast:
			//err := pool.AntsPool.Submit(func() {
			//	manager.LaunchMessage(message) // 协程池任务调度 抢占式消息下发
			//})
			//if err != nil {
			//zaplog.Error("LaunchMessage submit fail: %s", err.Error())
			//}
		}
	}
}
