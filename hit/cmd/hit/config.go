package main

import (
	"fmt"
	"strconv"
	"strings"
)

type parseFunc func(string) error

type config struct {
	url string
	n   int
	c   int
	rps int
}

func parseArgs(c *config, args []string) error {
	flagSet := map[string]parseFunc{
		"-url": stringVar(&c.url),
		"-n":   intVar(&c.n),
		"-c":   intVar(&c.c),
		"-rps": intVar(&c.rps),
	}

	for _, arg := range args {
		flag, val, _ := strings.Cut(arg, "=")

		setVar, ok := flagSet[flag]
		if !ok {
			return fmt.Errorf("flag provided but not defined: %s", flag)
		}
		if err := setVar(val); err != nil {
			return fmt.Errorf("invalid value %q for flag %s: %w", val, flag, err)
		}
	}

	return nil
}

func stringVar(p *string) parseFunc {
	return func(s string) error {
		*p = s
		return nil
	}
}

func intVar(p *int) parseFunc {
	return func(s string) error {
		var err error
		*p, err = strconv.Atoi(s)
		return err
	}
}
