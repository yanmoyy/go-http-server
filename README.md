# GO-HTTP-SERVER

- Simple HTTP Server in Go (Chirpy) ([boot.dev](https://boot.dev))

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

### Chirpy

In this project, I built `"Chirpy"`. Chirpy is a social network similar to
Twitter. (`tweet` -> `chirp`)

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

## API

this is the sample documentation practice.

You can find each endpoints and functions from `main.go`

- `/app/`: endpoint for file directory.
- `GET` `/app/assets/logo.png` : logo file

- `GET` `/api/healthz` : just check server is OK (status 200)

### User

`User Data`

```json
{
  "id": "16577ee6-5cbf-4ff9-a720-ec2c4c7b644c",
  "created_at": "2025-05-27 23:32:23.378207",
  "updated_at": "2025-05-27 23:32:23.378207",
  "email": "walt@breakingbad.com",
  "hashed_password": "$2a$10$Xxe/PF6Nph/5DO9pO637v.jDQmPoQaFzGxXa1tU.mrJcFuO2fkBvq",
  "is_chirpy_red": false
}
```

```go
type User struct {
	ID          uuid.UUID `json:"id"`
	CreateAt    time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Email       string    `json:"email"`
	IsChirpyRed bool      `json:"is_chirpy_red"` // check user is upgraded for red chirpy!
}
```

- `POST` `/api/users`: create user

  - request body
    ```json
    {
      "email": "walt@example.com",
      "password": "password123"
    }
    ```
  - response: `User Data`
  - status code: `201` `Created`

- `PUT` `/api/users`: update user

  - request body
    ```json
    {
      "email": "walt@example.com",
      "password": "password123"
    }
    ```
  - response: `User Data`
  - status code: `200` `OK`

- `POST` `/api/login`: login user

  - request body
    ```json
    {
      "email": "walt@example.com",
      "password": "password123"
    }
    ```
  - response: `User Data` + `Token` + `RefreshToken`
  - status code: `200` `OK`

  ```json
  {
    "id": "16577ee6-5cbf-4ff9-a720-ec2c4c7b644c",
    "created_at": "2025-05-27 23:32:23.378207",
    "updated_at": "2025-05-27 23:32:23.378207",
    "email": "walt@breakingbad.com",
    "hashed_password": "$2a$10$Xxe/PF6Nph/5DO9pO637v.jDQmPoQaFzGxXa1tU.mrJcFuO2fkBvq",
    "is_chirpy_red": false,
    "token": "{jwt_token}",
    "refresh_token": "{random_32_byte_string}"
  }
  ```

### Chirp

`Chirp Data`

```json
{
  "id": "16577ee6-5cbf-4ff9-a720-ec2c4c7b644c",
  "created_at": "2025-05-27 23:32:23.378207",
  "updated_at": "2025-05-27 23:32:23.378207",
  "body": "{the_content_of_chirp}",
  "user_id": "{author's uuid}"
}
```

```go
type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}
```

- `POST` `/api/chirps`: create chirp

  - request body
    ```json
    {
      "body": "password123"
    }
    ```
  - response: `Chirp`
  - status code: `201` `Created`

- `GET` `/api/chirps`: get chirp list

  - Query Parameters
    - `author_id` (uuid): get list of chirps that only user_id = author_id
    - `sort` ("asc"|"desc"): sort list order (default: ascending)
  - response: `List` of `Chirp`
  - status code: `200` `OK`

- `GET` `/api/chirps/{chirp_id}`: get chirp by chirp_id

  - response: `Chirp`
  - status code: `200` `OK`

- `DELETE` `api/chirps/{chirp_id}`: delete chirp by chirp_id
  - response: no data
  - status code: `204` `No Content`

### Token

- `POST` `/api/refresh`: refresh access token with `refresh_token`
  - request:
    - add `Authorization` header, value should be `"Bearer {refresh_token}"`
  - response:
    ```json
    {
      "token": {jwt_token}
    }
    ```
  - status code: `200` `OK`
- `POST` `/api/revoke`:
  - request: same as `/api/refresh` end point
  - response: no data
  - status code: `204` `No Content`

### Admin features

- `POST /admin/reset`: reset `users` DB (automatically reset `chirps`,
  `refresh_tokens`) and `hit count` of server
  - `hit count`: automatically incremented when client hit the server with
    requests.
  - `PLATFORM` value in `.env` should be "dev"
  - response:
  ```
  "Hits reset to 0, users successfully deleted."
  ```
  - status code: `200` `OK`

### Webhooks

Used `Polka` as a example 3rd party service for webhooks.

- `POST /api/polka/webhooks`: handle polka's request based on the events.
  - request body
  ```json
  {
    "event": "user.upgraded",
    "data": {
      "user_id": "{user's uuid, gonna be upgraded}"
    }
  }
  ```
  - add `Authorization` header, value should be `"ApiKey {polka_api_key}"`
