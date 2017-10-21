package main

import (
	"flag"
	"log"

	"github.com/loivis/feiliu/awslogs"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

var (
	logGroupName string
)

func parseFlags() {
	flag.StringVar(&logGroupName, "g", "", "name or prefix of log group")
	flag.Parse()
}

func main() {
	parseFlags()
	awslogs.Run(logGroupName)
}
