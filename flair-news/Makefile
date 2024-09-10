build:
	docker build -t flair-news .
run:
	docker run -p 1337:1337 flair-news
heartbeat:
	curl http://localhost:1337/heartbeat
test:
	curl -X POST -H "Content-Type: application/json" -d '{"text":"Apple Inc. is planning to open a new store in New York City."}' http://localhost:1337/ner
