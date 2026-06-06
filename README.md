# PromptStore API

A REST API for sharing, discovering, and managing AI prompts.

Users can create prompts, organize them with tags and categories, and share them publicly. Visitors can browse prompts without authentication, while authenticated users can create and manage their own prompts.

---

## Features

### Public Features

* View all public prompts
* View prompt details
* Search prompts
* Filter by category
* Filter by tags
* View prompt author

### Authenticated Features

* Register account
* Login
* Create prompt
* Update own prompt
* Delete own prompt
* View own profile

### Future Features

* Like prompts
* Bookmark prompts
* Comments
* Prompt collections
* User profiles
* Prompt version history
* Trending prompts
* API rate limiting
* Admin moderation

---

## Tech Stack

### Backend

* Go
* Gin
* SQLite (initially)
* GORM

### Authentication

* JWT

### Future Infrastructure

* PostgreSQL
* Redis
* Docker
* Swagger/OpenAPI

---

## Project Structure

```text
promptstore-api/

├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│
│   ├── auth/
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── middleware.go
│   │
│   ├── prompt/
│   │   ├── model.go
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── routes.go
│   │
│   ├── user/
│   │   ├── model.go
│   │   ├── handler.go
│   │   ├── service.go
│   │   ├── repository.go
│   │   └── routes.go
│   │
│   ├── middleware/
│   └── shared/
│
├── pkg/
│   ├── database/
│   ├── logger/
│   └── jwt/
│
├── data/
│   └── app.db
│
├── configs/
├── migrations/
├── tests/
│
├── go.mod
└── README.md
```

---

## Data Model

### User

```json
{
  "id": 1,
  "username": "john",
  "password_hash": "...",
  "created_at": "2026-01-01T00:00:00Z"
}
```

### Prompt

```json
{
  "id": 1,
  "content": "Act as a senior Go developer...",
  "category": "Programming",
  "tags": [
    "golang",
    "backend",
    "api"
  ],
  "user_id": 1,
  "created_at": "2026-01-01T00:00:00Z",
  "updated_at": "2026-01-01T00:00:00Z"
}
```

---

## Future Data Model

### Prompt Like

```json
{
  "id": 1,
  "prompt_id": 10,
  "user_id": 5
}
```

### Bookmark

```json
{
  "id": 1,
  "prompt_id": 10,
  "user_id": 5
}
```

---

## API Endpoints

### Authentication

```http
POST /api/v1/auth/register
POST /api/v1/auth/login
```

### Public Prompt Routes

```http
GET /api/v1/prompts
GET /api/v1/prompts/:id
```

### Protected Prompt Routes

```http
POST /api/v1/prompts
PUT /api/v1/prompts/:id
DELETE /api/v1/prompts/:id
```

### User Routes

```http
GET /api/v1/users/:username
GET /api/v1/me
```

---

## Authorization Rules

### Guest

Can:

* View prompts
* Search prompts

Cannot:

* Create prompts
* Edit prompts
* Delete prompts

### Authenticated User

Can:

* Create prompts
* Update own prompts
* Delete own prompts
* View own profile

Cannot:

* Modify prompts owned by others

---

## Scalability Notes

The project starts with SQLite for simplicity.

The architecture separates:

* Handlers (HTTP layer)
* Services (business logic)
* Repositories (data access)

This allows migration to PostgreSQL and Redis later with minimal changes to business logic.

Future high-scale additions:

* Pagination
* Redis caching
* Full-text search
* Background jobs
* CDN for assets
* PostgreSQL
* Horizontal API scaling

---

## Development Roadmap

### Phase 1

* User registration
* Login
* JWT authentication
* Create prompt
* List prompts
* View prompt

### Phase 2

* Update prompt
* Delete prompt
* Ownership checks
* Pagination

### Phase 3

* Likes
* User profiles
* Prompt statistics

### Phase 4

* Bookmarks
* Comments
* Collections

### Phase 5

* PostgreSQL migration
* Redis caching
* Docker deployment
* Swagger documentation

```
```
