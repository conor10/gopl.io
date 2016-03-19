package main

import (
	"gopl.io/ch2/distanceconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, val := range os.Args {
		distance, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Errorf("Unable to process value: %v", err)
		}
		fmt.Printf("%f feet == %f metres\n", distance, distanceconv.FToM(distanceconv.Feet(distance)))
		fmt.Printf("%f metres == %f feet\n", distance, distanceconv.MToF(distanceconv.Metres(distance)))
	}
}
