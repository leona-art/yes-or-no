package main

import (
	"context"
	"fmt"
	yesornov1 "gomini/gen/yesorno/v1"
	"gomini/gen/yesorno/v1/yesornov1connect"
	"gomini/infra"
	"gomini/usecase"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type YesOrNoServer struct {
	AskService *usecase.AskService
}

// Ask implements yesornov1connect.YesOrNoServiceHandler.
func (y *YesOrNoServer) Ask(ctx context.Context, req *connect.Request[yesornov1.AskRequest]) (*connect.Response[yesornov1.AskResponse], error) {
	answer, err := y.AskService.Ask(ctx, "イーロンマスク", req.Msg.Question)
	if err != nil {
		return nil, err
	}
	log.Printf("answer: %s", answer)
	res := connect.NewResponse(&yesornov1.AskResponse{
		Answer: answer.Answer,
		Finish: answer.Finish,
	})
	return res, nil
}

func main() {
	ctx := context.Background()
	ai, err := infra.NewGemini(ctx)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ai.Close()

	askService := usecase.NewAskService(ai)

	server := &YesOrNoServer{
		AskService: askService,
	}
	mux := http.NewServeMux()
	path, handler := yesornov1connect.NewYesOrNoServiceHandler(server)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
