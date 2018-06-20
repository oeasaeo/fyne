package test

import "github.com/fyne-io/fyne"

var dummyCanvas fyne.Canvas

// Canvas returns an un-rendered canvas that is used for behavioural testing
func Canvas() fyne.Canvas {
	return dummyCanvas
}

type testCanvas struct {
	content fyne.CanvasObject
	focused fyne.FocusableObject
}

func (c *testCanvas) Content() fyne.CanvasObject {
	return c.content
}

func (c *testCanvas) SetContent(content fyne.CanvasObject) {
	c.content = content
}

func (c *testCanvas) Refresh(fyne.CanvasObject) {
}

func (c *testCanvas) Contains(fyne.CanvasObject) bool {
	return true
}

func (c *testCanvas) Focus(obj fyne.FocusableObject) {
	c.focused = obj
	obj.OnFocusGained()
}

func (c *testCanvas) Focused() fyne.FocusableObject {
	return c.focused
}

func (c *testCanvas) Size() fyne.Size {
	return fyne.NewSize(10, 10)
}

func (c *testCanvas) Scale() float32 {
	return 1.0
}
func (c *testCanvas) SetScale(float32) {
}

func (c *testCanvas) SetOnKeyDown(func(*fyne.KeyEvent)) {
}

// GetTestCanvas returns a reusable in-memory canvas used for testing
func GetTestCanvas() fyne.Canvas {
	if dummyCanvas == nil {
		dummyCanvas = &testCanvas{}
	}

	return dummyCanvas
}
