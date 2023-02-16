package ws

// 定义的一些状态码
const (
	heartBeat = 0    // 心跳
	connOut   = 5000 // 离线
	connOk    = 1000 // 上线
	SendOk    = 200  // 消息投递成功
	CrowdedOk = 4001 // 已在别处登录
)

type ImClientMessage struct {
	Code        int    `json:"code,omitempty"`
	FromId      int    `json:"from_id,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ToId        int    `json:"to_id,omitempty"`
	Status      int    `json:"status,omitempty"`
	MsgType     int    `json:"msg_type,omitempty"`
	ChannelType int    `json:"channel_type"`
}

type ImServerMessage struct {
	Sender    uint   `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
	Mes       *ImClientMessage
}

type ImOnlineMessage struct {
	Code        int    `json:"code,omitempty"`
	Msg         string `json:"msg,omitempty"`
	ID          uint   `json:"id,omitempty"`
	ChannelType int    `json:"channel_type"` // 1 私聊 2 群聊 3 广播
}
