from spacy.util import minibatch, compounding
from spacy.training.example import Example
import spacy
import json
import random

with open("symptoms_training_data.json", 'r', encoding='utf-8') as f:
    data = json.load(f)

def create_training_data(data, nlp):
    training_data = []
    for item in data:
        if "text" in item and "cats" in item:
            doc = nlp.make_doc(item["text"])
            example = Example.from_dict(doc, {"cats": item["cats"]})
            training_data.append(example)
    return training_data

nlp = spacy.blank("fr")

if "textcat" not in nlp.pipe_names:
    textcat = nlp.add_pipe("textcat")
else:
    textcat = nlp.get_pipe("textcat")

for label in data[0]["cats"].keys():
    textcat.add_label(label)

training_data = create_training_data(data, nlp)

n_iter = 20

optimizer = nlp.begin_training()

for i in range(n_iter):
    random.shuffle(training_data)
    losses = {}
    batches = minibatch(training_data, size=compounding(4.0, 32.0, 1.001))
    for batch in batches:
        nlp.update(batch, sgd=optimizer, losses=losses)
    print(f"Iteration {i + 1} - Losses: {losses}")

nlp.to_disk("mon_modele_categorisation")
