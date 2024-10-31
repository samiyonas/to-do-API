# TO-DO API

# Description

to-do-API is a simple task management API written in Go. The API allows users to:

* Create users by providing their name and email.
* Add tasks associated with a specific user.
* Retrieve tasks by their unique ID or fetch all tasks for a specific user.
This API provides endpoints for creating, retrieving, and managing tasks for users in a simple and efficient way. It’s designed to be easy to integrate into any task management system or to serve as a backend for a productivity app.

# Key features:

* Create user accounts and associate tasks with them.
* Track task completion status (done field).
* Retrieve individual tasks or all tasks for a user.

# Prerequisites
Before setting up and running the project, ensure you have the following installed:

Go (1.18 or later)
Git 
To confirm that Go is installed, you can run:

```bash
go version
```
# Installation
To get started with this project, follow the steps below:

Clone the repository:

First, clone the project repository to your local machine using Git:

```bash
git clone https://github.com/your-username/your-repo-name.git
```
Navigate to the project directory:

Change into the project’s root directory:

```bash
cd to-do-API
```
# Install dependencies:

The project uses Go modules. Download the required dependencies by running:

```bash
go mod tidy
```
This command ensures that all dependencies are resolved and downloaded.

# Running the Application
Once dependencies are installed, you can run the application in the following ways:

Option 1: Run Directly with go run
You can run the application directly from the source code:

```bash
go run main.go
```
Option 2: Build and Execute
Alternatively, you can build an executable binary and run it:

```bash
go build -o to-do-API
./to-do-API
```

# API Endpoints
1. Create a User
Use this endpoint to create a new user.

```bash
curl -X POST localhost:8080/register \
-H "Content-Type: application/json" \
-d '{"name": "<Your name>", "email": "<Your email>"}'
```
Note:
After creating a user, you will receive a unique user_id. Store this ID safely, as you will need it to add and manage tasks.

2. Create a Task
Use this endpoint to add a new task for a specific user.

```bash
curl -X POST localhost:8080/addtask \
-H "Content-Type: application/json" \
-d '{"user_id": <user ID>, "title": "<Title>", "content": "<Content>", "done": <true or false>}'
```
user_id: Use the ID provided when you registered the user.
done: Set this to true if the task is already completed; otherwise, set it to false.
Note:
After creating a task, you will receive a unique task_id. Store this ID safely, as you will need it to retrieve or manage tasks.

3. Retrieve a Task by Task ID
Use this endpoint to retrieve a task by its unique task_id.

```bash
curl localhost:8080/taskbyid/<taskID>
```
You will need the specific task_id to fetch the task details.

4. Retrieve All Tasks of a User
Use this endpoint to retrieve all tasks associated with a particular user.

```bash
curl localhost:8080/alltasks/<user_id>
```
You will need the user_id provided during registration to fetch the user's tasks.
