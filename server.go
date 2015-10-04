package godoauth

import "net/http"

// TODO: Implement graceful shutdown?
type Server struct {
	*http.ServeMux
}

// NewServer returns a new instance of Server built from a config.
func NewServer(authHandler *TokenAuthHandler) *Server {
	// BUG(dejan) add support to write logs to a text file
	//	if c.Log.File != "" {
	//		// BUG(dejan): Implement file handler
	//		s.Handler = handlers.CombinedLoggingHandler(os.Stdout, s.Handler)
	//	} else {
	//		s.Handler = handlers.CombinedLoggingHandler(os.Stdout, s.Handler)
	//	}
	s := &Server{
		ServeMux: http.NewServeMux(),
	}

	s.Handle("/auth", authHandler)
	s.HandleFunc("/server-ping", s.ping)
	return s
}

// ping is an health check handler, so we can use ELB/HA proxy health check
func (Server) ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Save the Whales !\"}\n\r"))
}
