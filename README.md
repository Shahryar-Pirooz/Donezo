# Donezo

Donezo is a lightweight task management application designed to help you organize and track your tasks effectively. With its simple and efficient architecture, Donezo is perfect for individuals or small teams who need a reliable tool to manage tasks.

## Features
- **Task Management**: Create, view, update, and delete tasks.
- **User Authentication**: Secure access to your tasks using JWT-based authentication.
- **Fast and Scalable**: Built with Go for performance and PostgreSQL for reliable data storage.

## Technologies Used
- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Containerization**: Docker

### Prerequisites
- Go
- PostgreSQL
- Docker

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/donezo.git
   cd donezo
   ```

2. Set up the environment variables:
   Create a `.env` file in the project root with the following:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_postgres_user
   DB_PASSWORD=your_postgres_password
   DB_NAME=donezo
   JWT_SECRET=your_secret_key
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run database migrations (if any):
   ```bash
   go run migrate.go
   ```

5. Start the server:
   ```bash
   go run main.go
   ```

6. Access the API at `http://localhost:8080`.

## API Endpoints
| Method | Endpoint       | Description            |
|--------|----------------|------------------------|
| GET    | `/tasks`       | Get all tasks          |
| POST   | `/tasks`       | Create a new task      |
| PUT    | `/tasks/{id}`  | Update an existing task|
| DELETE | `/tasks/{id}`  | Delete a task          |

## Docker Setup (Optional)
1. Build and run the app with Docker:
   ```bash
   docker-compose up --build
   ```
2. Access the API at `http://localhost:8080`.

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests to improve Donezo.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
Start organizing your tasks with **Donezo** today!

