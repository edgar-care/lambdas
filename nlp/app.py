from fastapi import FastAPI, Request, HTTPException
from handlers.process import process
from handlers.request import Req
from dotenv import load_dotenv
import spacy
import os
from spacy.matcher import Matcher
from handlers.graphql import get_symptoms

app = FastAPI()

load_dotenv()

loaded_spacy_package = spacy.load('fr_core_news_lg')
matcher = Matcher(loaded_spacy_package.vocab)
computed_symptoms = []

def init():
    symptoms = get_symptoms()
    for symptom in symptoms:
        words = [loaded_spacy_package(word) for word in symptom['symptom']]
        computed_symptoms.append({**symptom, 'symptom': words})
    duration_pattern = [
        {"LIKE_NUM": True},
        {"LOWER": {"IN": ["jour", "jours", "semaine", "semaines", "mois", "annee", "annees", "an", "ans"]}}
    ]
    matcher.add("DURATION", [duration_pattern])

init()

@app.post("/nlp")
async def predictions(req: Req, request: Request):
    admin_key = request.headers.get("admin-key")
    
    if admin_key != os.environ.get("ADMIN_TOKEN"):
        raise HTTPException(status_code=401, detail="Unauthorized")
    
    return process(req, loaded_spacy_package, computed_symptoms, matcher)

@app.post("/{env}/nlp")
async def predictions_with_env(req: Req, request: Request, env: str):
    admin_key = request.headers.get("admin-key")
    
    if admin_key != os.environ.get("ADMIN_TOKEN"):
        raise HTTPException(status_code=401, detail="Unauthorized")
    
    return process(req, loaded_spacy_package, computed_symptoms, matcher)
