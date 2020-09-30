package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	//	"log"
	"os"
	"regexp"
	"strings"

	"haii.or.th/api/util/filepathx"
	"haii.or.th/api/util/pqx"
	"haii.or.th/api/util/shell"
)

var verbose = flag.Bool("verbose", false, "Show more message")
var configFile = flag.String("config", "", "server configuration file (required)")
var outFile = flag.String("out", "", "GO lang output file name (required)")
var envPrefix = flag.String("envprefix", "", "Environment variable name prefix")

var dbCfg = regexp.MustCompile(`export\s+(.*)_DB\s*=\s*(.*)\s*`)
var keyCfg = regexp.MustCompile(`export\s+(.*)_KEY\s*=\s*(.*)\s*`)

var BuildVersion = ""

func initDB() {
	flag.Usage = showUsage

	if shell.ParseVersionFlag(BuildVersion) {
		os.Exit(0)
	}

	//	if *outFile == "" {
	//		showUsage()
	//		os.Exit(1)
	//	}

	c, _, err := readConfig(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stdout, "can not read configuration ...%v\n", err)
		os.Exit(1)
	}

	pqx.SetDefaultDbConnector(c)
}

func showUsage() {
	fmt.Fprintf(os.Stderr, "Usage of %s\n",
		filepathx.BaseName(os.Args[0]))
	flag.PrintDefaults()
}

func readConfig(fname string) (string, []byte, error) {
	fn, err := os.Open(fname)
	if err != nil {
		return "", nil, err
	}
	defer fn.Close()

	var dbconn string
	var key string

	rx := bufio.NewScanner(fn)
	for rx.Scan() {
		l := strings.TrimSpace(rx.Text())
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		if a := dbCfg.FindStringSubmatch(l); len(a) > 1 {
			if *envPrefix != "" && *envPrefix != a[1] {
				continue
			}

			*envPrefix = a[1]
			dbconn = strings.Trim(a[2], "\"'")
			continue
		}
		if a := keyCfg.FindStringSubmatch(l); len(a) > 1 {
			if *envPrefix != "" && *envPrefix != a[1] {
				continue
			}

			*envPrefix = a[1]
			key = strings.Trim(a[2], "\"'")
			continue
		}
	}

	if dbconn == "" || key == "" {
		return "", nil, nil
	}

	b, err := hex.DecodeString(key)
	if err != nil {
		return "", nil, err
	}

	return dbconn, b, nil
}
