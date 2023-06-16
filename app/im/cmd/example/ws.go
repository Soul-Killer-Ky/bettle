package main

import (
	pb "beetle/api/im/service/v1"
	"errors"
	"flag"
	"fmt"
	"github.com/Soul-Killer-Ky/kratos/websocket"
)

var testClient *websocket.Client

func handleClientChatMessage(message *pb.PersonalChatMessage) error {
	fmt.Printf("Payload: %v\n", message)
	//_ = testClient.SendMessage(api.MessageTypeChat, message)
	return nil
}

var token string

func init() {
	flag.StringVar(&token, "token", "", "jwt token")
}

func main() {
	flag.Parse()
	cli := websocket.NewClient(
		websocket.WithEndpoint("ws://localhost:10000/ws"),
		websocket.WithClientCodec("json"),
		websocket.WithRequestHeader("Authorization", fmt.Sprintf("Bearer %s", token)),
	)
	defer cli.Disconnect()

	testClient = cli

	cli.RegisterMessageHandler(websocket.MessageType(pb.MessageType_PersonalChat),
		func(payload websocket.MessagePayload) error {
			switch t := payload.(type) {
			case *pb.PersonalChatMessage:
				return handleClientChatMessage(t)
			default:
				return errors.New("invalid payload type")
			}
		},
		func() websocket.Any { return &pb.PersonalChatMessage{} },
	)

	err := cli.Connect()
	if err != nil {
		panic(err)
	}

	var msg string
	for {
		fmt.Println("Please enter string: ")
		fmt.Scanln(&msg)
		if msg == "break" {
			break
		}
		cli.SendMessage(1, pb.PersonalChatMessage{
			Content: &pb.Content{
				Type: pb.ContentType_CT_Text,
				Body: msg,
			},
			Sender: 2,
		})
	}
}
