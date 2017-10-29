package main

import (
	"os"

	"github.com/loivis/feiliu/awslogs"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("feiliu", "cli tool to stream aws cloudwatch logs").Version("0.0.1").Author("loivis")
	prefix     = app.Flag("group-prefix", "prefix of log groups").PlaceHolder("kinase").Short('g').String()
	apiGateway = app.Flag("api-gateway", "id of api gateway").PlaceHolder("k2i0n1a5se").Short('a').String()
	stage      = app.Flag("stage", "name of api gateway stage").PlaceHolder("prod").Short('s').String()
	lambda     = app.Flag("lambda", "name or prefix of lambda function").PlaceHolder("function-name").Short('l').String()
	match      = app.Flag("match", "substring to match name of log groups").PlaceHolder("something").Short('m').String()
	start      = app.Flag("start", "when to start streaming, default 1 minute").PlaceHolder("10m").Short('t').Default("1m").Duration()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	awslogs.Start(match, prefix, apiGateway, stage, lambda, start)
}
