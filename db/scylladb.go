package db

import (
	"fmt"
	"golang-scyllacloud/log"

	"github.com/gocql/gocql"
	"github.com/gocql/gocql/scyllacloud"
	"go.uber.org/zap"
)

var Session *gocql.Session

const (
	connectionBundlePath = "/Users/preethamu/Desktop/code/golang/scyllacloudgo/db/conn.yaml"
)

func Scyllaconnect() {
	logger := log.CreateLogger("info")

	cluster, err := scyllacloud.NewCloudCluster(connectionBundlePath)
	if err != nil {
		logger.Error("Failed to create cloud cluster config: %s", zap.Error(err))
	}

	// cluster.NumConns = 10
	// cluster.Timeout = 5000

	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("us-east-1")

	session, err := cluster.CreateSession()
	if err != nil {
		logger.Error("Failed to connect to cluster: %s", zap.Error(err))
	}

	Session = session
	fmt.Println("printing session %s ", Session)
	//Session = session
	//defer session.Close()
	//return Session
}
