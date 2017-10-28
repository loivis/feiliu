package awslogs

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func listGroups(prefix *string) []*cloudwatchlogs.LogGroup {
	var logGroups []*cloudwatchlogs.LogGroup
	client := client()
	input := &cloudwatchlogs.DescribeLogGroupsInput{}
	if *prefix != "" {
		input.SetLogGroupNamePrefix(*prefix)
	}
	fn := func(output *cloudwatchlogs.DescribeLogGroupsOutput, hasMore bool) bool {
		logGroups = append(logGroups, output.LogGroups...)
		for output.NextToken != nil {
			return true
		}
		return false
	}
	client.DescribeLogGroupsPages(input, fn)
	return logGroups
}

func matchGroups(match *string) {
	groups := listGroups(aws.String(""))
	fmt.Println(strings.Repeat("#", 10), len(groups), "groups", strings.Repeat("#", 10))
	for _, group := range groups {
		if ok, _ := regexp.MatchString(*match, *group.LogGroupName); ok {
			fmt.Println(*group.LogGroupName)
		}
	}
}
func prefixGroups(prefix *string) {
	groups := listGroups(prefix)
	fmt.Println(strings.Repeat("#", 10), len(groups), "groups", strings.Repeat("#", 10))
	for _, group := range groups {
		fmt.Println(*group.LogGroupName)
	}
}
