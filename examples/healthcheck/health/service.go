package health

import (
	"github.com/upitau/goinbigdata/examples/healthcheck/mongo"
	"github.com/upitau/goinbigdata/examples/healthcheck/elastic"
	"math/rand"
	"fmt"
)

type HealthStatus struct {
	Nodes   map[string]string `json:"nodes"`
	Mongo   string `json:"mongo"`
	Elastic string `json:"elastic"`
}

type Service interface {
	Health() HealthStatus
}

type service struct {
	nodes   []string
	mongo   mongo.Service
	elastic elastic.Service
}

func New(nodes []string, mongo mongo.Service, elastic elastic.Service) Service {
	return &service{
		nodes: nodes,
		mongo: mongo,
		elastic: elastic,
	}
}

func (s *service) Health() HealthStatus {
	nodesStatus := make(map[string]string)
	for _, n := range s.nodes {
		if rand.Intn(10) > 7 {
			nodesStatus[n] = "Node ERROR: Node not responding"
		} else {
			nodesStatus[n] = "OK"
		}
	}

	mongoStatus := "OK"
	if err := s.mongo.Health(); err != nil {
		mongoStatus = fmt.Sprintf("Mongo ERROR: %s", err)
	}

	elasticStatus := "OK"
	if err := s.elastic.Health(); err != nil {
		elasticStatus = fmt.Sprintf("Elastic ERROR: %s", err)
	}

	return HealthStatus{
		Nodes: nodesStatus,
		Mongo: mongoStatus,
		Elastic: elasticStatus,
	}
}
