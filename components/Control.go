package components

type DirectionMap map[string]bool

type Control struct {
	MoveSpeed float64
	Direction DirectionMap
}

func NewControl(moveSpeed float64) Control {
	return Control{
		MoveSpeed: moveSpeed,
		Direction: map[string]bool{
			"Up":    false,
			"Down":  false,
			"Right": false,
			"Left":  false,
		},
	}
}
