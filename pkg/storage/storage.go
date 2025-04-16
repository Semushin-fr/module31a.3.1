package storage

type Post struct {
	ID      int
	Title   string
	Content string
	PubTime int64
	Link    string
}

type Interface interface {
	Posts(n int) ([]Post, error)
	StorePosts([]Post) error
}
