package main

import (
	"encoding/base64"
	"github.com/chzyer/readline"
	log "github.com/sirupsen/logrus"
	"github.com/u-speak/poc/chain"
	d "github.com/u-speak/poc/distribution"
	"github.com/u-speak/poc/util"
	context "golang.org/x/net/context"
	"io"
	"strconv"
	"strings"
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func listNodes(ns *NodeServer) func(string) []string {
	return func(_ string) []string {
		ret := []string{}
		for r := range ns.remoteConnections {
			ret = append(ret, r)
		}
		return ret
	}
}

func formatHash(hash [32]byte) string {
	if Config.Logger.PrintEmoji {
		return util.CompactEmoji(hash)
	}
	return base64.URLEncoding.EncodeToString(hash[:])
}

func repl(c *chain.Chain, ns *NodeServer) {
	var hosts = readline.PcItemDynamic(listNodes(ns))
	var completer = readline.NewPrefixCompleter(
		readline.PcItem("block",
			readline.PcItem("get"),
			readline.PcItem("add"),
		),
		readline.PcItem("mine"),
		readline.PcItem("chain", readline.PcItem("print")),
		readline.PcItem("node",
			readline.PcItem("add"),
			readline.PcItem("status", hosts),
			readline.PcItem("sync", hosts),
			readline.PcItem("list"),
		),
	)
	l, err := readline.NewEx(&readline.Config{
		Prompt:          Config.NodeNetwork.Interface + ":" + strconv.Itoa(Config.NodeNetwork.Port) + " \033[31m»\033[0m ",
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
		case strings.HasPrefix(line, "block add "):
			content := line[10:]
			n := mine(content, c.LastHash())
			ns.Push(&chain.Block{Content: content, Nonce: n})
		case strings.HasPrefix(line, "block get "):
			content := line[10:]
			h, err := base64.URLEncoding.DecodeString(content)
			if err != nil {
				log.Error(err)
				break
			}
			log.Debug(h)
			var hash [32]byte
			copy(hash[:], h)
			block := c.Get(hash)
			if block == nil {
				log.Info("No block found")
				break
			}
			log.Info(block.Content)
		case strings.HasPrefix(line, "node add "):
			content := line[9:]
			err := ns.Connect(content)
			if err != nil {
				log.Error(err)
				continue
			}
		case strings.HasPrefix(line, "node status "):
			content := line[12:]
			if _, contained := ns.remoteConnections[content]; !contained {
				log.Errorf("You are not connected to node %s", content)
				continue
			}
			client := d.NewDistributionServiceClient(ns.remoteConnections[content])
			info, err := client.GetInfo(context.Background(), &d.StatusParams{Host: Config.NodeNetwork.Interface + ":" + strconv.Itoa(Config.NodeNetwork.Port)})
			if err != nil {
				log.Error(err)
				continue
			}
			log.Debugf("%#v", info)
		case strings.HasPrefix(line, "node sync "):
			content := line[10:]
			if _, contained := ns.remoteConnections[content]; !contained {
				log.Errorf("You are not connected to node %s", content)
				continue
			}
			err := ns.SynchronizeChain(content)
			if err != nil {
				log.Error(err)
			}
		case strings.HasPrefix(line, "node list"):
			for r := range ns.remoteConnections {
				log.Debug(r)
			}
		case strings.HasPrefix(line, "chain print"):
			dump, err := c.DumpChain()
			if err != nil {
				log.Error(err)
				break
			}
			for _, b := range dump {
				log.WithFields(log.Fields{
					"hash": formatHash(b.Hash()),
					"prev": formatHash(b.PrevHash),
				}).Debug(b.Content)
			}
		case line == "emj":
			log.Info("Toggling Emoji mode...")
			Config.Logger.PrintEmoji = !Config.Logger.PrintEmoji
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
