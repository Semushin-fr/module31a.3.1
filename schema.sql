DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  link TEXT NOT NULL,
  created_at BIGINT NOT NULL
);

INSERT INTO posts (id, title, content, link, created_at) VALUES (0, 'Статья', 'Содержание статьи', 'https://google.com', 0);