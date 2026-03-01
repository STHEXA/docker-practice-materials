# マルチステージビルドの確認用ファイル(Go の場合)

## ビルド
### マルチステージビルド
```sh
docker build -t golang-sample:multi-stage .
```

### マルチステージビルドでは無い場合
```sh
docker build -t golang-sample:single-stage -f ./Dockerfile.single-stage .
```


## サイズの確認
```sh
docker image ls --format 'table {{.Repository}}\t{{.Tag}}\t{{.Size}}' | grep golang-sample
```


## 実行
```sh
docker run --rm --name golang-app -p 8080:8080 golang-sample:multi-stage
```

あるいは以下

```sh
docker run --rm --name golang-app -p 8080:8080 golang-sample:single-stage
```

## あとかたづけ
コンテナを停止 & 削除(デタッチでコンテナを起動している場合)
```sh
docker rm golang-app -f 
```

イメージを削除
```sh
docker image rm golang-sample:single-stage golang-sample:multi-stage
```


