import re
import unidecode
import string


def lower(s: str):
    return s.lower()


def slug(s: str):
    return re.sub(r'[&/,;.:=+^$*%`<>@_\-\s]+', "-", s.lower())


abbreviations = [
    ["(^|)f\.\s*c\.", "fc"],
    ["(^|)r\.\s*c\.", "rc"],
    ["(^|)r\.\s*s\.", "as"]
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