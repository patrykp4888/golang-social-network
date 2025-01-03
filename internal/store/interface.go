package store

import "context"

type Posts interface {
	Create(ctx context.Context) error
}

type Users interface {
	Create(ctx context.Context) error
}
