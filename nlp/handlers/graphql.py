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

def create_nlp_report(version, input_symptoms, input_sentence, output, computation_time):
    gql_url = os.environ.get('GRAPHQL_URL')
    if gql_url is None:
        return None
    
    query = '''
    mutation CreateNlpReport($input: CreateNlpReportInput!) {
        createNlpReport(input: $input) {
            id
            version
            input_symptoms
            input_sentence
            output {
                symptom
                present
                days
            }
            computation_time
            createdAt
            updatedAt
        }
    }
    '''

    variables = {
        'input': {
            'version': version,
            'input_symptoms': input_symptoms,
            'input_sentence': input_sentence,
            'output': output,
            'computation_time': computation_time
        }
    }

    headers = {
        'Content-Type': "application/json",
        os.environ.get('API_KEY'): os.environ.get('API_KEY_VALUE')
    }

    data = {
        'query': query,
        'variables': variables
    }

    response = requests.post(gql_url, headers=headers, json=data)
    response.raise_for_status()

    return response.json()