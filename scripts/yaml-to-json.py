#!/usr/local/bin/python3

# pip3 install pyyaml

import os
import yaml
import json

data_directory = "../data"

for filename in os.listdir(data_directory):
    # If you want to convert one single file among many,
    # you can simply change `.yml` with the full file name.
    if filename.endswith('.yml'):
        file_path = os.path.join(data_directory, filename)
        with open(file_path, 'r', encoding='utf8') as file:
            bible_books = yaml.safe_load(file)
            #print(bible_books['books']['Jude'][1][2])
            json_file_name = filename.removesuffix('.yml')+'.json'
            with open(data_directory + '/' + json_file_name, 'w', encoding='utf8') as json_file:
                json.dump(bible_books, json_file, indent=2)
