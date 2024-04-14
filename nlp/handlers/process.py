from .request import Req
from .graphql import create_nlp_report
from .clean import clean
from fastapi.exceptions import HTTPException
import threading
import os
import time

similarities = {}


def calcul_similarity(sentence, symptom, spacy_package):
    global similarities
    similarity = 0.0
    sentence_vector = spacy_package(sentence)
    for word in symptom['symptom']:
        similarity = max(similarity, sentence_vector.similarity(word))

    similarities[symptom['code']] = similarity

def calcul_similarities(sentence, symptoms, spacy_package):
    global similarities
    similarities = {}
    threads = []

    for symptom in symptoms:
        threads.append(threading.Thread(target=calcul_similarity, args=(sentence, symptom, spacy_package)))

    for thread in threads:
        thread.start()

    for thread in threads:
        thread.join()

    return (list(reversed(sorted(similarities.items(), key=lambda item: item[1]))))

def is_present(context, symptom):
    for i in context:
        if i["symptom"] == symptom:
            return True
    return False


def process(req: Req, loaded_spacy_package, symptoms) -> dict:
    start_time = time.time()
    input = req.input
    context: list = []
    for symptom in req.symptoms:
        if "oui" in input:
            context.append({ "symptom": symptom, "present": True })
        elif "non" in input:
            context.append({ "symptom": symptom, "present": False })
        else:
            if symptom != "":
                context.append({ "symptom": symptom, "present": None })
    if symptoms == None:
        HTTPException(500, "No symptom in database")
    for sentence in input.split("."):
        splitted = sentence.split(" et ")
        for symptom in splitted:
            results = calcul_similarities(clean(symptom), symptoms, loaded_spacy_package)
            if is_present(context, results[0][0]) == False:
                context.append({ "symptom": results[0][0], "present": True })
    create_nlp_report(int(os.environ.get('VERSION')), req.symptoms, input, context, int((time.time() - start_time) * 1000))
    similarities = {}
    return {
        "context": context
    }