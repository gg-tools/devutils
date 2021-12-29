package main

import (
	"github.com/gg-tools/devutils/commands/timestamp"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "Dev Utils"
	app.Compiled = time.Now()
	app.Usage = "Developer Utilities"
	app.UsageText = `devutils [command] [args...]`
	app.Commands = []*cli.Command{timestamp.Time, timestamp.Timestamp}
	app.Flags = timestamp.Timestamp.Flags

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
