# 📝 Todo API with JWT Authentication (Go)

A secure RESTful **Todo API** built using **Golang**, **JWT authentication**, and **PostgreSQL**.  
This project follows a clean, layered architecture with explicit database migrations and startup commands documented in `cmd.txt`.

---

## 🚀 Features

- User Registration & Login
- JWT-based Authentication & Authorization
- CRUD Operations for Todo Items
- Repository-based data access layer
- PostgreSQL integration
- SQL migration support (up/down)
- Environment-based configuration
- Hot reload support using Air

---

## 🏗️ Actual Project Structure

```text
Todo_api_with_JWT_go_lang-main/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── config/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   └── repository/
├── migrations/
│   ├── 000001_create_todos_table.up.sql
│   ├── 000001_create_todos_table.down.sql
│   ├── 000002_create_user_table.up.sql
│   ├── 000002_create_user_table.down.sql
│   ├── 000003_add_user_id_to_todos_table.up.sql
│   └── 000003_add_user_id_to_todos_table.down.sql
├── scripts/
│   └── migrate.ps1
├── cmd.txt
├── .env
├── .air.toml
├── go.mod
├── go.sum
└── README.md
```

---

## 🔧 Prerequisites

- Go 1.20+
- PostgreSQL
- golang-migrate CLI

---

## ⚙️ Environment Variables

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todo_db

JWT_SECRET=your_secret_key
JWT_EXPIRE_HOURS=24
```

---

## 🗄️ Database Migration

### Run Migrations (Up)

```bash
migrate -path migrations -database "postgres://user:password@localhost:5432/todo_db?sslmode=disable" up
```

### Rollback Migrations (Down)

```bash
migrate -path migrations -database "postgres://user:password@localhost:5432/todo_db?sslmode=disable" down
```

---

## ▶️ Running the Application

```bash
go mod tidy
go run cmd/api/main.go
```

Application will start on:

```
http://localhost:8080
```

---

## 🔐 API Endpoints

### 🧑 Authentication APIs

#### Register User
```
POST /register
```

**Request Body**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response**
```json
{
  "message": "user registered successfully"
}
```

---

#### Login User
```
POST /login
```

**Request Body**
```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response**
```json
{
  "token": "jwt_token_here"
}
```

---

### ✅ Todo APIs (JWT Protected)

All Todo APIs require the following header:

```
Authorization: Bearer <JWT_TOKEN>
```

---

#### Get All Todos
```
GET /todos
```

**Response**
```json
[
  {
    "id": 1,
    "title": "Learn Go",
    "completed": false
  }
]
```

---

#### Create Todo
```
POST /todos
```

**Request Body**
```json
{
  "title": "Build Todo API"
}
```

---

#### Update Todo
```
PUT /todos/{id}
```

**Request Body**
```json
{
  "title": "Build Secure Todo API",
  "completed": true
}
```

---

#### Delete Todo
```
DELETE /todos/{id}
```

**Response**
```json
{
  "message": "todo deleted successfully"
}
```

---

## 🛡️ Security Notes

- JWT validation implemented using middleware
- Secrets loaded from environment variables
- Password hashing should be done using bcrypt

---


## 👨‍💻 Author

Dattatray Narhe
