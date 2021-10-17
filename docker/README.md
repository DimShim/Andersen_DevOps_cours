# Flask application in Docker container
### Commands:
```sh
# build image
$ docker build -t emojis_loopback:final -f docker/Dockerfile .
# run image
$ docker run --name emojis_loopback -p 8080:80 -it -d emojis_loopback:final
```
