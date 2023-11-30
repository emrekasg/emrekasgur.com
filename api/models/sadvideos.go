package models

import "github.com/emrekasg/personal-website-api/components"

type SadVideosResponse struct {
	ID        int    `json:"id"`
	VideoLink string `json:"video_link"`
}

func GetRandomSadVideo() (string, error) {
	var sadVideo string

	query := `
		SELECT
			video_link
		FROM
			sad_videos
		ORDER BY RAND()
		LIMIT 1
	`

	err := components.DB.QueryRow(query).Scan(&sadVideo)

	if err != nil {
		return "", err
	}

	return sadVideo, nil
}
