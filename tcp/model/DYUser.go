/** * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 * group chat room
 * generate by DavidYang 2017.9.5
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package model

import "time"

const (
	UserStatusOffline = 1 //user offline
	UserStatusOnline  = 2 //user online
)

// im user info model
type User struct {
	ID             uint      `gorm:"primary_key"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	LastSeenAtTime time.Time `json:"lastSeenAtTime"`
	UserId         string    `json:"userId" gorm:"type:varchar(32)"`    //user id unique
	UserName       string    `json:"userName" gorm:"type:varchar(255)"` //user name
	Phone          string    `json:"phone"`
	PwsSalt        []byte    `json:"-"`
	PwsHash        []byte    `json:"-"`
	ImageUrl       string    `json:"lastSeenAtTime" gorm:"type:varchar(64)"` //user head image url
	UserStatus     int       `json:"userStatus"`                             //UserStatusOffline or UserStatusOnline
}
