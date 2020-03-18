package main

import "fmt"

type books struct {
	book_name string
	isbn      int
}

func main() {
	var book1 books
	var book2 books

	var book3 books
	var book4 books
	book3.book_name = " aflkdsajfkld"
	book3.isbn = 32
	book4.book_name = " khjghkgj"
	book4.isbn = 23

	book1.book_name = "GO Programming Language"
	book2.book_name = "Structure topic"

	book1.isbn = 12345
	book2.isbn = 123

	//ar book4 books = books{"testing", 999}
	//book3 := books{"Neway", 456}
	/*fmt.Println("Book 3 is ", book3.isbn)
	fmt.Println("Book 4 is ", book4.isbn)
	fmt.Println(book1.book_name)
	fmt.Println(book1.isbn)
	fmt.Println(book2.book_name)
	fmt.Println(book2.isbn)
	fmt.Println(book1, book2, book3, book4)
	*/
	//for i := 0; i < 4; i++ {
	//var name []string = []string{"book1", "book2", "book3", "book4"}
	//name := []string{"book1", "book2", "book3", "book4"}
	//temp := name[i]
	//var name string
	//bookname := name +i
	//fmt.Println(temp)
	printBook(book1)
	printBook(book2)
	printBook(book3)
	printBook(book4)
	var abc int = 5
	fmt.Println("answer is ", factorial(abc))

	//}
	//printBook(book1 book2 book3 book4)

}
func printBook(bookA books) {
	fmt.Println("Books itle is", bookA.book_name)
	fmt.Println("isbn Number is", bookA.isbn)

}

func factorial(i int) int {
	if i <= 1 {
		return i
	}
	return i * factorial((i - 1))
}
