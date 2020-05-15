docker run --rm -it --net=host -v /var/run/docker.sock:/var/run/docker.sock -v `pwd`:`pwd` -w `pwd` docker/compose "$@"
