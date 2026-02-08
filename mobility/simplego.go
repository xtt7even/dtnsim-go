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

func UpdateConnections(peers []peer.Peer) {
	for i := range peers {
        peers[i].ConnectedTo = peers[i].ConnectedTo[:0]
    }

	for i := 0; i < len(peers); i++ {
        root := &peers[i]

        for j := i + 1; j < len(peers); j++ {   
            neighbour := &peers[j]

            if arePeersOverlapping(*root, *neighbour) {
                root.ConnectedTo      = append(root.ConnectedTo,      neighbour.Id)
                neighbour.ConnectedTo = append(neighbour.ConnectedTo, root.Id)
            }
        }
    }
}

func arePeersOverlapping(root, neighbour peer.Peer) bool {
    return !(root.X + root.ConnectionRadius < neighbour.X - neighbour.ConnectionRadius ||
             root.X - root.ConnectionRadius > neighbour.X + neighbour.ConnectionRadius ||
             root.Y + root.ConnectionRadius < neighbour.Y - neighbour.ConnectionRadius ||
             root.Y - root.ConnectionRadius > neighbour.Y + neighbour.ConnectionRadius)
}

func getIfCoordInPeerArea(x int, y int, peer peer.Peer) bool {
	if x < (peer.X - peer.ConnectionRadius) || x > (peer.X + peer.ConnectionRadius) || y < (peer.Y - peer.ConnectionRadius) || y > (peer.Y + peer.ConnectionRadius) {
		return false
	}
	return true
}