package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Comments  []Comment `json:"comments"`
	Version   int
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags)
	VALUES ($1,$2,$3,$4) RETURNING id, created_at, updated_at, version
	`
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
		&post.Version,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetByID(ctx context.Context, postID int64) (*Post, error) {
	var post Post

	query := `
	SELECT id, title, user_id, content, created_at, tags, updated_at, version FROM posts WHERE id = $1 LIMIT 1;
	`

	err := s.db.QueryRowContext(ctx, query, postID).
		Scan(
			&post.ID,
			&post.Title,
			&post.UserID,
			&post.Content,
			&post.CreatedAt,
			pq.Array(&post.Tags),
			&post.UpdatedAt,
			&post.Version,
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

func (s *PostStore) DeleteByID(ctx context.Context, postID int64) error {

	query := `DELETE FROM posts WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query, postID)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	query := `
	UPDATE posts
		SET
			title = $1,
			content = $2,
			tags = $3,
			updated_at = NOW(),
			version = version + 1
		WHERE
			id = $4 AND version = $5
	RETURNING updated_at, version
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		pq.Array(post.Tags),
		post.ID,
		post.Version,
	).Scan(
		&post.UpdatedAt,
		&post.Version,
	)

	return err
}
