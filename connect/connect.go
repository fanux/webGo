package connect

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/fanux/lhttp"
	"github.com/fanux/webGo/gochess"
)

var chesses map[string]*gochess.ChessManual

//PlayProcessor is
type PlayProcessor struct {
	*lhttp.BaseProcessor
}

//LaziAPI is
type LaziAPI struct {
	X     int
	Y     int
	Color int
	Eats  []gochess.DropPoint
}

//OnMessage is
func (p *PlayProcessor) OnMessage(h *lhttp.WsHandler) {
	/*
		log.Print("on OnMessage: ", h.GetBody())
		h.AddHeader("content-type", "image/png")
		h.SetCommand("auth")
		h.Send(h.GetBody())
	*/
	if manual, ok := chesses["go-playroom1"]; ok {
		point := gochess.DropPoint{}
		fmt.Println("=========get body: ", h.GetBody())
		s := strings.SplitN(h.GetBody(), ",", 2)
		if len(s) != 2 {
			return
		}
		point.X, _ = strconv.Atoi(s[0])
		point.Y, _ = strconv.Atoi(s[1])
		fmt.Println("DropPoint length is: ", len(manual.DropPoints))
		if len(manual.DropPoints)%2 == 0 {
			point.State = gochess.PointStateBlack
		} else {
			point.State = gochess.PointStateWhite
		}

		eats := gochess.Lazi(manual, point)
		if eats != nil {
			resp := LaziAPI{point.X, point.Y, point.State, *eats}
			res, _ := json.Marshal(resp)
			h.Send(string(res))
			fmt.Println("send to client: ", resp)
		} else {
			fmt.Println("eats is nil: ", eats)
		}
	}

}

//StartConnection is
func StartConnection(websocketPort string) {
	chesses = make(map[string]*gochess.ChessManual)
	chesses["go-playroom1"] = gochess.NewDefaultChessManual()

	lhttp.Regist("play", &PlayProcessor{&lhttp.BaseProcessor{}})
	http.Handle("/", lhttp.Handler(lhttp.StartServer))
	http.ListenAndServe(websocketPort, nil)
}
