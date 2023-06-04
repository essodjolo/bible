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
- E.g: `["KJV", "LSG"]`

### `GET /books`

- Retrieve a version's book names
- E.g (for `LSG`): `["Genèse", "Exode", ... , "Apocalypse"]`
- `"version": "<version_name>"` is mandatory for this route

Examples:

```js
GET /books
{
  "version": "lsg"
}
```

### `GET /{book}`

- Retrieve the content of a book
- Search throuh a book

Examples:

```js
GET /John
```

```js
GET /Psalms
{
  "version": "kjv";
  "lookup": "The LORD is my shepherd"
}
```

### `GET /{book}/{chapter}`

- Retrieve the content of a chapter
- Search throuh a chapter

Examples:

```js
GET /Matthew/5
{
  "version": "kjv"
}
```

```js
GET /Psalms/23
```

### `GET /{book}/{chapter}/{verse}`

- Retrieve the content of a verse
- Search throuh a verse

Examples:

```js
GET /John
```

## Request Body type

```go
type Body struct {
  // Default version is KJV.
  Version string

  // Default behaviour is to return the whole passage,
  // no lookup for a keyword or sentence.
  Lookup string
}
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

var book []string

var versions []string
```
