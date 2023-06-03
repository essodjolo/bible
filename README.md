# BIBLE

The Bible in computer readable formats (YAML, JSON, etc.).
We worked on YAML version, but we also include a Python script to a JSON from the YAML. You can also use any other tool to conver from YAML or JSON to your desired format.

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

## Examples of how to access the content

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
import os
import yaml

data = "../../data/kjv.yml"

with open(data, 'r') as kjv_bible:
    kjv = yaml.safe_load(kjv_bible)
    print(kjv['books']['John'][3][16])
```

### NodeJS

```js
const fs = require('fs');

const bible_file = "../../data/kjv.json"

fs.readFile(bible_file, (err, data) => {
    if (err) throw err;
    let bible = JSON.parse(data);
    console.log(bible.books.John[3][16]);
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

// Bible is a nested struct that helps to load the Bible content.
type Bible struct {
  Books struct {
    // Will capture only the books I want to work with.
    Psalms map[string]map[string]string `yaml:"Psalms"`
    John   map[string]map[string]string `yaml:"John"`
    Jude   map[string]map[string]string `yaml:"Jude"`
  } `yaml:"books"`
}

func main() {
  // Path to the Bible file.
  const data = "../../data/kjv.yml"

  // Read the Bible content.
  kjv_bible, err := os.ReadFile(data)
  if err != nil {
    log.Fatal(err)
  }

  //Unmarshal the Bible YAML data into a Bible struct varibale
  var bible Bible
  err = yaml.Unmarshal(kjv_bible, &bible)
  if err != nil {
    log.Fatal(err)
  }

  // Printing Jean 3:16
  chapter := "3"
  verse := "16"
  fmt.Println("John " + chapter + ":" + verse)
  fmt.Println(bible.Books.John[chapter][verse])

  // Printing Psalms 23
  fmt.Println("\n\nPsalms 23")
  fmt.Println(bible.Books.Psalms["23"])

  // Printing Jude
  fmt.Println("\n\nJude:")
  fmt.Println(bible.Books.Jude)
}

```

## Supported languages/versions

- 🇬🇧 King James Version
- 🇫🇷 Louis Second 1910

## Contribute

You can contribute to this project in various ways including, not limited to:

- review the existing data and reported any error
- add support for another Bible version/language
- submit new ideas on how to improve this project
