package server

import (
	"fmt"
)

func Endpoint(s *Server, version string) {
	ver := fmt.Sprintf("/api/%s", version)

	v := s.srv.Group(ver)

	payment := v.Group("card")
	payment.POST("/notification", s.handler.EventNotification)
}
