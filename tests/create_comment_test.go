package tests

import (
	"fmt"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/MosesAliev/comment-service/graph"
	"github.com/go-playground/assert"
)

func TestCreateComment(t *testing.T) {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	c := client.New(h)
	t.Run("Create Comment", func(t *testing.T) {
		// Структура сохраняет тело ответа
		var respPost struct {
			CreatePost struct {
				commentAllowed bool
				Text           string
				ID             string
			}
		}

		c.MustPost(`
            mutation create {
  				createPost(input: { text: "Text", CommentAllowed:true}) {
					id
	    			text
  				}
			}`,
			&respPost,
		)

		var respComment struct {
			CreateComment struct {
				Text   string `json:"text"`
				UserId string `json:"userId"`
				PostId string `json:"postId"`
			}
		}

		c.MustPost(fmt.Sprintf(`
            mutation create {
  				createComment(input: { text: "Text", userId: "1", postId: "%s"}) {
	    			text
  				}
			}`, respPost.CreatePost.ID),
			&respComment,
		)
		assert.Equal(t, "Text", respComment.CreateComment.Text)
	})
}
