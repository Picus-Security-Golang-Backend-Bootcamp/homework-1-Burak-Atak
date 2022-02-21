package main

import (
	"bufio"
	"contains"
	"fmt"
	"os"
	"strings"
)

func create_book_list() []string {
	var book_list []string
	file, _ := os.Open("books.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		book_list = append(book_list, scanner.Text())
	}

	defer file.Close()

	return book_list
}

func take_input() []string {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name_list := strings.Split(name, " ")
	name_list[len(name_list)-1] = strings.TrimSpace(name_list[len(name_list)-1])

	return name_list
}

func main() {
	book_list := create_book_list()

	for true {
		fmt.Print("Kitapları listelemek için 'list' aramak için 'search -kitap adi-' yaziniz: ")
		name_list := take_input()

		if name_list[0] == "list" {
			if len(name_list) > 1 {
				fmt.Println("Kitaplari listelemek için sadece 'list' yazin.")
				continue
			}
			for _, element := range book_list {
				fmt.Println(element)
			}
			break

		} else if name_list[0] == "search" {
			if len(name_list) == 1 {
				fmt.Println("Aramak için 'search'den sonra kitap ismi yaziniz.")
				continue
			}
			book_name := strings.Join(name_list[1:], " ")
			books, is_true := contains.Contains(book_list, book_name)
			fmt.Println(books)
			if is_true {
				fmt.Printf("İcerisinde '%s' bulunan kitapların listesi:\n", book_name)
				for k, v := range books {
					fmt.Printf("Kitap adi: '%s' listedeki sirasi: '%d'\n", k, v)
				}
			} else {
				fmt.Println("Kitap mevcut degil.")
			}
			break

		}
	}

	fmt.Print("Tekrar arama yapmak istiyorsanız 'y' yazınız: ")
	use_again := take_input()

	if use_again[0] == "y" {
		main()
	}

}
