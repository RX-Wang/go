package pkg1

/**
	时间类型 
*/

/* import (
	"fmt"
	"time"
) */

/* var week time.Duration

func init() {
	t := time.Now()
	fmt.Println("1--:", t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("2---:%4d-%02d-%02d\n", t.Year(), t.Month(), t.Day())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println("3--",t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println("4--",time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(week)
	fmt.Println("5--",week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println("6--",t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println("7--",t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	fmt.Println("8--",t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println("9--", t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
} */