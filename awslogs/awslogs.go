package awslogs

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// Start ...
func Start(match, prefix, apiGateway, stage, lambda *string, start *time.Duration) {
	var result []*cloudwatchlogs.LogGroup
	var groupPrefix string

	switch {
	case *apiGateway != "":
		groupPrefix = "API-Gateway-Execution-Logs_" + *apiGateway + "/" + *stage
	case *lambda != "":
		groupPrefix = "/aws/lambda/" + *lambda
	default:
		groupPrefix = *prefix
	}

	groups := listGroups(&groupPrefix)

	if *match == "" {
		result = groups
	} else {
		for _, group := range groups {
			if ok, _ := regexp.MatchString(*match, *group.LogGroupName); ok {
				result = append(result, group)
			}
		}
	}

	switch len(result) {
	case 0:
		fmt.Println(strings.Repeat("#", 10), "no group matched", strings.Repeat("#", 10))
	case 1:
		fmt.Println(strings.Repeat("#", 10), *result[0].LogGroupName, strings.Repeat("#", 10))
		streaming(result[0].LogGroupName, start)
	default:
		fmt.Println(strings.Repeat("#", 10), len(result), "groups", strings.Repeat("#", 10))
		for _, group := range result {
			fmt.Println(*group.LogGroupName)
		}
	}
}
