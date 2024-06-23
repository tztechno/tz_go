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
