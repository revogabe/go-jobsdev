# JobsDev GO

<p align="center">
  <img src="./public/home.png" alt="JobsDev GO - Homepage">
</p>

-----------------------------------

This project is an API in Golang and Frontend in Next.js with the intention of studying more about Go and also helping the community with the job search in the market.

I used in this project techs like: Go-Gin as a router, GoORM for database communication, MongoDB Atlas as a database. The project follows a modern package structure to keep the code base organized and maintainable. I also created an authenticated area so that I can approve or delete jobs posted by users.

## Preview

<p align="center">
  <img src="./public/home-info.png" alt="JobsDev GO - Homepage - Info">
</p>

<p align="center">
  <img src="./public/dashboard.png" alt="JobsDev GO - Dashboard">
</p>

<p align="center">
  <img src="./public/dashboard-delete.png" alt="JobsDev GO - Dashboard Info">
</p>

---

## Features

- Introduction to Golang and building modern APIs
- Development environment setup for creating the API
- Using Go-Gin as a router for route management
- Implementing MongoDB as the database for the API
- Using MongoDB Go Driver for communication with the database
- Creating a modern package structure for organizing the project
- Implementing a complete job opportunities API with endpoints for searching, creating, editing, and deleting opportunities
- Implementing automated tests to ensure API quality

## Live Version

You can check the documentation and perform tests to the API by visiting the live version hosted at [JobsDev GO - Website](https://jobsdev-go.vercel.app/).

## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/username/repo-name.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Makefile Commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application without generating API documentation.
- `make run-with-docs`: Generate the API documentation using Swag, then run the application.
- `make build`: Build the application and create an executable file named `gopportunities`.
- `make test`: Run tests for all packages in the project.
- `make docs`: Generate the API documentation using Swag.
- `make clean`: Remove the `gopportunities` executable and delete the `./docs` directory.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```sh
make run
```

## Docker and Docker Compose

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment. Here are the most common Docker and Docker Compose commands you may want to use:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Run a container based on the built image. Replace `your-image-name` with the name you used when building the image. You can change the port number if necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Build the services defined in the `docker-compose.yml` file.
- `docker compose up`: Run the services defined in the `docker-compose.yml` file.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```sh
docker-compose down
```

For more information on Docker and Docker Compose, refer to the official documentation:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver) for database communication
- [MongoDB Atlas](https://www.mongodb.com/atlas/database) as the database

## Contributing

To contribute to this project, please follow these guidelines:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them using Conventional Commits
4. Push to the branch: `git push origin feature/your-feature-name`
5. Submit a pull request

---

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.

## Credits

This project was created by [Daniel Gabriel](https://github.com/revogabe).
