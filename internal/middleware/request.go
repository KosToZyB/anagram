package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HTTPRequestLog struct {
	Method        string
	Path          string
	UserAgent     string
	RemoteAddress string
	Payload       string
	Header        http.Header
}

func (m *Middleware) LoggerRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer next.ServeHTTP(w, r)

		buf := new(bytes.Buffer)
		if r.Body != nil {

			bodyBuf, err := ioutil.ReadAll(r.Body)
			if err != nil {
				m.log.Errorf("error read body: %s\n", err)
				return
			}

			temp1 := ioutil.NopCloser(bytes.NewBuffer(bodyBuf))
			temp2 := ioutil.NopCloser(bytes.NewBuffer(bodyBuf))

			_, err = buf.ReadFrom(temp1)
			if err != nil {
				m.log.Errorf("error read buffer: %s\n", err)
				return
			}
			r.Body = temp2
		}

		m.log.Infof("Request: %+v\n", &HTTPRequestLog{
			Method:        r.Method,
			Path:          r.URL.Path,
			UserAgent:     r.UserAgent(),
			RemoteAddress: r.RemoteAddr,
			Payload:       buf.String(),
			Header:        r.Header,
		})
	})
}
