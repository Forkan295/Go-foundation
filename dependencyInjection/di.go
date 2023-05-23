package main

import "fmt"

type RockClimber struct {
	rockClimbed int
}

func (rc *RockClimber) climbRock() {
	rc.rockClimbed++
	if rc.rockClimbed == 10 {
		rc.climbSafe()
	}
}

func (rc *RockClimber) climbSafe() {
	fmt.Println("climbing with brake", rc.rockClimbed)
}

func main() {
	rc := &RockClimber{}
	for i := 1; i < 11; i++ {
		rc.climbRock()
	}
}
