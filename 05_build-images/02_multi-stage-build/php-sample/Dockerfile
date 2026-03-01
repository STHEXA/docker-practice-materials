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
# 2. ビルドステージ(composer)
##############################
FROM composer:2 AS composer-build

# 作業ディレクトリを設定
WORKDIR /app

# composer.jsonとcomposer.lockをコピーして依存関係をインストール
COPY composer.json composer.lock ./
RUN composer install --no-dev --optimize-autoloader


##############################
# 3. 実行環境(php-fpm)
##############################
FROM php:8.3-fpm-alpine

# 必要なパッケージをインストール
RUN apk --no-cache add \
    freetype \
    libjpeg-turbo \
    libpng \
    libzip \
    zip \
    unzip \
    && apk add --no-cache --virtual .build-deps \
    freetype-dev \
    libjpeg-turbo-dev \
    libpng-dev \
    libzip-dev \
    && docker-php-ext-configure gd --with-freetype --with-jpeg \
    && docker-php-ext-install gd zip \
    && apk del .build-deps

# ソースをコピー
COPY . /var/www/html

# Composerビルドステージから生成物をコピー
COPY --from=composer-build /app/vendor /var/www/html/vendor

# Node.jsビルドステージから生成物をコピー
COPY --from=node-build /app/public/dist /var/www/html/public/dist

