from fastapi import FastAPI
from handlers.process import process
from handlers.request import Req
from dotenv import load_dotenv
import spacy
from handlers.graphql import get_symptoms


app = FastAPI()

load_dotenv()

loaded_spacy_package =  spacy.load('fr_core_news_lg')
computed_symptoms = []

def init():
    symptoms = get_symptoms()
    for symptom in symptoms:
        words = [loaded_spacy_package(word) for word in symptom['symptom']]
        computed_symptoms.append({**symptom, 'symptom': words})

init()

@app.post("/nlp")
def predictions(req: Req):
    return process(req, loaded_spacy_package, computed_symptoms)

@app.post("/{env}/nlp")
def predictions(req: Req):
    return process(req, loaded_spacy_package, computed_symptoms)