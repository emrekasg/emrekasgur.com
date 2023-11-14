package models

import (
	"github.com/emrekasg/personal-website-api/components"
)

type PostContentResponse struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Brief     string  `json:"brief"`
	Language  string  `json:"language"`
	PostLink  string  `json:"post_link"`
	CreatedAt []uint8 `json:"created_at"`
	UpdatedAt []uint8 `json:"updated_at"`
}

func GetPosts(limit, offset int, language string) ([]PostContentResponse, error) {
	var postResponse []PostContentResponse
	var postContent []PostContent

	query := `
		SELECT
			pc.id,
			pc.title,
			pc.content,
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
			pc.lang = ?
		ORDER BY
			pc.created_at DESC
		LIMIT ?
		OFFSET ?
	`

	rows, err := components.DB.Query(query, language, limit, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var content PostContent
		err := rows.Scan(
			&content.ID,
			&content.Title,
			&content.Content,
			&content.Brief,
			&content.Language,
			&content.PostLink,
			&content.CreatedAt,
			&content.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		postContent = append(postContent, content)
	}

	for _, content := range postContent {
		postResponse = append(postResponse, PostContentResponse{
			ID:        content.ID,
			Title:     content.Title,
			Content:   content.Content,
			Brief:     content.Brief,
			Language:  content.Language,
			PostLink:  content.PostLink,
			CreatedAt: content.CreatedAt,
			UpdatedAt: content.UpdatedAt,
		})
	}

	return postResponse, nil
}
