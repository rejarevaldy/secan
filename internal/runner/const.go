package runner

const (
	version = "0.0.0"
	author  = "rejarevaldy"
	banner  = `
   _____ _______________    _   __
  / ___// ____/ ____/   |  / | / /
  \__ \/ __/ / /   / /| | /  |/ / 
 ___/ / /___/ /___/ ___ |/ /|  /  
/____/_____/\____/_/  |_/_/ |_/   

	v` + version + ` - @` + author
	usage = `
  	secan [options]`
	options = `
	-t, --target <FILE>           Define single URL to fuzz
	-l, --list <FILE>         Fuzz URLs within file
	-o, --output <FILE>       Save output in file
	-threads <int>            Number of concurrent threads (default 10)
	-s, --silent              Silent mode
	-v, --verbose             Verbose mode
	-V, --version             Show current CRLFuzz version
	-h, --help                Display its help
`
)