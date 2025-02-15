package storage

type Post struct {
	ID         int    `json:"id"`
	AuthorID   int    `json:"author_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	AuthorName string `json:"author_name"`
	CreatedAt  int64  `json:"created_at"`
}

type Interface interface {
	Posts() ([]Post, error)
	AddPost(Post) error
	UpdatePost(Post) error
	DeletePost(Post) error
}
