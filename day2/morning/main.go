package main

import "fmt";

func main() {
	text := "Anak Kecil Lompat Lompat";
	fmt.Println(AlayGen(text));
	soal2 := fibonancci(6)
	fmt.Println(soal2);
}

func AlayGen(words ...string) string {
	result := ""
	for _, word := range words {
		result += word + " ";
	}
	return changeChar(result);
}

func changeChar(text string) string {
	result := ""
	for _, char := range text {
		switch string(char) {
		case "a":
			result += "4"
		case "e":
			result += "3"
		case "i":
			result += "!"
		case "l":
			result += "1"
		case "n":
			result += "N"
		case "s":
			result += "5"
		case "x":
			result += "*"
		default:
			result += string(char)
		}
	}
	return result;
}

func fibonancci(n int) int {
	a , b := 0 , 1;
	result := 0;
	for i := 0; i < n + 1; i++ {
		if i < 2 {
			result = i;
			continue;
		} 
		result = a + b;

		if i % 2 == 0 {
			a = result;
			continue;
		}
		b = result;
	}
	return result;
}
