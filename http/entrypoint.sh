#!/bin/sh

cat > /usr/share/nginx/html/pxelinux.cfg/default <<EOF
default vesamenu.c32
prompt 0
timeout 50

label Imager
  menu default
  linux vmlinuz rancher.cloud_init.datasources=[url:${RANCHER_CONFIG_URL}] vga=773 consoleblank=0 rancher.password=${RANCHER_PASSWORD}
  initrd initrd
EOF

exec "$@"
