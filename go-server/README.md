Steps to run:
1. shetty@SIT-SMBP-672Q5P GoRepo % cd go-server
2. shetty@SIT-SMBP-672Q5P go-server % go build github.com/vijishetty29/Go/go-server
3. shetty@SIT-SMBP-672Q5P go-server % go run ./main.go
   Starting server on port 3000!
4. Visit http://localhost:3000
5. Visit http://localhost:3000/form.html

With Docker:
1. Docker build -t go_demo .
2. docker run -p 3000:3000 go_demo
