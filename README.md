User Service เป็น Backend Service ที่พัฒนาด้วยภาษา **Go (Fiber Framework)**  
รองรับระบบ **Authentication แบบ Session + Cookie**  
และออกแบบมาให้สามารถรันได้ทันทีด้วย **Docker Compose**  
โดยผู้ใช้งานไม่จำเป็นต้องติดตั้ง Go หรือ PostgreSQL เอง

---

## 🔧 Technology Stack

- **Go** (Fiber v2)
- **PostgreSQL**
- **Session-based Authentication**
- **Docker & Docker Compose**
- RESTful API

---

## 📦 Project Structure

```

.
├── LICENSE
├── README.md
├── cmd
│   └── user_service
│       └── main.go
├── deployments
│   ├── docker
│   │   └── Dockerfile
│   └── k8s
│       ├── Upload_depl.yaml
│       ├── Upload_impl.yaml
│       └── Upload_svc.yaml
├── docker-compose.yml
├── env
│   └── dev.env
├── go.mod
├── go.sum
├── internal
│   ├── config
│   │   └── config.go
│   ├── handler
│   │   └── user_handler.go
│   ├── middleware
│   │   ├── auth_middleware.go
│   │   └── error_handler.go
│   ├── model
│   │   ├── roles.go
│   │   └── user.go
│   ├── repository
│   │   ├── assign_role.go
│   │   ├── session_repo.go
│   │   └── user_repo.go
│   ├── router
│   │   └── router.go
│   └── service
│       ├── login.go
│       └── register.go
├── migrations
│   ├── 0000_drop_all_table.drop.sql
│   ├── 0001_create_users_table.up.sql
│   ├── 0002_create_roles_table.up.sql
│   ├── 0003_create_user_role.up.sql
│   ├── 0004_create_permissions_table.up.sql
│   ├── 0005_create_role_permissions_table.up.sql
│   ├── 0006_create_user_permissions_table.up.sql
│   ├── 0007_create_session_table.up.sql
│   ├── 0008_create_preference_table.up.sql
│   ├── 0009_create_location_table.up.sql
│   ├── 0010_create_source_list_table.up.sql
│   ├── 0011_seed_superadmin.up.sql
│   ├── 0012_seed_role.up.sql
│   ├── 0013_seed_permission.up.sql
│   ├── 0014_seed_role_permission.up.sql
│   ├── 0015_assign_superadmin_role.up.sql
│   └── 0016_create_user_profile.up.sql
├── pkg
│   └── db
│       └── postgres.go
├── scripts
│   ├── makefile
│   ├── migrate.bat
│   └── migrate.sh
├── services
│   └── user_service
│       └── migrations

````

โครงสร้างนี้ออกแบบมาเพื่อ:
- แยก service อย่างชัดเจน
- รองรับการขยายเป็นหลาย service ในอนาคต
- ใช้งานได้เหมือนกันทุกเครื่อง (Windows / macOS / Linux)

---

## 🚀 Quick Start (One Command)

### Requirements
- Docker
- Docker Compose

### Run Project

```bash
git clone https://github.com/Natsu2022/GO_Userservice.git
cd GO_Userservice/scipts
make dev
````

หลังจากรันเสร็จ ระบบจะพร้อมใช้งานที่:

```
http://localhost:3455
```

---

## ⚙️ Environment Configuration

โปรเจกต์ใช้ environment variables สำหรับการตั้งค่า

### `.env.example`

```env
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=myDB

SERVER_PORT=3455
```

> **หมายเหตุ:**
> ในการใช้งานจริง เพียงคัดลอกไฟล์นี้เป็น `.env`
> หรือใช้ค่า default ที่กำหนดไว้ใน `docker-compose.yml`

---

## 🐘 Database

* PostgreSQL จะถูกสร้างอัตโนมัติผ่าน Docker
* ใช้ Docker Volume สำหรับเก็บข้อมูล
* ไม่ต้องติดตั้ง PostgreSQL บนเครื่อง

Service ภายในจะเชื่อมต่อฐานข้อมูลผ่าน hostname:

```
postgres
```

---

