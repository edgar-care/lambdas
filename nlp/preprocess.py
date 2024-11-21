import spacy
import pickle
import boto3
from spacy.matcher import Matcher
from handlers.graphql import get_symptoms
from dotenv import load_dotenv

load_dotenv()

def preprocess_and_upload():
    nlp = spacy.load('fr_core_news_lg')
    matcher = Matcher(nlp.vocab)
    computed_symptoms = []

    symptoms = get_symptoms()
    for symptom in symptoms:
        words = [nlp(word) for word in symptom['symptom']]
        computed_symptoms.append({**symptom, 'symptom': words})

    duration_pattern = [
        {"LIKE_NUM": True},
        {"LOWER": {"IN": ["jour", "jours", "semaine", "semaines", "mois", "annee", "annees", "an", "ans"]}}
    ]
    matcher.add("DURATION", [duration_pattern])

    preprocessed_data = {
        'matcher': matcher,
        'computed_symptoms': computed_symptoms
    }

    with open('preprocessed_data.pkl', 'wb') as f:
        pickle.dump(preprocessed_data, f)

    s3 = boto3.client('s3')
    s3.upload_file('preprocessed_data.pkl', 'edgar-nlp', 'preprocessed/data.pkl')

if __name__ == "__main__":
    preprocess_and_upload()