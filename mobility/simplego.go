package simplego

import (
	"fmt"
	simmap "sandbox/sandbox/map"
	"sandbox/sandbox/peer"
)

type Axis int

const (
	X Axis = iota
	Y
)

/*
	Tick based - it will only move peer 1 step X,Y per call
*/
func SimpleMove(peer *peer.Peer) {
	if peer.X < 0 || peer.Y < 0 {
		fmt.Println("ERROR: Peer's coordinates invalid")
		fmt.Println("X:", peer.X, "Y:", peer.Y)
		return;
	} 

	if peer.WaypointX >= 0 {
		if peer.WaypointX < peer.X {
			peer.X -= 1
		} 
		
		if peer.WaypointX > peer.X {
			peer.X += 1
		} 
	}
	if peer.WaypointY > 0 {
		if peer.WaypointY < peer.Y {
			peer.Y -= 1
		} 
		if peer.WaypointY > peer.Y {
			peer.Y += 1
		} 
	}

}

func IsOnWaypoint(peer *peer.Peer) bool {
	if (peer.WaypointX < peer.X || peer.WaypointX > peer.X || peer.WaypointY < peer.Y || peer.WaypointY > peer.Y) {
		return false
	}
	return true
}

func OnWaypointReach(peer *peer.Peer) {
	peer.WaypointX = -1
	peer.WaypointY = -1
}

func PickRandomWaypointForPeer(peer *peer.Peer, m *simmap.Map) {
	randomPosition := simmap.RandomPosition(m) 
	peer.WaypointX = randomPosition[0]
	peer.WaypointY = randomPosition[1]
}