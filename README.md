# Train Ticket Booking System

## RESULT
![image](https://github.com/pranavnaikp/train_ticketing/assets/84633869/156eea73-5544-4db9-9c60-b3d2f1174430)


## Installation

### Prerequisites
Ensure you have the following installed:
- Go (version 1.15 or later)
- Protocol Buffers Compiler (protoc)

### Steps to Install

#### Server Setup
1. Navigate to the "server" folder in your terminal.
2. Run the following command to install server-side dependencies:
   ```bash
   go mod tidy
   ```

#### Client Setup
1. Navigate to the "client" folder in your terminal.
2. Run the following command to install client-side dependencies:
   ```bash
   go mod tidy

## Run Project

#### Compile proto file
1. Go to proto folder in Terminal.
2. Run the following command to compile proto file
  ```bash
  protoc --go_out=. --go-grpc_out=. ticket.proto
  ```
#### Run server
1. Go to server folder in your Terminal.
2. Run the following command.
 ```bash
 go run main.go
  ```

#### Run Client
1. Go to client folder in your Terminal.
2. Run the following command.
 ```bash
 go run main.go
  ```




