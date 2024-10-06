package handlers

import (
	"auto-clicker/dto"
	"context"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

// HandleRandomDelayCheck checks if the random delay checkbox is checked and shows or hides the extra container accordingly
func HandleRandomDelayCheck(checked bool, cont *fyne.Container) {
	if checked {
		cont.Show()
	} else {
		cont.Hide()
	}
}

// HandleInfiniteClicksCheck checks if the infinite clicks checkbox is checked and disables or enables the clicks entry accordingly
func HandleInfiniteClicksCheck(checked bool, cont *widget.Entry) {
	if checked {
		cont.Disable()
	} else {
		cont.Enable()
	}
}

// HandleStartStopButton changes the text of the start button to "Stop" and vice versa
func HandleStartStopButton(startButton *widget.Button) {
	if startButton.Text == "Start" {
		startButton.SetText("Stop")
	} else {
		startButton.SetText("Start")
	}
	startButton.Refresh()
}

// HandleKeyPressEvent starts a goroutine that listens for a key press event and calls the callback function when the - key is pressed
// TODO: Implement some label on the UI that allows users to set their own hotkey and reference it here
func HandleKeyPressEvent(ctx context.Context, callback func()) {
	hook.Register(hook.KeyDown, []string{}, func(e hook.Event) {
		if e.Keychar == 45 {
			callback()
		}
	})

	s := hook.Start()
	go func() {
		for {
			select {
			case <-ctx.Done():
				hook.End()
			case <-hook.Process(s):
			}
		}
	}()
}

//HandleStartAutoClicker starts the auto clicker and manages the stopChannel
func HandleStartAutoClicker(inputData dto.DelayValuesDto, stopChannel *chan struct{}, button *widget.Button) {
	if isChannelClosed(*stopChannel) {
		*stopChannel = make(chan struct{})
		delay, extraDelay := inputData.CalculateDelay()
		startAutoClicker(delay, extraDelay, inputData.Clicks, *stopChannel, button)
	} else {
		close(*stopChannel)
	}
}

// startAutoClicker starts the auto clicker with a time.NewTicker using the given delay and extraDelay
func startAutoClicker(delay int, extraDelay int, clicks int, stopChan chan struct{}, button *widget.Button) {
	go func() {
		ticker := time.NewTicker(time.Duration(delay))
		defer ticker.Stop()

		clickCount := 0

		for {
			select {
			case <-ticker.C:
				if extraDelay > 0 {
					randomDelay := rand.Intn(extraDelay)
					ticker.Stop() // Stop the ticker before sleep
					time.Sleep(time.Duration(randomDelay))
					ticker = time.NewTicker(time.Duration(delay))
				}
				robotgo.Click("left")
				clickCount++
				if clicks > 0 && clickCount >= clicks {
					HandleStartStopButton(button)
					close(stopChan)
					return
				}
			case <-stopChan:
				return
			}
		}
	}()
}

func isChannelClosed(ch <-chan struct{}) bool {
	if ch == nil {
		return true
	}
	select {
	case _, ok := <-ch:
		return !ok
	default:
		return false
	}
}
