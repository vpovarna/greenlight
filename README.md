Greenlight API
====================================================

APIs
====================================================
GET    /v1/healthcheck           -> Show application information
GET    /v1/movies                -> Show details of all movies
POST   /v1/movies                -> Create movie
GET    /v1/movies/:id            -> Get movie by id
PATH   /v1/movies/:id            -> Update movie
DELETE /v1/movies/:id            -> Delete movie
POST   /v1/users                 -> Register a new user
PUT    /v1/users/activated       -> Activate a specific user
POST   /v1/tokens/authentication -> Generate a new authentication token

Running the application locally
====================================================

```
$ exprt GREENLIGHT_DB_DSN='postgres://greenlight:<password>@localhost/greenlight?sslmode=disable'
```

Create database commands
====================================================
```
$ docker exec -it pid_id /bin/bash
root@aebae6d87e96:/# psql -U postgres -h localhost
```

```
postgres=# SELECT current_user;
postgres=# CREATE DATABASE greenlight;
postgres=# CREATE ROLE greenlight WITH LOGIN PASSWORD '<password>';
postgres=# CREATE EXTENSION IF NOT EXISTS citext;
postgres=# grant all privileges on database greenlight to greenlight;
postgres=# alter database greenlight owner to greenlight;
```

Migrations
====================================================
Install golang migrate
```
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Run migrations
```
$ export GREENLIGHT_DB_DSN='postgres://greenlight:<password>@localhost/greenlight?sslmode=disable'
$ migrate -path=./migrations -database=$GREENLIGHT_DB_DSN up
```

Create migration files
```
$ migrate create -seq -ext .sql -dir ./migrations add_movies_indexes
```


Testing commands
====================================================
Testing optimistic locking
```
$ xargs -I % -P8 curl -X PATCH -d '{"runtime": "97 mins"}' "localhost:4000/v1/movies/1" < <(printf '%s\n' {1..8})
```


Measure the response time
```
curl -w '\nTime: %{time_total}s \n' localhost:4000/v1/movies/1
```
