module golang-scyllacloud

go 1.21.3

require (
	github.com/gocql/gocql v1.6.0
	go.uber.org/zap v1.26.0
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.7.3

require (
	github.com/golang/snappy v0.0.3 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)
