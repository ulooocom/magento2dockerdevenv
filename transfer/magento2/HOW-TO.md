# HOW-TOs

### 1.download
```http request
https://magento.com/tech-resources/download#archive-releases
```

### 2. start containers
```shell
sh ./start
```

### 3. docker cp
```shell
docker cp ./www/magentoce243p1demo da585d08b66b:/data/www/
```

### 4.domain
```
docker.local

# add below to /etc/hosts
127.0.0.1 docker.local

```

### 5.database name
```
magentoee243p1demo
```

### 6.nginx cfg
```shell
./.docker/nginx/etc/
```

### 7.docker compose cfg
```shell

# nginx
    ports:
      - 80:80
```

### 8. recreate containers
```shell
sh ./restart
```

### 9.install script
```
php bin/magento setup:install --base-url=http://docker.local/ --db-host=mysql --db-name=magentoee243p1demo --db-user=root --db-password=fusc --admin-firstname=Magento --admin-lastname=User --admin-email=robinfu@shtag.com --admin-user=admin --admin-password=admin123 --language=zh_Hans_CN --currency=CNY --timezone=Asia/Shanghai --use-rewrites=1 --search-engine=elasticsearch7 --elasticsearch-host=elasticsearch --elasticsearch-index-prefix=magentoce243p1demo --elasticsearch-port=9200  --use-sample-data
```

### 10.session redis config
```
'session' =>
array (
  'save' => 'redis',
  'redis' =>
  array (
    'host' => 'redis',
    'port' => '6379',
    'password' => 'redis',
    'timeout' => '2.5',
    'persistent_identifier' => '',
    'database' => '2',
    'compression_threshold' => '2048',
    'compression_library' => 'gzip',
    'log_level' => '3',
    'max_concurrency' => '100',
    'break_after_frontend' => '5',
    'break_after_adminhtml' => '30',
    'first_lifetime' => '600',
    'bot_first_lifetime' => '60',
    'bot_lifetime' => '7200',
    'disable_locking' => '0',
    'min_lifetime' => '60',
    'max_lifetime' => '2592000'
  )
),
```
### 11.cache redis config
```
'cache' => [
        'frontend' => [
            'default' => [
                'id_prefix' => '483_',
                'backend' => 'Magento\\Framework\\Cache\\Backend\\Redis',
                'backend_options' => [
                    'server' => '127.0.0.1',
                    'password' => 'redis',
                    'database' => '0',
                    'port' => '6379',
                    'compress_data' => '1',
                    'compression_lib' => '',
                    'preload_keys' => [
                        '483_EAV_ENTITY_TYPES',
                        '483_GLOBAL_PLUGIN_LIST',
                        '483_DB_IS_UP_TO_DATE',
                        '483_SYSTEM_DEFAULT'
                    ]
                ]
            ],
            'page_cache' => [
                'id_prefix' => '483_',
                'backend' => 'Magento\\Framework\\Cache\\Backend\\Redis',
                'backend_options' => [
                    'server' => '127.0.0.1',
                    'password' => 'redis',
                    'database' => '2',
                    'port' => '6379',
                    'compress_data' => '0',
                    'compression_lib' => ''
                ]
            ]
        ],
        'allow_parallel_generation' => false
    ],
```
### 12.magento commands
```shell
## all commands
php bin/magento
## magento module status
php bin/magento module:status
## disable magento module 
php bin/magento module:disable Magento_TwoFactorAuth
php bin/magento c:f
php bin/magento setup:upgrade
php bin/magento setup:compile
php bin/magento indexer:reindex
...
```
