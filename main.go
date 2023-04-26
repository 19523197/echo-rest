package main

import (
	"belajar-echo/database"
	"belajar-echo/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load(".env")

	e := echo.New()
	db, err := database.SetupSQLDatabase()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	router.SetupRouter(e, db)
	setupLog(e)

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func setupLog(e *echo.Echo) {
	log, err := os.OpenFile("log/"+time.Now().Format("01-02-2006")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: log,
		Format: `{"time":"${time_rfc3339}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
	}))
}
