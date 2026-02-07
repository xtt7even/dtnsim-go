package peer

import (
	"math/rand/v2"

	"github.com/google/uuid"
)




type Peer struct {
	X, Y int
	VX, VY int
	WaypointX, WaypointY int
	Id string
	ConnectionRadius int //simulating connection radius
	Name string
	MaxShareSpeedPerSec int	//simulating max share speed in kb/s
}

func NewPeer(name string) *Peer {
	id := uuid.New();
	randomSpeed := rand.IntN(50)
	randomRadius := rand.IntN(10)
	if randomSpeed < 10 {
		randomSpeed = 10
	}
	p := Peer{Id: id.String(), Name: name, ConnectionRadius: randomRadius, MaxShareSpeedPerSec: randomSpeed }
	return &p
}
