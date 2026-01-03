# HealthTech Document Service PoC

A secure backend service for managing medical documents, managing uploads, listing, downloads, and deletions.

## Features
- **Upload**: PDF documents (Max 10MB).
- **Security**: Mock JWT Authentication, File Size/Type validation.
- **Storage**: Local filesystem.
- **Caching**: Redis for metadata caching.
- **Database**: MySQL for structured data.

## Prerequisites
- Docker & Docker Compose
- Go 1.21+ (for local dev without Docker)

## Setup & Run

1.  **Clone the repository**.
2.  **Start the services**:
    ```bash
    docker-compose up --build
    ```
    This will start:
    - `healthtech-app` (Go Backend) on port `8080`
    - `healthtech-db` (MySQL) on port `3306`
    - `healthtech-redis` (Redis) on port `6379`

3.  **Tear down**:
    ```bash
    docker-compose down -v
    ```

## API Usage

The service comes with a test user:
- **Username**: `patient`
- **Password**: `123`

### 1. Register
```bash
 POST http://localhost:8080/register \
  - "Content-Type: application/json" \
  -request: '{
  "register_id":1,
  "username": "patient", 
  "password": "123",
  "role_id":2
}'
```
*Response: User mapped to role*

### 2. Login
```bash
 POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
  "username": "patient", 
  "password": "123"
}'
```
*Response: Returns a JWT Token. Export this as `$TOKEN`.*

### 3. Upload Document
```bash
 POST http://localhost:8080/documents/upload \
  -H "Authorization: Bearer $TOKEN" 
  -F "file=@/path/to/test.pdf" 
  -F "description=My Report" 
  -F "document_type_id=1" 
  -F "userID=1"
```
*Response: File uploaded and its details*

### 4. List Documents
```bash
GET http://localhost:8080/documents \
  -H "Authorization: Bearer $TOKEN"
  -F "userID=1"
```
*Response: Fetch all pdf for the user*


### 5. Download Document
```bash
GET http://localhost:8080/documents/{id}/download \
  -H "Authorization: Bearer $TOKEN"
  -F "userID=1"
```
*Response: Download pdf for the user*

### 6. Delete Document
```bash
DELETE http://localhost:8080/documents/{id} \
  -H "Authorization: Bearer $TOKEN"
  -F "userID=1"
```
*Response: Delete pdf for the user*

## Assumptions
-   **Security**: The system uses a fixed mocking approach for Auth where `patient` is the only active user. In production, this would be replaced by an IDP.
-   **Storage**: Local storage is used. In production, S3 or Blob Storage would be required.
