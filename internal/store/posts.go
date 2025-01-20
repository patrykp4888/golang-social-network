package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	UserID    int64    `json:"userId"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {
	query := `INSERT INTO posts (content, title, user_id, tags)
              VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostsStore) GetByID(ctx context.Context, postId int64) (*Post, error) {
	query := `SELECT id, user_id, title, content, tags, created_at, updated_at FROM posts WHERE id = $1`

	var post Post
	err := s.db.QueryRowContext(ctx, query, postId).Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		pq.Array(&post.Tags),
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
