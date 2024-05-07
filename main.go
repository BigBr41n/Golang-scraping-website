package main 


import (
	"os"
	"encoding/csv"
	"strings"
	"fmt"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

//check error 
func checkError(err error){
	if err!= nil {
		//fmt.Println("Error: ", err.Error())
        panic(err)
    }
}


//write file 
/* func writeFile(filename , data string) {
    file, err := os.Create(filename)
    checkError(err)
    defer file.Close()
    file.WriteString(data)
} */



func main () {
	url := "https://techcrunch.com" ;

	//get request to the url and either the response or the error
	response , error := http.Get(url); 

	//check for any errors 
	checkError(error)

	//bad requests handling
	if response.StatusCode > 400 {
		fmt.Println("Status code: ", response.StatusCode)
	}

	//read the response body
	defer response.Body.Close()


	doc , error := goquery.NewDocumentFromReader(response.Body)
	checkError(error)


	//creating csv file to store posts
	file, err := os.Create("posts.csv")
	checkError(err)



	writer := csv.NewWriter(file)


	doc.Find("div.river").Find("div.post-block").Each(func (index int , item * goquery.Selection){
		h2 := item.Find("h2")
		title := strings.TrimSpace(h2.Text())
		url , _ := h2.Find("a").Attr("href")

		excerpt := strings.TrimSpace(item.Find("div.post-block__content").Text())

		fmt.Println(title , url , excerpt)

		posts := []string{title, url, excerpt}


		writer.Write(posts)
	}) 
	checkError(error)


	writer.Flush()

	//writeFile("river.html", river)

	//fmt.Println(river)

}