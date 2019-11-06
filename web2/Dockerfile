FROM node:latest AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN npm install && npm run build

FROM nginx:latest
RUN rm -rf /usr/share/nginx/html && mkdir /usr/share/nginx/html
COPY --from=build /app/dist/ /usr/share/nginx/html/