package gyperhttp

import (
	"github.com/valyala/fasthttp"
)

type Server struct {
	fasthttp.Server
}

func (s *Server) ListenAndServe(addr string) error {
	ln, err := Listen("tcp4", addr)
	if err != nil {
		return err
	}
	return s.Serve(ln)
}

func (s *Server) ListenAndServeTLS(addr string, certFile string, keyFile string) error {
	ln, err := Listen("tcp4", addr)
	if err != nil {
		return err
	}

	return s.ServeTLS(ln, certFile, keyFile)
}

func (s *Server) ListenAndServeTLSEmbed(addr string, certData []byte, keyData []byte) error {
	ln, err := Listen("tcp4", addr)
	if err != nil {
		return err
	}

	return s.ServeTLSEmbed(ln, certData, keyData)
}
