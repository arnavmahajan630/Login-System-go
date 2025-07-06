
# 🔐 Go Authentication API

A secure, RESTful authentication backend built with:

- 🧬 **Go (Golang)**
- ⚡ **Gin** web framework
- 🗃️ **MongoDB**
- 🔑 **JWT-based authentication** with refresh token support

---

## 📁 Project Structure

```
login-portal-go/
├── config/           # MongoDB connection
├── controllers/      # Signup, login, user handlers
├── middleware/       # JWT authentication middleware
├── models/           # User model
├── routes/           # API routes
├── services/         # Password hashing, token generation, etc.
├── .env              # Environment variables
├── main.go           # Entry point
└── go.mod
```

---

## ⚙️ Features

THIS PROJECT IS A PART OF GO LEARNING SERIES https://github.com/arnavmahajan630/Learn-Go

✅ Signup with validation  
✅ Login with JWT + refresh token  
✅ Hashed password storage (bcrypt)  
✅ Access control by role (`admin` check)  
✅ Protected routes (`/users`, `/users/:id`)  
✅ Refreshable access tokens  
✅ Clean modular structure  
✅ MongoDB integration

---

## 🚀 Getting Started

### 1. 📦 Install Dependencies

```bash
go mod tidy
```

### 2. 🌐 Create `.env` File

```env
PORT=3000
DATABASE_URL=mongodb://localhost:27017
JWT_SECRET=super-secret-key-01
JWT_EXPIRY=1h
JWT_REFRESH_EXPIRY=7h
GIN_MODE=debug
```

### 3. 🧪 Run the Server

#### Option 1: Regular run

```bash
go run main.go
```

#### Option 2: With live reload using [air](https://github.com/cosmtrek/air)

```bash
air
```

---

## 📌 API Endpoints

| Method | Endpoint         | Description                | Auth |
|--------|------------------|----------------------------|------|
| `POST` | `/signup`        | Register a new user        | ❌   |
| `POST` | `/login`         | Login & get JWT tokens     | ❌   |
| `GET`  | `/`              | Hello world (public test)  | ❌   |
| `GET`  | `/users`         | Get all users (admin only) | ✅   |
| `GET`  | `/users/:id`     | Get a single user (admin+owner only) | ✅   |

---

## 🔐 JWT Authentication

- Access tokens (`access_token`) and refresh tokens (`refresh_token`) are issued on login.
- Protected routes require a valid `Authorization: Bearer <access_token>` header.
- Middleware is used to extract and verify JWTs.

---

## 🧰 Technologies Used

| Tech       | Purpose               |
|------------|------------------------|
| Gin        | HTTP router / middleware |
| MongoDB    | User data persistence |
| JWT        | Token-based auth      |
| bcrypt     | Password hashing      |
| godotenv   | Environment variables |
| validator  | Struct validation     |

---

## 🧪 Sample Request

### Signup

```json
POST /signup
Content-Type: application/json

{
  "email": "arnav@example.com",
  "phone": "1234567890",
  "password": "password123",
  "role": "admin"
}
```

### Login

```json
POST /login
Content-Type: application/json

{
  "email": "arnav@example.com",
  "password": "password123"
}
```

Response:
```json
{
  "access_token": "jwt-here",
  "refresh_token": "jwt-here",
  "user": {
    "_id": "...",
    "email": "..."
  }
}
```

---

## ✅ To Do (Ideas for V2)

- Refresh token endpoint (`/refresh`)
- Email verification on signup
- Password reset logic
- Rate limiting + brute force protection
- Admin dashboard routes

---

## 🙌 Credits

Built with 💙 by [Arnav Mahajan](https://github.com/arnavmahajan630)  
Made for learning + production-grade architecture in mind.
