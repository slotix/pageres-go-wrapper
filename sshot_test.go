package sshot

import (
	"fmt"
	"os"
	"testing"
)

func TestGetSShot(t *testing.T) {
	shotsDir := "shots"
    os.Mkdir(shotsDir, 0777)
	params := Parameters{
		Command: "pageres",
		Sizes:   "1024x768",
		Crop:    "--crop",
		Scale:   "--scale 0.9",
		Timeout: "--timeout 30",
		Filename:  fmt.Sprintf(`--filename=%s/<%%= url %%>`, shotsDir),
		UserAgent: "",
	}
	urls := []string{
		"http://google.com",
		"https://dbconvert.com",
		"http://something-that-doesnot-exists.com",
	}
	GetShots(urls, params)
	DeleteZeroLengthFiles(shotsDir)
	return
}
