# Flask application in Docker container
### Commands:
```sh
# build image
$ docker build -t emojis_loopback:final -f docker/Dockerfile .
# run image
$ docker run --name emojis_loopback -p 8080:80 -it -d emojis_loopback:final
# run interactively sh on a container
$ docker exec -it emojis_loopback sh
# stop container
$ docker stop emojis_loopback
# start container
$ docker start emojis_loopback
# remove container
$ docker rm emojis_loopback
# remove image
$ docker rmi emojis_loopback:final
```

When the container is running, you can try running:
```sh
$ curl -XPOST -d'{"word":"example", "count": 3}' http://127.0.0.1:8080
```
