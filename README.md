# GO-HTTP-SERVER

- Simple HTTP Server in Go for study ([boot.dev](https://boot.dev))

## Goal of This Project

- Understand what web server are and how they power real-world web application
- Build a production-style HTTP server in Go, without the use of `framework`
- Use JSON, headers, and status codes to communicate with clients via a
  `RESTful API`
- Use type safe `SQL` to store and retrieve data from a Postgres database
- Implement a secure `authentication/authorization` system with well-tested
  cryptography libraries
- Build and understand webhooks and API keys
- Document the `REST API` with markdown

## Installation

- goose

## Scripts

- For better testability, I added some test scripts files. enjoy!
- look at the `scripts/` directory

```shell
❯ scripts/test_users.sh
===== Test 0: Reset, =====
URL: http://localhost:8080/admin/reset
Data:
Response: Hits reset to 0, users successfully deleted
Status Code: 200

-----------------------------------------
===== Test 1: Creating a valid user =====
URL: http://localhost:8080/api/users
Data: {
  "email": "john@example.com"
```

## Tests

- But there are some limitations for testing by .sh scripts.
- That's why I added integration tests files in `./tests/` dir.
- run this command: `go test ./... -v` (-v can show the log)
