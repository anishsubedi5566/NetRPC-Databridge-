# RPC CRUD Operations for Database Manipulation

## Overview
This project implements CRUD (Create, Read, Update, Delete) operations using Remote Procedure Calls (RPC) for manipulating a database. It provides a means of remotely accessing and modifying data in the database via RPC calls over a network.

## Features
- **CRUD Operations:** Implementations for Create, Read, Update, and Delete operations on the database.
- **RPC Calls:** Utilizes Remote Procedure Calls for communication between client and server.
- **HTTP Handling:** API methods are registered and configured for handling HTTP requests.
- **Network Communication:** The server serves RPC functionality on a designated port, enabling network-based communication for database operations.

## Setup and Usage
1. **Installation:**
   - Clone the repository to your local machine.
   - Ensure you have Go installed on your system.

2. **Configuration:**
   - Configure database connection settings in the `config.go` file if needed.
   - Specify the desired port for RPC server in the `main.go` file.

3. **Starting the Server:**
   - Navigate to the project directory.
   - Run `go run main.go` to start the RPC server.

4. **Using the API:**
   - Access the CRUD operations via RPC calls from your client application.
   - Refer to the code comments in `main.go` for details on available methods and parameters.

