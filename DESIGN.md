# System Design & Architecture Document - HealthTech Document Service

## 1. Tech Stack & Architecture

### Tech Stack
- **Language**: Golang (1.21+)
  - *Why*: Strong typing, excellent concurrency support, fast execution, and a robust standard library make it ideal for backend services.
- **Web Framework**: Gin
  - *Why*: High performance, simple API, and robust middleware support.
- **Database**: MySQL (8.0)
  - *Why*: Reliable relational database, conforms to the provided schema requirements.
- **Caching**: Redis (7.0)
  - *Why*: Fast in-memory key-value store for caching document metadata.
- **Containerization**: Docker & Docker Compose
  - *Why*: Ensures consistent environment across development and production.

### Architecture
The system follows a layered architecture (Clean Architecture inspired):
1.  **Handler Layer**: Handles HTTP requests, validation, and sending responses.
2.  **Service Layer**: content business logic, file processing, and caching strategies.
3.  **Repository Layer**: Direct database interactions.

## 2. Data Flow

### Upload Flow
1.  **Client** sends `POST /documents/upload` with `multipart/form-data` (File + Metadata).
2.  **Auth Middleware** validates the JWT/Token (Mocked).
3.  **Handler** validates file type (PDF only) and size (<10MB).
4.  **Service**:
    - Generates a unique filename (UUID) to prevent collisions.
    - Saves the file to the local mounted volume (`./uploads`).
    - Calls Repository to save metadata (original name, path, size, patient_id).
5.  **Repository** inserts record into `documentDetails`.
6.  **Response** returns the document ID and metadata.

### Retrieval Flow (List)
1.  **Client** sends `GET /documents`.
2.  **Service** checks.
    - * Queries **Repository** (MySQL)

### Download Flow
1.  **Client** sends `GET /documents/{id}/download`.
2.  **Service** retrieves file path from DB/Cache.
3.  **Handler** streams the file from disk to the response.

## 3. API Specification

| Method    | Endpoint                  | Description                   | Auth Required |

| `POST`    | `/register`               | Register user for login       | No    |
| `POST`    | `/login`                  | Mock login to get token       | No    |
| `POST`    | `/documents/upload`       | Upload a PDF document         | Yes   |
| `GET`     | `/documents`              | List all documents  user      | Yes   |
| `GET`     | `/documents/:id/download` | Download a specific document  | Yes   |
| `DELETE`  | `/documents/:id`          | Delete a document             | Yes   |

### Request/Response Examples

**Upload Response:**
```json
{
  "id": 1,
  "filename": "report.pdf",
  "size": 10240,
  "uploaded_at": "2024-01-01T12:00:00Z"
}
```

## 4. Key Considerations

-   **Scalability**: The service is stateless. For horizontal scaling, the local file storage should be replaced with object storage (S3, GCS).
-   **Security**:
    -   Validate MIME types (magic numbers) not just extensions.
    -   Sanitize filenames to prevent path traversal.
    -   Store files outside the web root.
-   **Concurrency**: Go routines handle concurrent requests efficiently.
-   **Cleanup**: A background job (cron) could remove orphaned files if DB delete fails (though we implement transactional deletion).

## 5. CI/CD Pipeline (Conceptual)

1.  **Lint**: Run `golangci-lint` to check code quality.
2.  **Test**: Run `go test ./...` with coverage.
3.  **Build**: `docker build -t healthtech-backend .`.
4.  **Push**: Push image to registry (e.g., Docker Hub, ECR).

## 6. Infrastructure as Code

-   **Dockerfile**: Multi-stage build (Builder -> Runner) for minimal image size.
-   **Docker Compose**: Orchestrates App and MySQL. Uses a bound volume for uploads persistence.

## 7. Design Justification

-   **File Storage**: Local disk is chosen for the PoC for simplicity and zero cost. For production, S3 is preferred for durability and availability.
-   **Production Readiness**:
    -   Add HTTPS/TLS.
    -   Replace Mock Auth with OAuth2/OIDC.
    -   Centralized logging (ELK/Prometheus).
-   **HIPAA Compliance**:
    -   Encryption at rest (DB & disk).
    -   Encryption in transit (TLS 1.3).
    -   Strict Access Control Logs (Audit Trails).