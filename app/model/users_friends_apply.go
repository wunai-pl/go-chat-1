package model

import "time"

// 用户添加好友申请表
type UsersFriendsApply struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"` // 申请ID
	UserId    int       `gorm:"column:user_id;default:0" json:"user_id"`        // 申请人ID
	FriendId  int       `gorm:"column:friend_id;default:0" json:"friend_id"`    // 被申请人
	Remark    string    `gorm:"column:remark" json:"remark"`                    // 申请备注
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`            // 申请时间
}

// 用户添加好友申请表
type ApplyListItem struct {
	Id        int       `gorm:"column:id" json:"id" form:"id"`                         // 申请ID
	UserId    int       `gorm:"column:user_id" json:"user_id" form:"user_id"`          // 申请人ID
	FriendId  int       `gorm:"column:friend_id" json:"friend_id" form:"friend_id"`    // 被申请人
	Remark    string    `gorm:"column:remark" json:"remark" form:"remark"`             // 申请备注
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"` // 申请时间
}
