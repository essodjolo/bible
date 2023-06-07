# Bible Go API

## Features

- Browse a version of the Bible using a REST API
- Ability to search
- File-based database. However you can use a NoSQL database or search engine (Eg: Elasticsearch), in which case, you have to adjust the provided code.

## Implmemented Routes

- GET /bible/:version
- GET /booklist/:version (alias GET /books)
- GET /:book/:version
- GET /versionlist (alias GET /versions)

## TODO

- GET /:book/:chapter/:version
- GET /:book/:chapter/:verse/:version
- GET /search/:version/<keyword>
- Generate OpenAPI specs

## Test the API

```bash
git clone https://github.com/essodjolo/bible.git
cd bible/go-api
go mod tidy
go run .
```
