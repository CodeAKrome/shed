build:
	docker build -t flair-ner .
run:
	docker run -p 1337:1337 flair-ner
heartbeat:
	curl http://localhost:1337/heartbeat
test:
	curl -X POST -H "Content-Type: application/json" -d '{"text":"Apple Inc. is planning to open a new store in New York City."}' http://localhost:1337/ner
