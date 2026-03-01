# Webサーバーのnginxをベースイメージにする
FROM nginx:1.27-alpine

# node と pnpm のインストール
RUN apk add --no-cache nodejs pnpm

# 作業ディレクトリ
WORKDIR /app

# パッケージをインストール
COPY package*.json ./
RUN pnpm i --frozen-lockfile

# 成果物をビルド
COPY . .
RUN pnpm run build

# 成果物を nginx のドキュメントルートに移動
RUN mv ./build/* /usr/share/nginx/html
