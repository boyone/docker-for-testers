FROM node:14.14.0-alpine3.12
WORKDIR /app
COPY *.json /app/
RUN npm install
COPY index.js .
CMD npm start