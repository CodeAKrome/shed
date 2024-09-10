# flair-news
Uses [flair](https://github.com/flairNLP/flair) with [ner-english-ontonotes-large](https://huggingface.co/flair/ner-english-ontonotes-large) to identify entities (NER) which [NewsMTSC](https://github.com/fhamborg/NewsMTSC) uses to perform targetted sentiment analysis.

Services default to port 1337, but can be configured via environment variables.

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
