FROM node:16

# App directory
WORKDIR /usr/src/app

# app deps
COPY package*.json ./
RUN npm install

COPY . .
EXPOSE 8888

CMD [ "node", "bot.js" ]