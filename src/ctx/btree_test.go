package ctx

import (
	"algs4/src/rabbit"
	"fmt"
	"testing"
)

func TestBTree(t *testing.T) {
	//fmt.Println(rabbit.Less(rabbit.ComparableString("www.cs.princeton.edu"), rabbit.ComparableString("www.princeton.edu")))
	//fmt.Println(rabbit.Less(rabbit.ComparableString("www.princeton.edu"), rabbit.ComparableString("www.cs.princeton.edu")))
	//fmt.Println(rabbit.Equal(rabbit.ComparableString("www.cs.princeton.edu"), rabbit.ComparableString("www.princeton.edu")))
	//fmt.Println(rabbit.Equal(rabbit.ComparableString("www.cs.princeton.edu"), rabbit.ComparableString("www.cs.princeton.edu")))

	btree := NewBTree()

	btree.Put(rabbit.ComparableString("www.yale.edu"), "130.132.143.21")
	btree.Put(rabbit.ComparableString("www.princeton.edu"), "128.112.128.15")
	btree.Put(rabbit.ComparableString("www.cs.princeton.edu"), "128.112.136.12")
	btree.Put(rabbit.ComparableString("www.simpsons.com"), "209.052.165.60")

	btree.Put(rabbit.ComparableString("www.apple.com"), "17.112.152.32")
	btree.Put(rabbit.ComparableString("www.amazon.com"), "207.171.182.16")
	btree.Put(rabbit.ComparableString("www.ebay.com"), "66.135.192.87")
	btree.Put(rabbit.ComparableString("www.cnn.com"), "64.236.16.20")
	btree.Put(rabbit.ComparableString("www.google.com"), "199.239.136.200")

	fmt.Println("size:", btree.n)
	fmt.Println("height:", btree.height)
	fmt.Println(btree)
	err := btree.Validate()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("cs.princeton.edu:  ", btree.Get(rabbit.ComparableString("www.cs.princeton.edu")))
	fmt.Println("simpsons.com:      ", btree.Get(rabbit.ComparableString("www.simpsons.com")))
	fmt.Println("amazon.com:        ", btree.Get(rabbit.ComparableString("www.amazon.com")))
	fmt.Println("cnn.com:           ", btree.Get(rabbit.ComparableString("www.cnn.com")))

	btree.Put(rabbit.ComparableString("www.google.com"), nil)
	btree.Put(rabbit.ComparableString("www.cnn.com"), "64.236.16.20:999999")
	btree.Put(rabbit.ComparableString("www.ebay.com"), nil)
	btree.Put(rabbit.ComparableString("www.ebay.com"), nil)
	btree.Put(rabbit.ComparableString("www.cs.princeton.edu"), nil)

	fmt.Println("cs.princeton.edu:  ", btree.Get(rabbit.ComparableString("www.cs.princeton.edu")))
	fmt.Println("simpsons.com:      ", btree.Get(rabbit.ComparableString("www.simpsons.com")))
	fmt.Println("amazon.com:        ", btree.Get(rabbit.ComparableString("www.amazon.com")))
	fmt.Println("cnn.com:           ", btree.Get(rabbit.ComparableString("www.cnn.com")))

	fmt.Println("size:", btree.n)
	fmt.Println("height:", btree.height)
	fmt.Println(btree)

	//st.put("www.cs.princeton.edu", "128.112.136.12");
	//st.put("www.cs.princeton.edu", "128.112.136.11");
	//st.put("www.princeton.edu",    "128.112.128.15");
	//st.put("www.yale.edu",         "130.132.143.21");
	//st.put("www.simpsons.com",     "209.052.165.60");
	//st.put("www.apple.com",        "17.112.152.32");
	//st.put("www.amazon.com",       "207.171.182.16");
	//st.put("www.ebay.com",         "66.135.192.87");
	//st.put("www.cnn.com",          "64.236.16.20");
	//st.put("www.google.com",       "216.239.41.99");
	//st.put("www.nytimes.com",      "199.239.136.200");
	//st.put("www.microsoft.com",    "207.126.99.140");
	//st.put("www.dell.com",         "143.166.224.230");
	//st.put("www.slashdot.org",     "66.35.250.151");
	//st.put("www.espn.com",         "199.181.135.201");
	//st.put("www.weather.com",      "63.111.66.11");
	//st.put("www.yahoo.com",        "216.109.118.65");
}
