package arrayProblem

import "fmt"

type Vendor struct {
	start int
	end   int
	price int
}

func TakeInput(in int) (vendors []Vendor) {
	vendors := make([]Vendor, in)
	for i := 0; i < in; i++ {
		var st, en, pr int
		_, _ = fmt.Scan(&st, &en, &pr)
		v := Vendor{start: st, end: en, price: pr}
		vendors = append(vendors, v)
	}

	return vendors
}
func FindMinimumSellerPrice(vendors []Vendor, st, en int) {
	filterVendors := make([]Vendor, len(vendors))
	durationMap := make(map[int]string, 0)
	for i := 0; i < len(vendors); i++ {
		for j := 1; j < len(vendors)-1; j++ {

		}
		fmt.Println(filterVendors[i].start, filterVendors[i].end)
	}
}
