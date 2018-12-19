FROM node:9.6.1

RUN mkdir -p /usr/src
ENV PATH /usr/src/app/node_modules/.bin:$PATH

RUN npm install react-scripts@1.1.1 -g

COPY . /usr/src/app
WORKDIR /usr/src/app

RUN npm install

CMD ["npm", "start"]
