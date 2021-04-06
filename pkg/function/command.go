package function

import (
	"log"

	"github.com/spf13/cobra"
	cliClient "github.com/tass-io/cli/pkg/client"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
)

var (
	fnName      string
	fnNamespace string
	domain      string
	code        string
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a function",
	Long:  "Create a function",
	Run:   Create,
}

func Create(cmd *cobra.Command, args []string) {
	cf := &CreateFunction{
		name:   fnName,
		ns:     fnNamespace,
		domain: domain,
		code:   code,
		client: *cliClient.GetCRDClient(),
		fn:     &serverlessv1alpha1.Function{},
	}
	err := cf.do()
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Function " + fnName + " created.")
}

func init() {
	CreateCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of function")
	CreateCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	CreateCmd.Flags().StringVarP(&domain, "domain", "d", "default", "Domain of the function")
	CreateCmd.Flags().StringVarP(&code, "code", "c", "", "Namespace of function")
	CreateCmd.MarkFlagRequired("name")
}
