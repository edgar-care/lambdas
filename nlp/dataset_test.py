#!/usr/bin/env python3

import sys
import csv
from handlers import process
from handlers import request
from dotenv import load_dotenv

load_dotenv()

def helper():
    print("USAGE:\n")
    print("\t./dataset_test dataset_path\n")
    print("\dataset_path: path to the dataset")

def resp_to_string(resp):
    is_first = True
    result = ""
    for symptom in resp["context"]:
        if symptom["present"] == True:
            if is_first != True:
                result += " "
            result += symptom["symptom"]
    return result

def main():
    is_first = True
    if (len(sys.argv) != 2):
        helper()
        exit(0)
    new_lines = []
    with open(sys.argv[1]) as csv_file:
        file = csv.reader(csv_file, delimiter=',')
        for line in file:
            if is_first == True:
                is_first = False
                continue
            req = request.Req(
                symptoms=[], input=line[0])
            resp = process.process(req)
            new_lines.append(line + [resp_to_string(resp)])
            print(line + [resp_to_string(resp)])
    with open('result_dataset.csv', mode='w') as output_file:
        output_writer = csv.writer(output_file, delimiter=',', quotechar='"', quoting=csv.QUOTE_MINIMAL)
        for line in new_lines:
            output_writer.writerow(line)


main()