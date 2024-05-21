from .request import Req
from .graphql import create_nlp_report
from .clean import clean
from fastapi.exceptions import HTTPException
import threading
import os
import time
from text_to_num import text2num
import re

similarities = {}

def detect_duration(sentence_vector, matcher):
    matches = matcher(sentence_vector)
    for match_id, start, end in matches:
        duration_value_token = sentence_vector[start]
        try:
            if (duration_value_token.text).isdigit() != True:
                duration_value = int(text2num(duration_value_token.text, 'fr'))
            else:
                duration_value = int(duration_value_token.text)
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
    answers = 0

    if symptoms == None:
        HTTPException(500, "No symptom in database")
    for sentence in input.split("."):
        separators = [" et ", " mais "]
        splitted = [subphrase.strip() for subphrase in re.split("|".join(separators), sentence)]
        for symptom in splitted:
            cleaned_symptom = clean(symptom)
            sentence_vector = loaded_spacy_package(cleaned_symptom)
            if answers < len(req.symptoms) and req.symptoms[0] != "":
                if req.isTime != False:
                    if " " not in cleaned_symptom:
                        if cleaned_symptom.isdigit() == True:
                            context.append({"symptom": req.symptoms[answers], "present": True, "days": int(cleaned_symptom)})
                        else:
                            try:
                                duration = int(text2num(cleaned_symptom, 'fr'))
                                context.append({"symptom": req.symptoms[answers], "present": True, "days": duration})
                            except ValueError:
                                context.append({"symptom": req.symptoms[answers], "present": True, "days": 0})
                    else:
                        context.append({"symptom": req.symptoms[answers], "present": True, "days": detect_duration(sentence_vector, matcher)})
                    answers += 1
                    continue
                if "oui" in cleaned_symptom:
                    context.append({ "symptom": req.symptoms[answers], "present": True, "days": detect_duration(sentence_vector, matcher) })
                elif "non" in cleaned_symptom:
                    context.append({ "symptom": req.symptoms[answers], "present": False, "days": 0 })
                else:
                    if req.symptoms[answers] != "":
                        context.append({ "symptom": req.symptoms[answers], "present": None, "days": 0 })
                answers += 1
                continue
            results = calcul_similarities(sentence_vector, symptoms)
            duration = detect_duration(sentence_vector, matcher)
            if is_present(context, results[0][0]) == False:
                context.append({ "symptom": results[0][0], "present": True, "days": duration})
    create_nlp_report(int(os.environ.get('VERSION')), req.symptoms, input, context, int((time.time() - start_time) * 1000))
    similarities = {}
    return {
        "context": context
    }