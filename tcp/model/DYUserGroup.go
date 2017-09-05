/** * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 *
 * user group info
 * generate by DavidYang 2017.9.5
 *
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */
package model

import "time"

// user group info
type UserGroup struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	UserId    string    `json:"userid" gorm:"type:varchar(32)"`   //user id unique
	GroupId   string    `json:"imageUrl" gorm:"type:varchar(64)"` //user name
}
