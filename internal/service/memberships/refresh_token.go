package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
	"github.com/xprasetio/go_simple_forum.git/pkg/jwt"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token fron database")
		return "", err
	}
	if existingRefreshToken == nil {
		return "", errors.New("refresh has expired")
	}
	if existingRefreshToken.RefreshToken != request.Token {
		return "", errors.New("invalid refresh token is invalid")
	}
	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return "", err
	}
	if user == nil {
		log.Error().Msg("User not found")
		return "", errors.New("user not found")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", err
	}
	return token, nil
}
