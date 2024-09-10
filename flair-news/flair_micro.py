import sys
from json import dumps, loads
from flask import Flask, request, jsonify
import argparse
from flair.nn import Classifier
from flair.splitter import SegtokSentenceSplitter
from NewsSentiment import TargetSentimentClassifier

# python flair_sentiment_service.py 5000
# curl http://localhost:5000/heartbeat
# curl -X POST -H "Content-Type: application/json" -d '{"text":"Apple Inc. is planning to open a new store in New York City."}' http://localhost:5000/ner

class FlairSentiment:
    NER_TAGGER = "flair/ner-english-ontonotes-large"

    def __init__(self):
        self.sentiment_tagger = Classifier.load("sentiment")
        self.ner_tagger = Classifier.load(self.NER_TAGGER)
        self.splitter = SegtokSentenceSplitter()
        self.tsc = TargetSentimentClassifier()

    def process_text(self, text: str) -> list:
        sentences = self.splitter.split(text)
        self.sentiment_tagger.predict(sentences)
        self.ner_tagger.predict(sentences)

        output = []
        for sentence in sentences:
            if sentence:
                try:
                    spans = []
                    sent = sentence.to_plain_string()
                    for span in sentence.get_spans("ner"):
                        l = sent[: span.start_position]
                        m = sent[span.start_position : span.end_position]
                        r = sent[span.end_position :]
                        sentiment = self.tsc.infer_from_text(l, m, r)
                        for label in span.labels:
                            val = "" if label.value == "<unk>" else label.value
                            spans.append({
                                "text": span.text,
                                "start": span.start_position,
                                "end": span.end_position,
                                "value": val,
                                "score": f"{label.score:.2f}",
                                "sentiment": sentiment[0]["class_label"],
                                "probability": f"{sentiment[0]['class_prob']:.2f}",
                            })
                    output.append({
                        "sentence": sent,
                        "tag": sentence.tag.lower(),
                        "score": f"{sentence.score:.2f}",
                        "spans": spans,
                    })
                except Exception as e:
                    sys.stderr.write(f"{e}\nsent:\n{sentence}")
        return output

app = Flask(__name__)
classifier = FlairSentiment()

@app.route('/heartbeat', methods=['GET'])
def selftest():
    return jsonify({"status": "ok"}), 200

@app.route('/ner', methods=['POST'])
def ner():
    data = request.json
    if not data or 'text' not in data:
        return jsonify({"error": "No text provided"}), 400
    
    text = data['text']
    result = classifier.process_text(text)
    return jsonify(result), 200

def start_service(port):
    app.run(host='0.0.0.0', port=port)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Flair Sentiment Microservice")
    parser.add_argument('port', type=int, help='Port number to run the service on')
    args = parser.parse_args()
    
    start_service(args.port)
