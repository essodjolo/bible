# Bible Go API

## Features

- Browse a version of the Bible using a REST API
- Ability to search
- File-based database. However you can use a NoSQL database or search engine (Eg: Elasticsearch), in which case, you have to adjust the provided code.

## Synopsis

Default version is KJV.

If another version needed, specify it in the body.

When `lookup` is specified:

- we search for the provided content.
- we return a slice of verses where the content if found
- we return only the first 100 matching verses
- if the user is retrieving one single verse, return the whole verse without any search

## Routes

### `GET /versions`

- Retrieve a list of available versions
- E.g:

```json
[
  {
    "code" : "kjv",
    "name" : "King James Version",
    "language" : "english"
  },
  {
    "code" : "lsg",
    "name" : "Louis Segond 1910",
    "language" : "french"
  }
]
```

### `GET /booklist`

- Retrieve a version's book names
- Parameter `version` is optional. When not set, `kjv` is used.
- E.g (for `version=lsg`): `["Genèse", "Exode", ... , "Apocalypse"]`

Examples:

```js
GET /books?version=lsg
```

### `GET /passage`

- Retrieve the content of a book, chapter, or verse.
- Parameter `version` is optional. When not set, `kjv` is used.

Examples:

```js
GET /passage?version=kjv&book=Jude

GET /passage?version=kjv&book=Psalms&chapter=23

GET /passage?version=kjv&book=John&chapter=3&verse=16
```

### `GET /search`

- Search a keyword in the whole Bible, or in a book, a chapter, or a verse.
- Parameter `version` is optional. When not set, `kjv` is used.

Examples:

```js
GET /search?keyword=For God so loved the world

GET /search?keyword=For God so loved the world&book=John

GET /search?keyword=For God so loved the world&book=John&chapter=3

GET /search?keyword=For God so loved the world&book=John&chapter=3&verse=16
```

## Response data types

```go
type Verse struct {
  // Ref holds the reference for the verse. Eg: "John 3:16".
  Ref string
  Content string
}

type Chapter struct {
  // Ref holds the reference for the chapter. Eg: "John 3".
  Ref string
  Content map[string]string
}

type Book struct {
  Name string
  Content map[string]map[string]string
}

type Bible struct {
  Books map[string]map[string]map[string]string
}

var booklist []string

var versions []string

type SearchReslut struct {
  Result []Verse
}
```
