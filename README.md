# Hotel Management System

A multi-tenant hotel management system built with Go, Gin, and GORM(PostgreSQL).

## Features

- Multi-tenant architecture with isolated data per business
- Hotel and room management
- RESTful API endpoints

## Prerequisites

- Go 1.25.4 or later
- PostgreSQL database
- Git

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Amizhthanmd/hotel-management.git
   cd hotel-management
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up environment variables:

   - Copy the `.env` file and update the database credentials:
     ```
     PORT=8000
     POSTGRES_DB_URL=postgresql://username:password@localhost:5432/
     DB_NAME=hotel_management
     ```

4. Ensure PostgreSQL is running and create the database:
   ```bash
   createdb hotel_management
   ```

## Running the Application

Start the server:

```bash
go run cmd/server/main.go
```

The server will start on port 8000 (or the port specified in `.env`).

## API Documentation

[View API Documentation](https://documenter.getpostman.com/view/23869296/2sB3dMxAme)

## Multi-Tenant Architecture

This application uses a **multi-tenant architecture**.

- **Multi-tenancy** means one application serves multiple business (tenants), keeping their data completely separate.
- **Schema-based** means each tenant gets their own database schema in PostgreSQL.
- Each business has its own isolated space for hotels and rooms data.
- The `X-Business-ID` header identifies which tenant's data to access.
- This ensures complete data isolation between different businesses using the same application.

For example:

- Business A creates hotels in schema `tenant_business_a`
- Business B creates hotels in schema `tenant_business_b`
- They cannot see or access each other's data.
