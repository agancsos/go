package main;
import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"io/ioutil"
	"sort"
	"strconv"
);

// Models
type Credit struct {
	Payer         string        `json:payer`
	Points        int           `json:points`
	Timestamp     time.Time     `json:timestamp`
}

type Transaction struct {
	Payer         string  `json:payer`
	Points        int     `json:points`
}

type Points struct {
	Points       int  `json:points`
}

type Repository struct {
	credits      []*Credit
}

func (x *Repository) AddCredit(credit *Credit) {
	for _, cursor := range x.Credits() {
		// Pay off any dues to the payer
		if cursor.Payer == credit.Payer && cursor.Points < 0 && credit.Points >= cursor.Points {
			var diff = credit.Points - (credit.Points + cursor.Points);
			cursor.Points = cursor.Points + diff;
			credit.Points = credit.Points - diff;
		}
	}
	x.credits = append(x.credits, credit);
}

func (x *Repository) Spend(points int) []*Transaction {
	var rawTransactions = map[string]int{};
	var result = []*Transaction{};
	for _, cursor := range x.Credits() {
		var diff = 0;
		if points < 1 { break; }
		// Calculate balance
		if cursor.Points > points {
			diff = cursor.Points - (cursor.Points - points);
		} else {
			diff = points - (points - cursor.Points);
		}
		rawTransactions[cursor.Payer] -= diff;
		cursor.Points -= diff;  // Update available credit
		points -= diff;
	}
	for payer, point := range rawTransactions {
		result = append(result, &Transaction{Payer:payer, Points:point});
	}
	return result;
}

func (x *Repository) Balance() map[string]int {
	var result = map[string]int{};
	for _, credit := range x.credits {
		result[credit.Payer] += credit.Points;
	}
	return result;
}

func (x *Repository) Credits() []*Credit {
	sort.Slice(x.credits, func(l int, r int) bool {
		return x.credits[r].Timestamp.Sub(x.credits[l].Timestamp) > 0;
	});
	return x.credits;
}
/*****************************************************************************/

// Globals
var VERSION      = "1.0.0.0";
var REPOSITORY   = &Repository{};
/*****************************************************************************/

// REST handlers
func Spend(w http.ResponseWriter, r *http.Request) {
	var body, err = ioutil.ReadAll(r.Body);
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{\"result\":\"Failed to extract points. %v\"}", err)));
		return;
	}
	var points *Points;
	json.Unmarshal(body, &points);
	var rsp = REPOSITORY.Spend(points.Points);
	result, err := json.Marshal(rsp);
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{\"result\":\"Failed to serialize result. %v\"}", err)));
		return;
	}
	w.Write(result);
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	var rsp = REPOSITORY.Balance();
	result, err := json.Marshal(rsp);
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{\"result\":\"Failed to serialize result. %v\"}", err)));
		return;
	}
	w.Write(result);
}

func AddCredit(w http.ResponseWriter, r *http.Request) {
	var credit *Credit;
	var body, err = ioutil.ReadAll(r.Body);
	if err != nil {
		w.Write([]byte(fmt.Sprintf("{\"result\":\"Failed to extract transaction. %v\"}", err)));
		return;
	}
	json.Unmarshal(body, &credit);
	REPOSITORY.AddCredit(credit);
	w.Write([]byte("{\"result\":\"1\"}"));
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("{\"result\":\"%s\"}", VERSION)));
}
/*****************************************************************************/

func main() {
	var servicePort = 4441;

	for i, _ := range os.Args {
		switch (os.Args[i]) {
			case "-p", "--port":
				servicePort, _ = strconv.Atoi(os.Args[i + 1]);
				break;
		}
	}

	http.HandleFunc("/credit", AddCredit);
	http.HandleFunc("/spend", Spend);
	http.HandleFunc("/balance", GetBalance);
	http.HandleFunc("/version", GetVersion);

	// Start listener
	http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil);

	os.Exit(0);
}
