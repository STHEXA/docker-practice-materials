###############################
# ビルドステージ(node)
###############################
FROM node:latest AS build-stage

# pnpm をインストール
RUN npm install -g pnpm

# 作業ディレクトリ
WORKDIR /app

# パッケージをインストール
COPY package*.json pnpm-lock.yaml ./
RUN pnpm i --frozen-lockfile

# 成果物をビルド
COPY . .
RUN pnpm run build

###############################
# 実行環境(nginx)
###############################
FROM nginx:1.27-alpine

# build-stage でできた成果物を nginx のドキュメントルートに移動
COPY --from=build-stage /app/dist /usr/share/nginx/html
