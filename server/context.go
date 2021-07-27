package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Context 思考这里为什么不用interface
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(data interface{}) error {

	body, err := ioutil.ReadAll(c.R.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {

	c.W.WriteHeader(code)

	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson)

	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		W: writer, R: request,
	}
}
