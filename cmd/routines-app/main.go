package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
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
	configs := configs.LoadConfig("../../.")

	loc, _ := time.LoadLocation(configs.TZ)
	time.Local = loc

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName, url.QueryEscape(configs.TZ)))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	taskRepository := database.NewTaskMysqlRepository(db)
	metricRepository := database.NewMetricMysqlRepository(db)

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
		r.Put("/{task_id}", taskHandler.Update)
		r.Delete("/{task_id}", taskHandler.Delete)
		r.Post("/{task_id}/archive", taskHandler.Archive)
		r.Post("/{task_id}/unarchive", taskHandler.Unarchive)
	})

	routineHandler := webhandler.NewRoutineHandler(taskRepository)
	r.Route("/routines", func(r chi.Router) {
		r.Get("/", routineHandler.ListRoutine)
		r.Post("/{task_id}/today-finish", routineHandler.FinishRoutine)
		r.Post("/{task_id}/today-unfinish", routineHandler.UnfinishRoutine)
		r.Delete("/{routine_id}", routineHandler.DeleteRoutine)
	})

	metricHandler := webhandler.NewMetricHandler(metricRepository)
	r.Route("/metrics", func(r chi.Router) {
		r.Get("/", metricHandler.GetMetric)
	})

	weekdayHandler := webhandler.NewWeekdayHandler(taskRepository)
	r.Route("/weekdays", func(r chi.Router) {
		r.Get("/", weekdayHandler.GetWeekday)
	})

	go cron.ExecuteCronJobs(configs.CreateTodayRoutineTaskCron, taskRepository)
	fmt.Printf("Starting web server on port %s", configs.WebServerPort)
	http.ListenAndServe(":"+configs.WebServerPort, r)
}
