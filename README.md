
# Movies API

A simple RESTful API for managing a collection of movies, built using Go (Golang) and the Gorilla Mux router.

## Features

- Get all movies
- Get a specific movie by ID
- Create a new movie
- Update an existing movie
- Delete a movie

## Endpoints

| Method | Endpoint         | Description                  |
|--------|------------------|------------------------------|
| GET    | `/movies`        | Get all movies              |
| GET    | `/movies/{id}`   | Get a specific movie by ID  |
| POST   | `/movies`        | Add a new movie             |
| PUT    | `/movies/{id}`   | Update an existing movie    |
| DELETE | `/movies/{id}`   | Delete a movie              |

## Request and Response Examples

### GET `/movies`
#### Response:
```json
[
  {
    "id": "1",
    "isbn": "4587227",
    "title": "movie 01",
    "director": {
      "firstname": "John",
      "lastname": "Cameron"
    }
  },
  {
    "id": "2",
    "isbn": "4587228",
    "title": "movie 02",
    "director": {
      "firstname": "Tim",
      "lastname": "Smith"
    }
  }
]
```

### POST `/movies`
#### Request:
```json
{
  "isbn": "1234567",
  "title": "New Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}
```
#### Response:
```json
{
  "id": "3",
  "isbn": "1234567",
  "title": "New Movie",
  "director": {
    "firstname": "Jane",
    "lastname": "Doe"
  }
}
```

## Prerequisites

- Go 1.18+ installed
- Gorilla Mux library installed (`go get -u github.com/gorilla/mux`)

## How to Run

1. Clone the repository:
   ```bash
   git clone https://github.com/<your-username>/<repo-name>.git
   cd <repo-name>
   ```

2. Build and run the server:
   ```bash
   go run main.go
   ```

3. The server will run on `http://localhost:8000`.

## Technologies Used

- **Go (Golang):** Programming language
- **Gorilla Mux:** Router for handling API routes

## License

This project is licensed under the MIT License. Feel free to use and modify it.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---

### Author

Developed by B.K Jeeban(https://github.com/Jeeban369).

