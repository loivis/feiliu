package awslogs

import (
	"fmt"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// FullEvent ...
type FullEvent struct {
	LogGroupName *string
	*cloudwatchlogs.FilteredLogEvent
}

func streaming(group string) {
	eventsChan := make(chan *FullEvent)
	go fetchEvents(group, eventsChan)
	for {
		event := <-eventsChan
		fmt.Println(*event.Message)
	}
}

func fetchEvents(group string, c chan<- *FullEvent) {
	client := client()
	startTime := time.Now().Add(time.Duration(-1*time.Minute)).UnixNano() / 1e6
	endTime := time.Now().Add(time.Duration(-10*time.Second)).UnixNano() / 1e6
	input := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: aws.String(group),
		StartTime:    aws.Int64(startTime),
		EndTime:      aws.Int64(endTime)}
	for {
		var events []FullEvent
		// fmt.Println(startTime, endTime)
		input.StartTime = aws.Int64(startTime)
		input.EndTime = aws.Int64(endTime)
		client.FilterLogEventsPages(input, func(output *cloudwatchlogs.FilterLogEventsOutput, hasMore bool) bool {
			for _, e := range output.Events {
				var f FullEvent
				f.LogGroupName = aws.String(group)
				f.FilteredLogEvent = e
				events = append(events, f)
			}
			if output.NextToken != nil {
				return true
			}
			return false
		})
		// fmt.Println("number of events:", len(events))
		if count := len(events); count > 0 {
			sort.Slice(events, func(i, j int) bool { return *events[i].Timestamp < *events[j].Timestamp })
			sleepTime := int(1 * 1e3 / count)
			for _, e := range events[:] {
				// fmt.Println(*e.LogStreamName, *e.Timestamp, *e.Message)
				c <- &e
				time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			}
		} else {
			if endTime-startTime > 10000 {
				fmt.Println("no log event in history")
			}
			time.Sleep(time.Duration(1 * time.Second))
		}

		startTime = endTime
		endTime = time.Now().Add(time.Duration(-10*time.Second)).UnixNano() / 1e6
	}
}
