package cmd

import (
	"context"
	"fmt"
	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/spf13/cobra"
	"log"
)

// testCmd represents the test command
//var testCmd = &cobra.Command{
//	Use:   "test",
//	Short: "A brief description of your command",
//	Long: `A longer description that spans multiple lines and likely contains examples
//and usage of using your command. For example:
//
//Cobra is a CLI library for Go that empowers applications.
//This application is a tool to generate the needed files
//to quickly create a Cobra application.`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("test called")
//	},
//}
//
//func init() {
//	rootCmd.AddCommand(testCmd)
//
//	// Here you will define your flags and configuration settings.
//
//	// Cobra supports Persistent Flags which will work for this command
//	// and all subcommands, e.g.:
//	// testCmd.PersistentFlags().String("foo", "", "A help for foo")
//
//	// Cobra supports local flags which will only run when this command
//	// is called directly, e.g.:
//	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
//}

func testComand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test",
		Short: "A test command",
		Long:  `A test command - but long`,
		Run: func(cmd *cobra.Command, args []string) {
			loadTestConfig()
		},
	}
	return cmd
}

func loadTestConfig() {
	composeFilePath := "testdata/compose/aks/docker-compose.yml"
	projectName := "testproject"
	ctx := context.Background()

	options, err := cli.NewProjectOptions(
		[]string{composeFilePath},
		cli.WithName(projectName),
	)

	if err != nil {
		log.Fatal(err)
	}

	project, err := options.LoadProject(ctx)
	if err != nil {
		log.Fatal(err)
	}

	projectYaml, err := project.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(projectYaml))

	fmt.Println(project.ServiceNames())
}
