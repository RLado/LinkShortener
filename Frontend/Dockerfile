FROM nginx:mainline-alpine-slim
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY StaticSite /usr/share/nginx/html
RUN chmod 444 /usr/share/nginx/html/*
