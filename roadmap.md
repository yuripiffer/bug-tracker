# roadmap.md

## 1. Project Setup
- **Initialize a Go Project**  
  - Create a new folder for your project (e.g., `bugtracker-backend`).  
  - Run `go mod init yourusername/bugtracker-backend` to set up Go modules.  
  - Add a clear directory structure, for example:
    - `cmd/bugtracker/` for the main entry point.
    - `internal/handlers/` for REST handlers.
    - `internal/db/` for database logic.
    - `internal/models/` for data models.

- **Frontend (React + Next.js)**  
  - Create a separate folder (e.g., `bugtracker-frontend`) and initialize a Next.js app using `npx create-next-app`.  
  - Organize your frontend code under:
    - `pages/` for routes and views.
    - `components/` for reusable UI components.
    - `services/` (optional) for API interaction logic.

## 2. Persistence & Database
- **Use an In-Memory Database**  
  - [HashiCorp’s `go-memdb`][1] allows you to store data without needing an external server.  
  - Implement a schema for the `Bug` model, including fields like `ID`, `Title`, `Description`, `Status`, `Priority`, etc.  
  - Provide CRUD functions (Create, Read, Update, Delete) to interact with the in-memory data.

## 3. Core Features
1. **Bug Model**  
   - Fields: ID, title, description, status, priority, timestamps (optional).
2. **Basic CRUD Operations**  
   - **Create**: Implement an HTTP POST endpoint (`/api/bugs`) to add new bugs.  
   - **Read**: Provide endpoints to list all bugs (`GET /api/bugs`) and view a single bug (`GET /api/bugs/{id}`).  
   - **Update**: Route to edit an existing bug (`PUT /api/bugs/{id}`).  
   - **Delete**: Endpoint to remove a bug (`DELETE /api/bugs/{id}`).
3. **User Interface**  
   - **Next.js** for the frontend with server-side rendering out-of-the-box.  
   - Build simple pages or components for listing bugs, viewing details, and creating/editing items.

## 4. CI/CD Integration
- **Version Control**  
  - Host the project on GitHub (or any Git provider).
- **Automated Builds**  
  - Use GitHub Actions (or another CI tool) to run:
    - `go build` and `go test` on the backend.
    - `npm install` and `npm run build` (plus tests) on the frontend.
- **Static Analysis**  
  - For Go, integrate `golangci-lint`.
  - For the frontend, use ESLint + Prettier.
- **Testing**  
  - **Backend**: Use the built-in `testing` package to unit test the database logic and handlers.
  - **Frontend**: Use Jest or React Testing Library to test components and pages.

## 5. Documentation & Instructions
- **README.md**  
  - Provide clear instructions for running the backend (`go run cmd/bugtracker/main.go`) and the frontend (`npm run dev` in the Next.js folder).
- **API Documentation**  
  - List endpoints (`/api/bugs`, etc.), required parameters, and example requests/responses.
- **Usage Guidelines**  
  - Summarize how to create, list, update, and delete bugs via the UI or via HTTP requests.

## 6. Optional Enhancements
- **User Authentication**  
  - Add basic login to safeguard your routes and demonstrate secure sessions/tokens.
- **Advanced Filtering and Search**  
  - Let users filter bugs by status, priority, or date properties.
- **Export/Import**  
  - Provide a feature to export bugs to CSV or JSON and import them back into the system.
- **Containerization**  
  - Include a `Dockerfile` for the Go backend and a Docker-based setup for the frontend to easily deploy anywhere.

## 7. Milestones
1. **MVP**  
   - Set up the Go in-memory database and implement basic CRUD for bugs.  
   - A simple Next.js page to list and create bugs.
2. **Backend Completion**  
   - Add validation, error handling, and comprehensive unit tests.  
   - Ensure the CI pipeline is running the tests automatically.
3. **Frontend Completion**  
   - Refine UI (potentially use a UI library like Chakra UI or Material UI).  
   - Add forms for editing existing bugs and polish the user experience.
4. **Testing & QA**  
   - Increase test coverage on both backend (integration tests) and frontend (component tests).
5. **Documentation & Final Demo**  
   - Finalize README, add screenshots, and demo the app’s key features.

---

**References**  
[1]: https://github.com/hashicorp/go-memdb (HashiCorp’s go-memdb)  
[Next.js Official Docs](https://nextjs.org/docs)
