package model

import (
	"github.com/AH-dark/apollo/dto"
	"gorm.io/gorm"
	"time"
)

type CommentService interface {
	// GetCommentByID returns a comment by ID.
	GetCommentByID(id uint) (*Comment, error)
	// GetCommentByStatusAndID returns a list of comments by status and ID.
	GetCommentByStatusAndID(commentStatus CommentStatus, id uint) (*Comment, error)
	// GetCommentByRequestID returns a comment by RequestID.
	GetCommentByRequestID(requestID string) (*Comment, error)

	// ListCommentsByStatus returns a list of comments by status.
	ListCommentsByStatus(commentStatus CommentStatus, before time.Time, limit int) ([]*Comment, error)
	// ListCommentsByIP returns a list of comments by IP.
	ListCommentsByIP(ip string, before time.Time, limit int) ([]*Comment, error)

	// UpdateComment updates a comment.
	UpdateComment(comment *Comment) error

	// CreateComment creates a comment.
	CreateComment(comment *Comment) error
	// CreateCommentByDTO creates a comment by DTO.
	CreateCommentByDTO(commentDTO dto.CommentDTO) (*Comment, error)

	// DeleteComment deletes a comment.
	DeleteComment(id uint) error

	// CountCommentsByStatus returns the number of comments by status.
	CountCommentsByStatus(commentStatus CommentStatus) int64
}

type commentService struct {
	db *gorm.DB
}

func NewCommentService(db *gorm.DB) CommentService {
	return &commentService{db: db}
}

func (s *commentService) GetCommentByID(id uint) (*Comment, error) {
	var comment Comment
	if err := s.db.Model(&Comment{}).Where("id = ?", id).First(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (s *commentService) GetCommentByStatusAndID(commentStatus CommentStatus, id uint) (*Comment, error) {
	var comment Comment
	if err := s.db.Model(&Comment{}).Where("status = ? AND id = ?", commentStatus, id).First(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (s *commentService) GetCommentByRequestID(requestID string) (*Comment, error) {
	var comment Comment
	if err := s.db.Model(&Comment{}).Where("request_id = ?", requestID).First(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func (s *commentService) ListCommentsByStatus(commentStatus CommentStatus, before time.Time, limit int) ([]*Comment, error) {
	var comments []*Comment
	if err := s.db.Model(&Comment{}).Where("status = ? AND created_at < ?", commentStatus, before).Limit(limit).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *commentService) ListCommentsByIP(ip string, before time.Time, limit int) ([]*Comment, error) {
	var comments []*Comment
	if err := s.db.Model(&Comment{}).Where("ip = ? AND created_at < ?", ip, before).Limit(limit).Find(&comments).Error; err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *commentService) UpdateComment(comment *Comment) error {
	return s.db.Model(&Comment{}).Where("id = ?", comment.ID).Updates(comment).Error
}

func (s *commentService) CreateComment(comment *Comment) error {
	comment.Status = CommentStatusPending
	return s.db.Model(&Comment{}).Create(comment).Error
}

func (s *commentService) CreateCommentByDTO(commentDTO dto.CommentDTO) (*Comment, error) {
	comment := &Comment{
		Author:    commentDTO.AuthorName,
		Email:     commentDTO.AuthorEmail,
		Content:   commentDTO.Content,
		IP:        commentDTO.IP,
		RequestID: commentDTO.RequestID,
	}

	if err := s.CreateComment(comment); err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) DeleteComment(id uint) error {
	return s.db.Model(&Comment{}).Where("id = ?", id).Delete(&Comment{}).Error
}

func (s *commentService) CountCommentsByStatus(commentStatus CommentStatus) int64 {
	var count int64 = 0
	s.db.Model(&Comment{}).Where("status = ?", commentStatus).Count(&count)
	return count
}
