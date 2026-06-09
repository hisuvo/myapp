যদি আপনার Structure এমন হয়:

```text
internal/
└── user/
    ├── dto/
    │   ├── request.go
    │   └── response.go
    ├── entity.go
    ├── repository.go
    ├── service.go
    └── handler.go
```

তাহলে `DTO` (Data Transfer Object) ব্যবহার করা হয় API Layer এবং Business Layer-এর মধ্যে ডাটা আদান-প্রদানের জন্য।

---

# request.go

ক্লায়েন্ট থেকে যে ডাটা আসে তা এখানে রাখা হয়।

```go
package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
```

### Client Request

```json
{
  "name": "Suvo",
  "email": "suvo@gmail.com",
  "password": "123456"
}
```

Handler এই DTO-তে Bind করবে।

```go
var req dto.CreateUserRequest

if err := c.Bind(&req); err != nil {
    return err
}
```

---

# response.go

ক্লায়েন্টকে যে ডাটা পাঠাবেন তা এখানে রাখা হয়।

```go
package dto

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
```

Response:

```json
{
  "id": 1,
  "name": "Suvo",
  "email": "suvo@gmail.com"
}
```

---

# entity.go

Database বা Business Entity।

```go
package user

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}
```

---

# DTO Flow

```text
Client
   │
   ▼
Request DTO
   │
   ▼
Handler
   │
   ▼
Service
   │
   ▼
Entity
   │
   ▼
Repository
   │
 Database
```

Response এর সময়:

```text
Database
   │
   ▼
Repository
   │
   ▼
Entity
   │
   ▼
Service
   │
   ▼
Response DTO
   │
   ▼
Client
```

---

# কেন DTO ব্যবহার করব?

## 1. Password Hide করা

Entity:

```go
type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}
```

যদি Entity সরাসরি Return করেন:

```json
{
  "id": 1,
  "name": "Suvo",
  "email": "suvo@gmail.com",
  "password": "123456"
}
```

এটা Security Risk।

Response DTO:

```go
type UserResponse struct {
	ID    uint
	Name  string
	Email string
}
```

Password যাবে না।

---

## 2. Validation সহজ হয়

```go
type CreateUserRequest struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}
```

Request Validation DTO-তে করা যায়।

---

## 3. Entity Clean থাকে

Entity শুধুমাত্র Business Model।

```go
type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}
```

JSON Tag, Validation Tag, API Logic Entity-তে ঢুকাতে হয় না।

---

## 4. API Change করলে Database Change লাগে না

আজ:

```json
{
  "name": "Suvo"
}
```

কাল:

```json
{
  "full_name": "Suvo Datta"
}
```

শুধু DTO পরিবর্তন করলেই হবে, Entity একই থাকতে পারে।

---

# Production-Level Flow

```text
handler
   ↓
request DTO
   ↓
service
   ↓
entity
   ↓
repository
   ↓
database

database
   ↓
repository
   ↓
entity
   ↓
response DTO
   ↓
handler
   ↓
client
```

### Rule

- `entity.go` → Business/Data Model
- `dto/request.go` → API Input
- `dto/response.go` → API Output
- `service.go` → Business Logic
- `repository.go` → Database Logic
- `handler.go` → HTTP Request/Response

এটি Large-Scale Go Backend Project-এ খুবই জনপ্রিয় একটি Clean Architecture Pattern।
