package post

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/holandalhs/APIGOLANG/internally"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	Conn *pgxpool.Pool
}

func (r *Repository) Insert(post internally.Post) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println(r.Conn)

	_, err := r.Conn.Exec(
		ctx,
		"INSERT INTO posts (username, body) VALUES ($1, $2)",
		post.Username,
		post.Body)

	return err
}

func (r *Repository) Delete(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tag, err := r.Conn.Exec(
		ctx,
		"DELETE FROM posts WHERE id = $1",
		id)

	if tag.RowsAffected() == 0 {
		return ErrPostNotFound
	}

	return err
}

func (r *Repository) FindOneByID(id uuid.UUID) (internally.Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var post internally.Post
	err := r.Conn.QueryRow(
		ctx,
		"SELECT username, body, created_at FROM posts WHERE id = $1",
		id).Scan(&post.Username, &post.Body, &post.CreatedAt)

	if err == pgx.ErrNoRows {
		return internally.Post{}, ErrPostNotFound
	}

	if err != nil {
		return internally.Post{}, err
	}

	return post, nil
}
