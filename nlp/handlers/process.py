from .request import Req
from .graphql import create_nlp_report
from .clean import clean
from fastapi.exceptions import HTTPException
import threading
import os
import time
from num2words import num2words

similarities = {}


def detect_duration(sentence_vector, matcher):

    matches = matcher(sentence_vector)
    for match_id, start, end in matches:
        duration_value_token = sentence_vector[start]
        if duration_value_token.like_num:
            duration_value = int(duration_value_token.text)
        else:
            try:
                duration_value = int(num2words(duration_value_token.text, lang='fr'))
            except ValueError:
                duration_value = None

        if duration_value is not None:
            duration_unit = sentence_vector[end - 1].text.lower()
            if "jour" in duration_unit:
                return duration_value
            elif "semaine" in duration_unit:
                return duration_value * 7
            elif "mois" in duration_unit:
                return duration_value * 30
            elif "annee" in duration_unit or "an" in duration_unit:
                return duration_value * 365
    return 0

def calcul_similarity(sentence_vector, symptom):
    global similarities
    similarity = 0.0
    for word in symptom['symptom']:
        similarity = max(similarity, sentence_vector.similarity(word))

    similarities[symptom['code']] = similarity

def calcul_similarities(sentence_vector, symptoms,):
    global similarities
    similarities = {}
    threads = []


    for symptom in symptoms:
        threads.append(threading.Thread(target=calcul_similarity, args=(sentence_vector, symptom)))

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


def process(req: Req, loaded_spacy_package, symptoms, matcher) -> dict:
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
            if symptom == 'oui' or symptom == 'non':
                continue
            sentence_vector = loaded_spacy_package(clean(symptom))
            results = calcul_similarities(sentence_vector, symptoms)
            duration = detect_duration(sentence_vector, matcher)
            if is_present(context, results[0][0]) == False:
                context.append({ "symptom": results[0][0], "present": True, "days": duration})
    create_nlp_report(int(os.environ.get('VERSION')), req.symptoms, input, context, int((time.time() - start_time) * 1000))
    similarities = {}
    return {
        "context": context
    }