#!/bin/bash
trap exit INT
[ $(tty) == '/dev/tty1' ] && imager
[ $(tty) == '/dev/tty1' ] && exit
