package main

import (
	"encoding/base64"
	"github.com/kpashka/echo-logrusmiddleware"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/u-speak/poc/chain"
	d "github.com/u-speak/poc/distribution"
	context "golang.org/x/net/context"
	"net/http"
	"strconv"
)

type Webserver struct {
	ch *chain.Chain
	ns *NodeServer
}

type reqBlock struct {
	Content string `json:"content" form:"content" query:"content"`
	Nonce   uint   `json:"nonce" form:"nonce" query:"nonce"`
}

func startWeb(ch *chain.Chain, ns *NodeServer) {
	e := echo.New()
	e.Logger = logrusmiddleware.Logger{logrus.StandardLogger()}
	w := &Webserver{ch: ch, ns: ns}
	e.GET("/nodes", w.GetNodes)
	e.GET("/nodes/:node", w.GetNode)
	e.GET("/blocks/", w.DumpChain)
	e.GET("/blocks/:hash", w.GetBlock)
	e.POST("/blocks/", w.AddBlock)
	e.Logger.Error(e.Start(":4000"))
}

// GetNodes dumps all connected nodes
func (s *Webserver) GetNodes(c echo.Context) error {
	return c.JSON(http.StatusOK, s.ns.remoteConnections)
}

// GetNode returns the status of a specific node
func (s *Webserver) GetNode(c echo.Context) error {
	content := c.Param("node")
	if _, contained := s.ns.remoteConnections[content]; !contained {
		return c.NoContent(http.StatusNotFound)
	}
	client := d.NewDistributionServiceClient(s.ns.remoteConnections[content])
	info, err := client.GetInfo(context.Background(), &d.StatusParams{Host: Config.NodeNetwork.Interface + ":" + strconv.Itoa(Config.NodeNetwork.Port)})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, info)
}

// GetBlock returns the contents of a single block
func (s *Webserver) GetBlock(c echo.Context) error {
	h, err := base64.URLEncoding.DecodeString(c.Param("hash"))
	if err != nil {
		return err
	}
	var hash [32]byte
	copy(hash[:], h)
	block := s.ch.Get(hash)
	c.Logger().Debug(block)
	return c.JSON(http.StatusOK, block)
}

// DumpChain returns the whole chain for debugging purposes
func (s *Webserver) DumpChain(c echo.Context) error {
	ch, err := s.ch.DumpChain()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ch)
}

// AddBlock adds a block to the chain
func (s *Webserver) AddBlock(c echo.Context) error {
	r := new(reqBlock)
	if err := c.Bind(r); err != nil {
		return err
	}
	err := s.ch.AddData(r.Content, r.Nonce)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Message string `json:"message"`
		}{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}
