
FROM registry.cn-hangzhou.aliyuncs.com/ulooo/nodejs:16.14.2-alpine3.15
MAINTAINER Robin Fu <robinfu@shtag.com>
ENV HOST 0.0.0.0
RUN mkdir -p /data/www_frontend/nuxtjsdemoproj && chown -R www-data:www-data /data/www_frontend
COPY --chown=www-data:www-data ./www_frontend/nuxtjsdemoproj/ /data/www_frontend/nuxtjsdemoproj
WORKDIR /data/www_frontend/nuxtjsdemoproj
EXPOSE 3000
RUN npm config set registry https://registry.npm.taobao.org
RUN npm cache clean --force
RUN npm install
#RUN npm run build
#CMD ["npm", "start"]
CMD ["npm", "run", "dev"]
