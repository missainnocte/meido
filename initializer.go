package main

import (
	"fmt"
	"github.com/ivanh/meido/persistence"
	"github.com/ivanh/meido/server"
	log "github.com/sirupsen/logrus"
	"os"
)

func run() {
	// Do work here
	fmt.Print(`
███╗░░░███╗███████╗██╗██████╗░░█████╗░
████╗░████║██╔════╝██║██╔══██╗██╔══██╗
██╔████╔██║█████╗░░██║██║░░██║██║░░██║
██║╚██╔╝██║██╔══╝░░██║██║░░██║██║░░██║
██║░╚═╝░██║███████╗██║██████╔╝╚█████╔╝
╚═╝░░░░░╚═╝╚══════╝╚═╝╚═════╝░░╚════╝░ system is launching...
`)
}

type StdFormatter struct{}

func (f StdFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%v[%v] %v\n", entry.Time.Local(), entry.Level, entry.Message)), nil
}

func initLogger() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(new(StdFormatter))
}

func initServer(err chan error) {
	server.Init(err)
}

func main() {
	persistence.Init()
	defer persistence.Close()
	//log.Info(SQL_DRIVER)
	err := make(chan error)
	initLogger()
	run()
	initServer(err)
	log.Error(<-err)
}
