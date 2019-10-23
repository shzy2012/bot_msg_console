package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/common/log"
	"github.com/laiye-ai/wulai-openapi-sdk-golang/services/wulai"
)

//Bot bot
type Bot struct {
}

// botMsgRoute 消息路由 handles
func botMsgRoute(hub *Hub, w http.ResponseWriter, r *http.Request) {

	inBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Infof("%s\n", err)
		return
	}
	log.Infof("[消息路由数据]]=>%s\n", inBytes)

	if inBytes != nil {
		//PUSH:消息路由数据-->前端
		hub.botMsgQueue <- inBytes
		log.Info("PUSH LOG")
	}

	//处理收到的消息
	msgBody := &wulai.MessageRoute{}
	if err := json.Unmarshal(inBytes, msgBody); err != nil {
		log.Errorf("%s\n", err)
	}

	respBody := &wulai.MessageRouteResponses{}
	respBody.IsDispatch = false                            //不转人工
	respBody.SuggestedResponse = msgBody.SuggestedResponse //不处理,直接将消息传回

	outBytes, _ := json.Marshal(respBody)

	log.Info("返回处理后的结果给机器人")
	w.Write(outBytes)
}

//ServeBotMsg bot message handle
func ServeBotMsg(hub *Hub, w http.ResponseWriter, r *http.Request) {

	//request log
	log.Infof("[/]=>remote=>%s host=>%s   url=>%s   method=>%s\n", r.RemoteAddr, r.Host, r.URL, r.Method)

	url := strings.ToLower(r.URL.String())
	switch {
	case url == "/bot/message_route":
		botMsgRoute(hub, w, r)
	default:
		w.Write([]byte("Unknown Pattern"))
	}
}
