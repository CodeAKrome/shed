build:
	docker build -t flair-news .
run:
	docker run -p 1337:1337 flair-news
heartbeat:
	curl http://localhost:1337/heartbeat
test:
	curl -X POST -H "Content-Type: application/json" -d '{"text":"The 1st thing to do on Jan 1 at 5am is attend the Viking Funeral of Darth Vader in the School Gym. 1000 Rastafarians will chant 50% of the description of The Thinker for 10 dollars."}' http://localhost:1337/ner
