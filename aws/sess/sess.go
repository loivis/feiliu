package sess

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/loivis/godis/utils"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func Sess() *session.Session {
	region := "eu-west-1"
	sess, err := session.NewSession(&aws.Config{Region: &region})
	utils.CheckError(err)
	return sess
	// svc := cloudwatchlogs.New(session)
}
