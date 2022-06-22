# go_auth_sf
Golang version of auth_sf

# run
To run project you can:

`go run cmd/main.go` - to run it by your system GO

`make up`/`make down` - to use make tool for building and running project containers

`docker-compose up --build`/`docker-compose down --remove-orphans` - for running project directly with compose tool

When running project from containers make sure that you use .env.dev default file, instead of .env.
Use .env only for debugging and developing in case you run project with GO runner

# docs
`http://localhost:8000/swagger/index.html` - swagger API documentation
