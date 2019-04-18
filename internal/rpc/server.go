package rpc

import (
    "os"
    "net"
    "net/rpc"
    "net/http"

    api "github.com/coditva/bunker/internal/api"
)


type Server struct {
    sockAddr string
}

func NewServer(sockAddr string) *Server {
    return &Server{
        sockAddr: sockAddr,
    }
}

func (server Server) Serve(api *api.Api) error {
    if err := os.RemoveAll(server.sockAddr); err != nil {
        return err
    }

    rpc.Register(api)
    rpc.HandleHTTP()

    sock, err := net.Listen("unix", server.sockAddr)
    if err != nil {
        return err
    }
    http.Serve(sock, nil)
    return nil
}

func (server Server) Close() error {
    server.Close()
    return nil
}
