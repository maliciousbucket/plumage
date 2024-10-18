package chaos

import (
	"fmt"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
	"log"
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
	accountName := account
	if accountName == "" {
		accountChart := cdk8s.NewChart(app, jsii.String("account"), &cdk8s.ChartProps{
			DisableResourceNameHashes: jsii.Bool(true),
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
