## 💻 Hello World Project with Go and Docker 💻

This project creates a basic web application in Go that displays a "¡Hola Mundo!" message. The application is also dockerized to facilitate deployment to platforms such as Docker Hub or Render.

## 📰 **Description**  

The project uses Go to create a web application that displays the "Hello, World!" message. It is also dockerized to make it easy to deploy to container services like Docker Hub.

## 🎪 **Project Structure**  

The project structure is as follows:

Go_holamundo/

Dockerfile

go.mod

go.sum

main.go

## 📖 **Requirements**  

- **Go**: Make sure you have Go installed to run the application locally.
  
- **Docker**: You need Docker installed to build and run the application inside a container.
  
- **Docker Hub (Optional)**: If you want to push the image to Docker Hub, you need an account on [Docker Hub](https://hub.docker.com/).
  
- **Render (Optional)**: If you want to deploy the application on Render, you need an account on [Render](https://render.com/).

## 🔨 **Installation**  

1. Clone the repository:
   
```bash
git clone https://github.com/Carlosdhc333/Go_holamundo.git
```

2.- Navigate to the project directory:

```bash
cd Go_holamundo
```

3.- Initialize the Go module and download the dependencies:

```bash
go mod tidy
```

## ✈️ Running the Application

To run the application locally, use the following command:

```bash
go run main.go
```

The application will run on port 8080 on your local machine. You can access it at http://localhost:8080.

## 🐳 Building the Docker Image

To build the Docker image, use the following command inside the project directory:

```bash
docker build -t carlosdhc333/go_holamundo:v1 .
```

Once the Docker image is built, you can run the container with the following command:

```bash
docker run --rm -p 8080:8080 carlosdhc333/go_holamundo:v1
```

## 🎈 Pushing to Docker Hub (Optional)

If you want to share the container on Docker Hub, follow these steps:

1.- Log in to Docker Hub:

```bash
docker login
```

2.- Tag the image:

```bash
docker tag carlosdhc333/go_holamundo:v1 your_username/go_holamundo
```

3.- Push the image:

```bash
docker push your_username/go_holamundo
```

## 📈 Deploying to Render (Optional)

To deploy this application on Render, follow these steps:

1.- Upload the repository to GitHub.

2.- On Render, select "New" > "Web Service" and connect the repository.

3.- Set the execution command to go run main.go in the Render settings.

Render will handle building the image and running the application, displaying the message in the service log output.

## 🎳 Contributing

If you want to improve the project, feel free to fork and send a pull request!

## ©️ License ©️

This English version mirrors the structure and instructions from the original markdown, tailored to a Go and Docker application. Let me know if you need any further changes!
