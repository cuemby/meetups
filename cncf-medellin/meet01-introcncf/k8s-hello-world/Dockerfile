FROM node:10-alpine

RUN mkdir -p /home/node/code

WORKDIR /home/node/code

COPY . /home/node/code

RUN npm i --production

CMD npm start

EXPOSE 8080