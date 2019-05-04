package service

import (
	"github.com/influxdata/influxdb/client/v2"
	"graduation/internal/log"
	"graduation/pkg/modules/influxdb"
)

type LogService struct {
	client *influxdb.Client
}

func NewLogService(client *influxdb.Client) *LogService {
	return &LogService{
		client: client,
	}
}

func (l *LogService) CreateArchive(point client.Point) error {
	return l.client.Write(&point)
}

func (l *LogService) GetArchive() (archives []client.Result, err error) {
	archives, err = l.client.QueryDB("select * from graduation")
	log.Infof("%#v", archives)
	if err != nil {
		return nil, nil
	}
	return
}