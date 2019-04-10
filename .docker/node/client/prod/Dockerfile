# build stage
FROM node:10.13.0-alpine as builder

ENV NODE_ENV=production

RUN mkdir -p /app
COPY . /app
WORKDIR /app

RUN npm run build

# final stage
FROM node:10.13.0-alpine
LABEL maintainer="junxiang"

ENV HOST=0.0.0.0

RUN mkdir -p /app
WORKDIR /app

ADD package.json ./
ADD nuxt.config.js ./

COPY --from=builder ./app/server ./server/
COPY --from=builder ./app/node_modules ./node_modules/
COPY --from=builder ./app/.nuxt ./.nuxt/
COPY --from=builder ./app/static ./static/

EXPOSE 3000

CMD ["npm", "run", "start"]