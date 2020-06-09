#!/bin/sh

cat > /srv/tftp/pxelinux.cfg/default <<EOF
default vesamenu.c32
prompt 0
timeout 50

label Imager
  menu default
  linux vmlinuz rancher.cloud_init.datasources=[url:${RANCHER_CONFIG_URL}] vga=773 consoleblank=0
  initrd initrd
EOF

exec "$@"
