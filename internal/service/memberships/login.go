package memberships

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
	"github.com/xprasetio/go_simple_forum.git/pkg/jwt"
	tokenUtil "github.com/xprasetio/go_simple_forum.git/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, string, error) { //asumsi string pertama adalah token, string kedua adalah refresh token

	user, err := s.membershipRepo.GetUser(ctx, req.Email, "", 0)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return "", "", err
	}
	if user == nil {
		log.Error().Msg("User not found")
		return "", "", errors.New("user not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("Email or password invalid")
		return "", "", errors.New("Email or password invalid")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create token")
		return "", "", err
	}

	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, user.ID, time.Now())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get refresh token")
		return "", "", err
	}
	if existingRefreshToken != nil {
		return token, existingRefreshToken.RefreshToken, nil
	}

	refreshToken := tokenUtil.GenerateRefreshToken()
	if refreshToken == "" {
		log.Error().Msg("Failed to generate refresh token")
		return token, "", errors.New("failed to generate refresh token")
	}

	err = s.membershipRepo.InsertRefreshToken(ctx, memberships.RefreshTokenModel{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		ExpiredAt:    time.Now().Add(time.Hour * 24),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    strconv.FormatInt(user.ID, 10),
		UpdatedBy:    strconv.FormatInt(user.ID, 10),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to insert refresh token")
		return token, "", err
	}

	return token, refreshToken, nil
}
