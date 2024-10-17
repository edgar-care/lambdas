from .request import Req
from .clean import clean
from text_to_num import text2num

similarities = {}

def get_duration(text):
    print("duration text:" + text )
    if "avant-hier" in text or "avant hier" in text:
        return 2
    if "hier" in text:
        return 1

    try:
        value = text.split(" ")[0]
        if value.isdigit() != True:
            duration_value = int(text2num(value, 'fr'))
        else:
            duration_value = int(value)
    except ValueError:
        duration_value = None

    if duration_value is not None:
        if "jour" in text:
            return duration_value
        elif "semaine" in text:
            return duration_value * 7
        elif "mois" in text:
            return duration_value * 30
        elif "annee" in text or "an" in text:
            return duration_value * 365
        else:
            return duration_value
    return 0

def detect_symptom(custom_categorizer, text):
    doc = custom_categorizer(text)
    categories = doc.cats

    categories = sorted(doc.cats.items(), key=lambda x: x[1], reverse=True)
    return categories


def is_present(context, symptom):
    for i in context:
        if i["symptom"] == symptom:
            return True
    return False

def detect_entities(sentence, custom_ner):
    doc = custom_ner(sentence)
    entities = []

    for ent in doc.ents:
        entities.append({
            "text": ent.text,
            "label": ent.label_,
            "position": (ent.start_char, ent.end_char)
        })
    return entities

def get_first_duration(entities):
    first_duration = None
    for entity in entities:
        if entity['label'] == 'DURATION':
            first_duration = entity
            break
    return first_duration

def process(req: Req, custom_categorizer, custom_ner) -> dict:
    # start_time = time.time()
    input = clean(req.input)
    context: list = []

    if req.isMedicine:
        return {
            "context": context
        }

    entities = detect_entities(input, custom_ner)

    if req.isTime:
        first_duration = get_first_duration(entities)
        if first_duration is None:
            input_duration = get_duration(input)
            if input_duration == 0:
                return {"context": context, "answered": False}
            context.append({ "symptom": req.symptoms[0], "present": True, "days": input_duration})
        else:
            context.append({"symptom": req.symptoms[0], "present": True, "days": get_duration(first_duration["text"])})
            entities.remove(first_duration)

    if len(req.symptoms) > 0 and req.symptoms[0] != "" and req.isTime == False:
        if "oui" in input.lower() and "non" in input.lower():
            context.append({ "symptom": req.symptoms[0], "present": None})
        elif "oui" in input.lower():
            first_duration = get_first_duration(entities)
            if first_duration is None:
                context.append({ "symptom": req.symptoms[0], "present": True})
            else:
                context.append({ "symptom": req.symptoms[0], "present": True, "days": get_duration(first_duration["text"])})
                entities.remove(first_duration)
        else:
            context.append({ "symptom": req.symptoms[0], "present": False})

    for entity in entities:
        if entity['label'] == 'DURATION':
            continue
        print(entity)
        if entity['text'].isdigit():
            continue
        detected_symptoms = detect_symptom(custom_categorizer, entity['text'])
        print(detected_symptoms[:3])
        if detected_symptoms[0][1] < 0.80:
            continue
        first_duration = get_first_duration(entities)
        if first_duration is None:
            context.append({ "symptom": detected_symptoms[0][0], "present": True})
        else:
            context.append({ "symptom": detected_symptoms[0][0], "present": True, "days": get_duration(first_duration["text"])})
            entities.remove(first_duration)

    #create_nlp_report(int(os.environ.get('VERSION')), req.symptoms, input, context, int((time.time() - start_time) * 1000))
    return {
        "context": context
    }