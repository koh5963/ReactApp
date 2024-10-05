FROM node:20.13.1-alpine

WORKDIR /usr/src/frontend

# package.jsonとpackage-lock.jsonをコピー
COPY package*.json ./

# 依存関係のインストール
RUN npm install
RUN yarn install


# # react-router-domをインストール
# RUN npm install react-router-dom
# RUN npm install axios