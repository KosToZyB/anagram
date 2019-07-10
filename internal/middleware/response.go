package middleware

import (
	"net/http"
	"net/http/httptest"
)

type HTTPResponseLog struct {
	Method  string
	Payload string
	Header  http.Header
}

func (m *Middleware) LoggerResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := httptest.NewRecorder()

		next.ServeHTTP(c, r)

		for k, v := range c.Result().Header {
			w.Header()[k] = v
		}
		w.WriteHeader(c.Code)

		m.log.Infof("Response: %+v\n", &HTTPResponseLog{
			Method:  r.Method,
			Payload: c.Body.String(),
			Header:  w.Header(),
		})

		_, err := c.Body.WriteTo(w)
		if err != nil {
			m.log.Infof("error write data to response: %s\n", err)
		}
	})
}
