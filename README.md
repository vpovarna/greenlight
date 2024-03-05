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
export GREENLIGHT_DB_DSN='postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable'
```