# Info-OS

## Description

Api-infoOs is a Api for managing a fixing shop. It aims to provide a simple and easy way to manage your shop. It is a REST API that can be used to create, update, and delete shops. The API also provides endpoints for managing products, orders, and customers.
The API is built using golang and MySQL.

## Features

- authentication
- authorization
- client management
- shop management
- product management
- order management
- customer management
- inventory management

## Requirements

- golang
- mysql
- docker
- docker-compose

## Installation

To install Api-infoOs, follow these steps:

1. Clone the repository: `git clone https://github.com/info-os/api-infoOs.git`
2. Navigate to the project directory: `cd api-infoOs`
3. Install the dependencies: `go mod download`
4. Create a .env file in the root directory of the project with the following content:

```
URL_DATABASE=databaseUrl
JWT_SECRET=jwtSecret
```
5. Run database in container: `docker-compose up -d`
6. Run Api-infoOs: `go run main.go`
7. If you want to use hot reload install air: `go install github.com/air-verse/air@latest`
8. Access the Swagger documentation at http://localhost:8080/swagger/index.html


## Usage

To use Api-infoOs, you can access the Swagger documentation at the following URL: http://localhost:8080/swagger/index.html

## Contributing

If you would like to contribute to Api-infoOs, please follow these steps:

1. Fork the repository
2. Create a new branch for your changes
3. Make your changes and commit them
4. Push your changes to your forked repository
5. Create a pull request to the main repository

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Contact

If you have any questions or suggestions, please contact me at josiel.jcc@hotmail.com.
