# 概要
筋トレ実績を記録するサービス。

# 開発環境構築
- golang
- yarn

上記を準備してmakeすればビルド、サーブまでできる。\
クライアントサイドの開発時には、clientディレクトリで
```
# yarn run start
```
するのが良いかと思います。webpack-dev-serverでホットリロードできるしソースマップも効くので。

## db
ユーザ名とパスワード、DB名はすべてworkoutとして環境作ってください。
下記参考にしました。\
https://qiita.com/asylum/items/2bb69fee5fc8ad932e37 \
マイグレーションをするにはdbディレクトリで下記コマンド実行。
```
# sql-migrate up
# sql-migrate down
```