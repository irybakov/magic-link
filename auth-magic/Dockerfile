FROM node:alpine
WORKDIR /app 
COPY package.json /app 
COPY yarn.lock /app
RUN yarn install 
COPY . /app 
CMD yarn start 
EXPOSE 8081