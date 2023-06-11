# Email OTP Service

This project implements an email OTP (One-Time Password) service that simulates sending OTP emails to the specified email addresses. Instead of actually sending emails, it prints the OTP details to the console.

## Assumptions

- Email addresses are limited to a maximum of 100 characters.
- The project uses an in-memory database for simplicity. Any data stored will be lost upon restarting the application.

## Prerequisites

- Go programming language (go version go1.19.4 darwin/amd64
  )

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/jainritik/email-otp.git
   ```
   
2. Navigate to the project directory:
    ```shell
    cd email-otp
   ```

3. Run the following command to download the project dependencies:
    ```shell
    go mod tidy && go mod vendor
    ```

4. Start the application by running the main.go file:

    ```shell
    go run main.go
    ```

5. The application should now be running and ready to accept requests.



