#!/usr/bin/env python3
import os
import re
import json
import csv

INF_DIR = "infoblox"
OUT_DIR = "attribute_mappings"
RESOURCE_FILE_PATTERN = re.compile(r"^resource_.*\.go$")
ASSIGNMENT_PATTERN = re.compile(r'(\w+):\s*d\.Get\("([^"]+)"\)')

def extract_mappings_from_file(filepath):
    with open(filepath, "r", encoding="utf-8") as f:
        content = f.read()

    matches = re.findall(ASSIGNMENT_PATTERN, content)
    return {tf_attr: api_field for api_field, tf_attr in matches}

def save_json(resource_name, mapping):
    with open(os.path.join(OUT_DIR, f"{resource_name}.json"), "w", encoding="utf-8") as f:
        json.dump(mapping, f, indent=2)

def save_csv(resource_name, mapping):
    with open(os.path.join(OUT_DIR, f"{resource_name}.csv"), "w", encoding="utf-8", newline="") as f:
        writer = csv.writer(f)
        writer.writerow(["Terraform Attribute", "API Field"])
        for tf_attr, api_field in mapping.items():
            writer.writerow([tf_attr, api_field])

def main():
    os.makedirs(OUT_DIR, exist_ok=True)
    result = {}
    for filename in os.listdir(INF_DIR):
        if RESOURCE_FILE_PATTERN.match(filename) and not filename.endswith("_test.go"):
            filepath = os.path.join(INF_DIR, filename)
            resource_name = filename.replace(".go", "")
            mapping = extract_mappings_from_file(filepath)
            if mapping:
                result[resource_name] = mapping
                save_json(resource_name, mapping)
                save_csv(resource_name, mapping)

    print(json.dumps(result, indent=2))

if __name__ == "__main__":
    main()
