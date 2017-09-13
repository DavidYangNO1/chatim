/** * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 * net manager stream
 * generate by DavidYang 2017.9.13
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

package net

import "encoding/json"

//body = version+type+content
const (
	NetP2PMsgCMD           = 30001 // chat to chat
	NetP2PMsgAckCMD        = 30002 // chat to chat ack
	NetGroupMsgCMD         = 30003 // group chat
	NetHeatBeatRequestCMD  = 30004 // heat beat msg request cmd
	NetHeatBeatResponseCMD = 30005 // heat beat msg response cmd
	NetSyncMsgCMD          = 30006 // synnc msg cmd
	NetNotifyMsgCMD        = 30007 // notify msg cmd
	Version                = 1.0
)

// net msg pack
type netMsgPack struct {
	Version float64 `json:"version"`
	MsgType int     `json:"msgType"`
	Content string  `json:"content"`
}

func NewNetMsgPack() *netMsgPack {
	return &netMsgPack{}
}

func (msgPack *netMsgPack) BuildMsgPack(msg string, msgType int) string {
	msgPack.Version = Version
	msgPack.MsgType = msgType
	msgPack.Content = msg
	msgByte, _ := json.Marshal(msgPack)
	return string(msgByte)
}

func (msgPack *netMsgPack) ParseMsg(body string) (msg *string, msgType int, errMsg error) {
	err := json.Unmarshal([]byte(body), &msgPack)
	if err != nil {
		return nil, 0, err
	}
	msg = &msgPack.Content
	msgType = msgPack.MsgType
	errMsg = nil
	return
}
