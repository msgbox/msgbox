MsgBox Binary
===========================

This is the main binary for running MsgBox. It's written in Go and should compile on most platforms. It
brings in all the various MsgBox pieces and spawns their listeners. If you would like to customize a piece
you can edit this to import your own version of a certain piece.

For more information on what MsgBox is or how it works you can read the [Spec](https://github.com/msgbox/spec).

The binary reads in a config file at runtime and spawns:

  * Relay
  * Submission Agent
  * (n) Incoming Workers
  * (n) Outgoing Workers

## How To Install

In order to get MsgBox up and running it must be compiled using Go for your platform.

**Prerequisites**

  * RabbitMQ server
  * Postgresql server
  * Go language installed in order to compile
  * Go path setup as explained in the Go [Docs](http://golang.org/doc/code.html#tmp_2)
  * _optional_ Git and Mercurial for cloning packages

Once all that is setup you will probably want to look over the [Storage](https://github.com/msgbox/storage)
piece and setup your database tables and environment variables.

**Installing**

In your Go path run `go install github.com/msgbox/msgbox`

Copy the msgbox.conf file to `/etc/msgbox` and edit as needed

Now if the bin directory is in your Path you can start the service by running: `msgbox`