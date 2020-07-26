ARG PHP_IMAGE=php:7.4-fpm-alpine
ARG NGINX_IMAGE=nginx:1.19-alpine

# php stage
FROM ${PHP_IMAGE} AS php

COPY ./docker/build/php/ /etc/php7/

ARG APP_DIR=/var/www/app
WORKDIR ${APP_DIR}

CMD ["php-fpm"]

# nginx stage
FROM ${NGINX_IMAGE} AS nginx

COPY ./docker/build/nginx/conf.d/ /etc/nginx/conf.d/
COPY ./docker/build/nginx/ /etc/nginx/

CMD ["nginx", "-g", "daemon off;"]

EXPOSE 80
