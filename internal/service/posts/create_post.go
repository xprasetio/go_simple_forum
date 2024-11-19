package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/xprasetio/go_simple_forum.git/internal/model/posts"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHastags, ",") // convert array to string ['hashtag1', 'hashtag2'] -> 'hashtag1,hashtag2'
	now := time.Now()
	model := posts.PostModel{
		UserID:      userID,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHashtags,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   strconv.FormatInt(userID, 10),
		UpdatedBy:   strconv.FormatInt(userID, 10),
	}
	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create post")
		return err
	}
	return nil
}
