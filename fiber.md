`fiber`はGoのフレームワークであり、非常に高速かつ簡単にWebサーバを構築することができます。以下の手順に従って、ルートホルダ`myapp`にfiberをインストールし、`index.html`をホストする方法を説明します。

### 1. プロジェクトのセットアップ

まず、新しいGoプロジェクトのディレクトリを作成し、その中に移動します。

```bash
mkdir myapp
cd myapp
```

### 2. Goモジュールの初期化

次に、Goモジュールを初期化します。

```bash
go mod init myapp
```

### 3. fiberのインストール

次に、fiberをインストールします。

```bash
go get -u github.com/gofiber/fiber/v2
```

### 4. プロジェクト構成

以下のようなプロジェクト構成を作成します。

```
myapp/
├── go.mod
├── go.sum
├── main.go
└── public/
    └── index.html
```

`public`ディレクトリに`index.html`ファイルを作成し、その中に適当なHTMLコンテンツを追加します。

### 5. main.goの作成

次に、`main.go`ファイルを作成し、以下のコードを追加します。

```go
package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // 静的ファイルのホスティング
    app.Static("/", "./public")

    // サーバの起動
    app.Listen(":3000")
}
```

### 6. サーバの起動

最後に、サーバを起動します。

```bash
go run main.go
```

これで、Fiberサーバがポート3000で起動し、`index.html`をホストします。ブラウザで`http://localhost:3000`にアクセスすると、`index.html`の内容が表示されます。

### まとめ

以上の手順に従って、ルートホルダ`myapp`にfiberをインストールし、`index.html`をホストすることができます。Fiberの設定は非常にシンプルであり、少ないコードでWebサーバを立ち上げることができるため、非常に便利です。

---
```

関数は再帰関数にして計算速度の遅いやつが良い
backでmicro秒の値を返す
frontでmili秒（小数点以下３桁）に変換して表示する

cd myapp
go run main.go
http://127.0.0.1:3000   
```

---