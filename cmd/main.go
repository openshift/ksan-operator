/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	ksanv1alpha1 "openshift/ksan-operator/api/v1alpha1"
	"openshift/ksan-operator/cmd/nodedaemon"
	"openshift/ksan-operator/cmd/operator"
	// +kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(ksanv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

func main() {
	log := ctrl.Log.WithName("setup")
	if err := NewCmd(log).Execute(); err != nil {
		log.Error(err, "fatal error encountered")
		os.Exit(1)
	}
}

func NewCmd(setupLog logr.Logger) *cobra.Command {

	zapOpts := zap.Options{}
	zapFlagSet := flag.NewFlagSet("zap", flag.ExitOnError)
	zapOpts.BindFlags(zapFlagSet)

	klogFlagSet := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(klogFlagSet)

	cmd := &cobra.Command{
		Use:           "ksan",
		Short:         "Commands for running ksan storage",
		Long:          "Contains commands that control various components reconciling of the main cluster resources within ksan storage",
		SilenceErrors: false,
		SilenceUsage:  true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			zapLogger := zap.New(zap.UseFlagOptions(&zapOpts))
			ctrl.SetLogger(zapLogger)
			klog.SetLogger(zapLogger)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	cmd.PersistentFlags().AddGoFlagSet(klogFlagSet)
	cmd.PersistentFlags().AddGoFlagSet(zapFlagSet)

	cmd.AddCommand(
		operator.NewCmd(scheme),
		nodedaemon.NewCmd(scheme),
	)

	return cmd
}
