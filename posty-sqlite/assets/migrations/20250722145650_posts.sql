-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME NOT NULL
);

INSERT INTO posts (title, content, created_at) VALUES
('Post 1', 'This is the content of post 1.', '2025-07-22 10:00:00'),
('Post 2', 'This is the content of post 2.', '2025-07-22 10:05:00'),
('Post 3', 'This is the content of post 3.', '2025-07-22 10:10:00'),
('Post 4', 'This is the content of post 4.', '2025-07-22 10:15:00'),
('Post 5', 'This is the content of post 5.', '2025-07-22 10:20:00'),
('Post 6', 'This is the content of post 6.', '2025-07-22 10:25:00'),
('Post 7', 'This is the content of post 7.', '2025-07-22 10:30:00'),
('Post 8', 'This is the content of post 8.', '2025-07-22 10:35:00'),
('Post 9', 'This is the content of post 9.', '2025-07-22 10:40:00'),
('Post 10', 'This is the content of post 10.', '2025-07-22 10:45:00');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
