package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/grafana/pyroscope/pkg/pprof"
)

var suffix = flag.String("suffix", "", "Suffix to append to the processed files. If empty, original file is overwritten")
var normalize = flag.Bool("normalize", false, "Normalize profile, see https://pkg.go.dev/github.com/grafana/pyroscope/pkg/pprof#Profile.Normalize")

type profileFix func(*pprof.Profile)

var fixes = []profileFix{
	fixMaxStackSize,
}

func main() {
	flag.Parse()
	if *normalize {
		fixes = append(fixes, fixNormalize)
	}

	if err := run(os.Args[len(os.Args)-flag.NArg():]); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run(files []string) error {
	if len(files) == 0 {
		flag.Usage()
		return nil
	}

	var errs []error
	for _, p := range files {
		if err := fixFile(p, fixes); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func fixFile(p string, fixes []profileFix) error {
	profile, err := pprof.OpenFile(p)
	if err != nil {
		return err
	}

	for _, f := range fixes {
		f(profile)
	}

	f, err := os.Create(p + *suffix)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = profile.WriteTo(f)
	return err
}

func fixMaxStackSize(profile *pprof.Profile) {
	profile.Profile = pprof.FixGoProfile(profile.Profile)
}

func fixNormalize(profile *pprof.Profile) {
	profile.Normalize()
}
