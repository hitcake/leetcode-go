package discountPrices

import (
	"testing"
)

func TestDiscountPrices(t *testing.T) {
	var tests = []struct {
		sentence string
		discount int
		want     string
	}{
		{"there are $1 $2 and $5 candies in the shop", 50, "there are $0.50 $1.00 and $2.50 candies in the shop"},
		{"1 2 $3 4 $5 $6 7 8$ $9 $10$", 100, "1 2 $0.00 4 $0.00 $0.00 7 8$ $0.00 $10$"},
		{"$$ $10 123 good", 50, "$$ $5.00 123 good"},
		{"ka3caz4837h6ada4 r1 $602", 9, "ka3caz4837h6ada4 r1 $547.82"},
		{"$7383692 5q $5870426", 64, "$2658129.12 5q $2113353.36"},
		{"$2$3 $10 $100 $1 200 $33 33$ $$ $99 $99999 $9999999999", 0, "$2$3 $10.00 $100.00 $1.00 200 $33.00 33$ $$ $99.00 $99999.00 $9999999999.00"},
		{"7mozebb9smpwz$$yqr4mox3uae1a210prubb5zp1dykv$ffezb4jpfpcv5hldxnuob66bmu17g$aceym5vszi1$re1v4ttspua6b8yxtbfwnmyk8tudx1qj1ahxbseidrauclql3$ph$pj g5q61b01ho2k9c8fzdasxqvufyms66stvq2 $3238691891 ph3fw6mw $8130 t1 $8001261 $yrdnj9pek7yr1ccujc756i44qk5mr6h64zizbazgyx0k0$0 vmhb4r31ee2futqh76co5eff",
			21,
			"7mozebb9smpwz$$yqr4mox3uae1a210prubb5zp1dykv$ffezb4jpfpcv5hldxnuob66bmu17g$aceym5vszi1$re1v4ttspua6b8yxtbfwnmyk8tudx1qj1ahxbseidrauclql3$ph$pj g5q61b01ho2k9c8fzdasxqvufyms66stvq2 $2558566593.89 ph3fw6mw $6422.70 t1 $6320996.19 $yrdnj9pek7yr1ccujc756i44qk5mr6h64zizbazgyx0k0$0 vmhb4r31ee2futqh76co5eff"},
	}
	for _, test := range tests {
		got := discountPrices(test.sentence, test.discount)
		if got != test.want {
			t.Errorf("sentence %s\n discount %d\n got: %s\n want: %s\n", test.sentence, test.discount, got, test.want)
		}
	}
}
