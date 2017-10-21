package awslogs

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

func client() *cloudwatchlogs.CloudWatchLogs {
	// region := "eu-west-1"
	// config := &aws.Config{Region: aws.String(region)}
	// sess := session.Must(session.NewSession(config))
	options := session.Options{SharedConfigState: session.SharedConfigEnable}
	sess := session.Must(session.NewSessionWithOptions(options))
	return cloudwatchlogs.New(sess)
}
