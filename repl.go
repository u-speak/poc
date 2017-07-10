package main

import (
	"github.com/chzyer/readline"
	log "github.com/sirupsen/logrus"
	"github.com/u-speak/poc/chain"
	"io"
	"strings"
)

var completer = readline.NewPrefixCompleter(
	readline.PcItem("mine"),
	readline.PcItem("append"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func repl(c *chain.Chain) {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)
		switch {
		case strings.HasPrefix(line, "mine "):
			content := line[5:]
			mine(content, c.LastHash())
		case strings.HasPrefix(line, "add "):
			content := line[4:]
			mineAndAdd(content, c)
		case line == "print":
			err := c.PrintChain()
			if err != nil {
				log.Error(err)
			}
		case line == "exit":
			return
		case strings.HasPrefix(line, "get "):
			//hash := line[4:]
		default:
			log.Warnf("Command `%s' not found", line)
			log.Warn("Please check if you specified the correct number of arguments")
		}
	}
}
