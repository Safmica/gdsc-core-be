# GDSC API

REST API that will manage the internal member activity dan administration

## Features (Create, read, update, and delete)
- **User Management**
- **Activity Management**
- **Batch Management**
- **Division Management**
- **Final Project Management**
- **Member Management**
- **Role Management**
- **Participant Management**
- **Batch Configure Management**
- **JWT Authentication (Access and Refresh Token)**

## Tech Stack
- ![VSCode](https://img.shields.io/badge/VSCode-0078D4?style=for-the-badge&logo=visual%20studio%20code&logoColor=white) **Visual Studio Code** - Used as the IDE for developing the Restful API.
- ![XAMPP](https://img.shields.io/badge/Xampp-F37623?style=for-the-badge&logo=xampp&logoColor=white) **XAMPP** - Used for running local servers and managing databases.
- ![MySQL](https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white) **MySQL** - Used as the database management system.
- ![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=Postman&logoColor=white) **Postman** - Used for testing and documenting the API endpoints.
- ![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white) **Golang** - Used as the programming language for building the API.

## How to Run
This API can be run on your local development system using the following methods:

### Prerequisites
- Golang
- MySQL

### Environment Variables
Set up your environment variables as follows:
```bash
export PORT=<port>
export DBCONN="<username>:<password>@tcp(<hostname>:<port>)/<dbname>?charset=utf8&parseTime=True&loc=Local"
```

### Steps to Run
1. Clone the repository:
```bash
git clone https://github.com/Safmica/gdsc-core-be
cd gdsc-core-cb
```
2. Install dependencies:
```bash
go mod tidy
```
3. Run the application:
```bash
go run main.go
```
4. Use Postman to interact with the API endpoints.


## Documentation
You can access the documentation each endpoint from this API here:
> https://documenter.getpostman.com/view/36503501/2sAYBYgque
