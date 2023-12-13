package interfaces

import "github.com/go-toast/toast"

func ShowIPToast(ipAddrs []string) {
	if len(ipAddrs) == 0 {
		toastMessage := toast.Notification{
			AppID:   "PiDash",
			Title:   "PiDash",
			Message: "No IP address found",
		}
		toastMessage.Push()
	} else {
		stringBuilder := "Server running on one of the following addresses: "
		for _, ipAddr := range ipAddrs {
			stringBuilder += ipAddr + ":8080, "
		}
		toastMessage := toast.Notification{
			AppID:   "PiDash",
			Title:   "PiDash",
			Message: stringBuilder[:len(stringBuilder)-2],
		}
		toastMessage.Push()
	}
}
