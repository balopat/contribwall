FROM node:9.6.1

RUN mkdir -p /usr/src
ENV PATH /usr/src/app/node_modules/.bin:$PATH

RUN npm install react-scripts@1.1.1 -g
RUN npm install

COPY . /usr/src/app
WORKDIR /usr/src/app

CMD ["npm", "start"]
