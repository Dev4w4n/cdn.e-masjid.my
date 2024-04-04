package model

type Response struct {
	ID        int64  `json:"id"`
	Path      string `json:"path"`
	CreatedAt int64  `json:"created_at"`
}
