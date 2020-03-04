package main

import (
	"fyne.io/fyne"
	app2 "fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	app := app2.NewWithID("yuchanns Demo")
	app.Settings().SetTheme(theme.LightTheme())

	w := app.NewWindow("Demo")
	w.SetMaster() // App will exit with closing this window

	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("New", func() {
				println("New Menu")
			}),
			fyne.NewMenuItem("Open", func() {
				println("Open Folders")
			}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Undo", func() {
				println("Undo Command")
			}),
			fyne.NewMenuItem("Redo", func() {
				println("Redo Command")
			}),
		),
	))

	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Explore", theme.HomeIcon(), fyne.NewContainer(
			widget.NewHBox(
				widget.NewLabel("Exploring your favorite."),
				widget.NewVBox(
					widget.NewLabel("Popular things."),
					widget.NewLabel("Your trending."),
				),
			),
		)),
		widget.NewTabItemWithIcon("Profile", theme.InfoIcon(), fyne.NewContainer()),
	)
	tabs.SetTabLocation(widget.TabLocationBottom)
	tabs.SelectTabIndex(app.Preferences().Int("currentTab"))

	w.SetContent(tabs)

	w.ShowAndRun()

	app.Preferences().SetInt("currentTab", tabs.CurrentTabIndex())
}
