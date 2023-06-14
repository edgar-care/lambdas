import re
import unidecode
import string

abbreviations = [
    ["slt", "salut"]
]


def clean(sentence: str) -> str:
    cleaned_sentence = unidecode.unidecode(sentence.lower())
    punctuation = string.punctuation.replace("-", "")

    for abbreviation in abbreviations:
        cleaned_sentence = re.sub(abbreviation[0], abbreviation[1], cleaned_sentence)

    cleaned_sentence = re.sub("[" + punctuation + "]", "", cleaned_sentence)
    cleaned_sentence = re.sub(r"\n", " ", cleaned_sentence)
    cleaned_sentence = re.sub(r"\\n", " ", cleaned_sentence)

    return cleaned_sentence