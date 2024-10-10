package chaos

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"log"
)

const (
	envFlag = "-e"
	outFlag = "-o"
	runCmd  = "k6 run"
)

func SynthTemplateFile(file, outDir, ns, alloy string, account string) error {
	if outDir == "" {
		return fmt.Errorf("no output dir specified")
	}
	outPut := fmt.Sprintf("%s/%s", outDir, "tests")
	template, err := loadTemplate(file, ns)
	if err != nil {
		return err
	}

	if template.Scripts == nil || len(template.Scripts) == 0 {
		return fmt.Errorf("no test configurations found in template")
	}

	app := cdk8s.NewApp(&cdk8s.AppProps{
		Outdir:         jsii.String(outPut),
		YamlOutputType: cdk8s.YamlOutputType_FOLDER_PER_CHART_FILE_PER_RESOURCE,
	})
	var accountName string
	if account == "" {
		accountChart := cdk8s.NewChart(app, jsii.String("id"), &cdk8s.ChartProps{
			DisableResourceNameHashes: jsii.Bool(true),
			Namespace:                 jsii.String(ns),
		})
		_, newAccount := NewTestRunRBAC(accountChart, ns)
		accountName = newAccount
	}
	template.ServiceAccount = accountName

	for _, test := range template.Scripts {
		_, err = NewTestRunFromTemplate(app, test.Name, ns, alloy, template, &test)
		if err != nil {
			return err
		}
		log.Println("Synthed files for: ", test.Name)
	}
	app.Synth()
	log.Println("Success! Files can be found at ", outPut)
	return nil

}

func stringSliceToK8s(source []string) []*string {
	res := make([]*string, len(source))
	for i := range source {
		res[i] = &source[i]
	}
	return res
}
