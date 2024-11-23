package memberships

import (
	"context"
	"database/sql"
	"time"

	"github.com/xprasetio/go_simple_forum.git/internal/model/memberships"
)

func (r *repository) InsertRefreshToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*memberships.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expired_at, created_at, updated_at, created_by, updated_by FROM refresh_tokens
    WHERE user_id = ? AND expired_at >= ?`

	var result memberships.RefreshTokenModel
	row := r.db.QueryRowContext(ctx, query, userID, now)
	err := row.Scan(&result.ID, &result.UserID, &result.RefreshToken, &result.ExpiredAt, &result.CreatedAt, &result.UpdatedAt, &result.CreatedBy, &result.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}
