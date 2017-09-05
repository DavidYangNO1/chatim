/** * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 * chat message
 * generate by DavidYang 2017.9.5
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package model

import "time"

const (
	MsgReceived = 1 // msg has received
	MsgReaded   = 2 // msg has readed
)

// im Message info model
type Message struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	To        string    `json:"to" gorm:"type:varchar(255)"`     //The message send to who, is a userid
	MsgId     string    `json:"msgId" gorm:"type:varchar(32)"`   //The msg id, unique
	GroupId   string    `json:"groupId" gorm:"type:varchar(32)"` //  server groupid unique
	Msg       string    `json:"msg" gorm:"type:varchar(1024)"`   //the chat msg
	MsgType   int       `json:"msgType"`                         //msg type
	FileURL   string    `json:"fileURL" gorm:"type:varchar(64)"` //file url
	MsgStatus int       `json:"msgStatus"`                       //message statuss
}
