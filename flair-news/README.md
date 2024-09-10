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

## Example output
```
The 1st thing to do on Jan 1 at 5am is attend the Viking Funeral of Darth Vader in the School Gym. 1000 Rastafarians will chant 50% of the description of The Thinker for 10 dollars.
```

```json
[
  {
    "score": "0.96",
    "sentence": "The 1st thing to do on Jan 1 at 5am is attend the Viking Funeral of Darth Vader in the School Gym.",
    "spans": [
      {
        "end": 7,
        "probability": "0.94",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 4,
        "text": "1st",
        "value": "ORDINAL"
      },
      {
        "end": 28,
        "probability": "0.96",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 23,
        "text": "Jan 1",
        "value": "DATE"
      },
      {
        "end": 35,
        "probability": "0.96",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 32,
        "text": "5am",
        "value": "TIME"
      },
      {
        "end": 79,
        "probability": "0.89",
        "score": "0.98",
        "sentiment": "neutral",
        "start": 46,
        "text": "the Viking Funeral of Darth Vader",
        "value": "EVENT"
      },
      {
        "end": 97,
        "probability": "0.96",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 83,
        "text": "the School Gym",
        "value": "FAC"
      }
    ],
    "tag": "negative"
  },
  {
    "score": "1.00",
    "sentence": "1000 Rastafarians will chant 50% of the description of The Thinker for 10 dollars.",
    "spans": [
      {
        "end": 4,
        "probability": "0.73",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 0,
        "text": "1000",
        "value": "CARDINAL"
      },
      {
        "end": 17,
        "probability": "0.63",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 5,
        "text": "Rastafarians",
        "value": "NORP"
      },
      {
        "end": 32,
        "probability": "0.69",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 29,
        "text": "50%",
        "value": "PERCENT"
      },
      {
        "end": 66,
        "probability": "0.89",
        "score": "0.85",
        "sentiment": "neutral",
        "start": 55,
        "text": "The Thinker",
        "value": "WORK_OF_ART"
      },
      {
        "end": 81,
        "probability": "0.79",
        "score": "1.00",
        "sentiment": "neutral",
        "start": 71,
        "text": "10 dollars",
        "value": "MONEY"
      }
    ],
    "tag": "negative"
  }
]
```

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
