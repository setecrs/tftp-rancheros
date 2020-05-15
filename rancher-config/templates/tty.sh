#!/bin/bash
[ $(tty) == '/dev/tty1' ] && tty1
[ $(tty) == '/dev/tty2' ] && tty2
[ $(tty) == '/dev/tty3' ] && tty3
