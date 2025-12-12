package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func main() {
    // Create a new Fyne app.
    a := app.New()

    // Create a new window with the title "Hello" in the app.
    w := a.NewWindow("Hello")

    // Create a new label.
    hello := widget.NewLabel("Hello Fyne!")

    // Assign a container as the content of the window.
    // The container contains a vertical box, which contains
    // a label and a button. When the button is pressed,
    // the text of the label will be changed.
    w.SetContent(container.NewVBox(
        hello,
        widget.NewButton("Hi!", func() {
            hello.SetText("Welcome :)")
        }),
    ))

    // Display the window.
    w.ShowAndRun()
}
