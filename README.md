# Likes Me 
Likes Me is a web application for managing members and products. It provides a RESTful API for performing CRUD (Create, Read, Update, Delete) operations on member data and products data.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [API Documentation](#api-documentation)
- [Usage](#usage)
  - [Members](#members)
    - [Get all members](#get-all-members)
    - [Get member by ID](#get-member-by-id)
    - [Add a new member](#add-a-new-member)
    - [Update an existing member](#update-an-existing-member)
    - [Delete a member](#delete-a-member)
  - [Products](#products)
    - [Get product with reviews](#get-product-with-reviews)
    - [Like a review](#like-a-review)
    - [Cancel like on a review](#cancel-like-on-a-review)
- [Contributing](#contributing)
- [License](#license)

## Features

- Get all members
- Get member by ID
- Add a new member
- Update an existing member
- Delete a member
- Get product with reviews
- Like a review
- Cancel like on a review

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
git clone https://github.com/adisusilayasa/social_media.git

2. Change to the project directory:
cd social_media


3. Install the dependencies:
go mod download

4. Create database. You can find the sql in ./config/db 

5. Build the application:
go build

6. Run the application:
You can use makefile:
make run



The application should now be running on `http://localhost:8080`.

## API Documentation

API documentation is available at `http://localhost:8080/swagger/index.html`. You can use this documentation to explore the available API endpoints, view request and response examples, and test the API using the interactive Swagger UI.

## Usage

### Members

#### Get all members

Endpoint: `GET /members/all`

This endpoint returns a list of all members.

#### Get member by ID

Endpoint: `GET /members/{id}`

This endpoint returns a member based on the provided ID.

#### Add a new member

Endpoint: `POST /members/`

This endpoint adds a new member to the system.

#### Update an existing member

Endpoint: `PUT /members/{id}`

This endpoint updates an existing member with the provided ID.

#### Delete a member

Endpoint: `DELETE /members/{id}`

This endpoint deletes a member based on the provided ID.

### Products

#### Get product with reviews

Endpoint: `GET /products/{id}`

This endpoint returns a product along with its reviews based on the provided ID.

#### Like a review

Endpoint: `POST /products/reviews/{userId}/{id}/like`

This endpoint allows a user to like a review by review ID and user ID.

#### Cancel like on a review

Endpoint: `DELETE /products/reviews/{userId}/{id}/like`

This endpoint cancels the like on a review by review ID and user ID.

## Contributing

Contributions to Likes Me are welcome and encouraged! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## License

[MIT License](https://opensource.org/licenses/MIT)




