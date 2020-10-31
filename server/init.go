package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

//func routes() []Route {
//	return []Route{
//		{
//			"/",
//			func(res http.ResponseWriter, req *http.Request) {
//				log.Info(log.Fields{
//					"remote_addr": req.RemoteAddr,
//					"url":         req.URL,
//				})
//				_, err := res.Write([]byte("hello world!"))
//				if err != nil {
//					log.Error(err)
//				}
//			},
//		},
//	}
//}

func Init(err chan error) {
	mux := http.NewServeMux()
	for _, rt := range routes {
		log.Info(log.Fields{
			"action": "mount route on server",
			"path":   rt.path,
		})
	}
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		for _, rt := range routes {
			if strings.HasPrefix(request.URL.Path, rt.path) {
				log.Info(log.Fields{
					"Url":        request.URL,
					"Method":     request.Method,
					"RemoteAddr": request.RemoteAddr,
				})
				rt.handlerFunc(writer, request)
			}
		}
	})

	server := http.Server{
		Addr:    ":9000",
		Handler: mux,
	}
	go func(err chan error) {
		err <- server.ListenAndServe()
	}(err)
}
