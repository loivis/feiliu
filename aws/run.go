package aws

import (
	"github.com/loivis/feiliu/aws/cwl"
)

// Run ...
func Run(groupName string) {
	// cwl.LogGroups()
	// cwl.LogGroupsPages()
	cwl.Streaming(groupName)
}
