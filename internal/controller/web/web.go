package web

import (
	"github.com/dmRusakov/monkeysmoon-admin/pkg/logrus"
)

type webServer struct {
	log logrus.Logrus
}

func newWebRouter(log logrus.Logrus) *webServer {
	return &webServer{
		log: log,
	}
}
