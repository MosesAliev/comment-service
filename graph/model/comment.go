package model

type Comment struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	UserID string `json:"userId"`
	User   *User  `json:"user"`
	PostID string `json:"postId"`
	Post   *Post  `json:"post"`
}
