package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"emails/model/bo"
	"emails/model/dto"
	"emails/usecase"
)

func Emails(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Read body failed"))
		return
	}

	req := dto.EmailReq{}
	if err := json.Unmarshal(reqBytes, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Read body parse failed"))
		return
	}

	err = usecase.SendEmails(r.Context(), bo.EmailReq{
		To:      req.To,
		Subject: req.Subject,
		Body:    req.Body,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(req)
}
