# 概要
筋トレ実績を記録するサービス。

# 開発環境構築
ローカルで開発する際には下記の準備をする必要がある
- golang
- yarn

上記を準備して、次のコマンドを実行すればビルド及びコンテナ起動が開始してサーブされる。(localhost:8082でLISTEN)
```
$ make build
$ make docker/up
```
クライアントサイドの開発時には、clientディレクトリで
```
$ yarn run start
```
するのが良いかと思います。webpack-dev-serverでホットリロードできるしソースマップも効くので。(localhost:3000でLISTEN)

# docker
コンテナ起動するときは
```
$ make docker/up
```
停止するときは
```
$ make docker/down
```
イメージをリビルドするときは(ganyoginkoがすることはあまりないかもしれない)
```
$ make docker/rebuild
```

public/ とか server/bin/server はコンテナにマウントしているのでソース変更の度にイメージをリビルドしなきゃいけないとかはないです。