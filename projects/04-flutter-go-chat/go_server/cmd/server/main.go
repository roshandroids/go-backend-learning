// Command server is Project 7 onward: the flagship Flutter+Go chat backend.
// Builds on Project 3's Hub, adding persistence (challenge 13), auth
// (challenge 14), and eventually multi-instance + Redis (challenges 15-17).
//
// Not wired yet, deliberately: internal/chat has the new ack/auth
// pieces scaffolded (see its package doc), but wiring them into a
// running server needs your own working Hub copied in from
// projects/03-chat-server/internal/hub (internal packages can't be
// imported across modules) plus a live Postgres/Flutter client to
// actually exercise end-to-end — neither is available in the session
// that scaffolded this. Once your Hub is copied in:
//  1. In ServeWS, before upgrading: call chat.ValidateToken on a token
//     from the query string or Authorization header; reject with 401
//     if it errors.
//  2. In ReadPump (or wherever incoming messages are handled), decode
//     each message into a chat.IncomingMessage, call
//     chat.HandleIncoming, and send the returned AckMessage back to
//     that specific client (not broadcast).
package main

import "fmt"

func main() {
	fmt.Println("TODO: wire a copy of Project 3's Hub, then use chat.ValidateToken/HandleIncoming per the package doc above")
}
