package awslogs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/loivis/feiliu/utils"
)

func validateGroup(group string) {
	client := client()
	input := &cloudwatchlogs.DescribeLogGroupsInput{LogGroupNamePrefix: aws.String(group)}
	output, err := client.DescribeLogGroups(input)
	utils.CheckError(err)
	if count := len(output.LogGroups); count == 0 {
		fmt.Println("log group doesn't exist:", group)
		listGroups("")
	} else if count > 1 {
		for i := range output.LogGroups {
			if *output.LogGroups[i].LogGroupName == group {
				return
			}
		}
		fmt.Println("more groups found for prefix:", group)
		listGroups(group)
	}
}

func listGroups(prefix string) {
	var logGroups []*cloudwatchlogs.LogGroup
	client := client()
	input := &cloudwatchlogs.DescribeLogGroupsInput{}
	if prefix != "" {
		input.SetLogGroupNamePrefix(prefix)
	}
	fn := func(output *cloudwatchlogs.DescribeLogGroupsOutput, hasMore bool) bool {
		logGroups = append(logGroups, output.LogGroups...)
		for output.NextToken != nil {
			return true
		}
		return false
	}
	client.DescribeLogGroupsPages(input, fn)
	fmt.Println("number of log groups:", len(logGroups))
	for _, group := range logGroups {
		fmt.Println(*group.LogGroupName)
	}
	os.Exit(0)
}
