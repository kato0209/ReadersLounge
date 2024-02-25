# dev
FROM node:19.8.1-buster-slim AS dev
RUN apt-get update && apt-get install -y openjdk-11-jre-headless && rm -rf /var/lib/apt/lists/*
RUN apt-get update && apt-get install make
WORKDIR /readers-lounge/frontend
COPY package*.json ./

ARG UID=1000
ARG GID=1000

RUN groupmod -g $GID node && \
    usermod -u $UID -g $GID node && \
    chown -R $UID:$GID /home/node && \
    chown -R $UID:$GID /readers-lounge/frontend
USER node


RUN npm install
EXPOSE 5173
CMD [ "npm", "run", "dev" ]


#release
FROM node:19.8.1-buster-slim AS release
RUN apt-get update && apt-get install -y openjdk-11-jre-headless && rm -rf /var/lib/apt/lists/*

WORKDIR /readers-lounge/frontend
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 80
ENTRYPOINT [ "npm", "run", "build" ]

CMD ["npm", "run", "start"]
