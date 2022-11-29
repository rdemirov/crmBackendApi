# Customer CRM Backend API

## Description

This is the final project of the Go Language nanodegree. It implements backend API for customer CRM.
The project code is housed in 2 folders:

- /src/crm-api - contains the code of the project entry point and configuration. The .env file allows
the configuration of the port used by the application. It contains the following properties:

```
APP_PORT=3000
```
If no configuration is present, the default port is set to 3000.

- /src/customers - contains the struct and methods for storing and amanging customer data

## Configuration

The only configuration needed by the application is the option to set the application's port

## Starting the project

The project is started with the following commands in the project folder:

```
cd src/crm-api
go run main.go
```
## Using the project

Customer backend api can be invoked via curl or postman and supports the following methods:

- Getting a single customer through a /customers/{id} path (GET Request)
- Getting all customers through a the /customers path (GET Request)
- Creating a customer through a /customers path (POST Request)
- Updating a customer through a /customers/{id} path (PATCH Request)
- Deleting a customer through a /customers/{id} path (DELETE Request)

By default, the API URL is: 

```
http://127.0.0.1:3000
```
API postman collection can be found in the project root folder, in file "Customers CRM.postman_collection.json"