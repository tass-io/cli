package route

import (
	"github.com/tass-io/cli/pkg/logging"

	"github.com/spf13/cobra"
	cliClient "github.com/tass-io/cli/pkg/client"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
)

var (
	wfName      string
	wfNamespace string
	path        string
)

var client = *cliClient.GetNetClinet()

var log = logging.Log

var CreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"open", "op", "cret"},
	Short:   "Create a route for a workflow",
	Long:    "Create a route for a workflow specified by workflow wfn and namespace. For now one route can only match one workflow. Has alias of 'open', 'op' and 'cret'",
	Run:     Create,
}

var DeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "remove", "rm"},
	Short:   "Delete the route of a workflow",
	Long:    "Delete the only route of a workflow specified by wfn and namespace of the workflow. For now one route can only match one workflow. Has alias of 'rm', 'del', 'remove'",
	Run:     Delete,
}

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the route of a workflow",
	Long:  "Update the only route of a workflow specified by wfn and namespace of the workflow. For now one route can only match one workflow.",
	Run:   Update,
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all routes",
	Long:  "List all routes in a namespace if specified, else, list routes in default namespace",
	Run:   List,
}

func Create(cmd *cobra.Command, args []string) {
	cr := &CreateRoute{
		wfn:  wfName,
		ns:   wfNamespace,
		path: path,
		ig:   &networkingv1beta1.Ingress{},
	}
	if err := cr.do(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Route " + path + " for workflow " + wfName + " created.")
}

func Delete(cmd *cobra.Command, args []string) {
	dr := &DeleteRoute{
		wfn: wfName,
		ns:  wfNamespace,
	}
	if err := dr.do(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Route" + wfName + " deleted.")
}

func Update(cmd *cobra.Command, args []string) {
	ur := &UpdateRoute{
		wfn: wfName,
		ns:  wfNamespace,
		ig:  &networkingv1beta1.Ingress{},
	}
	if err := ur.do(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Route" + wfName + " updated.")
}

func List(cmd *cobra.Command, args []string) {
	lr := &ListRoutes{
		ns:     wfNamespace,
		igList: &networkingv1beta1.IngressList{},
	}
	if err := lr.do(); err != nil {
		log.Error(err)
	}
}

func init() {
	// Create command
	CreateCmd.Flags().StringVarP(&wfName, "name", "n", "", "Name of the workflow")
	CreateCmd.Flags().StringVarP(&wfNamespace, "ns", "", "default", "Namespace of the workflow")
	CreateCmd.Flags().StringVarP(&path, "path", "p", "", "Path of this route")
	CreateCmd.MarkFlagRequired("name")
	CreateCmd.MarkFlagRequired("path")
	// Delete command
	DeleteCmd.Flags().StringVarP(&wfName, "name", "n", "", "Name of the workflow")
	DeleteCmd.Flags().StringVarP(&wfNamespace, "ns", "", "default", "Namespace of the workflow")
	DeleteCmd.MarkFlagRequired("name")
	// Update command
	UpdateCmd.Flags().StringVarP(&wfName, "name", "n", "", "Name of the workflow")
	UpdateCmd.Flags().StringVarP(&wfNamespace, "ns", "", "default", "Namespace of the workflow")
	UpdateCmd.Flags().StringVarP(&path, "path", "p", "", "Path of this route")
	UpdateCmd.MarkFlagRequired("name")
	UpdateCmd.MarkFlagRequired("path")
	// list command
	ListCmd.Flags().StringVarP(&wfNamespace, "ns", "", "default", "Namespace of the route")
}
