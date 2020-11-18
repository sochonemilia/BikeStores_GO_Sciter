package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	//Set initial window dimention and cordinats
	winRect := sciter.NewRect(400, 400, 700, 400)

	//Create sciter window
	window, windowError := window.New(sciter.SW_MAIN|sciter.SW_TITLEBAR, winRect)
	//window, windowError := window.New(sciter.SW_MAIN|sciter.SW_TITLEBAR, sciter.DefaultRect)

	if windowError != nil {
		color.RedString("Failed to generate sciter window", windowError.Error())
	}

	//LoadUI
	uiLoadingError := window.LoadFile("./mainwindow.html")
	if uiLoadingError != nil {
		color.RedString("Failed to generate sciter window", windowError.Error())
	}

	//Set title bar for window
	window.SetTitle("Okienko sciter")
	rootEl, _ := window.GetRootElement()

	buttonSelect, _ := rootEl.SelectById("selectbtn")
	buttonClose, _ := rootEl.SelectById("closebtn")
	buttonClear, _ := rootEl.SelectById("clearbtn")
	scriptArea, _ := rootEl.SelectById("sqlscriptarea")

	buttonSelect.OnClick(func() {
		buttonSelect.SetStyle("border", "3px solid black")
		scriptArea.SetText("SELECT")
		scriptArea.SetStyle("background-color", "royalblue")
	})
	buttonClear.OnClick(func() {
		scriptArea.Clear()
	})
	buttonClose.OnClick(func() {
		os.Exit(0)
	})

	//Launch
	window.Show()

	//Start
	window.Run()
}
