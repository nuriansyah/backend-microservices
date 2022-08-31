package repository

import (
	"database/sql"
	"errors"
	"time"
)

type Posts struct {
	ID            int            `db:"id"`
	AuthorID      int            `db:"author_id"`
	AuthorName    string         `db:"author_name"`
	AuthorRole    string         `db:"author_role"`
	AuthorAvatar  sql.NullString `db:"author_avatar"`
	AuthorProgram sql.NullString `db:"author_program"`
	AuthorCompany sql.NullString `db:"author_company"`
	AuthorBatch   sql.NullInt32  `db:"author_batch"`
	Title         string         `db:"title"`
	Description   string         `db:"desc"`
	CreatedAt     time.Time      `db:"created_at"`
}

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

var (
	ErrPostNotFound = errors.New("Post not found")
)

func (p *PostRepository) InserPost(authorID int, title, description string) (int64, error) {
	sqlStatement := "INSERT INTO posts (author_id, title, desc, created_at) VALUES ( ?, ?, ?, ?);"
	tx, err := p.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(sqlStatement, authorID, title, description, time.Now())
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		return 0, err
	}
	return id, err
}

func (p *PostRepository) FetchPostByID(postID, authorID int) ([]Posts, error) {
	var (
		posts        []Posts
		sqlStatement string
	)
	sqlStatement = `SELECT 
							p.id as id,
							u.id as author_id,
							u.name as author_name,
							u.role as author_role,
							ud.program as author_program,
							ud.company as author_comapany,
							ud.batch as author_batch,
							p.title as title,
							p.desc as desc,
							p.created_at as created_at
					FROM posts p
					INNER JOIN users u ON p.author_id = u.id
					LEFT JOIN user_details ud ON u.id = ud.user_id
					WHERE p.id = ?`
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.Query(sqlStatement, postID, authorID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var post Posts
		err := rows.Scan(
			&post.ID, &post.AuthorID, &post.AuthorName, &post.AuthorRole,
			&post.AuthorCompany, &post.AuthorProgram, &post.AuthorBatch,
			&post.Title, &post.Description, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *PostRepository) FetchAuthorIDByPostID(postID int) (int, error) {
	sqlStatement := `
		SELECT author_id FROM posts WHERE id = ?;
	`

	tx, err := p.db.Begin()

	if err != nil {
		return 0, err
	}

	defer tx.Rollback()

	var authorID int
	err = tx.QueryRow(sqlStatement, postID).Scan(&authorID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrPostNotFound
		}

		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return authorID, nil
}
