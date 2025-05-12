## Project Description

This project implements a minimalist blog platform with a clean, text-focused interface. The system consists of:

1. **Database:** SQLite implementation with tables for admin users, blog posts, comments, and images
3. **Web Server:** Simple HTTP server rendering templates and handling form submissions
4. **Frontend:** Minimalist HTML templates with a dark theme

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
