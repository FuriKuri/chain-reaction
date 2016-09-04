# chain-reaction
## Usage
Start the first container.
```
$ docker run -d -v /var/run/docker.sock:/var/run/docker.sock  furikuri/chain-reaction
```
You can observe the behavior with:
```
$ watch 'docker ps'
Every 2,0s: docker ps                                                                                                                                                Sun Sep  4 10:31:36 2016

CONTAINER ID        IMAGE                     COMMAND                  CREATED             STATUS                  PORTS               NAMES
351e6bb9edc5        furikuri/chain-reaction   "/app/main --counter "   2 seconds ago       Up Less than a second   3000/tcp            chain-reaction-72
1b22233f7e8a        furikuri/chain-reaction   "/app/main --counter "   13 seconds ago      Up 12 seconds           3000/tcp            chain-reaction-73
```

### Cleanup
Remove all containers with based on this image:
```
docker run -v /var/run/docker.sock:/var/run/docker.sock  furikuri/chain-reaction --cleanup
```

## What will happend
The first container will start a new container form the same image with the name ```chain-reaction-10```. This container will remove the previous one and start a new container ```chain-reaction-9``` and so on ...
The last container ```chain-reaction-0``` will remove itself.