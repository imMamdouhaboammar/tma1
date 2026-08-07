package handler

import "testing"

func TestShouldStopTranscriptWatcher(t *testing.T) {
	cases := []struct {
		name        string
		agentSource string
		event       string
		want        bool
	}{
		{name: "claude stop can be rejected", agentSource: "claude_code", event: "Stop", want: false},
		{name: "claude session end is terminal", agentSource: "claude_code", event: "SessionEnd", want: true},
		{name: "codex stop is restart safe", agentSource: "codex", event: "Stop", want: true},
		{name: "codex session end is terminal", agentSource: "codex", event: "SessionEnd", want: true},
		{name: "ordinary hook keeps watcher running", agentSource: "claude_code", event: "PostToolUse", want: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := shouldStopTranscriptWatcher(tc.agentSource, tc.event); got != tc.want {
				t.Fatalf("shouldStopTranscriptWatcher(%q, %q) = %v, want %v", tc.agentSource, tc.event, got, tc.want)
			}
		})
	}
}
