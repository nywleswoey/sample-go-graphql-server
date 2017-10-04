package schema

import (
	"github.com/graphql-go/graphql"
)

// RootQuery is a pointer of the RootQuery graphql object
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootQuery",
  Fields: graphql.Fields{
    "lastTodo": &graphql.Field{
      Type: TodoType,
    },
  },
})
