package models

import "github.com/emrekasg/personal-website-api/components"

type Tags []Tag

type Tag struct {
	Name      string `json:"name"`
	PostCount int    `json:"post_count"`
}

func GetTags() (Tags, error) {
	var tags Tags

	query := `
		SELECT
			p.tag,
			COUNT(1) AS post_count
		FROM
			posts p		
		GROUP BY
			p.tag
		ORDER BY
			post_count DESC
	`

	rows, err := components.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.Name,
			&tag.PostCount,
		)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}
