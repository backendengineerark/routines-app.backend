package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/backendengineerark/routines-app/configs"
	"github.com/backendengineerark/routines-app/internal/cron"
	"github.com/backendengineerark/routines-app/internal/infra/database"
	webhandler "github.com/backendengineerark/routines-app/internal/infra/web/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	time.Local = loc

	configs := configs.LoadConfig("../../.")

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	taskRepository := database.NewTaskMysqlRepository(db)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	taskHandler := webhandler.NewTaskHandler(taskRepository)
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", taskHandler.Create)
		r.Get("/", taskHandler.FindAll)
		// r.Get("/{id}", productHandler.GetProduct)
		// r.Put("/{id}", productHandler.UpdateProduct)
		// r.Delete("/{id}", productHandler.DeleteProduct)
	})

	routineHandler := webhandler.NewRoutineHandler(taskRepository)
	r.Route("/routines", func(r chi.Router) {
		r.Get("/", routineHandler.ListRoutine)
		// r.Get("/{id}", productHandler.GetProduct)
		// r.Put("/{id}", productHandler.UpdateProduct)
		// r.Delete("/{id}", productHandler.DeleteProduct)
	})

	go cron.ExecuteCronJobs(configs.CreateRoutinesTaskCron, taskRepository)
	fmt.Printf("Starting web server on port %s", configs.WebServerPort)
	http.ListenAndServe(":"+configs.WebServerPort, r)
}
