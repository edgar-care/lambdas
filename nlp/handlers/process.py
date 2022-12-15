from .request import Req

symptoms = {
    "maux_de_tetes": [ "tête", "tete", "migraine" ],
    "vision_trouble": [ "vois", "yeux", "vision" ],
    "fievre": [ "froid", "température", "temperature", "fievre" ],
    "maux_de_ventre": [ "ventre", "estomac", "intestin" ],
    "vomissements": [ "vomir", "brassé", "vomi" ],
}

def process(req: Req) -> dict:
    context: list = []
    for symptom in req.symptoms:
        if "oui" in req.input.lower():
            context.append({ "symptom": symptom, "present": True })
        elif "non" in req.input.lower():
            context.append({ "symptom": symptom, "present": False })
        else:
            context.append({ "symptom": symptom, "present": None })

    for symptom in symptoms:
        for k in symptoms[symptom]:
            if k in req.input.lower():
                context.append({ "symptom": symptom, "present": True })
                break


    return {
        "context": context
    }