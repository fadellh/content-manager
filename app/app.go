package app

import (
	_blogHandler "content/api/blog"
	blogService "content/business/blog"
	blogRepo "content/modules/blog"

	"content/modules"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newDatabaseConnection() *gorm.DB {

	db, e := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	modules.InitMigrate(db)

	return db
}

func StartApp() {
	e := echo.New()
	dbConnection := newDatabaseConnection()

	br := blogRepo.NewGormDB(dbConnection)
	bs := blogService.NewService(br)
	_blogHandler.NewHandler(e, bs)

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	go func() {
		// address := fmt.Sprintf("localhost:%d", config.AppPort)
		address := fmt.Sprintf("0.0.0.0:2801")
		// address := fmt.Sprintf("localhost:2801")
		fmt.Println(address)
		if err := e.Start(address); err != nil {
			fmt.Println(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
