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

func (p *postgres) Posts(n int) ([]storage.Post, error) {
	if n == 0 {
		n = 10
	}
	rows, err := p.db.Query(context.Background(), `
	SELECT id, title, content, pub_time, link FROM posts
	ORDER BY pub_time DESC
	LIMIT $1
	`,
		n,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.PubTime,
			&p.Link,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (p *postgres) StorePosts(posts []storage.Post) error {
	for _, post := range posts {
		_, err := p.db.Exec(context.Background(), `
		INSERT INTO posts(title, content, pub_time, link)
		VALUES ($1, $2, $3, $4)`,
			post.Title,
			post.Content,
			post.PubTime,
			post.Link,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
