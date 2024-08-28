package tests

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/MosesAliev/comment-service/graph"
	"github.com/go-playground/assert"
)

func TestCreatePost(t *testing.T) {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	c := client.New(h)
	t.Run("Create Post", func(t *testing.T) {
		// struct that saves the response of the mutation below
		var resp struct {
			CreatePost struct {
				commentAllowed bool
				Text           string
				ID             string
			}
		}

		// Create the variable that is passed to addContext

		// Call the resolver via the client and modify the context via functional options
		c.MustPost(`
            mutation createTodo {
  				createPost(input: { text: "Text", CommentAllowed:true}) {
					id
	    			text
  				}
			}`,
			&resp,
		)

		assert.Equal(t, "Text", resp.CreatePost.Text)
	})
}
