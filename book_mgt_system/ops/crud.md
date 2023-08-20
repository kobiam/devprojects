# CRUD

## start the app

```bash
go run main.go
```

## create a record

`POST 127.0.0.1:9010/book/`

```json
{
  "Name":"test book",
  "Author":"test author",
  "Publication":"test publication"
}
```

## get records

`GET 127.0.0.1:9010/book/`

## update record

`PUT 127.0.0.1:9010/book/1`

```json
{
  "Name":"test book 2",
  "Author":"test author 2",
  "Publication":"test publication 2"
}
```

## delete record

`DELETE 127.0.0.1:9010/book/1`

```json
{
  "Name":"test book 2",
  "Author":"test author 2",
  "Publication":"test publication 2"
}
```