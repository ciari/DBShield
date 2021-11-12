package main

import (
	"flag"
	"log"
	_ "net/http/pprof"
	"runtime"

	"github.com/ciari/DBShield/dbshield"
)

func usage(showUsage bool) {
	print("DBShield " + dbshield.Version + "\n")
	if showUsage {
		flag.Usage()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // For Go < 1.5

	config := flag.String("c", "/etc/dbshield.yml", "config `file`")
	listPatterns := flag.Bool("l", false, "get list of captured patterns")
	listAbnormals := flag.Bool("a", false, "get list of abnormal queries")
	checkConfig := flag.Bool("k", false, "show parsed config and exit")
	showVersion := flag.Bool("version", false, "show version")
	showHelp := flag.Bool("h", false, "show help")
	//Parsing command line arguments
	flag.Parse()

	if *showHelp {
		usage(true)
		return
	}

	if *showVersion {
		usage(false)
		return
	}

	if err := dbshield.SetConfigFile(*config); err != nil {
		log.Println(err)
		return
	}

	if *listPatterns {
		dbshield.Patterns()
		return
	}

	if *listAbnormals {
		dbshield.Abnormals()
		return
	}

	if *checkConfig {
		err := dbshield.Check()
		log.Println(err)
		return
	}

	log.Println(dbshield.Start())
}
