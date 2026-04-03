## Plan: Full-Stack Book Store with Go Fiber, React, and SQLX

This plan outlines the creation of a full-stack bookstore application featuring CRUD operations for books. We will use Go Fiber for a fast backend, SQLX for clean PostgreSQL interactions, and React for a responsive frontend.

### Phase 1: Environment & Project Setup (Backend)

1. **Initialize Backend**: Navigate to [backend/](backend/) and install core dependencies.
   - Modules: `github.com/gofiber/fiber/v2`, `github.com/jmoiron/sqlx`, `github.com/lib/pq` (Postgres driver), and `github.com/joho/godotenv`.
2. **Directory Structure**: Create a clean structure:
   - `handlers/`: For Fiber controllers (CRUD logic).
   - `models/`: For `Book` structs and database logic.
   - `database/`: For connection setup.
   - `routes/`: To define API endpoints.
3. **Configuration**: Create `.env` to store database connection strings and server port.

### Phase 2: Database Layer

1. **Connection Logic**: Implement [backend/database/database.go](backend/database/database.go) to initialize the SQLX session.
2. **Schema Definition**: Create a simple `books` table:
   - `id` (Serial/UUID), `title` (Text), `author` (Text), `published_date` (Date), `price` (Numeric), `image_url` (Text).
3. **Model Methods**: In `models/book.go`, implement methods for:
   - `GetAllBooks()`, `GetBookByID(id)`, `CreateBook(book)`, `UpdateBook(id, book)`, `DeleteBook(id)`.

### Phase 3: API Endpoints (Fiber)

1. **CRUD Handlers**: Create logic in [backend/handlers/book_handler.go](backend/handlers/book_handler.go) to bind JSON bodies to structs and call model methods.
2. **Routing**: Define routes in [backend/routes/routes.go](backend/routes/routes.go):
   - `GET /api/books`, `GET /api/books/:id`, `POST /api/books`, `PUT /api/books/:id`, `DELETE /api/books/:id`.
3. **CORS Middleware**: Enable CORS in the main Fiber app to allow frontend access.

### Phase 4: Frontend Development (React)

1. **Project Scaffold**: Use Vite to create a React project in a new `frontend/` directory.
2. **API Client**: Install `axios` for HTTP requests.
3. **Components**:
   - `BookList`: Renders a grid of books with "Edit" and "Delete" actions.
   - `BookForm`: A reusable component for both "Add" and "Update" functionality.
4. **State Management**: Use standard React `useState` and `useEffect` for data fetching.

### Phase 5: Verification & Testing

1. **API Testing**: Use `curl` or a REST client (Postman/Thunder Client) to verify all CRUD endpoints.
2. **Frontend Integration**: Confirm the browser correctly reflects database changes.
3. **Error Handling**: Verify the UI handles API failures (e.g., database connection down).

### Relevant Files

- [backend/main.go](backend/main.go): Application entry point.
- [backend/database/database.go](backend/database/database.go): SQLX connection.
- [backend/models/book.go](backend/models/book.go): Book schema and SQLX queries.
- [backend/handlers/book_handler.go](backend/handlers/book_handler.go): Logic for request/response.
- `frontend/src/App.jsx`: Main routing and layout.

### Decisions & Assumptions

- **ORM**: SQLX will be used for a balance of speed and control over raw SQL.
- **Project Split**: Backend and Frontend will exist in separate root-level folders.
- **Styling**: Tailwind CSS is recommended for fast UI development on the frontend.
- **Database**: Assumes a running PostgreSQL instance is available.
