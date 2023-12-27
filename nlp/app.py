from fastapi import FastAPI
from handlers.process import process
from handlers.request import Req
from dotenv import load_dotenv

app = FastAPI()

load_dotenv()

@app.post("/nlp")
async def predictions(req: Req):
    return process(req)