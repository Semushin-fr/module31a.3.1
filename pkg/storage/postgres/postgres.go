package postgres

import (
	"GoNews/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	db *pgxpool.Pool
}

func New(constr string) (*postgres, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	p := postgres{
		db: db,
	}
	return &p, nil
}

func (p *postgres) Posts() ([]storage.Post, error) {
	rows, err := p.db.Query(context.Background(), `
		SELECT 
			posts.*,
			authors.name AS author_name
		FROM posts
		JOIN authors
		ON authors.id = posts.author_id;
	`)

	if err != nil {
		return nil, err
	}

	var posts []storage.Post

	for rows.Next() {
		var post storage.Post

		err := rows.Scan(
			&post.ID,
			&post.AuthorID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.AuthorName,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, rows.Err()
}

func (p *postgres) AddPost(post storage.Post) error {
	_, err := p.db.Exec(context.Background(), `
			INSERT INTO posts (author_id, title, content, created_at)
			VALUES ($1, $2, $3, $4);
		`,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *postgres) UpdatePost(post storage.Post) error {
	_, err := p.db.Exec(context.Background(), `
		UPDATE posts 
		SET author_id = $1, title = $2, content = $3, created_at = $4
		WHERE id = $5
	`,
		post.AuthorID,
		post.Title,
		post.Content,
		post.CreatedAt,
		post.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (p *postgres) DeletePost(post storage.Post) error {
	_, err := p.db.Exec(context.Background(), `
		DELETE FROM posts
		WHERE id = $1
	`, post.ID)

	if err != nil {
		return err
	}

	return nil
}
