package main

import (
	"online-indicator/pkg/db"
	"online-indicator/pkg/routes"

	"github.com/labstack/echo"
	"github.com/redis/go-redis/v9"
)

func dbMiddleware(db *redis.Client) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("redis", db)
			return next(c)
		}
	}
}

func main() {

	e := echo.New()
	rdc := db.NewRedisClient()
	e.Use(dbMiddleware(rdc))

	e.GET("/api/users", routes.UserHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
