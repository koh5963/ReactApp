# ベースイメージの指定
FROM golang:1.23.5-alpine

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