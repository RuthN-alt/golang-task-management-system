# Golang Task Management System

A robust task management system built with Golang, Gin, and GORM. This system handles user authentication, authorization, and access management, allowing users to create, update, and manage tasks securely.

## Key Features

- Secure user registration and authentication using JWT (JSON Web Tokens)
- Protection against vulnerabilities like SQL injection attacks
- Support for bulk upload of tasks using CSV

## Prerequisites

Make sure you have the following installed:

- Golang

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/RuthN-alt/golang-task-management-system.git
cd golang-task-management-system
````

2. Set up the environment variables:

Copy the `.env.example` file to `.env` and update the values with your own configuration:

```bash
copy .env.example .env
```

## API Endpoints

* **POST /user/register**: Register a new user

* **POST /user/login**: Log in and receive a JWT token

* **PUT /user/logout**: Log out and invalidate the JWT token

* **DELETE /user/delete**: Delete user

* **GET /tasks/:id**: Get a single task by ID

* **GET /tasks/**: Get all tasks

* **POST /tasks/create**: Create a new task

* **POST /tasks/bulkupload**: Bulk upload tasks using CSV

* **PUT /tasks/update/:id**: Update an existing task by ID

* **DELETE /tasks/delete/:id**: Delete a task by ID

## Bulk Upload

To bulk upload tasks from a CSV file, create `data.csv` with key `taskBulkUpload` using this format:

```
Title,Description,Planned Start Date,Planned Start Time,Planned End Date,Planned End Time,Seconds
Task 1,Description 1,2023-07-20,09:00:00,2023-07-20,18:00:00,120000
Task 2,Description 2,2023-07-21,09:30:00,2023-07-21,17:30:00,150000
...
```

<p align="center">
    With ❤️ by <a href="https://github.com/RuthN-alt" target="_blank">Katitu Ruth Naomi</a>
</p>
