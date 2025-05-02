# URL Shortening service with GO and Postgres

A simple and efficient URL shortener service built with Go, Gin, and PostgreSQL. This project allows users to shorten long URLs, retrieve the original URLs using the shortened ones, and update existing shortened URLs.

## Features
- Shorten long URLs into unique short URLs using very famous Snowflake ID generation.
- Redirect to the original URL using the short URL.
- Update the long URL associated with a short URL.
- Persistent storage using PostgreSQL.
- Scalable ID generation using Snowflake.

---

## Prerequisites
Before setting up the project, ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.19 or later)
- [Docker](https://www.docker.com/) (for PostgreSQL setup)
- [Postman](https://www.postman.com/) (optional, for testing APIs)

---

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/AshiishKarhade/url-shortner-go.git
cd url-shortner-go
```

### 2. Start PostgreSQL with Docker
Run the following command to start a PostgreSQL container:
```bash
docker run --name postgres-container -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=url_shortner -p 5432:5432 -d postgres
```

### 3. Configure the Database
The default database configuration is already set in `pkg/database/postgresql.go`. If needed, update the `DefaultConfig` function with your custom database credentials.

### 4. Install Dependencies
Run the following command to install all required Go modules:
```bash
go mod tidy
```

### 5. Run the Application
Start the server with:
```bash
go run cmd/main.go
```
The server will start on `http://localhost:8080`.

---

## API Endpoints

### 1. **Shorten URL**
- **Method**: `POST`
- **URL**: `http://localhost:8080/shorten`
- **Request Body**:
  ```json
  {
    "long_url": "https://example.com"
  }
  ```
- **Response**:
  ```json
  {
    "short_url": "abc1234",
    "long_url": "https://example.com"
  }
  ```

### 2. **Redirect URL**
- **Method**: `GET`
- **URL**: `http://localhost:8080/{shortURL}`
- Replace `{shortURL}` with the short URL (e.g., `abc1234`).
- **Response**: Redirects to the original URL.

### 3. **Update URL**
- **Method**: `PUT`
- **URL**: `http://localhost:8080/{shortURL}`
- Replace `{shortURL}` with the short URL (e.g., `abc1234`).
- **Request Body**:
  ```json
  {
    "long_url": "https://newexample.com"
  }
  ```
- **Response**:
  ```json
  {
    "message": "URL updated successfully"
  }
  ```

---

## How to Test
1. Use [Postman](https://www.postman.com/) or `curl` to test the API endpoints.
2. Example `curl` command to shorten a URL:
   ```bash
   curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"long_url": "https://example.com"}'
   ```

---

## Project Structure
- `cmd/main.go`: Entry point of the application.
- `internal/controller`: Contains API controllers.
- `internal/service`: Business logic for URL operations.
- `internal/repository`: Database interaction layer.
- `pkg/database`: PostgreSQL connection and schema initialization.
- `pkg/snowflake`: Snowflake ID generator.

---

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.
