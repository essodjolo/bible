#!/usr/local/bin/python3

# pip3 install pyyaml

import os
import yaml
import json

data_directory = "../data"

for filename in os.listdir(data_directory):
    if filename.endswith('.yml'):
        file_path = os.path.join(data_directory, filename)
        with open(file_path, 'r') as file:
            bible_books = yaml.safe_load(file)
            json_file_name = filename.removesuffix('.yml')+'.json'
            with open(data_directory + '/' + json_file_name, 'w') as json_file:
                json.dump(bible_books, json_file)
