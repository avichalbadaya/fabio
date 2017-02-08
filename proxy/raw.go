package proxy

import (
	"io"
	"github.com/eBay/fabio/logging"
	"net"
	"net/http"
	"net/url"

	"github.com/eBay/fabio/metrics"
)

// conn measures the number of open web socket connections
var conn = metrics.DefaultRegistry.GetCounter("ws.conn")

// newRawProxy returns an HTTP handler which forwards data between
// an incoming and outgoing TCP connection including the original request.
// This handler establishes a new outgoing connection per request.
func newRawProxy(t *url.URL) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn.Inc(1)
		defer func() { conn.Inc(-1) }()

		hj, ok := w.(http.Hijacker)
		if !ok {
			http.Error(w, "not a hijacker", http.StatusInternalServerError)
			return
		}

		in, _, err := hj.Hijack()
		if err != nil {
			logging.Error("[ERROR] Hijack error for %s. %s", r.URL, err)
			http.Error(w, "hijack error", http.StatusInternalServerError)
			return
		}
		defer in.Close()

		out, err := net.Dial("tcp", t.Host)
		if err != nil {
			logging.Error("[ERROR] WS error for %s. %s", r.URL, err)
			http.Error(w, "error contacting backend server", http.StatusInternalServerError)
			return
		}
		defer out.Close()

		err = r.Write(out)
		if err != nil {
			logging.Error("[ERROR] Error copying request for %s. %s", r.URL, err)
			http.Error(w, "error copying request", http.StatusInternalServerError)
			return
		}

		errc := make(chan error, 2)
		cp := func(dst io.Writer, src io.Reader) {
			_, err := io.Copy(dst, src)
			errc <- err
		}

		go cp(out, in)
		go cp(in, out)
		err = <-errc
		if err != nil && err != io.EOF {
			logging.Info("[INFO] WS error for %s. %s", r.URL, err)
		}
	})
}
