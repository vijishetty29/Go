Steps to run:
1. shetty@System GoRepo % cd go-crud-movies
2. shetty@System go-server % go build go-crud-movies
3. shetty@System go-server % go run ./main.go
   Starting server on port 3000!
4. Visit http://localhost:3000
5. Run the following postman collection for APIs **GO-CRUD-MOVIES.postman_collection.json**

With Docker:
1. Docker build -t go-crud-movies .
2. docker run -p 3000:3000 go-crud-movies
