package api;
import (
	"sort"
);

type Repository struct {
	transactions      []*Transaction
}

func (x *Repository) AddCredit(transaction *Transaction) {
	x.transactions = append(x.transactions, transaction);
	for _, y := range x.Transactions() {
		for _, z := range x.Transactions() {
			if y.Payer == z.Payer && y.Points < 0 && z.Timestamp.Sub(y.Timestamp).Seconds() < 0.0 {
				var diff = z.Points - (y.Points + z.Points);
				y.Points += diff;
				z.Points -= diff;
			}
		}
	}
}

func (x *Repository) Spend(points int) []*Credit {
	var rawCredits = map[string]int{};
	var result = []*Credit{};
	for _, cursor := range x.Transactions() {
		var diff = 0;
		if points < 1 { break; }
		if cursor.Points < 0 { continue; }
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
		return x.transactions[l].Timestamp.Sub(x.transactions[r].Timestamp).Seconds() < 0.0;
	});
	return x.transactions;
}
