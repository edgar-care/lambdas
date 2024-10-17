import spacy

nlp = spacy.load("mon_modele_categorisation")


while True:
    text = input("sentence: ")
    if text.lower() == "exit":
        break

    doc = nlp(text)

    categories = doc.cats
    print("results:")

    categories = sorted(doc.cats.items(), key=lambda x: x[1], reverse=True)

    for category, score in categories[:3]:
        print(f"{category}: {score:.2f}")
