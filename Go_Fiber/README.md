# Go Fiber API with PostgreSQL

This is a RESTful API built using Go Fiber framework with PostgreSQL as the database. The project is containerized using Docker for easy development and deployment.

## 🚀 Features

- RESTful API endpoints
- PostgreSQL database integration using GORM
- Docker and Docker Compose support
- Environment variable configuration
- Product management system

## 📋 Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.22.5 or higher)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## 🛠️ Tech Stack

- [Go Fiber](https://gofiber.io/) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Docker](https://www.docker.com/) - Containerization

## 🔧 Installation & Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/TaiChi112/Golang.git
   cd Go_Fiber
   ```

2. Create a `.env` file in the root directory with the following variables:
   ```env
   DB_HOST=postgres
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name
   DB_PORT=5432
   APP_PORT=3000
   ```

3. Start the application using Docker Compose:
   ```bash
   docker-compose up --build
   ```

The API will be available at `http://localhost:3000`

## 🌐 API Endpoints

The API provides endpoints for managing products:

- `GET /api/products` - Get all products
- `GET /api/products/:id` - Get a specific product
- `POST /api/products` - Create a new product
- `PUT /api/products/:id` - Update a product
- `DELETE /api/products/:id` - Delete a product

## 📝 Product Model

```go
type Product struct {
    ID          uint      `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    PriceCents  int64     `json:"price_cents"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## 🐳 Docker Support

The project includes:
- `Dockerfile` for building the application container
- `docker-compose.yml` for orchestrating the application and database containers

## 👥 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## ✨ Author

[TaiChi112](https://github.com/TaiChi112)
