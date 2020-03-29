
build:
	pip freeze > requirements.txt

run:
	. python/venv/bin/activate
	cd python && flask run --reload

mongo:
	docker run -p 27017:27017 --name mongodb --rm -d mongo


docker-build:
	docker build . -t chrome-ext

docker-run:
	docker run -it --rm -p 5000:5000 --name chrome-ext chrome-ext