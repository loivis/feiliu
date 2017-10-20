package cwl

import (
	"log"
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/loivis/feiliu/aws/sess"
)

// FullEvent ...
type FullEvent struct {
	LogGroupName string
	*cloudwatchlogs.FilteredLogEvent
}

// Streaming ...
func Streaming(groupName string) {
	var events []FullEvent
	client := cloudwatchlogs.New(sess.Sess())
	input := &cloudwatchlogs.FilterLogEventsInput{LogGroupName: aws.String(groupName),
		StartTime: aws.Int64(1508540633886),
		EndTime:   aws.Int64(1508540635888)}
	client.FilterLogEventsPages(input, func(output *cloudwatchlogs.FilterLogEventsOutput, hasMore bool) bool {
		log.Println(len(events))
		log.Println(len(output.Events))
		for _, e := range output.Events {
			var f FullEvent
			f.LogGroupName = groupName
			f.FilteredLogEvent = e
			events = append(events, f)
		}
		if output.NextToken != nil {
			return true
		}
		return false
	})
	log.Println("number of events:", len(events))
	sort.Slice(events, func(i, j int) bool { return *events[i].Timestamp < *events[j].Timestamp })
	for _, e := range events[:] {
		log.Println(*e.LogStreamName, *e.Timestamp, *e.Message)
	}
}
