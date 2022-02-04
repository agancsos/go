package api;
import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

var VERSION    = "1.0.0.0";
var repository = &Repository{};
var server     = gin.Default();

func spend(ctx *gin.Context) {
	var body, err = ioutil.ReadAll(ctx.Request.Body);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"result": fmt.Sprintf("{\"result\":\"Failed to extract points. %v\"}", err)});
		return;
	}
	var points *Points;
	json.Unmarshal(body, &points);
	var rsp = repository.Spend(points.Points);
	ctx.JSON(http.StatusOK, rsp);
}

func getBalance(ctx *gin.Context) {
	var rsp = repository.Balance();
	ctx.JSON(http.StatusOK, rsp);
}

func addCredit(ctx *gin.Context) {
	var transaction *Transaction;
	var body, err = ioutil.ReadAll(ctx.Request.Body);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"result": fmt.Sprintf("Failed to extract transaction. %v", err)});
		return;
	}
	json.Unmarshal(body, &transaction);
	repository.AddCredit(transaction);
	ctx.JSON(http.StatusOK, gin.H{"result": "1"});
}

func getVersion(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("{\"result\":\"%s\"}", VERSION)});
}

func StartServer(port int) {
	server.GET("/version", getVersion);
	server.GET("/balance", getBalance);
	server.POST("/credit", addCredit);
	server.POST("/spend", spend);
	server.Run(fmt.Sprintf("0.0.0.0:%d", port));
}

