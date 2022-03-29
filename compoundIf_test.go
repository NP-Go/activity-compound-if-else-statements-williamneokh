package main

import "testing"

func TestCheckYear(t *testing.T) {
	get := checkYear(2000)
	res := "2000 is a leap year"
	if get != res {
		t.Errorf("want %s but get %s", get, res)
	}
	get = checkYear(1912)
	res = "1912 is a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(2088)
	res = "2088 is a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(1972)
	res = "1972 is a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(2020)
	res = "2020 is a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(1988)
	res = "1988 is a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(2100)
	res = "2100 is not a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(1900)
	res = "1900 is not a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(2300)
	res = "2300 is not a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(2003)
	res = "2003 is not a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
	get = checkYear(2323)
	res = "2323 is not a leap year"
	if get != res {
		t.Errorf("want %s but get %s", res, get)
	}
}
