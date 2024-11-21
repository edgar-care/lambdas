import spacy
import unidecode

nlp_ner = spacy.load("./modele_ner")

def analyser_phrase(phrase):
    doc = nlp_ner(unidecode.unidecode(phrase))
    resultat = []

    for ent in doc.ents:
        resultat.append({
            "texte": ent.text,
            "label": ent.label_,
            "position": (ent.start_char, ent.end_char)
        })

    return resultat

# Boucle pour l'analyse de phrases en continu
while True:
    sentence = input("Entrez une phrase:")
    if sentence == "exit":
        break
    resultat = analyser_phrase(sentence)
    print(resultat)
