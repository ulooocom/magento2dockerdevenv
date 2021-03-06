#!/usr/bin/env bash

set -aeuo pipefail

echo "Run Tips:"
echo "********************************************************************"
echo "./go {container-name}"
echo "********************************************************************"

## whether skipped magento build
#echo SKIPPED_MAGENTO_BUILD_ACTION $1
if [ $# -ge 1 ] && [ -n "$1" ]
then
    CONTAINER_NAME=$1
    echo "********************************************************************"
    echo "Confirmed. Try to log go this container : "${CONTAINER_NAME}
    echo "********************************************************************"
else
    CONTAINER_NAME='sshd'
    echo "********************************************************************"
    echo "Confirmed. Try to log go the default container : "${CONTAINER_NAME}
    echo "********************************************************************"
fi

#mage_local_nginx
#mage_local_adminer
#mage_local_php_fpm
#mage_local_sshd
#mage_local_php81_fpm
#mage_local_mysql57
#mage_local_mysql
#mage_local_redis
#mage_local_php73_fpm
#mage_local_mariadb
#mage_local_es_head
#mage_local_es
#mage_local_sshd_ubuntu

case ${CONTAINER_NAME} in
 s|sshd)
       docker exec -ti mage_local_sshd sh
        ;;
 php73|php7.3)
       docker exec -ti mage_local_php73_fpm sh
        ;;
 php74|php7.4)
       docker exec -ti mage_local_php_fpm sh
        ;;
 php81|php8.1)
       docker exec -ti mage_local_php81_fpm sh
        ;;
 nginx)
       docker exec -ti mage_local_nginx sh
        ;;
 mysql57|mysql5.7)
       docker exec -ti mage_local_mysql57 sh
        ;;
 mysql80|mysql8.0)
       docker exec -ti mage_local_mysql sh
        ;;
 mariadb)
       docker exec -ti mage_local_mariadb sh
        ;;
 redis)
       docker exec -ti mage_local_redis sh
        ;;
 ubuntu)
       docker exec -ti mage_local_sshd_ubuntu sh
        ;;
 *)
       echo "Warning, Usage: ./go ${CONTAINER_NAME} [sshd|php73|php7.3|php74|php7.4|php81|php8.1|nginx|mysql57|mysql5.7|mysql80|mysql8.0|mariadb|redis|ubuntu]"
       exit 1
        ;;
esac
