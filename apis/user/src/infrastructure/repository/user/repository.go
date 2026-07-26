package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"

	domainuser "github.com/sukuname4976/portfolio/apis/user/src/domain/entities/user"
	userrepo "github.com/sukuname4976/portfolio/apis/user/src/domain/repository-interfaces/user"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/email"
	"github.com/sukuname4976/portfolio/apis/user/src/domain/value-objects/userid"
	sqlc "github.com/sukuname4976/portfolio/apis/user/src/infrastructure/auto-generated-by-sqlc"
)

// コンパイル時にインターフェース実装を検証
var _ userrepo.Repository = (*Repository)(nil)

// Repository sqlc 生成コードをラップして domain の Repository を実装する。
type Repository struct {
	queries *sqlc.Queries
}

// NewRepository Repository を生成する。db は pgxpool.Pool など DBTX を満たすもの。
func NewRepository(db sqlc.DBTX) *Repository {
	return &Repository{queries: sqlc.New(db)}
}

// Create 新規ユーザーを INSERT し、DB が発行した ID を含むユーザーを返す。
func (r *Repository) Create(ctx context.Context, name string, mail email.Email) (*domainuser.User, error) {
	row, err := r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Email: mail.Value(),
		Name:  name,
	})
	if err != nil {
		// email の UNIQUE 制約違反 (23505) はドメイン非依存のセンチネルに変換
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, userrepo.ErrDuplicateEmail
		}
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return toDomain(row.ID, row.Email, row.Name)
}

// FindByID ID でユーザーを取得する。存在しない場合は userrepo.ErrNotFound を返す。
func (r *Repository) FindByID(ctx context.Context, id userid.UserID) (*domainuser.User, error) {
	pgID, err := parseUUID(id.Value())
	if err != nil {
		// UUID として不正な ID は該当ユーザーなしとして扱う
		return nil, userrepo.ErrNotFound
	}

	row, err := r.queries.GetUser(ctx, pgID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, userrepo.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return toDomain(row.ID, row.Email, row.Name)
}

// toDomain DB の行データを domain エンティティに変換する。
func toDomain(id pgtype.UUID, mailStr, name string) (*domainuser.User, error) {
	uid, err := userid.New(uuid.UUID(id.Bytes).String())
	if err != nil {
		return nil, err
	}
	mail, err := email.New(mailStr)
	if err != nil {
		return nil, err
	}
	return domainuser.New(uid, name, mail), nil
}

// parseUUID 文字列の ID を pgtype.UUID に変換する。
func parseUUID(s string) (pgtype.UUID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return pgtype.UUID{}, err
	}
	return pgtype.UUID{Bytes: parsed, Valid: true}, nil
}