## 🔐 Authentication & Session

ระบบใช้ **Session-based Authentication** โดยมีหลักการดังนี้:

* เมื่อผู้ใช้ Login สำเร็จ

  * ระบบจะสร้าง `session_id`
  * เก็บ session ใน Database
  * ส่ง `session_id` กลับไปที่ Client ผ่าน Cookie

* Endpoint `/me`

  * ใช้ตรวจสอบว่า session ใน Cookie ยังใช้งานได้หรือไม่
  * ถ้า session ถูกต้อง → คืนข้อมูลผู้ใช้
  * ถ้าไม่ถูกต้อง → ตอบ `401 Unauthorized`

---

## 📡 API Overview

### Health Check

```
GET /
```

### Login

```
POST /api/v1/users/auth/login
```

### Get Current User (Session Check)

```
GET /api/v1/users/auth/me
```

---

## 🐳 Docker Compose Overview

`docker-compose.yml` จะประกอบด้วย:

* `postgres` : Database Service
* `user_service` : Go Backend Service

Docker Compose จะจัดการ:

* Network ระหว่าง service
* Environment variables
* Startup order

ผู้ใช้งานไม่ต้องรัน container ทีละตัว

---

## ✅ Why Docker Compose?

* ลดปัญหา "เครื่องผมรันได้ แต่เครื่องคุณรันไม่ได้"
* ไม่ผูกกับ OS หรือ environment ใด ๆ
* ใช้งานได้ด้วยคำสั่งเดียว
* เหมาะสำหรับการส่งงาน / demo / ใช้งานร่วมกับทีม

---

## 🧪 Development Mode

หากต้องการแก้ไขโค้ดและ rebuild:

```bash
docker compose up --build
```

หรือหยุดระบบ:

```bash
docker compose down
```

---

## 📌 Notes

* ห้าม commit ไฟล์ `.env` ที่มีข้อมูลจริง
* Repository นี้มี `.env.example` สำหรับอ้างอิง
* Session และ Database จะถูก reset หากลบ Docker volume

---

## 👤 Author

Developed by
**PHUMIN TONGLAR**
(Cooperative Education / Backend Development Project)

---

## 📄 License

This project is for educational and demonstration purposes.

ด้านล่างนี้คือ **API Documentation ที่ครบ ใช้ได้จริง และเหมาะทั้งส่งอาจารย์ + ให้เพื่อนใช้**
ผมเขียนให้ **ต่อจาก README เดิมได้ทันที** (คุณสามารถคัดลอกไปแปะต่อท้ายหัวข้อใหม่)

> โฟกัส: ชัดเจน, ไม่กำกวม, ใช้ทดสอบได้จริง (Postman / curl / frontend)

---

## 📘 API Documentation

Base URL (Development):

```
http://localhost:3455
```

API ทั้งหมดอยู่ภายใต้ namespace:

```
/api/v1/users
```

---

## 🔍 Authentication Overview

ระบบใช้ **Session-based Authentication (Cookie)**

Flow:

1. Client ส่ง username/password ไปที่ `/auth/login`
2. Server ตรวจสอบข้อมูล
3. ถ้าสำเร็จ:

   * สร้าง `session_id`
   * บันทึก session ลง Database
   * ส่ง `session_id` กลับไปใน Cookie
4. Client เรียก API ที่ต้องการ auth (เช่น `/auth/me`)

   * Cookie จะถูกส่งไปอัตโนมัติ
   * Server ตรวจสอบ session จาก Database

> ❗ Client **ไม่ต้อง** แนบ token เอง
> Cookie จะถูกจัดการโดย Browser / HTTP Client

---

## 🏥 Health Check

### GET /

ใช้ตรวจสอบว่า service ทำงานอยู่หรือไม่

**Request**

```
GET /
```

**Response**

```json
{
  "status": "ok",
  "service": "user_service"
}
```

**Status Codes**

* `200 OK` – Service ทำงานปกติ

---

---

## Register

### POST http://127.0.0.1:3455/api/v1/users/register

ใช้สำหรับสมัครเข้าระบบ

---

### Request

**Headers**

