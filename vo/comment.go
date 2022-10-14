package vo

import (
	"github.com/AH-dark/apollo/model"
	"github.com/AH-dark/apollo/pkg/hashids"
	"github.com/samber/lo"
	"time"
)

type CommentVO struct {
	ID          string  `json:"id" xml:"ID"`
	AuthorName  *string `json:"author_name" xml:"AuthorName"`
	AuthorEmail *string `json:"author_email" xml:"AuthorEmail"`
	Content     string  `json:"content" xml:"Content"`
}

func BuildCommentVO(comment *model.Comment) CommentVO {
	return CommentVO{
		ID:          hashids.Encode(comment.ID, hashids.CommentHash),
		AuthorName:  comment.Author,
		AuthorEmail: comment.Email,
		Content:     comment.Content,
	}
}

type ListCommentsVO struct {
	Comments []CommentVO `json:"comments" xml:"Comments"`
	Total    int64       `json:"total" xml:"Total"`
	Before   time.Time   `json:"before" xml:"Before"`
}

func BuildCommentVOList(comments []*model.Comment, total int64, before time.Time) ListCommentsVO {
	return ListCommentsVO{
		Comments: lo.Map(comments, func(comment *model.Comment, i int) CommentVO {
			return BuildCommentVO(comment)
		}),
		Total:  total,
		Before: before,
	}
}
