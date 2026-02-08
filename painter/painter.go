package painter

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	simmap "sandbox/sandbox/map"
	"sandbox/sandbox/peer"
)

func Draw(m *simmap.Map, peers []peer.Peer) {
	clearScreen()
	points := make([][2]int, len(peers))

	for i := 0; i < len(peers); i++ {
		points[i][0] = peers[i].X
		points[i][1] = peers[i].Y
	}
	for y := 0; y < m.Size; y++ {
		for x := 0; x < m.Size; x++ {
			var isPeerCoord bool
			for _, peerCoord := range points {
				if (peerCoord[0] == x && peerCoord[1] == y) {
					isPeerCoord = true
				}
			}
			if isPeerCoord {
				fmt.Print(" 0 ")
			} else {
				fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
}

func clearScreen() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run() 
}