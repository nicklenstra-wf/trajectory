package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyddonb/trajectory/db"
	"github.com/lyddonb/trajectory/pipe"
	"github.com/lyddonb/trajectory/rest"
)

const (
	TASK_PREFIX = "/api/tasks/"
	STAT_PREFIX = "/api/stats/"
)

func setupTasks(pool db.DBPool) {
	router := rest.SetupTaskRouter(pool, TASK_PREFIX)

	http.Handle(TASK_PREFIX, router)
}

func setupStats(pool db.DBPool) {
	router := rest.SetupStatRouter(pool, STAT_PREFIX)

	http.Handle(STAT_PREFIX, router)
}

func setupWeb() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/")))

	http.Handle("/", router)
}

func main() {
	// Stand up redis pool.
	pool := db.StartDB("127.0.0.1:6379", "")

	go func() {
		listener := pipe.MakeConnection("tcp", ":1300")
		taskPipeline := pipe.NewTaskPipeline(pool)
		pipe.Listen(listener, taskPipeline)
	}()

	setupStats(pool)
	setupTasks(pool)
	setupWeb()

	http.ListenAndServe(":3000", nil)
}
