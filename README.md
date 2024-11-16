## 🌍 Go Hello World Project 🌍

This project creates a basic web application using the Go programming language. The application displays "Hello World!" when accessed at the root URL. It is designed to be deployed on Heroku using Git.

## 📰 Description
This project uses Go to develop a simple web server that responds with "Hello World!" It is configured to run on Heroku without Docker, simplifying the deployment process.

## 🎪 Project Structure

The project structure is as follows:

Go_holamundo/
main.go       
go.mod         
go.sum         
Procfile

## 🔐 Key Files

main.go: Contains the HTTP server and core application logic.

go.mod and go.sum: Define the module and its dependencies.

Procfile: Tells Heroku how to start the application.

## 📖 Requirements

Go: Ensure Go is installed (preferably the latest version).

Git: Required for managing the source code and deploying to Heroku.

Heroku CLI: To deploy the application on Heroku.

## 🔨 Installation

### 1. Clone the Repository
   
Clone this repository to your local machine:

```bash
git clone https://github.com/Carlosdhc10/Go_holamundo.git
```

### 2. Navigate to the Project Directory

Enter the project directory:

```bash
cd Go_holamundo
```

### 3. Run Locally

Start the application locally:

```bash
go run main.go
```

Access the application at http://localhost:8080.

## ✈️ Deployment on Heroku

### 1. Log In to Heroku
   
Log in to Heroku using the following command:

```bash
heroku login
```

### 2. Create a Heroku Application

Create a new application on Heroku:

```bash
heroku create
```

### 3. Configure the Procfile

Ensure your project includes a Procfile with the following content:

```bash
web: go run main.go
```

This tells Heroku how to start your application.

### 4. Deploy to Heroku
   
Add changes to Git:

```bash
git add .
git commit -m "Initial deployment to Heroku"
git push heroku main
```

### 6. Open the Application

Once deployed successfully, open your application in a browser:

```bash
heroku open
```

## 🎈 View Logs

If you encounter any issues, view the Heroku logs using:

```bash
heroku logs --tail
```

## 🎳 Contributions
Feel free to fork this repository and submit pull requests to improve the project.

## 📜 License
This project is open source. You can use and adapt it according to your needs.
