package main

import (
	"net/http"
	"io/ioutil"

	"github.com/unrolled/render"
	"code_exercise/internal/message"
)

var r = render.New(render.Options{
	IndentJSON: true,
})

func MessageHandler(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.Text(rw, http.StatusBadRequest, err.Error())
		return
	}

	r.JSON(rw, http.StatusOK, message.Parse(string(body)))
}
