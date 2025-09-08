package cmd
import "time"

func getTimeFromZone(zone string) (string, error) {
	loc, err := time.LoadLocation(zone)

	if err != nil {
		return "", err
	}

	timeNow := time.Now().In(loc)
	return timeNow.Format(time.RFC1123), nil
}
