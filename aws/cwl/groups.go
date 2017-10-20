package cwl

import (
	"log"

	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/loivis/feiliu/aws/sess"
	"github.com/loivis/feiliu/utils"
)

// LogGroups ...
func LogGroups() {
	client := cloudwatchlogs.New(sess.Sess())
	input := &cloudwatchlogs.DescribeLogGroupsInput{}
	output, err := client.DescribeLogGroups(input)
	logGroups := output.LogGroups
	utils.CheckError(err)
	for output.NextToken != nil {
		input.SetNextToken(*output.NextToken)
		output, err = client.DescribeLogGroups(input)
		utils.CheckError(err)
		logGroups = append(logGroups, output.LogGroups...)
	}
	log.Println("number of log groups:", len(logGroups))
	// fmt.Println(logGroups)
}

// LogGroupsPages ...
func LogGroupsPages() {
	var logGroups []*cloudwatchlogs.LogGroup
	client := cloudwatchlogs.New(sess.Sess())
	input := &cloudwatchlogs.DescribeLogGroupsInput{}
	fn := func(output *cloudwatchlogs.DescribeLogGroupsOutput, morePages bool) bool {
		logGroups = append(logGroups, output.LogGroups...)
		for output.NextToken != nil {
			return true
		}
		return false
	}
	client.DescribeLogGroupsPages(input, fn)
	log.Println("number of log groups:", len(logGroups))
}
