##

# Magento 2 CE/EE DOCKER ENVIRONMENT

### suggestion : 
#### 1. China registry
```
 "https://registry.docker-cn.com",
 "https://k6djaujw.mirror.aliyuncs.com"
```

### 2. install git and docker-compose
```
sudo apt install git
sudo apt install docker-compose # if you install docker-desktop, it will be added too. not need to install independently.
```

### 3. annotation the unused container in docker-compose.dev.yml
`edit docker-compose.dev.yml, it will be copied from docker-compose.dev.yml.sample at the first running. # these configure codes`

### 4. Start Docker Compose
`sh ./start`

### 5. Restart
`sh ./restart`

### 6. Stop
`sh ./stop`

### 7. Remove all
`sh ./remove`

### 8. Upgrade
`sh ./upgrade`

### 9. log in to container
```
# mac&linux
sh ./go [sshd|php73|php7.3|php74|php7.4|php81|php8.1|nginx|mysql57|mysql5.7|mysql80|mysql8.0|mariadb|redis]
# windows
sh ./wingo [sshd|php73|php7.3|php74|php7.4|php81|php8.1|nginx|mysql57|mysql5.7|mysql80|mysql8.0|mariadb|redis]
```

### 10. sshd
```shell
# sshd configure, for phpStorm auto deployment when developing
ssh -p 2222 www-data@127.0.0.1 ##passwd:qwer
cd /data/www

# tools in sshd container : 
su root
# password is : admin
git --version
#git version 2.34.2
composer -v
#Composer version 2.3.5
node -v
#v16.14.2
npm -v
#8.1.3
cnpm -v
#cnpm@7.1.1 (/usr/local/lib/node_modules/cnpm/lib/parse_argv.js)
apidoc -v
#apidoc version: 0.51.1
ifconfig -a
# net-tools
ipaddr
# net-tools
mysql -h mysql -u fusc -p
# mysql-client
redis-cli -h redis
# redis-cli, redis pwd is redis
drill baidu.com
# drill
```

### 11. PHP7.3.27
```shell
http://localhost:7300/
```

### 12. PHP7.4.28
```shell
http://localhost:7400/
```

### 13. PHP8.1.5
```shell
http://localhost:8100/
```

### 14. ADMINER
```shell
http://localhost:8899/ #root user : root ; password : fusc
#visit mysql8.0, sever name : mysql
#visit mysql5.7, sever name : mysql57
#visit mariadb, sever name : mariadb
```
### 15. ES HEAD
```shell
http://localhost:9100/ # check ES data
```
