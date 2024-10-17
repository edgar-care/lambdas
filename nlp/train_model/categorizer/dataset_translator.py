import os
from dotenv import load_dotenv
import json
import requests

load_dotenv()

def get_symptoms():
    gql_url = os.environ.get('GRAPHQL_URL')
    if (gql_url == None):
        return None
    gql_query = """{
        getSymptoms {
            id
            name
            symptom
            question
            advice
        }
    }"""

    data = {"query" : gql_query}
    json_data = json.dumps(data)
    header = {'Content-Type': "application/json", os.environ.get('API_KEY'): os.environ.get('API_KEY_VALUE')}

    response = requests.post(url=gql_url, headers=header, data=json_data)
    response.raise_for_status()

    return response.json()['data']['getSymptoms']

symptoms = get_symptoms()
name_list = [item['name'] for item in symptoms]


with open("../generated_dataset.json", 'r', encoding='utf-8') as f:
    general_dataset = json.load(f)


dataset = []

for item in general_dataset:
    for entity in item['entities']:
        if entity['label'] == "SYMPTOM":
            labels = {label: 0.0 for label in name_list}
            labels[entity['key']] = 1.0
            dataset.append({"text": entity['text'], "cats": labels})
#
#
#
# for symptom in symptoms:
#     labels = {label: 0.0 for label in name_list}
#     labels[symptom['name']] = 1.0
#
#     print(symptom['name'], ": ", symptom['symptom'])
#
#     for word in symptom['symptom']:
#         dataset.append({"text": word, "cats": labels})

with open("symptoms_training_data.json", "w", encoding="utf-8") as f:
    json.dump(dataset, f, indent=2, ensure_ascii=False)

print("Successfully generated dataset")