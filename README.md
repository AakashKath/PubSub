# api-revenue-manager

Backend Microservice to handle pricing, ledger management, invoicing, payments, etc.

## Setup

### PostgreSQL
* *Installation Site:* [Postgres.app Downloads](https://postgresapp.com/downloads.html)
* *Create a schema with name "revenue" in DB:*
```
CREATE SCHEMA revenue;
GRANT ALL ON ALL TABLES IN SCHEMA revenue TO <user_name>;
GRANT ALL ON SCHEMA revenue TO <user_name>;
``` 

### Go lang
Follow all the installation steps as mentioned on the installation website
* *Installation Site:* [Go lang Install](https://golang.org/doc/install)

### Execute following commands
```cassandraql
go get
go build
./api-revenue-manager migrate // This command runs the migrations and creates relevant tables in DB
./api-revenue-manager run-server // This command starts REST server
```
