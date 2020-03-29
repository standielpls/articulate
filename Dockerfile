# The base image - ubuntu OS - with Python 3.7 installed.
FROM python:alpine3.7

# Copy everything here to inside the image at "/app".
COPY python /app

# WORKDIR is "set" where the home directory.
WORKDIR /app

# RUN executes the following command - install all the Python dependencies.
RUN pip install -r requirements.txt

# EXPOSE port 5000.
EXPOSE 5000

# CMD is the command that runs when the Docker container is started.
CMD flask run --host 0.0.0.0

# Docker
# 1. Builds the image.
# 1. a) Will get created one by one, and "cached".
# 2. Once image is built, it is done "building".
# 3. "docker run" to actually start the service.