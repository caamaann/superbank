# Customer Dashboard Application

A full-stack application with Next.js frontend, Go backend, and PostgreSQL database for customer data management.

## Features

- Customer search by ID, name, or email
- View customer details
- View bank account information
- View pocket information (savings goals)
- View term deposit information
- JWT-based authentication
- Containerized with Docker

## Tech Stack

### Frontend

- Next.js (React)
- ShadcnUI (Tailwind CSS-based components)
- TypeScript

### Backend

- Go (Golang)
- Gin Web Framework
- Clean Architecture
- JWT Authentication

### Database

- PostgreSQL

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Node.js (for local development)
- Go (for local development)
- PostgreSQL (for local development)

### Using Docker (recommended)

1. Clone the repository:

   ```bash
   git clone https://github.com/caamaann/superbank.git
   cd superbank
   ```

2. Start the application using Docker Compose:

   ```bash
   docker-compose up -d
   ```

3. Please make sure in your docker, that all service is Running (db, backend, frontend):

   ```bash
   docker ps
   ```

4. Access the application:
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080

### Manual Setup

#### Backend Setup

1. Navigate to the backend directory:

   ```bash
   cd backend
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file based on `.env.example`

4. Start the backend server:
   ```bash
   go run cmd/server/main.go
   ```

#### Frontend Setup

1. Navigate to the frontend directory:

   ```bash
   cd frontend
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Create a `.env.local` file based on `.env.example`

4. Start the development server:
   ```bash
   npm run dev
   ```

#### Database Setup

1. Create a PostgreSQL database:

   ```bash
   createdb superbank
   ```

2. Run the migration script:
   ```bash
   psql -d superbank -f migrations/01_create_tables.sql
   ```

## API Endpoints

### Authentication

- `POST /api/auth/login` - Login with username and password
- `POST /api/auth/register` - Add User with username and password

### Customer Data

- `GET /api/customers/search?q={query}` - Search for customers
- `GET /api/customers/{id}` - Get customer by ID

## Testing

### Backend Tests

Run the backend tests:

```bash
cd backend
go test -v ./...
```

## Frontend Components

The frontend is built with Next.js and ShadcnUI. The main components are:

- **Dashboard** - Main page with search functionality and customer data display
- **CustomerDetails** - Displays basic customer information
- **BankAccountInfo** - Displays customer bank accounts
- **PocketInfo** - Displays customer savings pockets with goals
- **TermDeposits** - Displays customer term deposits

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Next.js](https://nextjs.org/)
- [ShadcnUI](https://ui.shadcn.com/)
- [Gin Web Framework](https://gin-gonic.com/)
- [PostgreSQL](https://www.postgresql.org/)
