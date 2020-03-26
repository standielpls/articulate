
build:
	pip freeze > requirements.txt

run:
	. python/venv/bin/activate
	cd python && flask run --reload

mongo:
	docker run -p 27017:27017 --name mongodb --rm -d mongo