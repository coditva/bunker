package rpc

import (
    "os"
    "net"
    "net/rpc"
    "net/http"
)


type Server struct {
    sockAddr    string
    sock        net.Listener
}

func NewServer(sockAddr string) *Server {
    return &Server{
        sockAddr:   sockAddr,
        sock:       nil,
    }
}

func (server *Server) Serve(api interface{}) error {
    if err := os.RemoveAll(server.sockAddr); err != nil {
        return err
    }

    rpc.Register(api)
    rpc.HandleHTTP()

    var err error
    server.sock, err = net.Listen("unix", server.sockAddr)
    if err != nil {
        return err
    }
    http.Serve(server.sock, nil)
    return nil
}

func (server *Server) Close() error {
    server.sock.Close()
    return nil
}
