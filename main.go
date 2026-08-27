package main

import (
	"blog-aggregator/internal/database"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)
import (
	"blog-aggregator/internal/config"
	"log"
	//"os"
)

type state struct {
	cfg *config.JsonConfig
	db  *database.Queries
}

func main() {
	con, err := config.Read()
	if err != nil {
		log.Fatalf("error reading cfg: %v", err)
	}

	db, err := sql.Open("postgres", con.DBURL)
	if err != nil {
		log.Fatalf("error opening sql: %v", err)
	}

	dbQueries := database.New(db)
	s := state{db: dbQueries, cfg: &con}

	cmds := commands{
		CommandMap: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	err = cmds.run(&s, command{args[1], args[2:]})
	if err != nil {
		log.Fatalf("error running command:\n   %v", err)
	}

}
