FROM nginx:1.17-alpine

RUN apk add --no-cache syslinux 

WORKDIR /usr/share/nginx/html/
RUN cp -a /usr/share/syslinux/* ./

ENV RANCHER_VERSION=1.5.5
RUN wget http://releases.rancher.com/os/v${RANCHER_VERSION}/vmlinuz
RUN wget http://releases.rancher.com/os/v${RANCHER_VERSION}/initrd
RUN chmod a+rX . -R

COPY entrypoint.sh /
RUN mkdir pxelinux.cfg

ENV RANCHER_CONFIG_URL=http://localhost/Imager.yml
ENV RANCHER_PASSWORD=rancher

ENTRYPOINT ["/entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
