package util

import (
	"fmt"
	g "github.com/labstack/gommon/log"
	log "github.com/sirupsen/logrus"
)

const emj = "😄😃😀😊☺️😉😍😘😚😗😙😜😝😛😳😁😔😌😒😞😣😢😂😭😪😥😰😅😓😩😫😨😱😠😡😤😖😆😋😷😎😴😵😲😟😦😧😈👿😮😬😐😕😯😶😇😏😑👲👳👮👷💂👶👦👧👨👩👴👵👱👼👸😺😸😻😽😼🙀😿😹😾👹👺🙈🙉🙊💀👽💩🔥✨🌟💫💥💢💦💧💤💨👂👀👃👅👄👍👎👌👊✊✌️👋✋👐👆👇👉👈🙌🙏☝️👏💪🚶🏃💃👫👪👬👭💏💑👯🙆🙅💁🙋💆💇💅👰🙎🙍🙇🎩👑👒👟👞👡👠👢👕👔👚👗🎽👖👘👙💼👜👝👛👓🎀🌂💄💛💙💜💚❤️💔💗💓💕💖💞💘💌💋💍💎👤👥💬👣💭🐶🐺🐱🐭🐹🐰🐸🐯🐨🐻🐷🐽🐮🐗🐵🐒🐴🐑🐘🐼🐧🐦🐤🐥🐣🐔🐍🐢🐛🐝🐜🐞🐌🐙🐚🐠🐟🐬🐳🐋🐄🐏🐀🐃🐅🐇🐉🐎🐐🐓🐕🐖🐁🐂🐲🐡🐊🐫🐪🐆🐈🐩🐾💐🌸🌷🍀🌹🌻🌺🍁🍃"

type CustomLogger struct {
	log.Logger
}

func (c *CustomLogger) Debugj(j g.JSON) {}
func (c *CustomLogger) Errorj(j g.JSON) {}
func (c *CustomLogger) Infoj(j g.JSON)  {}
func (c *CustomLogger) Fatalj(j g.JSON) {}
func (c *CustomLogger) Level() g.Lvl {
	return g.Level()
}

//func (c *CustomLogger) Debugj(j g.JSON) {}

func toString(hash [32]byte) string {
	if hash == [32]byte{} {
		return "GENESIS"
	}
	return fmt.Sprintf("%x", hash)
}

// CompactEmoji transforms the hash into a compact 3 digit emoji representation
func CompactEmoji(hash [32]byte) string {
	emr := []rune(emj)
	f := []rune{}
	for i := 0; i < 4; i++ {
		f = append(f, emr[hash[i]])
	}
	f = append(f, '…')
	for i := 28; i < 32; i++ {
		f = append(f, emr[hash[i]])
	}
	return string(f)
}
