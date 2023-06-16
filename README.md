# Fitness workout tracking application

## Initial Setup
- Make sure Go is installed
- Make sure Docker is installed
- Make sure the Go Migrate library is installed. To install via homebrew, run `brew install golang-migrate`

## DB Setup
Use Docker to create and start a postgres DB using `docker compose up -d` command

Run migrations to initialize the DB with command `migrate -path db/migration -database "postgres://postgres:postgres@localhost:5433/fitness_app" -verbose up`

TODO Items
- 