package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestInitLatencyCmd(t *testing.T) {
	parent := &cobra.Command{}
	initLatencyCmd(parent)

	flags := latencyCmd.Flags()

	if got := flags.Lookup("limit"); got == nil {
		t.Fatal("expected limit flag; got nil")
	}
}

func TestRunLatency(t *testing.T) {
	// We're only interested in the first HTTP call, e.g., the one to get the test ID
	// to validate our parameters got passed properly.
	tr := &recordingTransport{}
	c, err := newTestPerfopsClient(tr)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	runLatency(c, "example.com", "From here", 123)
	if got, exp := tr.req.URL.Path, "/run/latency"; got != exp {
		t.Fatalf("expected %v; got %v", exp, got)
	}
	got := reqBody(tr.req)
	exp := `{"target":"example.com","location":"From here","limit":123}`
	if got != exp {
		t.Fatalf("expected %v; got %v", exp, got)
	}
}