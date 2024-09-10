# flair-news
Uses [flair](https://github.com/flairNLP/flair) with [ner-english-ontonotes-large](https://huggingface.co/flair/ner-english-ontonotes-large) to identify entities (NER) which [NewsMTSC](https://github.com/fhamborg/NewsMTSC) uses to perform targetted sentiment analysis.

- Services default to port 1337, but can be configured via environment variables.
  
## NER tags

| Tag | Meaning | Example |
|---|---|---|
| CARDINAL | cardinal value | 1, 2, 3, ... |
| DATE | date value | 2023-12-25, January 1st |
| EVENT | event name | Super Bowl, Olympics |
| FAC | building name | Empire State Building, Eiffel Tower |
| GPE | geo-political entity | United States, France |
| LANGUAGE | language name | English, Spanish |
| LAW | law name | Constitution, Copyright Act |
| LOC | location name | New York City, Paris |
| MONEY | money name | dollar, euro |
| NORP | affiliation | Republican, Democrat |
| ORDINAL | ordinal value | first, second, third |
| ORG | organization name | NASA, Google |
| PERCENT | percent value | 50%, 75% |
| PERSON | person name | John Doe, Jane Smith |
| PRODUCT | product name | iPhone, MacBook |
| QUANTITY | quantity value | 10, 20 |
| TIME | time value | 12:00 PM, 5:30 AM |
| WORK_OF_ART | name of work of art | Mona Lisa, Star Wars |

---

## Docker image for Flair News
`flair_docker.py` is a docker wrapper around flask.
By default, this docker image will download the latest versions of Flair and NewsMTSC and install all dependencies.

### Build
```sh
make build
```

### Run
```sh
make run
```

### Check heartbeat
```sh
make heartbeat
```

### Test processing
```sh
make test
```

## Flask app
`flair_micro.py` is a simple flask application that can be used to interact with the flair-news. It will serve an API endpoint at `/ner` where you can POST JSON data containing a `text` key with the data to process and it will return the analysis in JSON format.
