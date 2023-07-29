from .request import Req
from .graphql import get_symptoms
from .clean import clean
from fastapi.exceptions import HTTPException

def process(req: Req) -> dict:
    input = clean(req.input)
    context: list = []
    for symptom in req.symptoms:
        if "oui" in input:
            context.append({ "symptom": symptom, "present": True })
        elif "non" in input:
            context.append({ "symptom": symptom, "present": False })
        else:
            context.append({ "symptom": symptom, "present": None })

    symptoms = get_symptoms()
    if symptoms == None:
        HTTPException(500, "No symptom in database")
    for symptom in symptoms:
        for k in symptom["symptom"]:
            if k in input:
                context.append({ "symptom": symptom["code"], "present": True })
                break


    return {
        "context": context
    }