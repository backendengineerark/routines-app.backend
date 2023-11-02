package main

import (
	"database/sql"
	"fmt"

	"github.com/backendengineerark/routines-app/configs"
	"github.com/backendengineerark/routines-app/internal/infra/database"
	webhandler "github.com/backendengineerark/routines-app/internal/infra/web/handler"
	"github.com/backendengineerark/routines-app/internal/infra/web/webserver"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs := configs.LoadConfig("../../.")

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	taskRepository := database.NewTaskMysqlRepository(db)

	webserver := webserver.NewWebServer(configs.WebServerPort)

	taskHandler := webhandler.NewTaskHandler(taskRepository)
	webserver.AddHandler("/tasks", taskHandler.Create)

	fmt.Printf("Starting web server on port %s", configs.WebServerPort)
	webserver.Start()

}
