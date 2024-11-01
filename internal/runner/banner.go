package runner

import "github.com/projectdiscovery/gologger"

func showBanner() {
	gologger.Printf("%s\n\n", banner)
	gologger.Labelf("Please be responsible\n")
}