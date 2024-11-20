package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/xprasetio/go_simple_forum.git/internal/model/posts"
)

func (s *service) UpsertUserActivity(ctx context.Context, postID, userID int64, request posts.UserActivityRequest) error {
	now := time.Now().UTC()
	model := posts.UserActivityModel{
		PostID:    postID,
		UserID:    userID,
		IsLiked:   request.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepo.GetUserActivity(ctx, model)
	if err != nil {
		log.Error().Msg("error get user activity from database")
		return err
	}
	if userActivity == nil {
		if request.IsLiked {
			return errors.New("Anda Belum pernah like postingan ini")
		}
		err = s.postRepo.CreateUserActivity(ctx, model)
		if err != nil {
			log.Error().Msg("error create user activity")
			return err
		}
		return nil
	}
	err = s.postRepo.UpdateUserActivity(ctx, model)
	if err != nil {
		log.Error().Msg("error update user activity")
		return err
	}
	return nil
}
