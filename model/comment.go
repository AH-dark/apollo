package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Author  *string `gorm:"size:32"`           // Author 评论作者，允许匿名(<nil>)
	Email   *string `gorm:"size:64;index"`     // Email 评论邮箱，允许匿名(<nil>)
	Content string  `gorm:"size:256;not null"` // Content 评论内容

	Status    CommentStatus `gorm:"not null;index;default:0"` // Status 评论状态
	IP        string        `gorm:"not null;index"`           // IP 评论者 IP
	RequestID string        `gorm:"not null;unique_index"`    // RequestID 评论者 RequestID

	Reply     *string   `gorm:"type:varchar(255)"` // Reply 回复内容
	ReplyTime time.Time // ReplyTime 回复时间

	// TODO: 追问
}

type CommentStatus int

const (
	CommentStatusPending CommentStatus = iota
	CommentStatusSpam
	CommentStatusDeleted
	CommentStatusReplied
)
