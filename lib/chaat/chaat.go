package chaat

import (
	"log"

	"github.com/zeromq/goczmq"
)

// ServeConfig holds all of the configuration needed to run a chaat
// server.
type ServeConfig struct {
	ReqURL string
	PubURL string
}

// Serve runs a chaat server.
func Serve(c *ServeConfig) error {
	req, err := goczmq.NewRep(c.ReqURL)
	if err != nil {
		return err
	}
	defer req.Destroy()

	pub, err := goczmq.NewPub(c.PubURL)
	if err != nil {
		return err
	}
	defer pub.Destroy()

	return nil
}

// ChatConfig holds all of the configuration needed to run a chaat
// server.
type ChatConfig struct {
	URL string
}

// Chat runs a chaat server.
func Chat(c *ChatConfig) error {
	log.Println("Client")
	return nil
}
