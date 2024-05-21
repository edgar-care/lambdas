import re
import unidecode
import string

abbreviations = [
    ["slt", "salut"],
    ["ui", "oui"],
    ["ouais", "oui"],
    ["oauis", "oui"],
    ["yes", "oui"],
    ["yep", "oui"],
    ["no", "non"],
    ["nop", "non"],
]


def clean(sentence: str) -> str:
    cleaned_sentence = unidecode.unidecode(sentence.lower()).strip()
    punctuation = string.punctuation.replace("-", "")

    for abbreviation in abbreviations:
        if abbreviation[0] == cleaned_sentence:
            cleaned_sentence = abbreviation[1]
        else:
            cleaned_sentence = re.sub(abbreviation[0] + " ", abbreviation[1], cleaned_sentence)

    cleaned_sentence = re.sub("[" + punctuation + "]", "", cleaned_sentence)
    cleaned_sentence = re.sub(r"\n", " ", cleaned_sentence)
    cleaned_sentence = re.sub(r"\\n", " ", cleaned_sentence)

    return cleaned_sentence