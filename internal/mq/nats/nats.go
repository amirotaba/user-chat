package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"runtime"
)

type Client struct {
	cn *nats.Conn
}

// New initializes a connection to NATS server
func New() (*Client, error) {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to NATS: %v", err)
	}
	return &Client{cn: conn}, nil
}

func (c *Client) Sub() string {
	var out string
	c.cn.Subscribe("gen", func(msg *nats.Msg) {
		out = string(msg.Data)
	})

	runtime.Goexit()
	return out
}

func (c *Client) Pub(message []byte) error {
	err := c.cn.Publish("foo", message)
	if err != nil {
		return err
	}
	return nil
}
