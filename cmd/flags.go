package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/user"
	"sort"
	"strings"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

const appName = "lb-check"

var (
	// Version is the current tag
	Version = "development"

	// Commit is the current last commit
	Commit = "latest"

	dbPassword string
)

func init() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s checks if the PostgreSQL is running and up to date.\n\n", appName)
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  %s [OPTIONS]...\n\n", appName)
		fmt.Fprintf(os.Stderr, "General Options:\n")

		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nReport BUGs on https://github.com/sebastianwebber/lb-check.\n\n")
	}

	flag.String("pg-isready-bin", "/usr/bin/pg_isready", "pg_isready binary path")

	flag.StringP("username", "U", user.Username, "database user name")
	flag.StringP("host", "h", "local socket", "database server host or socket directory")
	flag.StringP("dbname", "d", user.Username, "database name to connect to")
	flag.IntP("port", "p", 5432, "database server port")
	flag.BoolP("password", "W", false, "force password prompt")

	flag.String("query-delay-seconds", "SELECT CASE WHEN pg_last_wal_receive_lsn() = pg_last_wal_replay_lsn() THEN 0 ELSE  EXTRACT(EPOCH FROM (now() - pg_last_xact_replay_timestamp()))::INTEGER END AS delay;", "query to check the replica delay in bytes")
	flag.String("query-recovery", "SELECT pg_is_in_recovery() as recovery;", "query to check if the replica is in recovery mode")

	flag.Duration("max-delay", 60*time.Second, "max delay allowed to a replica")

	var checkOnlyHelp bytes.Buffer

	checkOnlyHelp.WriteString("executes only the defined checks.")

	allChecks := []string{}

	for k := range checkList {
		allChecks = append(allChecks, k)
	}

	sort.Strings(allChecks)
	checkOnlyHelp.WriteString(fmt.Sprintf(" Checks avaliable:\n - %s\n", strings.Join(allChecks, "\n - ")))

	checkOnlyHelp.WriteString("By default all checks are enabled.\n")

	// "execute only defined checks -- by default all checks are enabled"

	flag.StringSlice("check-only", []string{}, checkOnlyHelp.String())

	flag.Parse()

	viper.BindPFlags(flag.CommandLine)
	viper.SetEnvPrefix(toPromLayout(appName))
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	if viper.GetBool("password") {
		fmt.Printf("Password: ")
		out, err := terminal.ReadPassword(int(syscall.Stdin))

		if err != nil {
			panic(err)
		}

		dbPassword = fmt.Sprintf("%s", out)
	}

	log.Printf("Starting %s %s (%s)\n", appName, Version, Commit)
}
