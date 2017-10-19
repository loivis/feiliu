package cwl

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/loivis/feiliu/aws/sess"
)

// Streaming ...
func Streaming(groupName string) {
	log.Println(groupName)
	var logStreams []*cloudwatchlogs.LogStream
	client := cloudwatchlogs.New(sess.Sess())
	dlsInput := &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName: &groupName,
		OrderBy:      aws.String("LastEventTime"),
		Descending:   aws.Bool(true)}
	dlsFn := func(output *cloudwatchlogs.DescribeLogStreamsOutput, hasMore bool) bool {
		for _, stream := range output.LogStreams {
			lastEventTime := time.Unix(*stream.LastEventTimestamp/1e3, 0)
			if lastEventTime.Add(time.Hour * time.Duration(2)).Before(time.Now()) {
				break
			}
			logStreams = append(logStreams, stream)
		}
		for output.NextToken != nil {
			return true
		}
		return false
	}
	client.DescribeLogStreamsPages(dlsInput, dlsFn)
	// log.Println(logStreams)
	for _, stream := range logStreams {
		log.Println(*stream.LogStreamName)
	}
	var logEvents []*cloudwatchlogs.OutputLogEvent
	gleInput := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  &groupName,
		LogStreamName: logStreams[0].LogStreamName}
	gleFn := func(output *cloudwatchlogs.GetLogEventsOutput, hasMore bool) bool {
		logEvents = append(logEvents, output.Events...)
		return false
	}
	client.GetLogEventsPages(gleInput, gleFn)
	log.Println(logEvents)
}
