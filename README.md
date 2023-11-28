# CRM Backend

## Description

This is a backend api for CRM application built using GO. It has the following api endpoints:

GET /customers - Returns all the available customers

POST /customers - Creates new customer. Example of a customer data is
`{
    "name": "Jack Smith",
    "role": "VP",
    "email": "jsmith@example.com",
    "phone": "1234567890",
    "contacted": false
}`

GET /customers/{id} - Fetches the customer that matches the request id. If not found, 404 error is returned

PUT /customers/{id} - Updates the customer that matches the request id. If not found, 404 error is returned. Example of customer data to updated is
`{
"name": "Jack Smith",
"role": "VP",
"email": "jsmith@example.com",
"phone": "1234567890",
"contacted": true
}``

DELETE /customers/{id} - Deletes the customer that matches the request id. If not found, 404 error is returned

## Running the backend server

### Clone the repo

Open terminal or command prompt or powershell in local machine and run git clone to clone the repository

### Run `go run main.go`

Change into the cloned repo and run the command to start the server. You should see a message `Server started on port 8000`
