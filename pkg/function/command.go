package function

import (
	"github.com/tass-io/cli/pkg/logging"

	"github.com/spf13/cobra"
	cliClient "github.com/tass-io/cli/pkg/client"
	serverlessv1alpha1 "github.com/tass-io/tass-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

var (
	fnName      string
	fnNamespace string
	code        string
	mockMode    bool
)

var client = *cliClient.GetCRDClient()

var log = logging.Log

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a function",
	Long:  "Create a function",
	Run:   Create,
}

var DeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "remove", "rm"},
	Short:   "Delete a function",
	Long:    "Delete a function",
	Run:     Delete,
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

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "Test a function",
	Long:  "Test a function",
	Run:   Test,
}

func Create(cmd *cobra.Command, args []string) {
	cf := &CreateFunction{
		name: fnName,
		ns:   fnNamespace,
		code: code,
		fn:   &serverlessv1alpha1.Function{},
		mock: mockMode,
	}
	if err := cf.do(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Function " + fnName + " created.")
}

func Delete(cmd *cobra.Command, args []string) {
	df := &DeleteFunction{
		name: fnName,
		ns:   fnNamespace,
	}
	if err := df.do(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Function" + fnName + " deleted.")
}

func Update(cmd *cobra.Command, args []string) {
	uf := &UpdateFunction{
		name: fnName,
		ns:   fnNamespace,
		fn:   &serverlessv1alpha1.Function{},
	}
	if err := uf.do(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Function" + fnName + " updated.")
}

func Get(cmd *cobra.Command, args []string) {
	gf := &GetFunction{
		name: fnName,
		ns:   fnNamespace,
		fn:   &serverlessv1alpha1.Function{},
	}
	if err := gf.do(); err != nil {
		log.Error(err)
		return
	}
}

func List(cmd *cobra.Command, args []string) {
	lf := &ListFunctions{
		ns:     fnNamespace,
		fnList: &serverlessv1alpha1.FunctionList{},
	}
	if err := lf.do(); err != nil {
		log.Error(err)
	}
}

func Test(cmd *cobra.Command, args []string) {
	tf := &TestFunction{
		name: fnName,
		ns:   fnNamespace,
		fn:   &serverlessv1alpha1.Function{},
		svc:  &corev1.Service{},
	}
	if err := tf.do(); err != nil {
		log.Error(err)
	}
}

func init() {
	// Create command
	CreateCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of the function")
	CreateCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	CreateCmd.Flags().StringVarP(&code, "code", "c", "", "Location of function code")
	CreateCmd.Flags().BoolVarP(&mockMode, "mock", "m", false, "Whether use mock mode")
	CreateCmd.MarkFlagRequired("name")
	// Delete command
	DeleteCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of the function")
	DeleteCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	DeleteCmd.MarkFlagRequired("name")
	// Update command
	UpdateCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of the function")
	UpdateCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	UpdateCmd.Flags().StringVarP(&code, "code", "c", "", "Location of function code")
	UpdateCmd.MarkFlagRequired("name")
	UpdateCmd.MarkFlagRequired("code")
	// Get command
	GetCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of the function")
	GetCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	GetCmd.MarkFlagRequired("name")
	// list command
	ListCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	// Test command
	TestCmd.Flags().StringVarP(&fnName, "name", "n", "", "Name of the function")
	TestCmd.Flags().StringVarP(&fnNamespace, "ns", "", "default", "Namespace of the function")
	TestCmd.MarkFlagRequired("name")
}
