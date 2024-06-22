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
