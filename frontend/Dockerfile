# dev
FROM node:19.8.1 AS dev
WORKDIR /readers-lounge/frontend
COPY package*.json ./
RUN npm install
EXPOSE 5173

ARG UID=1000
ARG GID=1000

RUN groupmod -g $GID node && \
    usermod -u $UID -g $GID node && \
    chown -R $UID:$GID /home/node && \
    chown -R $UID:$GID /readers-lounge/frontend

USER node
CMD [ "npm", "run", "dev" ]


#release
FROM node:19.8.1 AS release
WORKDIR /readers-lounge/frontend
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 5173
CMD [ "npm", "run", "build" ]
