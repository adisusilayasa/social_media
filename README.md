# social_media

# Project Name

Likes Me is a web application for managing members. It provides a RESTful API for performing CRUD (Create, Read, Update, Delete) operations on member data.

## Features

- Get all members
- Get member by ID
- Add a new member
- Update an existing member
- Delete a member

## Technologies Used

- Go: Programming language used to develop the application
- Echo: Web framework used for building the RESTful API
- OpenTracing: Distributed tracing system used for monitoring and debugging the application
- Zap: Logging library used for capturing application logs
- Swagger: API documentation tool used to generate API documentation

## Getting Started

### Prerequisites

- Go 1.16 or later installed on your machine

### Installation

1. Clone the repository:
git clone https://github.com/your-username/project-name.git

2. Change to the project directory:
cd social_media


3. Install the dependencies:
go mod download


4. Build the application:
go build

5. Run the application:
You can use makefile:
make run



The application should now be running on `http://localhost:8080`.

## API Documentation

API documentation is available at `http://localhost:8080/docs`. You can use this documentation to explore the available API endpoints, view request and response examples, and test the API using the interactive Swagger UI.

## Usage

### Get all members

Endpoint: `GET /members/all`

This endpoint returns a list of all members.

### Get member by ID

Endpoint: `GET /members/{id}`

This endpoint returns a member based on the provided ID.

### Add a new member

Endpoint: `POST /members/`

This endpoint adds a new member to the system.

### Update an existing member

Endpoint: `PUT /members/{id}`

This endpoint updates an existing member with the provided ID.

### Delete a member

Endpoint: `DELETE /members/{id}`

This endpoint deletes a member based on the provided ID.

## Contributing

Contributions to Project Name are welcome and encouraged! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## License

[MIT License](https://opensource.org/licenses/MIT)






