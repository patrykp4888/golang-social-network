package store

import "context"

type Posts interface {
	Create(ctx context.Context, post *Post) error
}

type Users interface {
	Create(ctx context.Context, user *User) error
}
