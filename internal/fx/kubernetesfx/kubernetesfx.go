package kubernetesfx

import (
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/yagehu/reactor/config"
)

var Module = fx.Provide(New)

type Params struct {
	fx.In

	Config config.Config
	Logger *zap.Logger
}

type Result struct {
	fx.Out

	Client kubernetes.Interface
}

func New(p Params) (Result, error) {
	conf := rest.Config{
		Host: "http://" + net.JoinHostPort(
			p.Config.Kubernetes.APIServer.Host,
			p.Config.Kubernetes.APIServer.Port,
		),
	}

	clientset, err := kubernetes.NewForConfig(&conf)
	if err != nil {
		return Result{}, err
	}

	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return Result{}, err
	}

	p.Logger.Info(
		"Kubernetes server discovered.",
		zap.String("version", version.String()),
	)

	return Result{
		Client: clientset,
	}, nil
}
