#FROM node:latest
FROM registry.cn-hangzhou.aliyuncs.com/ulooo/nodejs:16.14.2-alpine3.15
MAINTAINER Robin Fu <fusc@ulooo.com>
ENV HOST 0.0.0.0
RUN mkdir -p /data/www_frontend/nuxtjsdemoproj
COPY ./www_frontend/nuxtjsdemoproj/ /data/www_frontend/nuxtjsdemoproj
WORKDIR /data/www_frontend/nuxtjsdemoproj
EXPOSE 3000
RUN npm config set registry https://registry.npm.taobao.org
RUN npm install pm2 -g
RUN npm cache clean --force
RUN npm install
RUN npm run build
#RUN pm2 init
#CMD ["npm", "start"]
#CMD ["pm2-runtime", "--json", "process.yml"]
#"pm2-runtime start ecosystem.config.js --env production"
CMD ["pm2-runtime", "start", "ecosystem.config.js", "--env", "production"]
