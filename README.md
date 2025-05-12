# Minimalist Blog with SQLite Database

A simple, elegant blog system built with Go and SQLite as a final project for an introduction to databases class. This project demonstrates database design principles, SQL interactions, and basic web server functionality.

## Project Description

This project implements a minimalist blog platform with a clean, text-focused interface. The system consists of:

1. **Database Layer:** SQLite implementation with tables for admin users, blog posts, comments, and images
2. **Repository Layer:** Go interfaces and implementations for database operations
3. **Web Server:** Simple HTTP server rendering templates and handling form submissions
4. **Frontend:** Minimalist HTML templates with a dark theme

The project demonstrates essential database concepts including:
- Entity-relationship modeling
- Schema design with primary and foreign keys
- SQL operations (SELECT, INSERT, UPDATE, DELETE)
- Foreign key constraints and cascading deletes
- Repository pattern for data access abstraction

## Interesting Challenges

### 1. Repository Pattern Implementation

Implementing a clean repository pattern required careful consideration of how to abstract SQL operations while still maintaining type safety and error handling. The solution uses interfaces to define operations and concrete implementations to handle the actual SQL.

```go
type PostRepository interface {
    Create(post *models.Post) error
    GetByID(id int64) (*models.Post, error)
    GetAll() ([]*models.Post, error)
    Update(post *models.Post) error
    Delete(id int64) error
}
```

This approach allows for:
- Clear separation between data access and business logic
- Potential future swapping of database implementations
- Easier testing with mock implementations

### 2. Relational Data Handling

Handling one-to-many relationships (posts to comments) required careful consideration of query design and data retrieval patterns. This was solved by implementing specialized repository methods like `GetByPostID` that retrieve related records efficiently.

### 3. Minimalist UI Design

Creating a simple but effective UI that focuses on content while maintaining good usability required balancing minimalism with functionality. The design draws inspiration from text-focused blogs with clear typography and minimal distractions.

## How to Run the Code

### Prerequisites

- Go (version 1.15 or higher)
- SQLite3

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/blog.git
   cd blog
   ```

2. Install dependencies:
   ```
   go get github.com/mattn/go-sqlite3
   ```

### Running the Application

1. Initialize the database (first time only):
   ```
   cd /path/to/blog
   sqlite3 blog.db < database/schema.sql
   ```

2. Start the web server:
   ```
   go run server.go
   ```

3. Access the blog in your browser:
   ```
   http://localhost:8080
   ```

### Adding Content

To add your first blog post, you have two options:

1. **Using the SQLite CLI:**
   ```
   sqlite3 blog.db
   > INSERT INTO posts (title, content) VALUES ('My First Post', 'Hello world! This is my first blog post.');
   ```

2. **Using Go code:**
   Create a simple script that uses the PostRepository to add content.