```
Content-Type: application/json
```

**Body**

```json
{
    "email":"tester@example.com",
    "password":"test_register123",
    "first_name":"tester",
    "last_name":"register",
    "physical_gender":"male",
    "phone_number":"0987654321",
    "signup_source":"web"
}
```

---

### Response (Success)

**Status**

```
200 OK
```

## 🔐 Login

### POST http://127.0.0.1:3455/api/v1/users/auth/login

ใช้สำหรับเข้าสู่ระบบ และสร้าง session

---

### Request

**Headers**

```
Content-Type: application/json
```

**Body**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

---

### Response (Success)

**Status**

```
200 OK
```

**Set-Cookie**

```
session_id=<uuid>; HttpOnly; Path=/;
```

**Body**

```json
{
  "status": 1,
  "message": "login success"
}
```

---

### Response (Invalid Credentials)

**Status**

```
401 Unauthorized
```

**Body**

```json
{
  "status": 0,
  "error": "invalid email or password"
}
```

---

### Notes

* Cookie ถูกตั้งค่าเป็น `HttpOnly`
* Client ไม่สามารถอ่านค่า `session_id` ได้จาก JavaScript
* Session จะถูกใช้โดย API อื่นโดยอัตโนมัติ

---

## 👤 Get Current User (Session Validation)

### GET http://127.0.0.1:3455/api/v1/users/auth/me

ใช้ตรวจสอบว่า session ใน Cookie ยังใช้งานได้หรือไม่
และดึงข้อมูลผู้ใช้ปัจจุบัน

---

### Request

**Headers**

```
Cookie: session_id=<uuid>   (ส่งอัตโนมัติ)
```

---

### Response (Success)

**Status**

```
200 OK
```

**Body**

```json
{
  "status": 1,
  "data": {
    "id": "b1b0c2a1-9f9c-4c7a-9d77-2a6e6f93f3a1",
    "email": "user@example.com",
    "display_name": "Test User"
  }
}
```

---

### Response (Session Not Found / Expired)

**Status**

```
401 Unauthorized
```

**Body**

```json
{
  "status": 0,
  "error": "session not found"
}
```

---

### Response (No Cookie)

**Status**

```
401 Unauthorized
```

**Body**

```json
{
  "status": 0,
  "error": "unauthorized"
}
```

---

### Notes

* API นี้ใช้สำหรับ:

  * ตรวจสอบว่า user login อยู่หรือไม่
  * ใช้ตอนโหลดหน้าเว็บ (เช่น `/profile`)
* หาก session หมดอายุหรือไม่ถูกต้อง ระบบจะตอบ `401`

---

## 🚪 Logout (ถ้ามี / แนะนำให้เพิ่ม)

> หากคุณยังไม่ได้ implement สามารถเพิ่ม endpoint นี้ได้ในอนาคต

### POST http://127.0.0.1:3455/api/v1/users/auth/logout

**Behavior**

* ลบ session จาก Database
* ลบ cookie `session_id`

**Response**

```json
{
  "status": 1,
  "message": "logout success"
}
```

---

## 📊 HTTP Status Code Summary

| Status Code | Meaning               |
| ----------- | --------------------- |
| 200         | Success               |
| 400         | Bad Request           |
| 401         | Unauthorized          |
| 500         | Internal Server Error |

---

## 🧪 Example: Testing with curl

### Login

```bash
curl -i -X POST http://127.0.0.1:3455/api/v1/users/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"password123"}'
```

---

### Check Session

```bash
curl -i GET http://127.0.0.1:3455/api/v1/users/auth/me \
  --cookie "session_id=<uuid>"
```

---

## 🔒 Security Notes

* ใช้ Session + HttpOnly Cookie
* ไม่ส่ง token ผ่าน URL หรือ Header
* ป้องกัน XSS จากการเข้าถึง session
* เหมาะสำหรับ Web Application

---

## 📌 Summary

* API ออกแบบตาม RESTful principle
* Authentication ใช้ Session-based (Server-side)
* ใช้งานง่ายสำหรับ frontend
* รองรับ Docker Compose แบบ zero-config
