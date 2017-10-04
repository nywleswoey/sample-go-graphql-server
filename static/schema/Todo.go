package schema

import (
	"github.com/graphql-go/graphql"
)

// TodoType is a pointer of the Todo graphql object
var TodoType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Todo",
  Fields: graphql.Fields{
    "id": &graphql.Field{
      Type: graphql.String,
    },
    "text": &graphql.Field{
      Type: graphql.String,
    },
    "done": &graphql.Field{
      Type: graphql.Boolean,
    },
  },
})
