# pub-sub

Backend Microservice to handle pubsub.

## Setup

### PostgreSQL
* *Installation Site:* [Postgres.app Downloads](https://postgresapp.com/downloads.html) 

### Go lang
Follow all the installation steps as mentioned on the installation website
* *Installation Site:* [Go lang Install](https://golang.org/doc/install)

### Execute following commands
```cassandraql
go get
go build
./PubSub migrate // This command runs the migrations and creates relevant tables in DB
./PubSub run-server // This command starts REST server
./PubSub run-client // This command starts Client
```
