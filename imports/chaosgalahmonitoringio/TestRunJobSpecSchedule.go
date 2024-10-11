package chaosgalahmonitoringio

type TestRunJobSpecSchedule struct {
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	DayOfWeek  *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	Hour       *string `field:"optional" json:"hour" yaml:"hour"`
	Minute     *string `field:"optional" json:"minute" yaml:"minute"`
	Month      *string `field:"optional" json:"month" yaml:"month"`
}
