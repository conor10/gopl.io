package distanceconv

type Metres float64
type Feet float64

const convFactor = 3.2808

func MToF(m Metres) Feet {
	return Feet(convFactor * m)
}

func FToM(f Feet) Metres {
	return Metres(f / convFactor)
}

