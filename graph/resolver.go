package graph

import "github.com/MosesAliev/comment-service/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	comments     []*model.Comment
	posts        []*model.Post
	commentsById []*model.Comment
}
