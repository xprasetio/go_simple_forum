package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle   string   `json:"postTitle" binding:"required"`
		PostContent string   `json:"postContent" binding:"required"`
		PostHastags []string `json:"postHastags" binding:"required"`
	}
)

type (
	PostModel struct {
		ID          int64     `db:"id" json:"id"`
		UserID      int64     `db:"user_id" json:"user_id"`
		PostTitle   string    `db:"post_title" json:"post_title"`
		PostContent string    `db:"post_content" json:"post_content"`
		PostHastags string    `db:"post_hastags" json:"post_hastags"`
		CreatedAt   time.Time `db:"created_at" json:"created_at"`
		UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
		CreatedBy   string    `db:"created_by" json:"created_by"`
		UpdatedBy   string    `db:"updated_by" json:"updated_by"`
	}
)
