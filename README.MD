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

### 3. Annotation the unused container in docker-compose.dev.yml
`edit docker-compose.dev.yml, it will be copied from docker-compose.dev.yml.sample at the first running. # these configure codes`

### 4. Adjust the rights
```shell
# 第一次执行前先加权限：
chmod +x ./stop ./start ./restart ./remove ./go ./wingo ./upgrade ./docker-compose
chmod 777 -R .docker/php-fpm/log .docker/php81-fpm/log
```
### 4. Start Docker Compose
```shell
# 显示所有日志输出
./start
# 后台执行
./start -d
```
### 5. Restart
`./restart`

### 6. Stop
`./stop`

### 7. Remove all
`./remove`

### 8. Upgrade
`./upgrade`

### 9. log
```shell
./logs
#or
./logs |grep ubuntu
```

### 10. log in to container
```
# mac&linux
./go [sshd|php74|php7.4|php81|php8.1|nginx|mysql57|mysql5.7|mysql80|mysql8.0|mariadb|redis|ubuntu]
# windows
./wingo [sshd|php74|php7.4|php81|php8.1|nginx|mysql57|mysql5.7|mysql80|mysql8.0|mariadb|redis|ubuntu]
```

### 11. sshd
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
ifconfig -a
# net-tools
mysql -h mysql -u fusc -p
# mysql-client
redis-cli -h redis
# redis-cli, redis pwd is redis
drill baidu.com
# drill
dig baidu.com
# dig
ping baidu.com
# ping
nslookup baidu.com
# nslookup
traceroute baidu.com
# traceroute
# 部分环境因为开启了selinux会导致无法绑定22端口，需要改启动命令为：
entrypoint: ["sudo", "/usr/sbin/sshd", "-D", "-p", "2222"]
# 信息输出
entrypoint: ["sudo", "/usr/sbin/sshd", "-De", "-p", "2222"]
# 调式模式
entrypoint: ["sudo", "/usr/sbin/sshd", "-Dd", "-p", "2222"]
```

### 12. PHP7.4.28
```shell
http://php.demo/php74.php
```

### 13. PHP8.1.17
```shell
http://php.demo/php81.php
```

### 14. ADMINER
```shell
http://localhost:8899/ #root user : root ; password : fusc
#visit mysql8.0, sever name : mysql
#visit mysql5.7, sever name : mysql57
```
### 15. ES HEAD
```shell
http://localhost:9100/ # check ES data
```
### 16. 免密登录sftp
```shell
in windows, need to go container(./wingo ubuntu) and sudo chmod 755 /home/www-data/.ssh and sudo chmod 644 /home/www-data/.ssh/authorized_keys and sudo chown -R www-data:www-data authorized_keys at the first time.
ssh-copy-id -i ~/.ssh/id_rsa -p 3333 www-data@xx.xx.xx.xx # ubuntu container
```
### 17. mq management
```shell
http://192.168.236.170:15672/
fusc/fusc
```
