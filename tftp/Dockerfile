FROM alpine:3.7

RUN apk add --no-cache tftp-hpa syslinux 

WORKDIR /srv/tftp/
RUN cp -a /usr/share/syslinux ./
RUN ln -s syslinux/pxelinux.0

ENV RANCHER_VERSION=1.5.5
RUN wget http://releases.rancher.com/os/v${RANCHER_VERSION}/vmlinuz
RUN wget http://releases.rancher.com/os/v${RANCHER_VERSION}/initrd
RUN chmod a+rX . -R

COPY entrypoint.sh /
RUN mkdir pxelinux.cfg

ENV RANCHER_CONFIG_URL=http://localhost/Imager.yml
ENV RANCHER_PASSWORD=rancher

EXPOSE 69/udp

ENTRYPOINT ["/entrypoint.sh"]
CMD ["/usr/sbin/in.tftpd", "-L", "-vvv", "-s", "/srv/tftp"]
