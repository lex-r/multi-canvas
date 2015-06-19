package main

import (
	"github.com/lex-r/multi-canvas/messages"
	"github.com/golang/protobuf/proto"
	"log"
)

func registerFunctions(s *Service) {
	s.register("ping", func(c *connection, req *messages.ServerRequest) {
		resp := &messages.ClientRequest{}
		resp.Method = proto.String("pong")
		resp.RequestPong = &messages.ClientRequestPong{Text:proto.String("Pong")}

		response, err := proto.Marshal(resp)
		if err != nil {
			log.Print("Marshal error: ", err)
			return
		}

		c.send <- response
	})

	s.register("register", func(c *connection, req *messages.ServerRequest) {
		register := req.RequestRegister

		monitor := monitors[c]
		monitor.Width = *register.W
		monitor.Height = *register.H
		monitor.VirtX = 0
		monitor.VirtY = 0
		world := worlds[monitor.WorldId]
		if world == nil {
			world = &World{
				Width:monitor.Width,
				Height:monitor.Height,
				Ball:Ball{10,10,20,20},
				Connections:make(map[*connection]bool),
				stop:make(chan bool),
			}
			world.Connections[c] = true
			worlds[monitor.WorldId] = world
			go world.updateWorld()
		} else {
			monitor.VirtX = world.Width
			world.Width += monitor.Width;
			world.Connections[c] = true
		}

		monitorResp := makeMonitorRequest(monitor)
		response := makeCreateWorldRequest(world)

		c.send <- monitorResp
		c.send <- response

		log.Printf("Register monitor %vx%v", *register.W, *register.H)
		log.Printf("Monitor %v", monitor)
	})
}

func makeCreateWorldRequest(world *World) []byte {
	resp := &messages.ClientRequest{}
	resp.Method = proto.String("createWorld")
	resp.RequestCreateWorld = &messages.ClientRequestCreateWorld{
		Width:proto.Int32(world.Width),
		Height:proto.Int32(world.Height),
		Ball:&messages.Ball{
			X:proto.Int32(world.Ball.X),
			Y:proto.Int32(world.Ball.Y),
			DirX:proto.Int32(world.Ball.XDir),
			DirY:proto.Int32(world.Ball.YDir),
		},
	}

	log.Printf("makeCreateWorldRequst: resp %v", resp)

	response, err := proto.Marshal(resp)
	if err != nil {
		log.Printf("MakeCreateWorldRequest marshal error %v", err)
	}

	log.Printf("makeCreateWorldRequst: response %v", response)

	return response
}

func makeMonitorRequest(m *Monitor) []byte {
	resp := &messages.ClientRequest{}
	resp.Method = proto.String("monitor")
	resp.RequestMonitor = &messages.ClientRequestMonitor{
		X: &m.VirtX,
		Y: &m.VirtY,
	}
	response, err := proto.Marshal(resp)
	if err != nil {
		log.Printf("MakeMonitorRequest marshal error %v", err)
	}

	return response
}