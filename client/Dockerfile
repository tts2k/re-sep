FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
RUN npm run build

FROM node:20-alpine
USER node
WORKDIR /app
COPY --from=builder --chown=node:node /app/build build/
COPY --from=builder --chown=node:node /app/node_modules node_modules/
COPY --chown=node:node package.json .
EXPOSE 3000
ENV NODE_ENV=production
CMD [ "node", "build" ]
