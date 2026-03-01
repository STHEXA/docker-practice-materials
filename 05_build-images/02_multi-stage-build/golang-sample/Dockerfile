##############################
# 1. ビルドステージ(Node.js)
##############################
FROM node:22.5.1-alpine3.20 AS node-build

# 作業ディレクトリを設定
WORKDIR /app

# package.jsonとpackage-lock.jsonをコピーして依存関係をインストール
COPY package*.json ./
RUN npm ci

# ソースコードをコピーしてビルド(npm ciで用意されたtailwindcssのパッケージを使ってビルド)
COPY . .
RUN npm run build


#####################
# 1. ビルドステージ
#####################
FROM golang:1.23-alpine AS golang-build
WORKDIR /app

# モジュールの依存を先にダウンロードしてキャッシュを効かせる
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN go build -o main .


#####################
# 2. 実行ステージ
#####################
FROM alpine:latest
WORKDIR /root/

# ビルドしたcssをコピー
COPY --from=node-build /app/public/dist ./public/dist

# ビルドしたバイナリをコピー
COPY --from=golang-build /app/main .

# HTMLファイルをコピー
COPY ./templates ./templates

# 必要に応じて環境変数を設定する
ENV PORT=8080

# デフォルトの公開ポート
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]
