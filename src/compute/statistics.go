/*
Statisitcs transaction data.
*/

package compute

import (
		"fmt"
		"container/list"
		"strings"
        "strconv"
		"os"
	   )

func Statistics(contents *list.List) map[string]int {
    var statistics = make(map[string]int)
	statistics["small_sale"] = 0
	statistics["middle_sale"] = 0
	statistics["large_sale"] = 0
	statistics["small_buy"] = 0
	statistics["middle_buy"] = 0
	statistics["large_buy"] = 0
	for content := contents.Front(); content != nil; content = content.Next() {
        var c = content.Value
		if strings.Contains(c.(string), "S") {
            ct := strings.Split(c.(string), ",")
			handers, err1 := strconv.ParseFloat(strings.Split(ct[1], "\"")[1], 64)
			price, err2 := strconv.ParseFloat(strings.Split(ct[2], "\"")[1], 64)
    		if (err1 != nil) {
        		fmt.Println(err1)
        		os.Exit(2)
    		} else if (err2 != nil) {
        		fmt.Println(err2)
        		os.Exit(2)
			}
           	sum_sale := int(handers * price * 100)
			if sum_sale < 200000 {
				statistics["small_sale"] = statistics["small_sale"] + sum_sale
			} else if sum_sale > 1000000 {
				statistics["large_sale"] = statistics["large_sale"] + sum_sale
			} else {
				statistics["middle_sale"] = statistics["middle_sale"] + sum_sale
			}
		} else if strings.Contains(c.(string), "B") {
            ct := strings.Split(c.(string), ",")
			handers, err1 := strconv.ParseFloat(strings.Split(ct[1], "\"")[1], 64)
			price, err2 := strconv.ParseFloat(strings.Split(ct[2], "\"")[1], 64)
    		if (err1 != nil) {
        		fmt.Println(err1)
        		os.Exit(2)
    		} else if (err2 != nil) {
        		fmt.Println(err2)
        		os.Exit(2)
			}
           	sum_buy := int(handers * price * 100)
			if sum_buy < 200000 {
				statistics["small_buy"] = statistics["small_buy"] + sum_buy
			} else if sum_buy > 1000000 {
				statistics["large_buy"] = statistics["large_buy"] + sum_buy
			} else {
				statistics["middle_buy"] = statistics["middle_buy"] + sum_buy
			}
		}
    }
	statistics["large"] = statistics["large_buy"] - statistics["large_sale"]
	statistics["middle"] = statistics["middle_buy"] - statistics["middle_sale"]
	statistics["small"] = statistics["small_buy"] - statistics["small_sale"]
	statistics["total"] = statistics["small_buy"] + statistics["middle_buy"] + statistics["large_buy"] - statistics["small_sale"] - statistics["middle_sale"] - statistics["large_sale"]
	return statistics
}
