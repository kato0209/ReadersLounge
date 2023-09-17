# dev
FROM node:19.8.1 AS dev
WORKDIR /readers-lounge/frontend
COPY package*.json ./
RUN npm install
EXPOSE 5173
CMD [ "npm", "run", "dev" ]


#release
FROM node:19.8.1 AS release
WORKDIR /readers-lounge/frontend
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 5173
CMD [ "npm", "run", "build" ]
