# roomserve
Room Reservation System API written in Go using the Chi router library and Gorm ORM.

## Build
Requires installation of `Go`. Create a file called `.env` in the parent directory and in the `/test` directory. Copy the contents of `example.env` into each. Populate the database credentials. Populate the 'secret' which can be anything and is used to sign JWTs for authentication. Populate the default admin email and password.

Run `go get` to install the dependencies used in the project. These dependencies are located in the `go.mod` file.

## Database
Requires installation of PostgreSQL.
- Create a user called roomserve
- As user roomserve, create a database called roomserve_dev and roomserve_test
- Fill in the `/.env` file and `/test/.env` file with the appropriate details from there corresponding databases.

## Dev
Ensure the Postgres server is running locally. Run `go run .` to run dev server.

## Test
Ensure the Postgres server is running locally. Run `go test ./test` to run the tests.

## Prod (WIP)
Ensure production Postgres server is running. Run `go run . prod` to run prod server.
