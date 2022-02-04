package api;
import (
	"sort"
);

type Repository struct {
	transactions      []*Transaction
}

func (x *Repository) AddCredit(transaction *Transaction) {
	for _, cursor := range x.Transactions() {
		// Pay off any dues to the payer
		if cursor.Payer == transaction.Payer && cursor.Points < 0 && transaction.Points >= cursor.Points {
			var diff = transaction.Points - (transaction.Points + cursor.Points);
			cursor.Points = cursor.Points + diff;
			transaction.Points = transaction.Points - diff;
		}
	}
	x.transactions = append(x.transactions, transaction);
}

func (x *Repository) Spend(points int) []*Credit {
	var rawCredits = map[string]int{};
	var result = []*Credit{};
	for _, cursor := range x.Transactions() {
		var diff = 0;
		if points < 1 { break; }
		// Calculate balance
		if cursor.Points > points {
			diff = cursor.Points - (cursor.Points - points);
		} else {
			diff = points - (points - cursor.Points);
		}
		rawCredits[cursor.Payer] -= diff;
		cursor.Points -= diff;  // Update available credit
		points -= diff;
	}
	for payer, point := range rawCredits {
		result = append(result, &Credit{Payer:payer, Points:point});
	}
	return result;
}

func (x *Repository) Balance() map[string]int {
	var result = map[string]int{};
	for _, credit := range x.transactions {
		result[credit.Payer] += credit.Points;
	}
	return result;
}

func (x *Repository) Transactions() []*Transaction {
	sort.Slice(x.transactions, func(l int, r int) bool {
		return x.transactions[r].Timestamp.Sub(x.transactions[l].Timestamp) > 0;
	});
	return x.transactions;
}
