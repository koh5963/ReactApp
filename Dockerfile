# ベースイメージの指定
FROM golang:1.20-alpine

# ginのインストール
RUN go install github.com/codegangsta/gin@latest

# 作業ディレクトリの設定
WORKDIR /usr/src/backend

# ソースコードをコピー
COPY . .

# ビルド
RUN go build -o main .

# ポートの公開
EXPOSE 8080

# アプリケーションの実行
CMD ["gin", "--port", "8080", "--path", "/usr/src/backend", "run", "main.go"]

FROM node:20.13.1-alpine

WORKDIR /usr/src/frontend

# package.jsonとpackage-lock.jsonをコピー
COPY package*.json ./

# 依存関係のインストール
RUN npm install

# # react-router-domをインストール
# RUN npm install react-router-dom
# RUN npm install axios