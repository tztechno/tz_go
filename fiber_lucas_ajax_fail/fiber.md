Go言語でFiberという軽量なWebフレームワークを使用してHTMLをホストする方法を説明します。Fiberは高速で効率的なルーティングとミドルウェア機能を提供するため、Webアプリケーション開発に適しています。

### 1. プロジェクトのセットアップ

まずは、Goのプロジェクトをセットアップします。以下の手順に従ってください。

#### a. プロジェクトディレクトリの作成

まず、プロジェクト用のディレクトリを作成します。例えば、`myapp` という名前のディレクトリを作成します。

```bash
mkdir myapp
cd myapp
```

#### b. Goモジュールの初期化

Goモジュールを初期化し、依存関係を管理します。

```bash
go mod init myapp
```

### 2. Fiberのインストール

次に、Fiberをインストールします。

```bash
go get github.com/gofiber/fiber/v2
```

### 3. HTMLファイルの作成

プロジェクトのルートディレクトリに `public` フォルダを作成し、その中に `index.html` を配置します。

```html
<!-- public/index.html -->
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>My Fiber App</title>
</head>
<body>
    <h1>Hello, Fiber!</h1>
</body>
</html>
```

### 4. Goコードの作成

`main.go` という名前でGoのエントリーポイントファイルを作成します。

```go
// main.go
package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Fiberアプリケーションのインスタンスを作成
	app := fiber.New()

	// 静的ファイル（HTML）をホストする
	app.Static("/", "./public")

	// ルートパスのハンドラ
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// サーバーをポート8080で起動
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
```

### 5. アプリケーションの実行

以下のコマンドでアプリケーションをビルドして実行します。

```bash
go run main.go
```

### 6. アプリケーションの確認

ブラウザで `http://localhost:8080` にアクセスすると、`index.html` の内容が表示されるはずです。

### まとめ

これで、Go言語でFiberフレームワークを使用してHTMLファイルをホストする方法が完了しました。Fiberは高速で効率的なルーティングと静的ファイルのホスティングをサポートし、シンプルで使いやすいAPIを提供します。

cd myapp
go run main.go
http://127.0.0.1:3000   
