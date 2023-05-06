package lorca

import (
	"os"
	"runtime"
)

// BraveExecutable returns a string which points to the preferred Chromium
// executable file.
var BraveExecutable = LocateBrave

// LocateBrave returns a path to the Chromium binary, or an empty string if
// the Chromium installation is not found.
func LocateBrave() string {

	// If env variable "LORCABRAVE" specified and it exists
	if path, ok := os.LookupEnv("LORCABRAVE"); ok {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	var paths []string
	switch runtime.GOOS {
	case "darwin":
		paths = []string{
			"/usr/bin/brave",
		}
	case "windows":
		paths = []string{
			os.Getenv("LocalAppData") + "/BraveSoftware/Brave-Browser/Application/brave.exe",
			os.Getenv("ProgramFiles") + "/BraveSoftware/Brave-Browser/Application/brave.exe",
			os.Getenv("ProgramFiles(x86)") + "/BraveSoftware/Brave-Browser/Application/brave.exe",
		}
		
		
	default:
		paths = []string{
			"/usr/bin/brave",
			"/snap/bin/brave",
		}
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return path
	}
	return ""
}
