# ベースにするイメージを指定
FROM ubuntu:latest

# RUN(コマンドの実行)でcowsayコマンドをインストール
RUN apt-get update

# RUN(コマンドの実行)でcowsayコマンドをインストール
RUN apt-get install cowsay -y

#
# Case1. ENTRYPOINT のみを書くケース
#
# ENTRYPOINT を 以下のように書くと CMD ["Hello, Docker!"] が不要
# （その代わりコンテナの起動時に cowsay に指定する第一引数を変更できない。第二引数以降を追加することは可能。)
#
# ENTRYPOINT ["/usr/games/cowsay", "Hello, Docker!"]

#
# Case2. CMD のみを書くケース
#
# 以下の場合は ENTRYPOINT ["/usr/games/cowsay"] が不要になる
# コンテナの起動時に実行するコマンド自体を上書き可能(cowsay 自体の実行を強制できない)
#
# CMD ["/usr/games/cowsay", "Hello, Docker!"]

#
# Case3. ENTRYPOINT と CMD を組み合わせて書くケース
#
# 以下の場合は起動時に cowsay が実行される。かつ、cowsay　の引数に "Hello, Docker!" が指定される。
# "Hello, Docker!" の部分はコンテナの起動時に上書きできる
#
ENTRYPOINT ["/usr/games/cowsay"]
CMD ["Hello, Docker!"]
