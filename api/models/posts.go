package models

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/emrekasg/personal-website-api/components"
)

type PostsResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Brief     string    `json:"brief"`
	Tag       string    `json:"tag"`
	Language  string    `json:"language"`
	PostLink  string    `json:"post_link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

const (
	selectPostsQuery = `
		SELECT
			pc.id,
			pc.title,
			pc.brief,
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
			pc.lang = $1 AND p.visible = true
	`

	selectPostQuery = `
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
			p.post_link = $1
		AND
			pc.lang = $2
	`
)

// GetPosts retrieves a paginated list of posts with optional tag filtering
func GetPosts(limit, offset int, language, tag string) ([]PostsResponse, error) {
	var postResponse []PostsResponse

	tagWhere := "AND 1=$2"
	args := []interface{}{language, 1}
	if tag != "" {
		tagWhere = "AND p.tag = $2"
		args = []interface{}{language, tag}
	}

	query := selectPostsQuery + tagWhere + `
		ORDER BY
			pc.created_at DESC
		LIMIT $3
		OFFSET $4
	`

	args = append(args, limit, offset)

	ctx := context.Background()
	rows, err := components.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post PostsResponse
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Brief,
			&post.Language,
			&post.PostLink,
			&post.Tag,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("failed to scan post row: %w", err)
		}
		postResponse = append(postResponse, post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
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

// GetPost retrieves a single post by its link and language
func GetPost(postLink, language string) (PostResponse, error) {
	var post PostResponse

	ctx := context.Background()
	err := components.DB.QueryRowContext(ctx, selectPostQuery, postLink, language).Scan(
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
		if err == sql.ErrNoRows {
			return post, fmt.Errorf("post not found: %w", err)
		}
		return post, fmt.Errorf("failed to query post: %w", err)
	}

	return post, nil
}
