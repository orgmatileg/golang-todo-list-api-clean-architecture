package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// ResponseBody struct
type ResponseBody struct {
	StatusCode    int         `json:"status_code"`
	StatusMessage string      `json:"status_message"`
	Description   string      `json:"description"`
	Count         int64       `json:"count"`
	Offset        int64       `json:"offset"`
	Limit         int64       `json:"limit"`
	Href          string      `json:"href"`
	Payload       interface{} `json:"payload"`
}

// Response struct
type Response struct {
	Body ResponseBody
	Err  error
}

// ServeJSON main method for response
func (c *Response) ServeJSON(w http.ResponseWriter, r *http.Request) {

	defer func() {
		b, err := json.Marshal(c.Body)

		if err != nil {
			log.Println(err)
		}

		_, err = w.Write(b)

		if err != nil {
			log.Println(err)
		}
	}()

	w.Header().Add("Content-Type", "application/json")
	c.Body.Href = r.RequestURI

	if c.Err != nil {

		c.Body.StatusMessage = "Error"

		switch c.Err.Error() {
		case "sql: no rows in result set":
			c.Body.Description = "Data tidak ditemukan"
			c.Body.StatusCode = 404
			w.WriteHeader(404)
		default:
			c.Body.Description = c.Err.Error()
			c.Body.StatusCode = 403
			w.WriteHeader(403)
		}

	} else {

		c.Body.StatusMessage = "Success"
		c.Body.StatusCode = 200
		c.Body.Limit = 10

		if v := r.URL.Query().Get("offset"); v != "" {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				log.Println(err)
			}
			c.Body.Offset = int64(vInt)
		}

		if v := r.URL.Query().Get("limit"); v != "" {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				log.Println(err)
			}
			c.Body.Limit = int64(vInt)
		}

	}

}
