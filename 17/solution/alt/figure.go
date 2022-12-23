package alt

type figure struct {
	top  int
	left int
	form Form
}

func newFigure(top, left int, form Form) *figure {
	return &figure{
		top:  top,
		left: left,
		form: form,
	}
}

func (f *figure) height() int {
	return f.form.height()
}

func (f *figure) width() int {
	return f.form.width()
}

func (f *figure) moveDown() {
	f.top--
}

// func (f *figure) moveLeft() {
// 	if f.left >= 0 {
// 		f.left--
// 	}
// }

// func (f *figure) moveRight() {
// 	if
// 	f.left++
// }
