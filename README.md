# roomserve
Room Reservation 

## Database
Requires installation of PostgreSQL. Make note of your database name, username, password, and hostname / port that the server runs on.

## Build
Requires installation of `Go`. Create a file called `.env`. Copy the contents of `example.env`. Populate the database credentials. Populate the 'secret' which can be anything and is used to sign JWTs for authentication. Populate the default admin email and password.

Run `go get` to install the dependencies used in the project. These dependencies are located in the `go.mod` file.

## Dev
Ensure Postgres server is running locally. Run `go run .` to run dev server.

## Test
Ensure Postgres server is running locally. Run `go test ./test` to run the tests.

## Prod (WIP)
Ensure production Postgres server is running. Run `go run . prod` to run prod server.
