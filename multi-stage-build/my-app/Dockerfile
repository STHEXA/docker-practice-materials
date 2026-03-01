# Webサーバーのnginxをベースイメージにする
FROM nginx:1.27-alpine

# node と pnpm のインストール
RUN apk add --no-cache nodejs pnpm

# 作業ディレクトリ
WORKDIR /app

# パッケージをインストール
COPY package*.json pnpm-lock.yaml ./
RUN pnpm i --frozen-lockfile

# 成果物をビルド
COPY . .
RUN pnpm run build

# 成果物を nginx のドキュメントルートに移動
RUN mv ./dist/* /usr/share/nginx/html
