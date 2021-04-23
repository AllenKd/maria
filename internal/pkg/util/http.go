package util

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"maria/internal/pkg/db"
	"maria/internal/pkg/error"
	"strconv"
)

func SendHttpRequest(req *http.Request) (response []byte, err *error.Error) {
	client := http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		log.Error("fail to send request: ", httpErr.Error())
		err = error.HttpOperationFail.WithDetailAndStatus(httpErr.Error(), http.StatusInternalServerError)
		return
	}

	response, _ = ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Error("unexpected response code: ", resp.StatusCode)
		log.Error("response: ", string(response))
		err = error.HttpUnexpectedResponseCode.WithDetailAndStatus(resp.Status, http.StatusInternalServerError)
		return
	}
	return
}

func ParseQueryOption(c *gin.Context) db.QueryOption {
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	skip, _ := strconv.ParseInt(c.Query("skip"), 10, 64)
	return db.QueryOption{
		Limit:  limit,
		Skip:   skip,
		SortBy: c.Query("sortBy"),
		Desc:   parseDesc(c.Query("desc")),
	}
}

func parseDesc(desc string) bool {
	return desc == "true"
}
