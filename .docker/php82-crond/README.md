# Reademe

## 1. Where is the crond config file
```apacheconf
vi ./.docker/php82-crond/crontabs/www-data
```

## 2. How to build image
```apacheconf
cd ./magent2dockerdevenv
docker build -t registry.cn-shanghai.aliyuncs.com/ulooocom/magento2:php-8.2.16-fpm-alpine3.18-cron -f ./.docker/php82-crond/Dockerfile .
```

# 3. How to use in docker-compose
```apacheconf
  php82-cron:
    image: ${ALIYUN_CONTAINER_URL}/magento2:php-8.2.16-fpm-alpine3.18-cron
    container_name: ${COMPOSE_PROJECT_NAME}_php82_cron
    cap_add:
      - SYS_PTRACE
    restart: always
    volumes:
      - ./.docker/php82-crond/etc/php/php.local.ini:/usr/local/etc/php/php.ini:ro
      - data-www:/data/www
    dns:
      - 223.5.5.5 #https://www.alidns.com/
      - 119.29.29.29 #https://www.dnspod.cn/Products/Public.DNS
      - 114.114.114.114
```

# 4. How to got to the container
```apacheconf
docker exec -ti -u www-data mage_local_php81_cron sh
```
