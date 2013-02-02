// MsgBox Binary
//
// This is the main binary for running MsgBox.
//
// It reads in a config file at runtime and spawns:
//    - Relay
//    - Submission Agent
//    - (n) Incoming Workers
//    - (n) Outgoing Workers
//
package main

import (
	"fmt"
	"github.com/hgfischer/goconf"
	"github.com/msgbox/relay"
	"github.com/msgbox/submission-agent"
	"github.com/msgbox/workers"
)

func main() {

	// Read Config File
	c, err := conf.ReadConfigFile("msgbox.conf")
	if err != nil {
		// Handle Error
	}

	// Startup Relay
	relay_external_port, _ := c.GetString("", "relay-external-port")
	relay_internal_port, _ := c.GetString("", "relay-internal-port")

	go bindIncoming(string(relay_external_port))
	go bindOutgoing(string(relay_internal_port))

	// Start Workers
	incoming_workers, _ := c.GetInt("", "incoming-workers")
	outgoing_workers, _ := c.GetInt("", "outgoing-workers")

	for i := 0; i < incoming_workers; i++ {
		createIncomingWorker(string(i))
	}

	for i := 0; i < outgoing_workers; i++ {
		createOutgoingWorker(string(i), relay_internal_port)
	}

	// Start Submission Agent
	submission_agent_port, _ := c.GetString("", "submission-port")
	createSubmissionAgent(submission_agent_port)

	for {
	}
}

func bindIncoming(port string) {
	inLn := relay.ListenIncoming(port)
	defer inLn.Close()
}

func bindOutgoing(port string) {
	outLn := relay.ListenOutgoing(port)
	defer outLn.Close()
}

func createIncomingWorker(tag string) {
	fmt.Println("Creating an Incoming Worker")
	workers.CreateIncoming(tag)
}

func createOutgoingWorker(tag string, port string) {
	fmt.Println("Creating an Outgoing Worker")
	workers.CreateOutgoing(tag, port)
}

func createSubmissionAgent(port string) {
	submission_agent.CreateAgent(port)
}
