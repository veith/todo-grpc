package main

import (
	"../../internal/task"
	"github.com/prometheus/common/log"
	"golang.org/x/exp/errors/fmt"
	"strconv"

	"upper.io/db.v3/sqlite"
)

// ConnectionURL implements a SQLite connection struct.
type ConnectionURL struct {
	Database string
	Options  map[string]string
}

var settings = sqlite.ConnectionURL{
	Database: `data/task.db`, // Path to database file.
}

func main() {
	dbSession, err := sqlite.Open(settings)
	if err != nil {

		log.Fatalf("db.Open(): %q\n", err)
	}
	defer dbSession.Close() // Remember to close the database session.
	// DB session weitergeben
	task.ConnectDatabase(dbSession)

	var options task.QueryOptions

	list, meta, _ := task.ListTaskItems(options)

	fmt.Println("Page " + strconv.FormatUint(uint64(meta.CurrentPage), 10) + " of " + strconv.FormatUint(uint64(meta.LastPage), 10))
	for _, value := range list {
		fmt.Println(value)
	}

}
