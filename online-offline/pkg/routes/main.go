package routes

import (
	"context"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func UserHandler(c echo.Context) error {

	db := c.Get("redis").(redis.Client)

	err := db.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

}
