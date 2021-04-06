package function

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	cliClient "github.com/tass-io/cli/pkg/client"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
)

var (
	fnName      string
	fnNamespace string
	code        string
)

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a function",
	Long:  "Create a function",
	Run:   Create,
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a function",
	Long:  "Delete a function",
	Run:   Delete,
}

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a function",
	Long:  "Update a function",
	Run:   Update,
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a function source code",
	Long:  "Get a function source code",
	Run:   Get,
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List functions",
	Long:  "List all functions in a namespace if specified, else, list functions in default namespace",
	Run:   List,
}

func Create(cmd *cobra.Command, args []string) {
	cf := &CreateFunction{
		name:   fnName,
		ns:     fnNamespace,
		code:   code,
		client: *cliClient.GetCRDClient(),
		fn:     &serverlessv1alpha1.Function{},
	}
	if err := cf.do(); err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Function " + fnName + " created.")
}

func Delete(cmd *cobra.Command, args []string) {
	df := &DeleteFunction{
		name:   fnName,
		ns:     fnNamespace,
		client: *cliClient.GetCRDClient(),
	}
	if err := df.do(); err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Function" + fnName + " deleted.")
}

func Update(cmd *cobra.Command, args []string) {
	uf := &UpdateFunction{
		name:   fnName,
		ns:     fnNamespace,
		client: *cliClient.GetCRDClient(),
		fn:     &serverlessv1alpha1.Function{},
	}
	if err := uf.do(); err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Function" + fnName + " updated.")
}

func Get(cmd *cobra.Command, args []string) {
	gf := &GetFunction{
		name:   fnName,
		ns:     fnNamespace,
		client: *cliClient.GetCRDClient(),
		fn:     &serverlessv1alpha1.Function{},
	}
	if err := gf.do(); err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Function" + fnName + " gotten.")
	fmt.Println(*gf.fn)
}

func List(cmd *cobra.Command, args []string) {
	lf := &ListFunctions{
		ns:     fnNamespace,
		client: *cliClient.GetCRDClient(),
		fnList: &serverlessv1alpha1.FunctionList{},
	}
	if err := lf.do(); err != nil {
		log.Fatalln(err)
		return
	}
	log.Println("Function list in " + fnNamespace + " namespace gotten.")
	fmt.Println(*lf.fnList)
}

func init() {
	// Create command
	CreateCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of function")
	CreateCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	CreateCmd.Flags().StringVarP(&code, "code", "c", "", "Namespace of function")
	CreateCmd.MarkFlagRequired("name")
	// Delete command
	DeleteCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of function")
	DeleteCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	DeleteCmd.MarkFlagRequired("name")
	// Update command
	UpdateCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of function")
	UpdateCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	UpdateCmd.MarkFlagRequired("name")
	// Get command
	GetCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of function")
	GetCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	GetCmd.MarkFlagRequired("name")
	// list command
	ListCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
}
