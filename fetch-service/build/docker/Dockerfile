
FROM node:current-alpine3.16
ENV NODE_ENV=production

WORKDIR /app

COPY ../../. .

RUN yarn install --production=false && \
    yarn build && \
    rm -rf src/ && \
    rm -rf config/ && \
    rm -rf node_modules && \
    yarn install \
CMD [ "node", "./dist/server/server.js" ]