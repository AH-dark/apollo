package dto

type CommentDTO struct {
	SubmitCommentDTO
	IP        string `json:"ip" xml:"IP" form:"ip" binding:"required"`
	RequestID string `json:"request_id" xml:"RequestID" form:"request_id" binding:"required"`
}

type SubmitCommentDTO struct {
	AuthorName  *string `json:"author_name,omitempty" xml:"AuthorName,omitempty" form:"author_name" binding:"omitempty,max=32"`
	AuthorEmail *string `json:"author_email,omitempty" xml:"AuthorEmail,omitempty" form:"author_email" binding:"omitempty,email,max=64"`
	Content     string  `json:"content" xml:"Content" form:"content" binding:"required,max=256"`
}

type ListCommentsDTO struct {
	Before   *int64 `json:"before,omitempty" xml:"Before,omitempty" form:"before" query:"before" binding:"omitempty,gt=0"` // Before unix micro timestamp
	PageSize int    `json:"page_size" xml:"PageSize" form:"page_size" query:"page_size" binding:"required,gt=0,lte=50"`
}
