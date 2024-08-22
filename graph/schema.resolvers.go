package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/MosesAliev/comment-service/graph/model"
)

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	for _, post := range r.posts {
		if post.ID == input.PostID {
			if !post.CommentAllowed {
				var err error
				comment := &model.Comment{
					Text: "NotAllowed",
					ID:   "",
					User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
					Post: &model.Post{ID: input.PostID},
				}
				return comment, err
			}

		}

	}

	comment := &model.Comment{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randNumber),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
		Post: &model.Post{ID: input.PostID},
	}
	r.comments = append(r.comments, comment)
	return comment, nil
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	post := &model.Post{
		Text:           input.Text,
		ID:             fmt.Sprintf("T%d", randNumber),
		CommentAllowed: input.CommentAllowed,
	}
	r.posts = append(r.posts, post)
	return post, nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	// var comments []*model.Comment
	// for _, comment := range r.comments {
	// 	if comment.PostID == id {
	// 		comments = append(comments, comment)
	// 	}

	// }

	return r.comments, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

// CommentsByID is the resolver for the commentsById field.
func (r *queryResolver) CommentsByID(ctx context.Context, id string) ([]*model.Comment, error) {
	r.commentsById = nil
	for _, comment := range r.comments {
		if comment.Post.ID == id {
			log.Println(comment.ID)
			r.commentsById = append(r.commentsById, comment)
		}

	}
	return r.commentsById, nil
}

// Comment returns CommentResolver implementation.

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
