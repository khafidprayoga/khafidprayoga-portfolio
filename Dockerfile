FROM node:14-alpine
WORKDIR /app
COPY . /app
RUN yarn
CMD ["yarn","start"]