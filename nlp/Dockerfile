FROM python:3.9

WORKDIR /code

ADD docker.mk .
ADD requirements-local.txt .

RUN apt-get update && \
    apt-get install build-essential

RUN make -f docker.mk install

ENV TF_CPP_MIN_LOG_LEVEL "2"
ENV NLTK_DATA /tmp

RUN python -m spacy download fr_core_news_lg

ENTRYPOINT ["uvicorn"]

CMD ["--host", "0.0.0.0", "--port", "5999", "app:app", "--reload"]