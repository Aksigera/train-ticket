FROM --platform=linux/arm64/v8 openresty/openresty:trusty

RUN /bin/cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo 'Asia/Shanghai' >/etc/timezone

COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf

COPY static /usr/share/nginx/html
