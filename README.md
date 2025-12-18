User API (Go + Fiber + SQLC)
===============================

A production-ready RESTful API built with **Go**, **Fiber**, **SQLC**, and **PostgreSQL** to manage users with **date of birth (DOB)** and dynamically calculated **age**.

* * * * *
Features
----------

-   Create, update, delete, and fetch users

-   Store **DOB** in the database

-   Calculate **age dynamically** using Go's `time` package

-   Type-safe database access using **SQLC**

-   Input validation with **go-playground/validator**

-   Structured logging using **Uber Zap**

-   Middleware for:

    -   Request ID injection (`X-Request-ID`)

    -   Request duration logging

-   Clean HTTP status codes and error handling

-   Environment-based configuration

-   Unit test for business logic (age calculation)

* * * * *

🧱 Project Structure
--------------------

```
.
├── cmd/server/main.go
├── config/
│   └── config.go
├── db/
│   ├── migrations/
│   └── sqlc/
├── internal/
│   ├── handler/
│   ├── repository/
│   ├── service/
│   ├── routes/
│   ├── middleware/
│   ├── logger/
│   └── models/
├── logs/
├── go.mod
└── README.md

```

* * * * *

⚙️ Tech Stack
-------------

-   **Go**

-   **GoFiber**

-   **PostgreSQL**

-   **SQLC**

-   **Uber Zap**

-   **go-playground/validator**

* * * * *

🚀 Setup & Run (Local)
----------------------

### 1️⃣ Prerequisites

Make sure you have:

-   Go **1.22+**

-   PostgreSQL running

-   `sqlc` installed

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

```

Verify:

```
sqlc version

```

* * * * *

### 2️⃣ Clone the Repository

```
git clone https://github.com/romitdubey/user-api.git
cd user-api

```

* * * * *

### 3️⃣ Environment Variables


Create a `.env` file **locally**:

```
APP_PORT=8080
APP_ENV=development

DATABASE_URL=postgres://postgres:password@localhost:5432/userdb?sslmode=disable

```

* * * * *

### 4️⃣ Prevent `.env` from Being Pushed

Add this to `.gitignore`:

```
.env
logs/

```

✔ This keeps secrets safe\
✔ Required for production best practices

* * * * *

### 5️⃣ Create Database & Table(Manually Using psql shell)

```
CREATE DATABASE userdb;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL
);

```

* * * * *

### 6️⃣ Generate SQLC Code

```
sqlc generate

```

This generates type-safe Go code in `db/sqlc`.

* * * * *

### 7️⃣ Install Dependencies

```
go mod tidy

```

* * * * *

### 8️⃣ Create Logs Directory

Zap requires the log directory to exist.

```
mkdir logs

```

* * * * *

### 9️⃣ Run the Server

```
go run cmd/server/main.go

```

Server starts on:

```
http://localhost:8080

```

* * * * *

 API Endpoints
----------------

### ➕ Create User

**POST** `/users`

```
{
  "name": "Alice",
  "dob": "1990-05-10"
}

```

* * * * *

### 📄 Get User by ID

**GET** `/users/:id`

```
{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 34
}

```

* * * * *

### ✏️ Update User

**PUT** `/users/:id`

```
{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

```

* * * * *

### 🗑️ Delete User

**DELETE** `/users/:id`

**Response:** `204 No Content`

* * * * *

### 📋 List Users

**GET** `/users`

* * * * *

Unit Tests
-------------

Run unit tests:

```
go test ./...

```

Includes:

-   Business logic test for age calculation

* * * * *

🪵 Logging
----------

-   Structured logs using **Uber Zap**

-   Logs written to:

    -   Terminal (stdout)

    -   `logs/app.log`

-   Each request includes:

    -   Request ID

    -   Duration

    -   HTTP method & path



* * * * *

🎯 Why This Architecture?
-------------------------

-   Clear separation of concerns

-   Easy to test and maintain

-   Production-ready design

-   Follows Go backend best practices



* * * * *