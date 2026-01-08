ด้านล่างนี้คือ **README.md เวอร์ชันสมบูรณ์ “ปัจจุบันจริง”**
ผมจัดโครงสร้าง + เนื้อหาใหม่ให้ **เพื่อนร่วมทีม clone → run → dev ต่อได้ทันที**
คุณสามารถ **copy-paste ทั้งไฟล์** ไปแทน README.md เดิมได้เลย

---

```md
# User Service (OSCIRA MCS)

User Service คือบริการจัดการผู้ใช้งาน (Authentication / Authorization / Profile / User Option)  
พัฒนาด้วย **Go (Fiber v2)**, **PostgreSQL**, และ **golang-migrate**  
ออกแบบตามแนวคิด **Clean Architecture + Session-based Auth + RBAC**

เหมาะสำหรับ:
- ทำงานเป็นทีม
- ใช้พัฒนา (Dev) ต่อ
- ส่งงาน / สหกิจ / Demo อาจารย์

---

## 🔧 Technology Stack

- **Go 1.22+**
- **Fiber v2**
- **PostgreSQL**
- **pgx / pgxpool**
- **Session-based Authentication (Cookie)**
- **golang-migrate**
- **Docker & Docker Compose**
- RESTful API

---

## 📦 Project Structure

```

.
├── cmd
│   └── user_service
│       └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── constants
│   │   └── context.go
│   ├── handler
│   │   └── user_handler.go
│   ├── middleware
│   │   ├── auth_middleware.go
│   │   └── error_handler.go
│   ├── model
│   │   ├── user.go
│   │   ├── roles.go
│   │   └── profile.go
│   ├── repository
│   │   ├── user_repo.go
│   │   ├── session_repo.go
│   │   └── assign_role.go
│   ├── router
│   │   └── router.go
│   └── service
│       ├── auth_service.go
│       ├── user_service.go
│       └── profile_service.go
├── migrations
│   ├── 0001_create_users_table.up.sql
│   ├── 0002_create_roles_table.up.sql
│   ├── 0003_create_user_roles.up.sql
│   ├── 0007_create_session_table.up.sql
│   ├── 0016_create_user_profile.up.sql
│   └── ...
├── pkg
│   └── db
│       └── postgres.go
├── scripts
│   ├── migrate.sh
│   ├── migrate.bat
│   └── Makefile
├── docker-compose.yml
├── env
│   └── dev.env
├── go.mod
└── README.md

```

---

## 🧠 Architecture Overview

โครงสร้างใช้แนวคิด **Clean Architecture**

```

Handler  →  Service  →  Repository  →  Database

````

หน้าที่แต่ละ layer:

- **Handler**  
  รับ HTTP request / response
- **Service**  
  Business logic
- **Repository**  
  ติดต่อ Database เท่านั้น
- **Middleware**  
  Authentication / Authorization
- **Constants**  
  Key กลาง เช่น ContextUserID

---

## 🚀 Quick Start (สำหรับเพื่อนร่วมทีม)

### Requirements

- Docker
- Docker Compose
- Git

---

### 1️⃣ Clone Project

```bash
git clone https://github.com/Natsu2022/GO_Userservice.git
cd GO_Userservice
````

---

### 2️⃣ ตั้งค่า Environment

```bash
cp env/dev.env .env
```

ตัวอย่าง `.env`

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=oscira_user

SERVER_PORT=3455
```

---

### 3️⃣ Run ด้วย Docker Compose

```bash
docker compose up --build
```

Service จะพร้อมใช้งานที่:

```
http://localhost:3455
```

---

## 🐘 Database & Migration

### ใช้ golang-migrate

**Up migration**

```bash
make migrate-up
```

**Down migration**

```bash
make migrate-down
```

> ✔ สามารถพัฒนาแบบ **ไม่ต้องใช้ Docker ก็ได้**
> แค่มี PostgreSQL และตั้งค่า `.env` ให้ถูกต้อง

---

## 🔐 Authentication (สำคัญ)

ระบบใช้ **Session-based Authentication**

### Flow

1. User login
2. Server สร้าง `session_id`
3. เก็บ session ลง DB
4. ส่ง `session_id` กลับใน **HttpOnly Cookie**
5. ทุก request ที่ protected จะใช้ cookie นี้อัตโนมัติ

---

## ⚠️ สำคัญมาก: Context Key

ไฟล์ `internal/constants/context.go`

```go
package constants

const ContextUserID = "userID"
```

**ต้องใช้ key นี้ตรงกันทุกที่**

Middleware:

```go
c.Locals(constants.ContextUserID, userID)
```

Handler:

```go
uid := c.Locals(constants.ContextUserID).(uuid.UUID)
```

❌ ถ้า key ไม่ตรง → `interface {} is nil`

---

## 🧩 API Routes

Base URL:

```
http://localhost:3455/api/v1
```

---

### 🔓 Public

#### Register

```
POST /users/register
```

```json
{
  "email": "tester@example.com",
  "password": "password123",
  "first_name": "Tester",
  "last_name": "User"
}
```

---

#### Login

```
POST /auth/login
```

```json
{
  "email": "tester@example.com",
  "password": "password123"
}
```

Response จะตั้ง Cookie:

```
session_id=<uuid>; HttpOnly
```

---

### 🔒 Protected (ต้อง Login)

#### Check Session

```
GET /me
```

```json
{
  "message": "You are authenticated",
  "user_id": "uuid"
}
```

---

#### Get My Profile

```
GET /me/profile
```

---

#### Update My Profile

```
PATCH /me/profile
```

```json
{
  "display_name": "New Name",
  "phone_number": "0999999999"
}
```

---

## 🧪 ตัวอย่าง curl

```bash
curl -i -X POST http://localhost:3455/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"tester@example.com","password":"password123"}'
```

---

## 🛠 Common Errors & Fix

### ❌ `interface conversion: interface {} is nil, not uuid.UUID`

**สาเหตุ**

* Context key ไม่ตรงกัน

**วิธีแก้**

* ใช้ `constants.ContextUserID = "userID"` ให้ตรงทุกไฟล์

---

### ❌ `userRepository does not implement UserRepository`

**สาเหตุ**

* struct ขาด method ที่ interface กำหนด

**วิธีแก้**

* ตรวจ method signature ให้ตรง 100%
* ชื่อ + parameter + return type ต้องตรง

---

## 👥 สำหรับเพื่อนร่วมทีม

สิ่งที่เพื่อนต้องรู้:

* Login ก่อนเสมอ
* Cookie ถูกจัดการให้อัตโนมัติ
* ใช้ `/me` เช็ค session
* อย่าเปลี่ยน Context key เอง

---

## 👤 Author

**PHUMIN TONGLAR**
Backend Development / Cooperative Education Project

---

## 📄 License

Educational & Demonstration Purpose Only

```

---

ถ้าคุณต้องการ:
- แยก **API Doc เป็นไฟล์ Postman**
- เพิ่ม **User Option / Preference API**
- เพิ่ม **RBAC Middleware**
- หรือทำ **Diagram Architecture**

บอกได้เลย ผมต่อให้เป็น production-grade ได้ทันที
```
