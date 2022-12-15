from fastapi import FastAPI
from handlers.process import process
from handlers.request import Req

app = FastAPI()

@app.post("/nlp")
async def predictions(req: Req):
    return process(req)
