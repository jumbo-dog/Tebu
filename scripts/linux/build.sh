docker build -t tebubot ../../.
docker run -it -p 10101:80 --env-file=../../.env tebubot