# go-auth

HTTP Basic Auth template with Go.

### How to run

1. Make .env file, set admin user and password. Template is available in .env.example   
2. Execute go run *.go

### Endpoints

1. `/user`: get list of user
2. `/user?id=a01`: get user with id a01

Example: execute `curl -X GET --user <username>:<password> http://localhost:9000/user` on command line/terminal