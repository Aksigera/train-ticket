package main

type News struct {
	Title   string `bson:"Title"`
	Content string `bson:"Content"`
}

func hello(val string) string {
	var str = []byte(`[
                       {"Title": "News Service Complete", "Content": "Congratulations:Your News Service Complete"},
                       {"Title": "Total Ticket System Complete", "Content": "Just a total test"}
                    ]`)
	return string(str)
}

func main() {
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:12862")
}
