package ntfyerror

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/AnthonyHewins/gotfy"
)

type NtfyServer struct {
	server  *url.URL
	tp      *gotfy.Publisher
	binName string
}

func New(ntfyURL string) *NtfyServer {
	server, err := url.Parse(ntfyURL)
	if err != nil {
		panic("bad url:" + err.Error())
	}

	tp, err := gotfy.NewPublisher(server, http.DefaultClient)
	if err != nil {
		panic("bad config:" + err.Error())
	}

	var binName string
	binaryName, err := os.Executable()
	if err == nil {
		binName = filepath.Base(binaryName)
	}

	return &NtfyServer{
		server:  server,
		tp:      tp,
		binName: binName,
	}
}

func (s NtfyServer) SendError(err error, tags ...string) error {
	_, errm := s.tp.SendMessage(context.Background(), &gotfy.Message{
		Topic:    "alert",
		Message:  err.Error(),
		Title:    s.binName + " error",
		Tags:     tags,
		Priority: gotfy.High,
	})

	return errm
}
