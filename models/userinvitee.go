package models

import (
	"fmt"
)

func AddUserInvitee(id, uid int64) {
	dt := make(map[string]interface{})
	dt["id"] = id
	dt["uid"] = uid
	fmt.Println(dt)
	DB.Table("user_invitees").Create(dt)
}

func RmvUserInvitee() {

}
