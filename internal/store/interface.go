//go:generate mockgen -source=interface.go -destination=./mocks/mock.go -package=mock
package store

import "context"

type Posts interface {
	Create(ctx context.Context, post *Post) error
}

type Users interface {
	Create(ctx context.Context, user *User) error
}
