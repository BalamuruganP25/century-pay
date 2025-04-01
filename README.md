# Century-Pay Bank API

A simple RESTful API for basic banking operations like transferring money, checking balances, and viewing transaction history, implemented in Go.

## Features

-   **Transfer Money**: Transfer money between users.
-   **Get User Balance**: Retrieve a user's balance.
-   **Get Transaction History**: Retrieve a user's transaction history.

## Endpoints

### 1.  **Transfer Money**

-   **POST** `/v1/transaction/transfer_money`
-   **Description**: Transfers money from one user to another.
-   **Request Body**:

    ```json
    {
      "sender": "sender_name",
      "receiver": "receiver_name",
      "amount": 100.0
    }
    ```

-   **Response**:
    -   `200 OK` if successful.
    -   `400 Bad Request` if invalid data or insufficient funds.

### 2.  **Get User Balance**

-   **GET** `/v1/transaction/{user}/balance`
-   **Description**: Fetches the balance of a specific user.
-   **Request Parameters**:
    -   `user`: The username of the user whose balance you want to check.
-   **Response**:

    ```json
    {
      "name": "user_name",
      "balance": 150.0
    }
    ```

-   **Response Codes**:
    -   `200 OK` if the user exists and the balance is retrieved.
    -   `404 Not Found` if the user does not exist.

### 3.  **Get Transaction History**

-   **GET** `/v1/transaction/{user}/transaction_history`
-   **Description**: Retrieves a list of all transactions for a specific user.
-   **Request Parameters**:
    -   `user`: The username of the user whose transaction history you want to retrieve.
-   **Response**:

    ```json
    [
      {
        "sender": "user1",
        "receiver": "user2",
        "amount": 50.0,
        "timestamp": "2025-04-01T12:00:00Z"
      }
    ]
    ```

-   **Response Codes**:
    -   `200 OK` if the user exists and the transaction history is retrieved.
    -   `404 Not Found` if the user does not exist.

## Setup and Run

## Using the Makefile

The `Makefile` helps automate Docker operations and Go dependency management. You can use the following commands:

1.  **Build the Docker image:**

    ```bash
    make build
    ```

2.  **Run the service (builds and runs the Docker container):**

    ```bash
    make run
    ```

3.  **Stop and remove the running container:**

    ```bash
    make stop
    ```

4.  **Manage Go dependencies:**

    ```bash
    make dep
    ```

## Makefile Breakdown

-   `run`: Builds and runs the Docker container, exposing the service on port 8089.
-   `build`: Builds the Docker image for the service.
-   `stop`: Stops and removes the Docker container.
-   `dep`: Manages Go dependencies by running `go mod tidy` and `go mod vendor`.

