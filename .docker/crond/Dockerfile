FROM registry.cn-shanghai.aliyuncs.com/ulooocom/magento2:php-8.1.7-fpm

USER root
#change apk source
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# set time zone
RUN apk add --no-cache tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >  /etc/timezone

# https://github.com/krallin/tini
RUN apk add tini
# Tini is now available at /sbin/tini
ENTRYPOINT ["/sbin/tini", "--"]

# copy PHP configuration to image
COPY ./.docker/crond/etc/php/php.ini /usr/local/etc/php/php.ini

# add alias
#RUN alias php="/usr/local/bin/php"
RUN ln -s /usr/local/bin/php /usr/bin/php

ARG CRON_FILENAME=www-data
ENV CRON_FILENAME ${CRON_FILENAME}
# Add cron job
ADD ./.docker/crond/etc/crontabs/www-data /etc/cron.d/www-data
RUN chmod 0644 /etc/cron.d/www-data && crontab -u www-data /etc/cron.d/www-data

WORKDIR /data/www

#CMD ["crond", "-f"]
CMD ["/bin/bash", "-c", "set -e && crond -f"]