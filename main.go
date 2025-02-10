package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Mode:
type TextModel struct {
	content string
}

// ViewModel:
type TextViewModel struct {
	model            *TextModel
	text_binding     *widget.Entry
	save_button_text string
}

func NewTextViewModel(model *TextModel) *TextViewModel {
	entry := widget.NewMultiLineEntry()
	entry.SetText(model.content)

	return &TextViewModel{
		model:            model,
		text_binding:     entry,
		save_button_text: "submit",
	}
}

func (self *TextViewModel) Save() {
	self.model.content = self.text_binding.Text
	println("Content saved:", self.model.content)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Text Editor")

	model := &TextModel{"Siii"}
	view_model := NewTextViewModel(model)

	// View:
	save_button := widget.NewButton(view_model.save_button_text, func() {
		view_model.Save()
	})

	// View:
	content := container.NewVBox(
		view_model.text_binding,
		save_button,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	// myWindow.SetFixedSize(true)
	myWindow.ShowAndRun()
}
