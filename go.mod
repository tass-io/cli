module github.com/tass-io/cli

go 1.15

require (
	github.com/go-redis/redis/v8 v8.8.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	github.com/tass-io/tass-operator v0.0.0-20210409032653-6fc20e8df6b0
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
	sigs.k8s.io/controller-runtime v0.8.3
)
