from fastapi import FastAPI, Request, HTTPException
from handlers.process import process
from handlers.request import Req
from dotenv import load_dotenv
import spacy
import os
import time
app = FastAPI()

load_dotenv()

start_loading_spacy = time.time()
custom_ner = spacy.load("./custom_ner_model")
custom_categorizer = spacy.load('./custom_categorizer_model')
end_loading_spacy = time.time()

@app.post("/nlp")
async def predictions(req: Req, request: Request):
    admin_key = request.headers.get("admin-key")
    print("adminkey: ", admin_key)

    if admin_key != os.environ.get("ADMIN_TOKEN"):
        raise HTTPException(status_code=401, detail="Unauthorized")

    return process(req, custom_categorizer, custom_ner)

@app.post("/{env}/nlp")
async def predictions_with_env(req: Req, request: Request, env: str):
    admin_key = request.headers.get("admin-key")
    if admin_key != os.environ.get("ADMIN_TOKEN"):
        raise HTTPException(status_code=401, detail="Unauthorized")

    return process(req, custom_categorizer, custom_ner)

@app.get("/nlp/status")
async def status():
    return {
        "status": "running",
    }

@app.get("/{env}/nlp/status")
async def status():
    return {
        "status": "running",
    }