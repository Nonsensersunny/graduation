package influxdb

import (
	"fmt"
	"graduation/internal/log"
	"strconv"
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

type Client struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string

	c client.Client
}

func (c *Client) Write(ps ...*client.Point) error {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database: c.DbName,
	})
	if err != nil {
		return err
	}
	for _, p := range ps {
		bp.AddPoint(p)
	}
	return c.c.Write(bp)
}

func (c *Client) close() error { return c.c.Close() }

func (c *Client) NewInfluxClient() (err error) {
	host := c.Host
	port := c.Port
	addr := "http://" + host + ":" + strconv.Itoa(port)
	dbName := c.DbName
	username := c.Username
	password := c.Password
	c.c, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	_, _, err = c.c.Ping(time.Second)
	if err != nil {
		return err
	}
	_, err = c.c.Query(client.Query{
		Command: fmt.Sprintf("create database %s", dbName),
	})
	defer c.c.Close()
	return
}

func (c * Client) QueryDB(cmd string) (res []client.Result, err error) {
	log.Infof("Query influxdb:%s", cmd)
	q := client.Query{
		Command: cmd,
		Database: c.DbName,
	}
	if resp, e := c.c.Query(q); e == nil {
		if resp.Error() != nil {
			return res, resp.Error()
		}
		res = resp.Results
	} else {
		return res, err
	}
	return res, nil
}


