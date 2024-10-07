package chaos

const (
	envFlag = "-e"
	outFlag = "-o"
	runCmd  = "k6 run"
)

func stringSliceToK8s(source []string) []*string {
	res := make([]*string, len(source))
	for i := range source {
		res[i] = &source[i]
	}
	return res
}
