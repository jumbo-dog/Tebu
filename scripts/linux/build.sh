docker build -t tebubot ../../
docker run -it -p 80:10101 --env-file=../../.env tebubot
