package orb

import (
	"fmt"
	"time"
)

func Time() {
	fmt.Println("Calculating point in time based on Universal to Atomic time drift..")
	if time.Now().Second()%2 == 0 {
		fmt.Println("Warning: anomaly detected in planet's rotational pattern - point of origin is not on Earth.")
	}
}
