# GolangのDDDに沿ったAPI Template

## 📌 概要
<details><summary>ディレクトリ構成について</summary><br><pre>

```
.
├── _sql
│     * マイグレーション対象のクエリを格納する
│
├── _tools
│     * goのコード以外の設定値ファイルなどを格納する
│
├── application
│     * DDDでいうアプリケーションサービスの実装を格納する
│     * libraryやinfrastructure層の呼び出しを行う
│
├── config
│     * 設定値を読み込み構造体に格納する
│
├── cmd
│   └── api
│       └── main.go
│             * エントリーポイント
│ 
├── domain
│   │ * DDDでいうドメイン
│   │ * 他の層に依存してはいけない
│   │ * 共通のものをまとめるというより、entity定義やそれに付随するロジックを置く場所
│   │
│   ├── repository
│   │     * repositoryのinterface定義などを置く
│   │
│   └── service
│         * application serviceのinterface定義などを置く
│ 
├── handler
│   │ * HTTPのハンドラ
│   │
│   └── middleware
│        * HTTPハンドラの前後の処理を行うミドルウェアを置く
│
├── infrastructure
│   │ * DDDでいうリポジトリの実装を格納する
│   │ * DBやAPIなど外部とのやり取りを行う処理を置く
│   │
│   └── db
│         * DBコネクション生成処理などを格納
│
└── library
      * 汎用性の高いプログラムを１つのパッケージとして提供できるようにしたもの
```

</pre></details>

### API仕様書
[こちらをご覧ください](https://documenter.getpostman.com/view/2534584/RWTpswZX)

## 🌐 動作環境 
* Golang v1.11.0
* Docker Latest Version

## ▶️ 実行方法
### 環境立ち上げ
```
$ make dstart
$ make dmigrate
```

### API呼び出し
Postmanから呼び出す

[こちらをご覧ください](https://documenter.getpostman.com/view/2534584/RWTpswZX)


## その他
### 参考リンク
* [標準的なパッケージのレイアウト](http://allishackedoff.hatenablog.com/entry/2016/08/23/015016)
* [Goのパッケージ構成の失敗遍歴と現状確認](https://medium.com/@timakin/go%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E6%A7%8B%E6%88%90%E3%81%AE%E5%A4%B1%E6%95%97%E9%81%8D%E6%AD%B4%E3%81%A8%E7%8F%BE%E7%8A%B6%E7%A2%BA%E8%AA%8D-fc6a4369337)
