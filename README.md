# BIBLE (YAML, JSON)

The Bible in computer readable formats (YAML, JSON, API, etc.).

We basically worked on YAML version, but we also provide a Python [script](scripts/yaml-to-json.py) to convert YAML versions to JSON.
You can also use any other tool to conver from YAML or JSON to your desired format.
Finally, you can use the provided **API** to explore the Bible (see [The Go REST API](#the-go-rest-api) section).

## Structure

The structure is organized in a way that you can access a verse using the path `book.chapter.verse`.

### Excerpts

#### YAML

```yml
books:
  Jude:
    1:
      1: "Jude, the servant of Jesus Christ, and brother of James, to them that are sanctified by God the Father, and preserved in Jesus Christ, and called:"
      2: "Mercy unto you, and peace, and love, be multiplied."
      3: "Beloved, when I gave all diligence to write unto you of the common salvation, it was needful for me to write unto you, and exhort you that ye should earnestly contend for the faith which was once delivered unto the saints."
```

#### JSON

```json
{
  "books": {
    "Jude": {
      "1": {
        "1": "Jude, the servant of Jesus Christ, and brother of James, to them that are sanctified by God the Father, and preserved in Jesus Christ, and called:",
        "2": "Mercy unto you, and peace, and love, be multiplied.",
        "3": "Beloved, when I gave all diligence to write unto you of the common salvation, it was needful for me to write unto you, and exhort you that ye should earnestly contend for the faith which was once delivered unto the saints."
      }
    }
  }
}
```

## How to access the content (API, Go, Python, NodeJS, etc.)

You can access the Bible content using **a tool** or **a programming language** of your choice (see exameples below).

In addition, you can also use the provided **API** to explore the Bible (see [The Go REST API](#the-go-rest-api) section).

### Yq

```bash
yq .John.3.16 data/kjv.yml

yq .Psalms.23 data/kjv.yml

yq .Jude data/kjv.yml
```

### Jq

```bash
jq '.John."3"."16"' data/kjv.json

jq '.Psalms."23"' data/kjv.json

jq '.Jude' data/kjv.json
```

### Python

```python
import os
import yaml

data = "../../data/kjv.yml"

with open(data, 'r', encoding='utf8') as kjv_bible:
    kjv = yaml.safe_load(kjv_bible)
    print(kjv['John'][3][16])
```

### NodeJS

```js
const fs = require('fs');

const bible_file = "../../data/kjv.json"

fs.readFile(bible_file, (err, data) => {
    if (err) throw err;
    let bible = JSON.parse(data);
    console.log(bible.John[3][16]);
});
```

### Go

```go
package main

import (
    "fmt"
    "log"
    "os"

    "gopkg.in/yaml.v3"
)

type Bible map[string]map[string]map[string]string

func main() {
    // Path to the Bible file.
    data_folder := "../../data/"

    // Path to the Bible file.
    bible_file_path := data_folder + "kjv.yml"

    // Read the Bible content.
    bible_content, err := os.ReadFile(bible_file_path)
    if err != nil {
        log.Fatal(err)
    }

    //Unmarshal the Bible YAML data
    var bible Bible
    err = yaml.Unmarshal(bible_content, &bible)
    if err != nil {
        log.Fatal(err)
    }

    // Printing Jean 3:16
    chapter := "3"
    verse := "16"
    fmt.Println("John " + chapter + ":" + verse)
    fmt.Println(bible["John"][chapter][verse])

    // Printing Psalms 23
    fmt.Println("\n\nPsalms 23")
    fmt.Println(bible["Psalms"]["23"])

    // Printing Jude
    fmt.Println("\n\nJude:")
    fmt.Println(bible["Jude"])
}

```

## The Go REST API

🚧 See [go-api](go-api/).

## Supported languages/versions

- 🇬🇧 King James Version
- 🇫🇷 Louis Second 1910

## Contribute

You can contribute to this project in various ways including, not limited to:

- review the existing data and reported any error
- add support for another Bible version/language
- submit new ideas on how to improve this project
