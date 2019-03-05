package pmmapitests

import (
	"context"
	"crypto/tls"
	"flag"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

var Context context.Context
var ServerURLF = flag.String("pmm.server-url", "https://127.0.0.1:8443/", "PMM Server URL.")

func init() {
	debugF := flag.Bool("pmm.debug", false, "Enable debug output.")
	flag.Parse()

	if *debugF {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Debug logging enabled.")
	}

	var cancel func()
	Context, cancel = context.WithCancel(context.TODO())

	// handle termination signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-signals
		signal.Stop(signals)
		logrus.Warnf("Got %s, shutting down...", unix.SignalName(s.(syscall.Signal)))
		cancel()
	}()

	u, err := url.Parse(*ServerURLF)
	if err != nil {
		logrus.Fatalf("Failed to parse PMM Server URL: %s.", err)
	}
	if u.Host == "" || u.Scheme == "" {
		logrus.Fatalf("Invalid PMM Server URL: %s", u.String())
	}
	if u.Path == "" {
		u.Path = "/"
	}
	logrus.Debugf("PMM Server URL: %#v.", u)

	// use JSON APIs over HTTP/1.1 (setting TLSNextProto to non-nil map disables automated HTTP/2)
	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.Transport.(*http.Transport).TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}

	l := logrus.WithField("component", "client")
	transport.SetLogger(l)
	transport.Debug = *debugF

	client.Default = client.New(transport, nil)
}
