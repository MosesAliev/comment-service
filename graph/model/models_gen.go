// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewComment struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
	PostID string `json:"postId"`
}

type NewPost struct {
	Text           string `json:"text"`
	CommentAllowed bool   `json:"CommentAllowed"`
}

type Post struct {
	ID             string `json:"id"`
	Text           string `json:"text"`
	CommentAllowed bool   `json:"CommentAllowed"`
}

type Query struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
