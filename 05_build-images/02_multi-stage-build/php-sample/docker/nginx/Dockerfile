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


##############################
# 2. 実行環境(nginx)
##############################
FROM nginx:latest

# Node.jsビルドステージから生成物をコピー
COPY --from=node-build /app/public/dist /var/www/html/public/dist

# nginx の設定を用意
COPY ./docker/nginx/conf.d /etc/nginx/conf.d

# public 以下のファイルをコピー
COPY ./public /var/www/html/public

# nginx をフォアグラウンドで起動
CMD ["nginx", "-g", "daemon off;"]

