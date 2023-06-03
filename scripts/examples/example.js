const fs = require('fs');

const bible_file = "../../data/kjv.json"

fs.readFile(bible_file, (err, data) => {
    if (err) throw err;
    let bible = JSON.parse(data);
    console.log(bible.books.John[3][16]);
});
