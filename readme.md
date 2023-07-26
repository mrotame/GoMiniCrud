# Getting started
This project is a GO Restful API with basic user management, Authorization, and Authentication system
You can run the server or run the tests by running the following commands

## Running locally
You can run this project locally or containerize

### Using Docker-Compose

##  Running
Run the command `docker-compose up -d`
Wait until it builds the image and starts the container
The app and the Postgres database will be automatically configured
After it finishes, you can test if the API is working by accessing `127.0.0.1:8001`
If you see the message `Golang API V1.0.0`, then the server is working

## Testing
If you want to test all the modules inside the container, you can:
- change the `TEST_ONLY` variable to `true` in the `.env` file
- Access the terminal inside the container by calling `docker exec -t <container-id> /bin/bash`
  And call `.\GoCrudApi --test` 

### Installing
This project was built to run on GO ver 1.20.5. It is highly recommended you use
the same version to avoid any problems and unexpected behaviors.
1. check your GO version running `go version`
2. install all the requirements by running `go mod download`

### Configuring Database
This project uses Postgresql as the main database. If you're running locally, 
You will need to configure it manually and change the `.env` file with the
correct data

### Testing
To run all the automated tests, from inside the root directory, you can:
- run the command `go run main.go --test` or `go run main.go -t`
- change the `TEST_ONLY` variable to `true` in the `.env` file
The tests status will be logged as they run

### Running
Start the web server by running `go run main.go` command or compile it with
command `go build` then execute the compiled file `GoCrudApi`
Check if the server is running by accessing `localhost:8001` or `127.0.0.1:8001`
if you see the message `Golang API V1.0.0`, then the server is working

# Endpoints
There are two main endpoints besides the index. The user's endpoint and the Auth endpoint.

## Auth
Here's where you can send the user's credentials to get an access token

* Endpoint: `localhost:8001/auth`
* Method: `POST`
* Request body: 
```JSON
    {
        "email": "<string>",
        "password": "<string>"
    }
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<null> or <string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```

## User

### POST method
Create a new user by sending the required data.
1. `Name`: User's name
2. `Age`: User's age
3. `Access level`: User's level of access (higher level can access/control users with lower access)
4. `Email`: User's login

* Endpoint: `localhost:8001/user`
* Method: `POST`
* Request body: 
```JSON
    {
        "name": "<string>",
        "age": "<integer>",
        "acceessLevel": "<integer>",
        "email": "<string>",
        "password": "<string>"
    }
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<null> or <string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```

## GET Method
Retrieve the user's information. 
 
- By Auth (The endpoint will use the received auth token to return all info for the user related to the token)
- By ID (the user must have a higher access level than the required user's info)

### By Auth
The user will receive his own data

* Endpoint: `localhost:8001/user`
* Method: `GET`
* Request Header:
```JSON
{
    "Authorization": "Bearer <token>"
}
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<null> or <string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```

### By ID
The user will receive someone else data. The user must have a higher access level to get the data

* endpoint: `localhost:8001/user/<ID>`
* method: `GET`
* Request Header:
```JSON
{
    "Authorization": "Bearer <token>"
}
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<null> or <string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```

## PUT Method
Alter user information. 

- By Auth (The endpoint will use the received auth token to identify the user and modify the info for the user related to the token)
- By ID (the user must have a higher access level than the required user to modify the data)

### By Auth
The user will modify his own data

* Endpoint: `localhost:8001/user`
* Method: `PUT`
* Request Header:
```JSON
{
    "Authorization": "Bearer <token>"
}
```

* Request body: 
```JSON
    {
        "name": "<string>",
        "age": "<integer>",
        "email": "<string>",
    }
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<null> or <string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```

### By ID
The user will modify someone else data. The user must have a higher access level than the requested user

* Endpoint: `localhost:8001/user/<ID>`
* Method: `PUT`
* Request Header:
```JSON
{
    "Authorization": "Bearer <token>"
}
```

* Request body: 
```JSON
    {
        "name": "<string>",
        "age": "<integer>",
        "email": "<string>",
    }
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<null> or <string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```


## DELETE Method
Delete user. 

- By Auth (The endpoint will use the received auth token to identify the user and delete the user related to the token)
- By ID (the user must have a higher access level than the user required to be deleted)

### By Auth
The user will delete his own account

* Endpoint: `localhost:8001/user`
* Method: `DELETE`
* Request Header:
```JSON
{
    "Authorization": "Bearer <token>"
}
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```

### By ID
The user will delete someone's else account. The user must have a higher access level than the user requested to be deleted

* Endpoint: `localhost:8001/user/<ID>`
* Method: `DELETE`
* Request Header:
```JSON
{
    "Authorization": "Bearer <token>"
}
```

* Response body: 
```JSON
    {
        "auth_token": "<string>",
        "user_data": {
            "CreatedAt": "<string>",
            "DeletedAt": "<string>",
            "ID": "<integer>",
            "UpdatedAt": "string",
            "accessLevel": "<integer>",
            "age": "<integer>",
            "email": "<string>",
            "name": "<string>",
            "password": "<null>"
        }
    }
```