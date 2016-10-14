package main

import (
	"github.com/kataras/iris"

	"../core/elements"
	"../core/random"
)

type BoardTicket struct {
	Ticket       string `json:"ticket"`
	Board        *elements.Board `json:"board"`
	BoardRequest *elements.BoardRequest `json:"boardRequest"`
}

type Check struct {
	Ticket string `json:"ticket"`
	Codes  []int `json:"codes"`
}

var ticketBoardMap = map[string]*BoardTicket{}

func main() {
	iris.Config.IsDevelopment = true // this will reload the templates on each request

	iris.StaticServe("../web", "/web")
	iris.Post("/api/create", create)
	iris.Post("/api/check", check)
	iris.Get("/api/tickets", tickets)
	iris.Get("/api/show/:ticket", show)

	iris.Listen(":8080")
}

func show(ctx *iris.Context) {
	if bTicket, ok := ticketBoardMap[ctx.Param("ticket")]; ok {
		ctx.JSON(iris.StatusOK, bTicket)
	} else {
		ctx.JSON(iris.StatusInternalServerError, "ticket not found")
	}
}

func tickets(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, ticketBoardMap)
}
func create(ctx *iris.Context) {
	bRequest, err := convertBoardRequest(ctx)
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, "erro ao converter")
		return
	}

	board, err := elements.Create(bRequest)
	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, "erro ao criar")
		return
	}

	ctx.JSON(iris.StatusOK, createTicketBoard(bRequest, board))
}

func check(ctx *iris.Context) {
	check := Check{}
	err := ctx.ReadJSON(&check)

	if err != nil {
		ctx.JSON(iris.StatusInternalServerError, "erro ao converter")
		return
	}

	if bTicket, ok := ticketBoardMap[check.Ticket]; ok {
		bResponse := bTicket.Board.Check(check.Codes)
		if bResponse.Error != nil {
			ctx.JSON(iris.StatusInternalServerError, bResponse)
		} else {
			ctx.JSON(iris.StatusOK, bResponse)
		}
	} else {
		ctx.JSON(iris.StatusInternalServerError, "ticket not found")
	}
}

func convertBoardRequest(ctx *iris.Context) (elements.BoardRequest, error) {
	bRequest := elements.BoardRequest{}
	err := ctx.ReadJSON(&bRequest)
	return bRequest, err
}

func generateTicket() (string) {
	for {
		ticket := random.RandAlphanum(20)
		if _, ok := ticketBoardMap[ticket]; !ok {
			return ticket
		}
	}
}

func createTicketBoard(bRequest elements.BoardRequest, board elements.Board) (*BoardTicket) {
	ticket := generateTicket()
	ticketBoardMap[ticket] = &BoardTicket{BoardRequest: &bRequest, Board: &board, Ticket: ticket}
	return ticketBoardMap[ticket]
}