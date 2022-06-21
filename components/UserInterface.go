package components

import "image/color"

type Text struct {
	Content  string
	Color    color.Color
	FontFace string
	FontSize int
}

func NewText(content string, fg color.Color) Text {
	return Text{
		Content: content,
		Color:   fg,
	}
}

type DebugText struct {
	Content  map[string]any
	Color    color.Color
	FontFace string
	FontSize int
}

func NewDebugText(fg color.Color) DebugText {
	return DebugText{
		Content: map[string]any{},
		Color:   fg,
	}
}
