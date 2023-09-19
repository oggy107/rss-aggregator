<h1 align="center"> rss-aggregator</h1>

## API documentation

click [here](https://documenter.getpostman.com/view/19350060/2s9YC8wrDS) to read API documentation

## Getting started

- Spin up a postgres server
- Clone the repo and create `.env` file in the root
- Here is the example `.env`. Edit the file to match your postgres connection string and desired port for the server

    ```txt
    PORT=3000
    POSTGRES_CONNECTION=postgres://postgres:root@localhost:5432/rssagg?sslmode=disable
    ```

- Build and run the server

    ```bash
    go build && ./rss-aggregator.exe
    ```
