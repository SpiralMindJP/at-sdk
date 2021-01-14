package auth

import (
	"context"
)

type userCtxKey struct{}

// UserFromContext は context.Context にセットされた User を取得します。
func UserFromContext(ctx context.Context) User {
	u, ok := ctx.Value(userCtxKey{}).(User)
	if !ok {
		return nil
	}

	return u
}

// NewContextWithUser は u を設定した parent のコピーを返します。
func NewContextWithUser(parent context.Context, u User) context.Context {
	return context.WithValue(parent, userCtxKey{}, u)
}
