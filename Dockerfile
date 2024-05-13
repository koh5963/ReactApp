FROM node:20.13.1-alpine

WORKDIR /usr/src/app

# react-router-domをインストール
RUN npm install react-router-dom
