package pusher

import (
	"github.com/pusher/pusher-http-go/v5"
)

func PusherInit() *pusher.Client {
	var PusherClient pusher.Client
	PusherClient = pusher.Client{
		AppID:   "1924815",
		Key:     "1862852b892468c862b2",
		Secret:  "1e7e510bc7d1594b37d0",
		Cluster: "ap2",
		Secure:  true,
	}
	return &PusherClient
	//
	//data := map[string]string{"message": "hello hh"} // can receive from post handler
	//err := pusherClient.Trigger("my-channel", "my-event", data)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("pusher trigger")
}

//client := pusher.New("1862852b892468c862b2")
//
//fmt.Println("conneteced")
//client.BindGlobal(func(channel, event string, data interface{}) {
//fmt.Println(channel, event, data)
//})
//
//client.Subscribe("my-channel")
//time.Sleep(1000 * time.Second)
//
//fmt.Println("pusher subscribed")
