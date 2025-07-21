package sqlite

import (
	"database/sql"
	"github.com/levintoo/posty/internal/models"
)

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title, content string) (error) {
	stmt := `INSERT INTO posts (title, content, created_at) VALUES (?, ?, datetime('now'))`
	_, err := m.DB.Exec(stmt, title, content)
	return err
}

func (m *PostModel) All() ([]models.Post, error) {
	stmt := `SELECT id, title, content, created_at FROM posts ORDER BY created_at DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	posts := []models.Post{}
	for rows.Next() {
		p := models.Post{}
		err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
