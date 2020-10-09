package main

import (
    "os"
    "github.com/urfave/cli/v2"
)

func main() {
    app := cli.NewApp()

    app.Name = "Nagitter"
    app.Usage = "Nagitter is Tweet client"
    app.Version = "0.0.1"

    app.Run(os.Args)
}
