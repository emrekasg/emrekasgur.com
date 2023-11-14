package models

type PostContent struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	PostLink  string  `json:"post_link"`
	Content   string  `json:"content"`
	Brief     string  `json:"brief"`
	Language  string  `json:"language"`
	CreatedAt []uint8 `json:"created_at"`
	UpdatedAt []uint8 `json:"updated_at"`
	PostId    int     `json:"post_id"`
}
