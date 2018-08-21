# GolangのDDDに沿ったAPI Template

## ディレクトリ構成について
.
├── _sql
│     * マイグレーション対象のクエリを格納する
├── _tools
│     * goのコード以外の設定値ファイルなどを格納する
├── application
│     * DDDでいうアプリケーションサービスを格納する
│     * libraryやinfrastructure層の呼び出しを行う
├── config
│     * 設定値を読み込み構造体に格納する
├── cmd
│   └── api
│       └── main.go
│             * エントリーポイント
├── domain
│     * DDDでいうドメイン
│     * どこにも依存してはいけない
├── handler
│   │ * HTTPのハンドラ
│   └── middleware
│        * HTTPハンドラの前後の処理を行う
├── infrastructure
│     * DDDでいうリポジトリを格納する
│     * DBやAPIなど外部とのやり取りを行う処理
└── library
      * 汎用性の高いプログラムを１つのパッケージとして提供できるようにしたもの
      
## 参考リンク
* [標準的なパッケージのレイアウト](http://allishackedoff.hatenablog.com/entry/2016/08/23/015016)
* [Goのパッケージ構成の失敗遍歴と現状確認](https://medium.com/@timakin/go%E3%81%AE%E3%83%91%E3%83%83%E3%82%B1%E3%83%BC%E3%82%B8%E6%A7%8B%E6%88%90%E3%81%AE%E5%A4%B1%E6%95%97%E9%81%8D%E6%AD%B4%E3%81%A8%E7%8F%BE%E7%8A%B6%E7%A2%BA%E8%AA%8D-fc6a4369337)