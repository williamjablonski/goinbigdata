package elastictest

import (
	"gopkg.in/olivere/elastic.v3"
	"reflect"
)

type Service interface {
	GetLog(app string, lines int) ([]string, error)
}

func NewService(url string) (Service, error) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return &service{elasticClient: client}, nil
}

type service struct {
	elasticClient *elastic.Client
}

type Log struct {
	Message string `json:"message"`
}

// GetLog returns limited tail of log sorted by time in ascending order
func (s *service) GetLog(app string, limit int) ([]string, error) {
	termQuery := elastic.NewTermQuery("app", app)

	res, err := s.elasticClient.Search("_all").
		Query(termQuery).
		Sort("@timestamp", false).
		Size(limit).
		Do()

	if err != nil {
		return nil, err
	}

	msgNum := len(res.Hits.Hits)
	if msgNum == 0 {
		return []string{}, nil
	}

	messages := make([]string, msgNum, msgNum)

	var l Log
	for i, item := range res.Each(reflect.TypeOf(l)) {
		l := item.(Log)
		messages[i] = l.Message
	}

	// Reversing messages
	for i := 0; i < msgNum/2; i++ {
		messages[i], messages[msgNum-(i+1)] = messages[msgNum-(i+1)], messages[i]
	}

	return messages, nil
}
