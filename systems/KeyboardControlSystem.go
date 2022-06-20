package systems

import (
	"github.com/airtonix/rpg/components"
	"github.com/airtonix/rpg/core/array"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
	"github.com/sedyh/mizu/pkg/engine"
)

type Axis int

const (
	AxisHorizontal = 1
	AxisVertical   = 2
)

var DirectionAxisMap = map[ebiten.Key]Axis{
	ebiten.KeyW: AxisVertical,
	ebiten.KeyS: AxisVertical,
	ebiten.KeyD: AxisHorizontal,
	ebiten.KeyA: AxisHorizontal,
}

var DirectionKeyMap = map[ebiten.Key]string{
	ebiten.KeyW: "Up",
	ebiten.KeyD: "Right",
	ebiten.KeyA: "Left",
	ebiten.KeyS: "Down",
}

type KeyboardControlSystem struct {
	KeyPressMap []ebiten.Key
}

func NewKeyboardControlSystem() *KeyboardControlSystem {
	return &KeyboardControlSystem{
		KeyPressMap: []ebiten.Key{},
	}
}

func (system *KeyboardControlSystem) Draw(
	world engine.World,
	screen *ebiten.Image,
) {
	controllables := world.View(
		components.Control{},
	)

	system.WatchKey(ebiten.KeyW)
	system.WatchKey(ebiten.KeyD)
	system.WatchKey(ebiten.KeyS)
	system.WatchKey(ebiten.KeyA)

	actionCount := len(system.KeyPressMap)
	firstKey := ebiten.Key0
	secondKey := ebiten.Key0

	if actionCount > 0 {
		firstKey = system.KeyPressMap[0]
	}
	if actionCount > 1 {
		secondKey = system.KeyPressMap[1]
	}

	firstAxis := DirectionAxisMap[firstKey]
	firstDirection := DirectionKeyMap[firstKey]
	secondAxis := DirectionAxisMap[secondKey]
	secondDirection := DirectionKeyMap[secondKey]

	controllables.Each(func(entity engine.Entity) {
		var control *components.Control
		entity.Get(&control)

		for direction := range control.Direction {
			isFirst := direction == firstDirection
			isSecond := direction == secondDirection
			isEnabled := isFirst || isSecond && firstAxis != secondAxis
			control.Direction[direction] = isEnabled
		}
	})
}

func (system *KeyboardControlSystem) WatchKey(key ebiten.Key) {
	if inpututil.IsKeyJustPressed(key) {
		system.AddKey(key)
	}

	if inpututil.IsKeyJustReleased(key) {
		system.RemoveKey(key)
	}
}

func (system *KeyboardControlSystem) AddKey(key ebiten.Key) {
	if !lo.Contains(system.KeyPressMap, key) {
		system.KeyPressMap = array.Prepend(system.KeyPressMap, key)
	}
}

func (system *KeyboardControlSystem) RemoveKey(key ebiten.Key) {
	system.KeyPressMap = array.Remove(system.KeyPressMap, key)
}
