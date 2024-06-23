package main

import (

    "github.com/gofiber/fiber/v2"
    "strconv"
    "time"
)

// Lucas数を計算する関数
func lucas(n int) int {
    if n == 0 {
        return 2
    }
    if n == 1 {
        return 1
    }
    return lucas(n-1) + lucas(n-2)
}

func main() {
    app := fiber.New()

    // 静的ファイルのホスティング
    app.Static("/", "./public")

    // /calculate エンドポイント
    app.Post("/calculate", func(c *fiber.Ctx) error {
        // リクエストのパース
        var requestData struct {
            N string `json:"n"`
        }
        if err := c.BodyParser(&requestData); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
        }

        // Nの変換
        n, err := strconv.Atoi(requestData.N)
        if err != nil || n < 0 {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid number"})
        }

        // Lucas数の計算と時間の測定
        start := time.Now()
        result := lucas(n)
        processTime := time.Since(start)

        // レスポンスの作成
        response := fiber.Map{
            "result":      result,
            "process_time": processTime.Microseconds(), // 時間をマイクロ秒で送信
        }

        return c.JSON(response)
    })

    // サーバの起動
    app.Listen(":3000")
}
