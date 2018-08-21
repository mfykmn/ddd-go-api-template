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
├── cmd
│   └── api
│       └── main.go
│             * エントリーポイント
├── domain
│     * DDDでいうドメイン
├── handler
│   │ * HTTPのハンドラ
│   └── middleware
│        * HTTPハンドラの前後の処理を行う
├── infrastructure
│     * DDDでいうリポジトリを格納する
│     * DBやAPIなど外部とのやり取りを行う処理
└── library
      * 汎用性の高いプログラムを１つのパッケージとして提供できるようにしたもの