// think of it as the root directory of our project
// when creating another package when we import it we use the module name as the root directory
// it's like the main path for our packages
module server

go 1.23.4

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/gorilla/mux v1.8.1
	github.com/lib/pq v1.10.9
)

require golang.org/x/crypto v0.31.0
