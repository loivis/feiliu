package main

import "github.com/loivis/feiliu/aws"

func main() {
	aws.Run("/var/log/messages")
	// aws.Run("cassandra")
}
