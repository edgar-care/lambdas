import random
import json
import unidecode

durations = [
    "hier",
    "avant hier",
    "la semaine dernière",
    "une dizaine de jours",
    "une douzaine de jours",
    "x jours",
    "x jours",
    "x jours",
    "x semaines",
    "x mois",
    "x ans",
]

sentences_templates = [
    "j'ai {symptom1}",
    "j'ai {symptom1} depuis {duration}",
    "j'ai {symptom1} et aussi j'ai {symptom2}",
    "j'ai {symptom1} depuis {duration} et aussi j'ai {symptom2}",
    "Je ressens {symptom1}.",
    "Je souffre de {symptom1} depuis {duration}.",
    "J'ai remarqué {symptom1} qui persiste.",
    "Je fais face à {symptom1} gênant.",
    "Je me plains de {symptom1} qui ne part pas.",
    "Je lutte contre {symptom1} douloureux depuis {duration}.",
    "{symptom1} m'accompagne depuis {duration}.",
    "Je vis avec {symptom1} intense.",
    "Je ressens souvent {symptom1}.",
    "{symptom1} léger a commencé hier, mais il s'intensifie.",
    "Depuis {duration}, je ressens {symptom1} persistant.",
    "Ca fait {duration} que je ressens {symptom1} persistant.",
    "{symptom1} a débuté la nuit dernière.",
    "Cela fait maintenant {duration} que je souffre de {symptom1}.",
    "{symptom1} a fait son apparition il y a trois jours.",
    "Je n'arrive pas à me débarrasser de {symptom1}.",
    "{symptom1} est apparu après une journée stressante depuis {duration}.",
    "Le {symptom1} que je ressens est de plus en plus fréquent.",
    "Je me sens épuisé à cause de {symptom1}.",
    "{symptom1} est très dérangeant.",
    "Depuis {duration}, je ressens {symptom1} qui ne part pas.",
    "Depuis {duration} aussi, je ressens {symptom1} qui ne part pas.",
    "{symptom1} discret s'est intensifié ces derniers jours.",
    "Je souffre {symptom1} et {symptom2} depuis {duration}.",
    "Depuis {duration}, je ressens des {symptom1} et des {symptom2}.",
    "Je me bats contre {symptom1} et {symptom2}.",
    "{symptom1} et {symptom2} durent depuis {duration}.",
    "Cela fait maintenant trois jours que je ressens des {symptom1} et {symptom2}.",
    "{symptom1} et {symptom2} ne cessent d'empirer.",
    "Je ressens {symptom1} et {symptom2} et je m'inquiète de leur intensité.",
    "Après avoir souffert de {symptom1}, je ressens maintenant aussi {symptom2}.",
    "Je suis préoccupé par mes {symptom1} qui s'accompagnent de {symptom2}.",
    "Depuis {duration}, j'ai des {symptom1} et des {symptom2} qui me dérangent.",
    "Cela fait deux semaines que je souffre de {symptom1}, et maintenant j'ai aussi {symptom2}.",
    "{symptom1} persistent, et maintenant les {symptom2} sont apparus.",
    "Hier, j'avais seulement {symptom1}, mais aujourd'hui j'ai aussi {symptom2} et {symptom3} depuis {duration}.",
    "{duration}",
    "{duration} aussi",
    "Ca va faire {duration}",
    "Ca fait {duration}",
    "Ca fait depuis {duration}",
    "{duration} et je souffre de {symptom1} depuis {duration} aussi"
]

with open("symptomes.json", "r", encoding="utf-8") as f:
    symptoms = json.load(f)

def choose_symptom_variant(symptom_key):
    return unidecode.unidecode(random.choice(symptoms[symptom_key]))

def adjust_articles(sentence):
    sentence = sentence.replace(" de une ", " d'une ")
    sentence = sentence.replace(" de un ", " d'un ")
    sentence = sentence.replace(" de des ", " de ")
    return sentence

def generate_sentence():
    template = random.choice(sentences_templates)
    i = 1
    text_entities = []
    entities = []
    old_symptom = ""

    while "{symptom" + str(i) + "}" in template:
        symptom_key = random.choice(list(symptoms.keys()))
        if old_symptom == symptom_key:
            continue
        old_symptom = symptom_key
        symptom_phrase = choose_symptom_variant(symptom_key)
        template = template.replace("{symptom" + str(i) + "}", symptom_phrase)
        text_entities.append({"label": "SYMPTOM", "text": symptom_phrase, "key": symptom_key})
        i += 1

    if "{duration}" in template:
        duration = random.choice(durations)
        if "x" in duration:
            duration = duration.replace("x", random.choice(["1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix", "onze", "douze", "treize", "quatorze", "quinze"]))
        template = template.replace("{duration}", duration)
        text_entities.append({"label": "DURATION", "text": duration})

    for entity in text_entities:
        text = entity["text"]
        if text.startswith("des") or text.startswith("une"):
            text = text[3:]
        if text.startswith("un") and "ans" not in text:
            text = text[2:]
        if text.startswith(" "):
            text = text[1:]

        template = adjust_articles(template)
        start = template.index(text)
        end = template.index(text) + len(text)
        
        entities.append({"start": start, "end": end, **entity, "text": template[start:end]})

    return {
        "text": template,
        "entities": entities
    }

def generate_dataset(size):
    dataset = []
    for _ in range(size):
        sentence = generate_sentence()
        dataset.append(sentence)
    return dataset

def save_dataset(filename, dataset):
    with open(filename, 'w') as f:
        json.dump(dataset, f, indent=2, ensure_ascii=False)

dataset = generate_dataset(5000)
save_dataset('generated_dataset.json', dataset)
print("Dataset généré et sauvegardé sous 'generated_dataset_varied_large.json'")
