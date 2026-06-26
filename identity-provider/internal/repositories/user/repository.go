package userrepo

import (
	"context"
	_ "embed" //
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	userdmn "github.com/uniqelus/opentrade/identity-provider/internal/domains/user"
	usersrv "github.com/uniqelus/opentrade/identity-provider/internal/services/user"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

var _ usersrv.UserRepository = &Repository{}

//go:embed queries/upsert_user.sql
var upsertUserQuery string

// SaveUser implements [usersrv.UserRepository].
func (r *Repository) SaveUser(ctx context.Context, user *userdmn.User) error {
	dto, err := toDTO(user)
	if err != nil {
		return fmt.Errorf("failed to map domain user to dto: %w", err)
	}

	_, err = r.pool.Exec(ctx, upsertUserQuery,
		dto.ID,
		dto.Email,
		dto.FirstName,
		dto.LastName,
		dto.State,
		dto.CreateTime,
		dto.UpdateTime,
	)
	if err != nil {
		return fmt.Errorf("failed to execute save user query: %w", err)
	}

	return nil
}
