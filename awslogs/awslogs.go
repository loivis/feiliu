package awslogs

// Run ...
func Run(group string) {
	validateGroup(group)
	streaming(group)
}
