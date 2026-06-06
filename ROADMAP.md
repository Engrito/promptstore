# 🚀 Project Roadmap

This document outlines the full backend development plan for the Go-based API system, structured in phased milestones and Git branches.

---

# 🧱 Phase 1 — Foundation + Auth

## 1. Project Bootstrap
**Branch:** `chore/init-project`  
**Ticket:** Initialize Go project structure  

**Description:**
- Create `go.mod`
- Setup Gin server in `cmd/api/main.go`
- Add `/health` route
- Define base folder structure

---

## 2. Database Setup (SQLite + GORM)
**Branch:** `chore/setup-database`  
**Ticket:** Initialize database layer  

**Description:**
- Add GORM setup in `pkg/database`
- Connect SQLite (`data/app.db`)
- Enable auto-migration hook

---

## 3. User Model + Migration
**Branch:** `feature/user-model`  
**Ticket:** Create User schema  

**Description:**
- Define `User` model
- Fields:
  - id
  - username
  - password_hash
  - timestamps
- Add GORM migration

---

## 4. User Repository Layer
**Branch:** `feature/user-repository`  
**Ticket:** Add DB access for users  

**Description:**
- Implement CRUD methods:
  - CreateUser
  - GetUserByUsername
  - GetUserByID

---

## 5. Password Hashing Utility
**Branch:** `feature/password-hash`  
**Ticket:** Secure password storage  

**Description:**
- Add bcrypt utilities in `pkg/shared`
- Functions:
  - HashPassword
  - CheckPassword

---

## 6. JWT Service
**Branch:** `feature/jwt-service`  
**Ticket:** Add authentication token system  

**Description:**
- Create JWT generator/validator in `pkg/jwt`
- Define claims:
  - user_id
  - username
  - exp

---

## 7. Auth Service (Business Logic)
**Branch:** `feature/auth-service`  
**Ticket:** Authentication business logic  

**Description:**
- Register user (hash password)
- Login user (verify + return JWT)

---

## 8. Auth Handler (HTTP Layer)
**Branch:** `feature/auth-handler`  
**Ticket:** Expose auth endpoints  

**Description:**
- `POST /auth/register`
- `POST /auth/login`
- Return JWT on login

---

## 9. Auth Middleware
**Branch:** `feature/auth-middleware`  
**Ticket:** Protect routes using JWT  

**Description:**
- Validate token
- Inject user into context
- Reject unauthorized requests

---

## 10. User Profile Endpoints
**Branch:** `feature/user-profile`  
**Ticket:** User info routes  

**Description:**
- `GET /users/:username`
- `GET /me` (protected)

---

# 🧠 Phase 2 — Prompt Core System

## 11. Prompt Model
**Branch:** `feature/prompt-model`  
**Ticket:** Create prompt schema  

**Description:**
- Fields:
  - content
  - category
  - tags
  - user_id
  - timestamps

---

## 12. Prompt Repository
**Branch:** `feature/prompt-repository`  
**Ticket:** Database layer for prompts  

**Description:**
- CreatePrompt
- GetPromptByID
- GetAllPrompts
- UpdatePrompt
- DeletePrompt

---

## 13. Prompt Service Layer
**Branch:** `feature/prompt-service`  
**Ticket:** Business logic for prompts  

**Description:**
- Ownership validation
- Input validation
- Tag handling

---

## 14. Create Prompt Endpoint
**Branch:** `feature/prompt-create`  
**Ticket:** Add prompt creation API  

**Description:**
- `POST /prompts`
- Auth required
- Auto-assign `user_id`

---

## 15. Get Prompt Endpoints
**Branch:** `feature/prompt-read`  
**Ticket:** Public prompt access  

**Description:**
- `GET /prompts`
- `GET /prompts/:id`

---

## 16. Update Prompt Endpoint
**Branch:** `feature/prompt-update`  
**Ticket:** Modify prompt  

**Description:**
- `PUT /prompts/:id`
- Only owner can update

---

## 17. Delete Prompt Endpoint
**Branch:** `feature/prompt-delete`  
**Ticket:** Remove prompt  

**Description:**
- `DELETE /prompts/:id`
- Ownership check required

---

## 18. Prompt Filtering + Search
**Branch:** `feature/prompt-search`  
**Ticket:** Add query support  

**Description:**
Filter prompts by:
- category
- tags
- keyword search (content)

---

## 19. Pagination System
**Branch:** `feature/pagination`  
**Ticket:** Add scalable listing  

**Description:**
- limit / offset support
- page-based metadata response

---

# 🧩 Phase 3 — User Enhancements

## 20. User Service Improvements
**Branch:** `feature/user-service`  
**Ticket:** Extend user logic  

**Description:**
- Improve profile handling
- Add future stats hooks (likes, prompt count)

---

## 21. Ownership Middleware Helper
**Branch:** `feature/ownership-helper`  
**Ticket:** Centralize permission checks  

**Description:**
- Reusable ownership validation logic for prompts

---

## 22. Basic Validation Layer
**Branch:** `feature/request-validation`  
**Ticket:** Input validation system  

**Description:**
- Validate:
  - login
  - register
  - prompt payloads

---

# ⭐ Phase 4 — Engagement Features

## 23. Likes Model
**Branch:** `feature/likes-model`  
**Ticket:** Add prompt likes schema  

**Description:**
- prompt_likes table
- user_id + prompt_id relation

---

## 24. Like/Unlike Endpoints
**Branch:** `feature/likes-api`  
**Ticket:** Toggle like system  

**Description:**
- `POST /prompts/:id/like`
- `DELETE /prompts/:id/like`

---

## 25. Bookmarks Model
**Branch:** `feature/bookmarks-model`  
**Ticket:** Save prompts  

**Description:**
- Bookmark schema

---

## 26. Bookmark Endpoints
**Branch:** `feature/bookmarks-api`  
**Ticket:** Bookmark system  

**Description:**
- Save/remove bookmarks

---

## 27. Prompt Stats
**Branch:** `feature/prompt-stats`  
**Ticket:** Engagement metrics  

**Description:**
- likes count
- bookmark count

---

# ⚙️ Phase 5 — Production Ready

## 28. Logging System
**Branch:** `chore/logger`  
**Ticket:** Structured logging  

**Description:**
- Add logger package
- Request logging middleware

---

## 29. Error Handling Standardization
**Branch:** `chore/error-handling`  
**Ticket:** Unified API errors  

**Description:**
- Consistent JSON error format

---

## 30. Swagger/OpenAPI Setup
**Branch:** `docs/swagger-setup`  
**Ticket:** API documentation  

**Description:**
- Generate Swagger docs
- Document all endpoints

---

## 31. Docker Support
**Branch:** `chore/docker-setup`  
**Ticket:** Containerize API  

**Description:**
- Dockerfile
- docker-compose (future DB-ready)

---

## 32. PostgreSQL Migration Prep
**Branch:** `chore/postgres-ready`  
**Ticket:** Prepare DB abstraction  

**Description:**
- Decouple SQLite-specific logic
- Enable config-based DB switching

---

## 33. Redis Caching Layer (Optional Prep)
**Branch:** `feature/cache-layer`  
**Ticket:** Add caching abstraction  

**Description:**
- Create cache interface
- Prepare for Redis integration

---

# 📌 Suggested Execution Order

Follow strictly:

### Phase 1
1 → 10 (Auth foundation)

### Phase 2
11 → 19 (Prompt system)

### Phase 3
20 → 22 (User improvements)

### Phase 4
23 → 27 (Engagement features)

### Phase 5
28 → 33 (Production readiness)

---