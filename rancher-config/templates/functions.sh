#!/bin/bash
waitfor(){
  wait-for-docker
  if ! docker inspect $1 >/dev/null 2>/dev/null
  then
    $1
  fi
}
udev(){
  docker run -d --name=udev \
    --privileged \
    --net=host \
    -v /dev/:/dev/ \
    setecrs/imager:2.4.0 udevd
}
imager(){
  waitfor udev
  docker run --rm -it \
    --privileged \
    --net=host \
    {{- range .Mounts}}
    -v {{ index . 1 }}:{{ index . 1 }} \
    {{- end}}
    -v /dev/:/dev/ \
    -e GRAPHQL_URL='http://192.168.2.79/graphql' \
    setecrs/imager:2.5.0 $@
}
udevadm-trigger(){
 docker exec udev udevadm trigger -s block
}
