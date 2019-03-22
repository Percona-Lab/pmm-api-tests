package pmmapitests

import (
	"context"
	"crypto/tls"
	"flag"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/percona/pmm/api/inventory/json/client"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

//nolint:gochecknoglobals
var (
	// Context is canceled on SIGTERM or SIGINT. Tests should cleanup and exit.
	Context context.Context

	// BaseURL contains PMM Server base URL like https://127.0.0.1:8443/.
	BaseURL *url.URL

	// Hostname contains local hostname that is used for generating test data.
	Hostname string
)

//nolint:gochecknoinits
func init() {
	debugF := flag.Bool("pmm.debug", false, "Enable debug output.")
	traceF := flag.Bool("pmm.trace", false, "Enable trace output.")
	serverURLF := flag.String("pmm.server-url", "https://127.0.0.1:8443/", "PMM Server URL.")
	flag.Parse()

	flag.VisitAll(func(i *flag.Flag) {
		envVar := strings.Replace(strings.ToUpper(i.Name), ".", "_", -1)
		envVar = strings.Replace(envVar, "-", "_", -1)
		env, ok := os.LookupEnv(envVar)
		if ok {
			err := i.Value.Set(env)
			if err != nil {
				logrus.Fatalf("Invalid ENV variable %s: %s", envVar, env)
			}
		}
	})

	if *debugF {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if *traceF {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(true)
	}

	var cancel context.CancelFunc
	Context, cancel = context.WithCancel(context.Background())

	// handle termination signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-signals
		signal.Stop(signals)
		logrus.Warnf("Got %s, shutting down...", unix.SignalName(s.(syscall.Signal)))
		cancel()
	}()

	var err error
	BaseURL, err = url.Parse(*serverURLF)
	if err != nil {
		logrus.Fatalf("Failed to parse PMM Server URL: %s.", err)
	}
	if BaseURL.Host == "" || BaseURL.Scheme == "" {
		logrus.Fatalf("Invalid PMM Server URL: %s", BaseURL.String())
	}
	if BaseURL.Path == "" {
		BaseURL.Path = "/"
	}
	logrus.Debugf("PMM Server URL: %s.", BaseURL)

	Hostname, err = os.Hostname()
	if err != nil {
		logrus.Fatalf("Failed to detect hostname: %s", err)
	}

	// use JSON APIs over HTTP/1.1
	transport := httptransport.New(BaseURL.Host, BaseURL.Path, []string{BaseURL.Scheme})
	transport.SetLogger(logrus.WithField("component", "client"))
	transport.Debug = *debugF || *traceF
	// disable HTTP/2
	transport.Transport.(*http.Transport).TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}
	client.Default = client.New(transport, nil)
}
