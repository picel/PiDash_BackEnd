package interfaces

import (
	"os"

	"github.com/getlantern/systray"
)

func SetUpSystray(ipAddrs []string, Version string) {
	systray.Run(func() {
		systray.SetTitle("PiDash")
		systray.SetTooltip("PiDash")

		mRunning := systray.AddMenuItem("PiDash Server "+Version, "")
		mRunning.Disable()

		systray.AddSeparator()

		for _, ipAddr := range ipAddrs {
			mIP := systray.AddMenuItem(ipAddr+":8080", "Open PiDash in browser")
			mIP.Disable()
		}

		mQuit := systray.AddMenuItem("Quit", "Quit PiDash")
		go func() {
			<-mQuit.ClickedCh
			systray.Quit()
			os.Exit(0)
		}()
	}, nil)
}
