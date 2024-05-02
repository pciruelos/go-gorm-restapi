## Project Overview
This Go project uses Gorilla Mux for routing and GORM for object-relational mapping, providing a RESTful API to manage users and tasks.
It features CRUD operations and automatic schema migration for a PostgreSQL database.

## Prerequisites
Go version 1.15 or later.
PostgreSQL.

## Setup and Installation
Clone the repository: Provide instructions on how to clone your repository.
Install dependencies: List any Go packages that need to be installed using go get.
Configure the database: Mention the need to update the database connection string in db/db.go.

## Running the Application
Provide a command to run the application, typically go run main.go, and note that it serves on http://localhost:3000.

## API Endpoints
Detail your API endpoints with their functions:

GET /users: Fetch all users.
POST /users: Create a user.
GET /user/{id}: Get a specific user.
DELETE /users/{id}: Delete a user.
Similar endpoints for tasks.

## Database Models
Describe your User and Task models, noting fields like:
User: FirstName, LastName, Email (unique)
Task: Title (unique), Description, Done (default: false)