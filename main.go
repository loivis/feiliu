package main

import (
	"os"

	"github.com/loivis/feiliu/awslogs"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("feiliu", "cli tool to stream aws cloudwatch logs")
	prefix     = app.Flag("prefix", "prefix of log groups").PlaceHolder("prod-nginx").Short('p').String()
	apiGateway = app.Flag("api-gateway", "id of api gateway").PlaceHolder("k2i0n1a5se").Short('g').String()
	stage      = app.Flag("stage", "name of api gateway stage").PlaceHolder("prod").Short('s').String()
	lambda     = app.Flag("lambda", "name of lambda function").PlaceHolder("function-name").Short('l').String()

	list  = app.Command("list", "list cloudwatch log groups")
	match = list.Flag("match", "substring to match name of log groups").PlaceHolder("something").Short('m').String()

	stream = app.Command("stream", "stream cloudwatch logs")
	name   = stream.Flag("name", "name of the log group").Required().Short('n').String()
	start  = stream.Flag("start", "when to start streaming, default 1 minute").PlaceHolder("10m").Short('t').Default("1m").Duration()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case list.FullCommand():
		awslogs.List(prefix, apiGateway, stage, lambda, match)
	case stream.FullCommand():
		awslogs.Stream(name, start)
	}
}
