// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package service

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {

	// Registered clients.
	clients map[*Client]bool

	//发送消息给用户
	echoMsgQueue chan string

	//机器人消息
	botMsgQueue chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

//NewHub New Hub
func NewHub() *Hub {
	hub := &Hub{
		register:    make(chan *Client),
		unregister:  make(chan *Client),
		clients:     make(map[*Client]bool),
		botMsgQueue: make(chan []byte),
	}

	return hub
}

//Run start service
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case echoMsg := <-h.botMsgQueue:
			for client := range h.clients {
				select {
				case client.send <- []byte(echoMsg):
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
