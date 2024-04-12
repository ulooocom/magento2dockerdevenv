#!/bin/bash
set -x
# 保存环境变量，开启crontab服务
env >> /etc/default/locale
/etc/init.d/cron start
