package awslogs

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// FullEvent ...
type FullEvent struct {
	LogGroupName *string
	*cloudwatchlogs.FilteredLogEvent
}

func validateGroup(group *string) {
	client := client()
	input := &cloudwatchlogs.DescribeLogGroupsInput{LogGroupNamePrefix: group}
	output, err := client.DescribeLogGroups(input)
	if err != nil {
		fmt.Println(err)
	}
	if count := len(output.LogGroups); count == 0 {
		fmt.Println("log group doesn't exist:", *group)
		listGroups(aws.String(""))
	} else if count > 1 {
		for i := range output.LogGroups {
			if *output.LogGroups[i].LogGroupName == *group {
				return
			}
		}
		prefixGroups(group)
		fmt.Println(strings.Repeat("#", 10), "exiting...", strings.Repeat("#", 10))
		os.Exit(0)
	}
}

func streaming(group *string, start *time.Duration) {
	eventsChan := make(chan *FullEvent)
	go fetchEvents(group, start, eventsChan)
	for {
		event := <-eventsChan
		fmt.Println(*event.Message)
	}
}

func fetchEvents(group *string, start *time.Duration, c chan<- *FullEvent) {
	client := client()
	duration := 1 * time.Minute
	if *start != duration {
		duration = *start
	}
	startTime := time.Now().Add(-duration).UnixNano() / 1e6
	endTime := time.Now().Add(-10*time.Second).UnixNano() / 1e6
	input := &cloudwatchlogs.FilterLogEventsInput{
		LogGroupName: group,
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
				f.LogGroupName = group
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
			sleepTime := 1 * 1e3 / count
			for _, e := range events[:] {
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
		endTime = time.Now().Add(-10*time.Second).UnixNano() / 1e6
	}
}
