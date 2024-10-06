package main

import (
	"auto-clicker/dto"
	"auto-clicker/handlers"
	"auto-clicker/validators"
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Auto Clicker")

	// make the stop channel
	stopChannel := make(chan struct{})
	close(stopChannel)

	// Create a context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	application.Lifecycle().SetOnStopped(func() {
		cancel()
		if _, ok := <-stopChannel; ok {
			close(stopChannel)
		}
	})

	// set the window size
	window.Resize(fyne.NewSize(800, 600))

	// set up bound values
	minutes := binding.NewInt()
	seconds := binding.NewInt()
	ms := binding.NewInt()
	minutes.Set(0)
	seconds.Set(0)
	ms.Set(0)

	randomMinutes := binding.NewInt()
	randomSeconds := binding.NewInt()
	randomMs := binding.NewInt()
	randomMinutes.Set(0)
	randomSeconds.Set(0)
	randomMs.Set(0)

	clicks := binding.NewInt()
	clicks.Set(1)

	// set up entry objects for UI
	minutesEntry := widget.NewEntryWithData(binding.IntToString(minutes))
	secondsEntry := widget.NewEntryWithData(binding.IntToString(seconds))
	msEntry := widget.NewEntryWithData(binding.IntToString(ms))

	randomMinutesEntry := widget.NewEntryWithData(binding.IntToString(randomMinutes))
	randomSecondsEntry := widget.NewEntryWithData(binding.IntToString(randomSeconds))
	randomMsEntry := widget.NewEntryWithData(binding.IntToString(randomMs))

	clicksEntry := widget.NewEntryWithData(binding.IntToString(clicks))

	// set up validation
	minutesEntry.Validator = validators.ValidateIntegerInput
	secondsEntry.Validator = validators.ValidateIntegerInput
	msEntry.Validator = validators.ValidateIntegerInput

	randomMinutesEntry.Validator = validators.ValidateIntegerInput
	randomSecondsEntry.Validator = validators.ValidateIntegerInput
	randomMsEntry.Validator = validators.ValidateIntegerInput

	clicksEntry.Validator = validators.ValidateIntegerInput

	// create a container for the random delay entries
	randomDelayContainer := container.NewVBox(
		widget.NewForm(
			&widget.FormItem{Text: "Random Minutes", Widget: randomMinutesEntry},
			&widget.FormItem{Text: "Random Seconds", Widget: randomSecondsEntry},
			&widget.FormItem{Text: "Random Milliseconds", Widget: randomMsEntry},
		),
	)
	randomDelayContainer.Hide() // hide the container by default

	// set up the random delay checkbox
	randomDelayCheckbox := widget.NewCheck("Random Delay?", func(checked bool) {
		handlers.HandleRandomDelayCheck(checked, randomDelayContainer)
	})

	// set up infinite clicks checkbox
	infiniteClicksCheckbox := widget.NewCheck("Infinite Clicks?", func(checked bool) {
		handlers.HandleInfiniteClicksCheck(checked, clicksEntry)
	})

	// set up the form layout
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Minutes", Widget: minutesEntry},
			{Text: "Seconds", Widget: secondsEntry},
			{Text: "Milliseconds", Widget: msEntry},
		},
	}

	// set up the clicks form
	clicksForm := &widget.Form{
		Items: []*widget.FormItem{
			{Widget: infiniteClicksCheckbox},
			{Text: "Clicks", Widget: clicksEntry},
		},
	}

	// start button
	startButton := widget.NewButton("Start", nil)
	startButton.OnTapped = func() {
		handlers.HandleStartStopButton(startButton)
		delayValues := getDelayValues(minutes, seconds, ms, randomMinutes, randomSeconds, randomMs, clicks, infiniteClicksCheckbox.Checked)
		handlers.HandleStartAutoClicker(*delayValues, &stopChannel, startButton)
	}

	// set up keypress handler
	go func() {
		handlers.HandleKeyPressEvent(
			ctx,
			func() {
				handlers.HandleStartStopButton(startButton)
				delayValues := getDelayValues(minutes, seconds, ms, randomMinutes, randomSeconds, randomMs, clicks, infiniteClicksCheckbox.Checked)
				handlers.HandleStartAutoClicker(*delayValues, &stopChannel, startButton)
			},
		)
	}()

	// add the form to the window
	window.SetContent(container.NewVBox(
		widget.NewLabel("Enter the delay between clicks"),
		form,
		randomDelayCheckbox,
		randomDelayContainer,
		clicksForm,
		startButton,
	))

	window.ShowAndRun()
}

func getDelayValues(min, sec, ms, rMin, rSec, rMs, clicks binding.Int, infinite bool) *dto.DelayValuesDto {
	minutes, _ := min.Get()
	seconds, _ := sec.Get()
	milliseconds, _ := ms.Get()
	randomMinutes, _ := rMin.Get()
	randomSeconds, _ := rSec.Get()
	randomMs, _ := rMs.Get()
	clickCount, _ := clicks.Get()

	delayValues := dto.NewDelayValues(minutes, seconds, milliseconds, randomMinutes, randomSeconds, randomMs)

	if infinite {
		delayValues.SetClicks(0)
	} else {
		delayValues.SetClicks(clickCount)
	}

	return delayValues
}
