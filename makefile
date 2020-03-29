
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
	docker run -it --rm --name=chrome-ext -p 5000:5000 chrome-ext