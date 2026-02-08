package main

import (
	"fmt"
	"math/rand/v2"
	simmap "sandbox/sandbox/map"
	simplego "sandbox/sandbox/mobility"
	"sandbox/sandbox/painter"
	"sandbox/sandbox/peer"
	"strings"
	"time"
)

func main() {
	simmap := simmap.NewMap(30)
	names := [5]string{"Alice", "John", "Matthew", "Gregory", "Maria"}
	var peers []peer.Peer
	for i := 0; i < len(names); i++ {
		peers = append(peers, *peer.NewPeer(names[i]))
		simplego.PickRandomWaypointForPeer(&peers[i], simmap)
	}

	for true {
		tick(peers, simmap)
	}
}

func tick(peers []peer.Peer, m *simmap.Map) {
	for i := 0; i < len(peers); i++ {
		peer := &peers[i]
		if simplego.IsOnWaypoint(peer) {
			simplego.OnWaypointReach(peer)
		}
		toCreateNewWaypoint := rand.IntN(100) < 1 //5% chance of creating a new waypoint
		if (toCreateNewWaypoint) {
			simplego.PickRandomWaypointForPeer(peer, m)
		}
		simplego.SimpleMove(peer)
		painter.Draw(m, peers)
		simplego.UpdateConnections(peers)		
		// debugOutput(peers, peers[i]);
		time.Sleep(50 * time.Millisecond)
	}
}

func debugOutput (peers []peer.Peer, p peer.Peer) {
		pos := fmt.Sprintf("%d,%d", p.X, p.Y)
		wp  := fmt.Sprintf("%d,%d", p.WaypointX, p.WaypointY)
		
		var conns string
		if len(p.ConnectedTo) == 0 {
			conns = "-"
		} else {
			var names []string
			for _, id := range p.ConnectedTo {
				for _, other := range peers {
					if other.Id == id {
						names = append(names, other.Name)
						break
					}
				}
			}
			conns = strings.Join(names, ", ")
		}

		fmt.Printf("%-12s %-20s %-20s %s\n", p.Name, pos, wp, conns)
}