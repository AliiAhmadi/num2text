/*
 * Ali Ahmadi
 * Dedicated to MHM
 */

package num2text

type collection struct {
	ones      map[int]string
	tens      map[int]string
	hundreds  map[int]string
	thousands map[int]string
	n         string
	v         string
	m         string
}

var coll = collection{
	ones: map[int]string{
		0:  "صفر",
		1:  "یک",
		2:  "دو",
		3:  "سه",
		4:  "چهار",
		5:  "پنج",
		6:  "شش",
		7:  "هفت",
		8:  "هشت",
		9:  "نه",
		10: "ده",
		11: "یازده",
		12: "دوازده",
		13: "سیزده",
		14: "چهارده",
		15: "پانزده",
		16: "شانزده",
		17: "هفده",
		18: "هجده",
		19: "نوزده",
	},
	tens: map[int]string{
		0: "",
		1: "ده",
		2: "بیست",
		3: "سی",
		4: "چهل",
		5: "پنجاه",
		6: "شصت",
		7: "هفتاد",
		8: "هشتاد",
		9: "نود",
	},
	hundreds: map[int]string{
		0: "",
		1: "صد",
		2: "دویست",
		3: "سیصد",
		4: "چهارصد",
		5: "پانصد",
		6: "ششصد",
		7: "هفتصد",
		8: "هشتصد",
		9: "نهصد",
	},
	thousands: map[int]string{
		0:  "",
		1:  "هزار",
		2:  "میلیون",
		3:  "میلیارد",
		4:  "تریلیون",
		5:  "کوادریلیون",
		6:  "کوینتیلیون",
		7:  "سکستیلیون",
		8:  "سپتیلیون",
		9:  "اکتیلیون",
		10: "نونیلیون",
		11: "دسیلیون",
	},
	n: "منفی ",
	v: " و ",
	m: "م",
}
