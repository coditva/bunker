package rpc

import (
    "net/rpc"
)

type Client struct {
    sockAddr    string
    sock        *rpc.Client
}

func NewClient(sockAddr string) *Client {
    return &Client{
        sockAddr:   sockAddr,
        sock:       nil,
    }
}

func (client *Client) Connect() error {
    var err error
    client.sock, err = rpc.DialHTTP("unix", client.sockAddr)
    return err
}

func (client *Client) Call(meth string, args interface{}, reply interface{}) error {
    err := client.sock.Call(meth, args, reply);
    return err
}

func (client *Client) Close() error {
    client.sock.Close()
    return nil
}
