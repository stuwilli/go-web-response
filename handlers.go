package webresponse

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	validate "github.com/stuwilli/govalidate"
)

//NewOkResponse ...
func NewOkResponse(data interface{}, path string) *Response {

	res := &Response{Message: "Ok", StatusCode: 200, Data: data,
		Timestamp: time.Now().Unix(), Path: path}

	return res
}

//NewCreatedResponse ...
func NewCreatedResponse(data interface{}, path string) *Response {

	res := &Response{Message: "Created", StatusCode: 201, Data: data,
		Timestamp: time.Now().Unix(), Path: path}

	return res
}

//NewNotFoundResponse ...
func NewNotFoundResponse(err interface{}, path string) *Response {

	res := &Response{Message: "NotFound", StatusCode: 404, Error: parseError(err),
		Timestamp: time.Now().Unix(), Path: path}

	return res
}

//NewServerErrorResponse ...
func NewServerErrorResponse(err interface{}, path string) *Response {

	res := &Response{Message: "ServerError", StatusCode: 500,
		Timestamp: time.Now().Unix(), Path: path, Error: parseError(err)}

	return res
}

//NewBadRequestResponse ...
func NewBadRequestResponse(err interface{}, path string) *Response {

	res := &Response{Message: "BadRequest", StatusCode: 400,
		Timestamp: time.Now().Unix(), Path: path, Error: parseError(err)}

	return res
}

//NewUnauthorizedResponse ...
func NewUnauthorizedResponse(err interface{}, path string) *Response {

	res := &Response{Message: "Unauthorized", StatusCode: 401,
		Timestamp: time.Now().Unix(), Path: path, Error: parseError(err)}

	return res
}

//NewForbiddenResponse ...
func NewForbiddenResponse(err interface{}, path string) *Response {

	res := &Response{Message: "Forbidden", StatusCode: 403,
		Timestamp: time.Now().Unix(), Path: path, Error: parseError(err)}

	return res
}

//ParseError ...
func parseError(err interface{}) map[string]string {

	m := make(map[string]string)

	switch err.(type) {

	case validate.ValidationError:

		fmt.Println(err.(error).Error())
		return validate.CastError(err.(error)).Errors()

	case *mysql.MySQLError:

		fmt.Println(err.(error).Error())
		m["error"] = "Database error"
		return m

	case string:

		fmt.Println(err.(string))
		m["error"] = err.(string)
		return m

	case error:

		fmt.Println(err.(error).Error())
		m["error"] = err.(error).Error()
		return m

	default:
		m["error"] = "Something went wrong"
		return m
	}
}
