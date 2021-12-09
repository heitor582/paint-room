FROM node:16.10 as node

WORKDIR /angular/src/front-paint-room

RUN npm install -g @angular/cli --force

COPY . .
RUN npm install

EXPOSE 4200
