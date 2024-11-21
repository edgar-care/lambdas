import json
import spacy
from spacy.training.example import Example
from spacy.training import offsets_to_biluo_tags

with open("../generated_dataset.json", "r", encoding="utf-8") as f:
    data = json.load(f)

TRAIN_DATA = []

nlp = spacy.load("fr_core_news_md")

def align_entities_to_tokens(doc, entities):
    aligned_entities = []
    
    for start, end, label in entities:
        token_start = None
        token_end = None

        for token in doc:
            if token.idx <= start < token.idx + len(token.text):
                token_start = token
            if token.idx < end <= token.idx + len(token.text):
                token_end = token
        
        if token_start is not None and token_end is not None:
            aligned_entities.append((token_start.idx, token_end.idx + len(token_end.text), label))
        else:
            print(f"Could not align entity {label} in '{doc.text[start:end]}'; sentence: {doc.text}")
    
    return aligned_entities

for entry in data:
    text = entry["text"]
    entities = []
    
    for ent in entry["entities"]:
        start = ent["start"]
        end = ent["end"]
        label = ent["label"]
        entities.append((start, end, label))
    
    doc = nlp.make_doc(text)
    corrected_entities = align_entities_to_tokens(doc, entities)
    biluo_tags = offsets_to_biluo_tags(doc, corrected_entities)
    TRAIN_DATA.append((text, {"entities": entities}))



if "ner" not in nlp.pipe_names:
    ner = nlp.add_pipe("ner")
else:
    ner = nlp.get_pipe("ner")

for _, annotations in TRAIN_DATA:
    for ent in annotations["entities"]:
        ner.add_label(ent[2])



other_pipes = [pipe for pipe in nlp.pipe_names if pipe != "ner"]
with nlp.disable_pipes(*other_pipes):
    optimizer = nlp.begin_training()
    print("training started")
    for itn in range(16):
        losses = {}
        for text, annotations in TRAIN_DATA:
            example = Example.from_dict(nlp.make_doc(text), annotations)
            nlp.update([example], drop=0.5, losses=losses)
        print(f"Iteration {itn}, Losses: {losses}")

nlp.to_disk("modele_ner")

