
# GoGinJSONResponse

This is a simple API built with the **Gin** framework and **GORM** for Go, which connects to an SQLite database to fetch and respond with JSON data.

## Requirements

- Go 1.18+ (or any compatible version)
- SQLite3 (used as the database)

## Setup

### 1. Clone the repository

First, clone the repository to your local machine:

```bash
git clone https://github.com/tofu639/GoGinJSONResponse.git
cd GoGinJSONResponse
```

### 2. Install dependencies

Make sure Go is installed on your machine. Then, run the following command to install the required dependencies:

```bash
go mod tidy
```

This will install the necessary dependencies, including:
- **Gin** (web framework)
- **GORM** (ORM for Go)
- **SQLite dialect for GORM**

### 3. Run the application

To run the application, use the following command:

```bash
go run main.go
```

This will start the server, and it will listen on `http://localhost:8080`.

### 4. Access the API

Once the server is running, you can interact with the API.

#### API Endpoints:

- **GET /users**: Fetches a list of users, along with their associated posts.

#### Example response from **GET /users**:

```json
[
  {
    "id": 1,
    "name": "Alice",
    "posts": [
      {
        "id": 1,
        "title": "Post 1",
        "body": "This is the first post",
        "user_id": 1
      },
      {
        "id": 2,
        "title": "Post 2",
        "body": "This is the second post",
        "user_id": 1
      }
    ]
  },
  {
    "id": 2,
    "name": "Bob",
    "posts": [
      {
        "id": 3,
        "title": "Bob's First Post",
        "body": "Bob's post content",
        "user_id": 2
      }
    ]
  }
]
```