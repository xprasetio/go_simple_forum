package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("username or email already exists")
	}
	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	now := time.Now()
	userModel := memberships.UserModel{
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(pass),
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: "system",
		UpdatedBy: "system",
	}
	return s.membershipRepo.CreateUser(ctx, userModel)
}
