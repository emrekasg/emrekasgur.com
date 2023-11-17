package models

import (
	"time"

	"github.com/emrekasg/personal-website-api/components"
)

type PostsResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Brief     string    `json:"brief"`
	Language  string    `json:"language"`
	PostLink  string    `json:"post_link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetPosts(limit, offset int, language, tag string) ([]PostsResponse, error) {
	var postResponse []PostsResponse

	tagWhere := ""
	args := []interface{}{language}
	if tag != "" {
		tagWhere = "AND p.tag = ?"
		args = append(args, tag)
	}

	query := `
		SELECT
			pc.id,
			pc.title,
			pc.brief,
			pc.lang,
			p.post_link,
			pc.created_at,
			pc.updated_at
		FROM
			posts p
		INNER JOIN
			post_contents pc ON p.id = pc.post_id
		WHERE
			pc.lang = ? ` + tagWhere + `
		ORDER BY
			pc.created_at DESC
		LIMIT ?
		OFFSET ?
	`

	args = append(args, limit, offset)
	rows, err := components.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post PostsResponse
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Brief,
			&post.Language,
			&post.PostLink,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		postResponse = append(postResponse, post)
	}

	return postResponse, nil
}

type PostResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Brief     string    `json:"brief"`
	Content   string    `json:"content"`
	Language  string    `json:"language"`
	PostLink  string    `json:"post_link"`
	Tag       string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetPost(postLink, language string) (PostResponse, error) {
	var post PostResponse

	query := `
		SELECT
			pc.id,
			pc.title,
			pc.brief,
			pc.content,
			pc.lang,
			p.post_link,
			p.tag,
			pc.created_at,
			pc.updated_at
		FROM
			posts p
		INNER JOIN
			post_contents pc ON p.id = pc.post_id
		WHERE
			p.post_link = ?
		AND
			pc.lang = ?
	`

	err := components.DB.QueryRow(query, postLink, language).Scan(
		&post.ID,
		&post.Title,
		&post.Brief,
		&post.Content,
		&post.Language,
		&post.PostLink,
		&post.Tag,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return post, err
	}

	return post, nil
}
