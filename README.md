Greenlight API
===================

APIs
===================
GET    /v1/healthcheck -> Show application information
POST   /v1/movies      -> Create movie
GET    /v1/movies/:id  -> Get movie by id
PUT    /v1/movies/:id  -> Update movie
DELETE /v1/movies/:id  -> Delete movie

Running the application locally
===================
```
export GREENLIGHT_DB_DSN='postgres://greenlight:<password>@localhost/greenlight?sslmode=disable'
```

Create database commands
===================
```
$ docker exec -it pid_id /bin/bash
root@aebae6d87e96:/# psql -U postgres -h localhost
```

```
postgres=# SELECT current_user;
postgres=# CREATE DATABASE greenlight
postgres=# CREATE ROLE greenlight WITH LOGIN PASSWORD '<password>';
postgres=# CREATE EXTENSION IF NOT EXISTS citext;
```
