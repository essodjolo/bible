# BIBLE

The Bible in computer readable formats (YAML, JSON, etc.).
We worked on YAML version, but we also include a Python script to a JSON from the YAML. You can also use any other tool to conver from YAML or JSON to your desired format.

## Structure

The structure is organized in a way that you can access a verse using the path `book.chapter.verse`.

## Examples

### Yq

```bash
yq .books.John.3.16 data/kjv.yml

yq .books.Psalms.23 data/kjv.yml

yq .books.Jude data/kjv.yml
```

### Jq

```bash
jq '.[].John."3"."16"' data/kjv.json

jq '.[].Psalms."23"' data/kjv.json

jq '.[].Jude' data/kjv.json
```

### Python

```python
# TODO
```

### Go

```go
// TODO
```

## Supported language/versions

- 🇬🇧 King James Version
- 🇫🇷 Louis Second 1910

## Contribute

// TODO
