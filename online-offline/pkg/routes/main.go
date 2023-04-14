package routes

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func UserHandler(c echo.Context) error {

	db := c.Get("redis").(*redis.Client)
	users_string := c.QueryParam("users")
	users := strings.Split(users_string, ",")

	val, err := db.MGet(ctx, users...).Result()

	if err != nil {
		return c.String(
			http.StatusInternalServerError, "server failed",
		)
	}

	return c.JSONPretty(http.StatusOK, val, "  ")

}

func UserHeartBeatHandler(c echo.Context) error {

	db := c.Get("redis").(*redis.Client)
	users_string := c.QueryParam("user")

	_, err := db.Set(ctx, users_string, time.Now(), 30*time.Second).Result()

	if err != nil {
		return c.String(
			http.StatusInternalServerError, "server failed",
		)
	}

	return c.String(http.StatusOK, "HB Set")

}
