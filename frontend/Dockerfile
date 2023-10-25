FROM node:19.0.0-bullseye-slim

COPY ./package.json ./

COPY ./ ./

EXPOSE 3000

WORKDIR /usr/app/frontend

RUN npm install

CMD ["npm", "start"]