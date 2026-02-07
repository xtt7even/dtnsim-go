package main

import (
	"fmt"
	"math/rand/v2"
	simmap "sandbox/sandbox/map"
	simplego "sandbox/sandbox/mobility"
	"sandbox/sandbox/peer"
	"time"
)

func main() {
	simmap := simmap.NewMap(250)
	names := [5]string{"Alice", "John", "Matthew", "Gregory", "Maria"}
	var peers []peer.Peer
	for i := 0; i < len(names); i++ {
		peers = append(peers, *peer.NewPeer(names[i]))
		simplego.PickRandomWaypointForPeer(&peers[i], simmap)
	}

	for i := 0; i < 1000; i++ {
		tick(peers, simmap)
	}
}

func tick(peers []peer.Peer, m *simmap.Map) {
	for i := 0; i < len(peers); i++ {
		peer := &peers[i]
		if simplego.IsOnWaypoint(peer) {
			simplego.OnWaypointReach(peer)
		}
		toCreateNewWaypoint := rand.IntN(100) < 5 //5% chance of creating a new waypoint
		if (toCreateNewWaypoint) {
			simplego.PickRandomWaypointForPeer(peer, m)
		}
		simplego.SimpleMove(peer)

		//? temp
		if i == 0 {
			fmt.Println("========================================")
			fmt.Println("Tick:", time.Now().Format("2006-01-02 15:04:05"))
			fmt.Printf("%-3s %-12s %-20s %-20s %s\n", "ID", "NAME", "POSITION", "WAYPOINT", "DETAILS")
		}
		p := peers[i]
		pos := fmt.Sprintf("pos=%v,%v", p.X, p.Y)
		wp := fmt.Sprintf("wp=%v,%v", p.WaypointX, p.WaypointY)
		fmt.Printf("%-3d %-12s %-20s %-20s %v\n", i, p.Name, pos, wp, p)
	}
}