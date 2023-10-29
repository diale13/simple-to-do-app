# Simple-to-do-app
📋 This project is a simple yet effective way to delve into the worlds of both Svelte and Golang. While this app may not be intended for real-world use, it serves as an excellent learning exercise to grasp the fundamentals of both technologies.

## Installation
🚀 Get the app up and running in no time with these simple steps.

### Backend
1. Clone the repository, install dependencies and run the server:
```shell
git clone https://github.com/diale13/simple-todo-app.git
cd simple-todo-app/backend
go mod tidy
go run .
```
### Frontend
```shell
cd ../frontend
npm i
yarn dev
```

## Usage
📝 Start managing your to-do list effortlessly with this app.

* Open your web browser and go to http://localhost:5173 to access the app.
* Use the "Tasks" component to create, edit, and delete your to-do items.
* Stay organized and productive!

 ![example](https://github.com/diale13/simple-to-do-app/assets/33520681/7cbd8c7f-f5fb-449a-91cc-c45caef749b0)


##  Structure
📁 This project is organized into two main folders: backend and frontend.
backend: Contains the API that handles CRUD requests and stores information in a simple .json file.
frontend: Houses two components:
* task: This component handles individual tasks.
* tasks: The main component for managing your to-do list.

## Docker Compose
🐳 If you prefer using Docker Compose, follow these steps to deploy the app.

Make sure you have Docker and Docker Compose installed on your system.

From the project's root directory, where the docker-compose.yml file is located, run the following command:
```shell
docker-compose up
```
Wait for the containers to start, and then access the app at http://localhost:5173.

That's it! You're now ready to start using the Simple To-Do App with Docker Compose.

Enjoy staying organized and productive! 😄📅



