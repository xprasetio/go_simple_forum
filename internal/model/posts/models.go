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

type (
	GetAllPostResponse struct {
		Data       []Post     `json:"data"`
		Pagination Pagination `json:"pagination"`
	}

	Post struct {
		ID          int64    `json:"id"`
		UserID      int64    `json:"userID"`
		Username    string   `json:"username"`
		PostTitle   string   `json:"postTitle"`
		PostContent string   `json:"postContent"`
		PostHastags []string `json:"postHastags"`
		IsLiked     bool     `json:"isLiked"`
	}
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
	}
	GetPostResponse struct {
		PostDetail Post      `json:"postDetail"`
		LikeCount  int       `json:"likeCount"`
		Comments   []Comment `json:"comments"`
	}
	Comment struct {
		ID             int64  `json:"id"`
		UserID         int64  `json:"user_id"`
		Username       string `json:"username"`
		CommentContent string `json:"commentContent"`
	}
)
