package awslogs

import (
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
