import os 
import requests  
import json

def get_symptoms():
    gql_url = os.environ.get('GRAPHQL_URL')
    if (gql_url == None):
        return None
    gql_query = """{
        getSymptoms {
            id
            code
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
