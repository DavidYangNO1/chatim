/** * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 * group chat room
 * generate by DavidYang 2017.9.5
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package model

import "time"

const (
	ChatRoomPrivate   = 1 //Private group : Other users are not able to join this group voluntarily
	ChatRoomPublic    = 2 //Public Group : Users are able to search and join the group
	ChatRoomBroadcast = 3 //Group : User can send personal message to a group of Users
	ChatRoomOpenGroup = 4 //Open Group : Any user can send the message to this group without being the member of this group
)

// chat room model
type ChatRoom struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	GroupId   string    `json:"groupId" gorm:"type:varchar(32)"`   //  server groupid unique
	GroupName string    `json:"groupName" gorm:"type:varchar(32)"` //  group name                  //
	GroupType int       `json:"groupType" gorm:"size:4"`           //  1 Private group; 2 Public Group; 3	Broadcast Group;  4 Open Group
	ImageUrl  string    `json:"imageUrl" gorm:"type:varchar(64)"`  //  chat room head pic
	UserCount int       `json:"userCount" gorm:"size:4"`           //  chat room of user counts
	OwnerId   string    `json:"ownerId" gorm:"type:varchar(32)"`   //  chat room owner
}
