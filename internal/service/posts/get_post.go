package posts

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/xprasetio/go_simple_forum.git/internal/model/posts"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get post by id")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error count like by post id")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("error get comment by post id")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:          postDetail.ID,
			UserID:      postDetail.UserID,
			Username:    postDetail.Username,
			PostTitle:   postDetail.PostTitle,
			PostContent: postDetail.PostContent,
			PostHastags: postDetail.PostHastags,
			IsLiked:     postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:   comments,
	}, nil
}
