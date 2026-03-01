# マルチステージビルドの確認用ファイル(PHPの場合)

## ビルド
### マルチステージビルド
```sh
docker build -t php-sample:multi-stage .
```

### マルチステージビルドでは無い場合
```sh
docker build -t php-sample:single-stage -f ./Dockerfile.single-stage .
```


## サイズの確認
```sh
docker image ls --format 'table {{.Repository}}\t{{.Tag}}\t{{.Size}}' | grep php-sample
```


## 実行
### 1. 専用のネットワークを作成

```sh
docker network create php-app-network
```

### 2. php-fpm のコンテナを起動
```sh
docker run -d --rm --name php-app --network php-app-network php-sample:multi-stage
```

あるいは以下

```sh
docker run -d --rm --name php-app --network php-app-network php-sample:single-stage
```

### 3. Webサーバー(nginx) を起動
イメージをビルド
```sh
docker build -f ./docker/nginx/Dockerfile -t my-web-server:latest .
```

コンテナを起動
```sh
docker run -d --rm --name web-server --network php-app-network --publish 8080:80 my-web-server:latest
```


## あとかたづけ
コンテナを停止 & 削除
```sh
docker rm web-server php-app -f 
```

イメージを削除
```sh
docker image rm php-sample:single-stage php-sample:multi-stage my-web-server:latest
```

ネットワークを削除
```sh
docker network rm php-app-network
```


