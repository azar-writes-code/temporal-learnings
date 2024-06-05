## Temporal worker poc

This is a temporal worker poc.

## How to use

- Install temporal server from [temporal.io](https://github.com/temporalio/temporal)
- Run the below command to start the server
  ```bash
  temporal server start-dev
  ```

- Run the below command to test the workflow
  ```bash
    go run starter/main.go
  ```

- check the workflow on `http://localhost:8233/namespaces/default/workflows`