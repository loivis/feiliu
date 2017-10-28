package awslogs

import (
	"time"
)

// List ...
func List(prefix, apiGateway, stage, lambda, match *string) {
	if *match != "" {
		matchGroups(match)
		return
	}

	var groupPrefix string
	if *apiGateway != "" {
		groupPrefix = "API-Gateway-Execution-Logs_" + *apiGateway + "/" + *stage
	} else if *lambda != "" {
		groupPrefix = "/aws/lambda/" + *lambda
	} else if *prefix != "" {
		groupPrefix = *prefix
	}
	prefixGroups(&groupPrefix)
}

// Stream ...
func Stream(name *string, start *time.Duration) {
	validateGroup(name)
	streaming(name, start)
}
