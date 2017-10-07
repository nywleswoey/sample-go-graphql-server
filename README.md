# GraphQL server using Go and Postgresql

This is a sample GraphQL server implemented using GO and Postgresql(hosted on Docker).
This implementation includes a GraphQL API explorer(GraphiQL) as well.

Credits
Chris Ramon and Hafiz Ismail for their good work in implementing [graphl-go](https://github.com/graphql-go/graphql) and [graphql-go-handler](https://github.com/graphql-go/handler)
Alexandru Topliceanu for his [article](http://alexandrutopliceanu.ro/post/graphql-with-go-and-postgresql) and [sample codes](https://github.com/topliceanu/graphql-go-example) for integrating GraphQL with Postgresql

## How to run this

To run this project, install Go: [https://golang.org/dl/](https://golang.org/dl/)
Then install [Masterminds/glide](https://github.com/Masterminds/glide) which is a package manager for golang projects.

Install all the dependencies for the project.

Install [Docker](https://docs.docker.com/engine/installation/) for Postgresql (or do a local installation if you prefer).
Create a database called test.

Install [sql-migrate](https://github.com/rubenv/sql-migrate) which is a tool to run migrations against sql databases.

The default settings for the database are
host: localhost
db name: test
user: postgres
password: not required

If you want to make changes, just edit in main.go and dbconfig.yml

To run the migrations which will create the database tables and indexes `$GOPATH/bin/sql-migrate up`. If you ever want to clean up the the database run `$GOPATH/bin/sql-migrate down` then `$GOPATH/bin/sql-migrate up` again.


