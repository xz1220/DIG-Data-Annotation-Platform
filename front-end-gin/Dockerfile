# From node:14.4.0-slim
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# RUN apk add nginx && mkdir /run/nginx/
# WORKDIR /usr/src/app
# COPY ["package.json", "package-lock.json*", "npm-shrinkwrap.json*", "./"]
# RUN npm install -g cnpm -registry=https://registry.npm.taobao.org 
# RUN cnpm install
# COPY . .
# RUN cnpm run build
# RUN ln -sf /dev/stdout /var/log/nginx/access.log \
# 	&& ln -sf /dev/stderr /var/log/nginx/error.log
# EXPOSE 80
# RUN cp -r dist/* /var/www/html \
#     && rm -rf /user/src/app
# CMD ["nginx","-g","daemon off;"]

FROM nginx
COPY ./docs /usr/share/nginx/html/
COPY ./nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx","-g","daemon off;"]
