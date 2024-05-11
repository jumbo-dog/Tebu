docker build -t leo5123/tebu-service ../../.
docker run -it -p 80:10101 --env-file=../../.env leo5123/tebu-service