package start

import "github.com/VandiKond/user-http-go/Backend/http_service"

func StartAll() error {
	http_service.Start()
	return nil
}
