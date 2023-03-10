
package main

import (
"bufio"
"fmt"
"os"
"strings"
)
func main() {

	var resr string
	var x, y int
	var a_t, b_t string
	var num = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "-1", "-2", "-3", "-4", "-5", "-6", "-7", "-8", "-9", "-10"}
	var numr = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "", "", "", "", "", "", "", "", "", "", ""}
	fmt.Println("Добро пожаловать в калькулятор Асимана Гулиева")
	fmt.Println()
	fmt.Println("Обрати внимание, что данный калькулятор может работать только с числами от 1 до 10 (как с арабскими так и с римскими)")
	fmt.Println()
	fmt.Println("Можешь ввести операцию:")
	reader := bufio.NewReader(os.Stdin)

	a, _ := reader.ReadString('\n')
	a = strings.TrimSpace(a)
	aa := strings.Split(a, " ")

	for i := 0; i < 20; i++ {
		if num[i] == aa[0] {
			a_t = "num"
		}
	}
	for j := 0; j < 20; j++ {
		if numr[j] == aa[0] {
			a_t = "numr"
		}
	}
	for k := 0; k < 20; k++ {
		if num[k] == aa[2] {
			b_t = "num"
		}
	}
	for h := 0; h < 20; h++ {
		if numr[h] == aa[2] {
			b_t = "numr"
		}
	}

	if aa[1] != "+" && aa[1] != "/" && aa[1] != "*" && aa[1] != "-" {
		fmt.Println("Неизвестная операция. Калькулятор работает только с операциями   +, /, *, -")
	}
	if b_t != a_t {
		fmt.Println("Ой! Произошла ошибка. Нужно выбрать с какими цифрами работаем (с арабскими или римскими). Цифры должны быть в интервале от единицы до десяти :)")
	}
	if len(aa) != 3 {
		fmt.Println("Введено более двух операндов")
	}
	if (a_t != "num" && a_t != "numr") || (b_t != "num" && b_t != "numr") {
		fmt.Println("Должны быть введены арабские или римские цифры от 1 до 10")
	}

	if (a_t == "num" && b_t == "num") || (a_t == "numr" && b_t == "numr") {
		if len(aa) == 3 {
			if b_t == a_t {
				if aa[1] == "+" || aa[1] == "/" || aa[1] == "*" || aa[1] == "-" {

					if b_t == "num" {
						if aa[0] == "1" {
							x = 1
						} else if aa[0] == "2" {
							x = 2
						} else if aa[0] == "3" {
							x = 3
						} else if aa[0] == "4" {
							x = 4
						} else if aa[0] == "5" {
							x = 5
						} else if aa[0] == "6" {
							x = 6
						} else if aa[0] == "7" {
							x = 7
						} else if aa[0] == "8" {
							x = 8
						} else if aa[0] == "9" {
							x = 9
						} else if aa[0] == "10" {
							x = 10
						} else if aa[0] == "-1" {
							x = -1
						} else if aa[0] == "-2" {
							x = -2
						} else if aa[0] == "-3" {
							x = -3
						} else if aa[0] == "-4" {
							x = -4
						} else if aa[0] == "-5" {
							x = -5
						} else if aa[0] == "-6" {
							x = -6
						} else if aa[0] == "-7" {
							x = -7
						} else if aa[0] == "-8" {
							x = -8
						} else if aa[0] == "-9" {
							x = -9
						} else if aa[0] == "-10" {
							x = -10
						}
						if aa[2] == "1" {
							y = 1
						} else if aa[2] == "2" {
							y = 2
						} else if aa[2] == "3" {
							y = 3
						} else if aa[2] == "4" {
							y = 4
						} else if aa[2] == "5" {
							y = 5
						} else if aa[2] == "6" {
							y = 6
						} else if aa[2] == "7" {
							y = 7
						} else if aa[2] == "8" {
							y = 8
						} else if aa[2] == "9" {
							y = 9
						} else if aa[2] == "10" {
							y = 10
						} else if aa[2] == "-1" {
							y = -1
						} else if aa[2] == "-2" {
							y = -2
						} else if aa[2] == "-3" {
							y = -3
						} else if aa[2] == "-4" {
							y = -4
						} else if aa[2] == "-5" {
							y = -5
						} else if aa[2] == "-6" {
							y = -6
						} else if aa[2] == "-7" {
							y = -7
						} else if aa[2] == "-8" {
							y = -8
						} else if aa[2] == "-9" {
							y = -9
						} else if aa[2] == "-10" {
							y = -10
						}
						if aa[1] == "+" {
							res := x + y
							fmt.Println("Вот что получилось: ", a, "=", res)

						} else if aa[1] == "-" {
							res := x - y
							fmt.Println("Вот что получилось: ", a, "=", res)
						} else if aa[1] == "*" {
							res := x * y
							fmt.Println("Вот что получилось: ", a, "=", res)
						} else if aa[1] == "/" {
							res := x / y
							fmt.Println("Вот что получилось: ", a, "=", res)
						}
					}
					if b_t == "numr" {
						if aa[0] == "I" {
							x = 1
						} else if aa[0] == "II" {
							x = 2
						} else if aa[0] == "III" {
							x = 3
						} else if aa[0] == "IV" {
							x = 4
						} else if aa[0] == "V" {
							x = 5
						} else if aa[0] == "VI" {
							x = 6
						} else if aa[0] == "VII" {
							x = 7
						} else if aa[0] == "VIII" {
							x = 8
						} else if aa[0] == "IX" {
							x = 9
						} else if aa[0] == "X" {
							x = 10
						}
						if aa[2] == "I" {
							y = 1
						} else if aa[2] == "II" {
							y = 2
						} else if aa[2] == "III" {
							y = 3
						} else if aa[2] == "IV" {
							y = 4
						} else if aa[2] == "V" {
							y = 5
						} else if aa[2] == "VI" {
							y = 6
						} else if aa[2] == "VII" {
							y = 7
						} else if aa[2] == "VIII" {
							y = 8
						} else if aa[2] == "IX" {
							y = 9
						} else if aa[2] == "X" {
							y = 10
						}
						if aa[1] == "+" {
							res := x + y
							if res == 1 {
								resr = "I"
							} else if res == 2 {
								resr = "II"
							} else if res == 3 {
								resr = "III"
							} else if res == 4 {
								resr = "IV"
							} else if res == 5 {
								resr = "V"
							} else if res == 6 {
								resr = "VI"
							} else if res == 7 {
								resr = "VII"
							} else if res == 8 {
								resr = "VIII"
							} else if res == 9 {
								resr = "IX"
							} else if res == 10 {
								resr = "X"
							} else if res == 11 {
								resr = "XI"
							} else if res == 12 {
								resr = "XII"
							} else if res == 13 {
								resr = "XIII"
							} else if res == 14 {
								resr = "XIV"
							} else if res == 15 {
								resr = "XV"
							} else if res == 16 {
								resr = "XVI"
							} else if res == 17 {
								resr = "XVII"
							} else if res == 18 {
								resr = "XVIII"
							} else if res == 19 {
								resr = "XIX"
							} else if res == 20 {
								resr = "XX"
							} else if res == 21 {
								resr = "XXI"
							} else if res == 22 {
								resr = "XXII"
							} else if res == 23 {
								resr = "XXIII"
							} else if res == 24 {
								resr = "XXIV"
							} else if res == 25 {
								resr = "XXV"
							} else if res == 26 {
								resr = "XXVI"
							} else if res == 27 {
								resr = "XXVII"
							} else if res == 28 {
								resr = "XXVIII"
							} else if res == 29 {
								resr = "XXIX"
							} else if res == 30 {
								resr = "XXX"
							} else if res == 31 {
								resr = "XXXI"
							} else if res == 32 {
								resr = "XXXII"
							} else if res == 33 {
								resr = "XXXIII"
							} else if res == 34 {
								resr = "XXXIV"
							} else if res == 35 {
								resr = "XXXV"
							} else if res == 36 {
								resr = "XXXVI"
							} else if res == 37 {
								resr = "XXXVII"
							} else if res == 38 {
								resr = "XXXVIII"
							} else if res == 39 {
								resr = "XXXIX"
							} else if res == 40 {
								resr = "XL"
							} else if res == 41 {
								resr = "XLI"
							} else if res == 42 {
								resr = "XLII"
							} else if res == 43 {
								resr = "XLIII"
							} else if res == 44 {
								resr = "XLIV"
							} else if res == 45 {
								resr = "XLV"
							} else if res == 46 {
								resr = "XLVI"
							} else if res == 47 {
								resr = "XLVII"
							} else if res == 48 {
								resr = "XLVIII"
							} else if res == 49 {
								resr = "XLIX"
							} else if res == 50 {
								resr = "L"
							} else if res == 51 {
								resr = "LI"
							} else if res == 52 {
								resr = "LII"
							} else if res == 53 {
								resr = "LIII"
							} else if res == 54 {
								resr = "LIV"
							} else if res == 55 {
								resr = "LV"
							} else if res == 56 {
								resr = "LVI"
							} else if res == 57 {
								resr = "LVII"
							} else if res == 58 {
								resr = "LVIII"
							} else if res == 59 {
								resr = "LIX"
							} else if res == 60 {
								resr = "LX"
							} else if res == 61 {
								resr = "LXI"
							} else if res == 62 {
								resr = "LXII"
							} else if res == 63 {
								resr = "LXIII"
							} else if res == 64 {
								resr = "LXIV"
							} else if res == 65 {
								resr = "LXV"
							} else if res == 66 {
								resr = "LXVI"
							} else if res == 67 {
								resr = "LXVII"
							} else if res == 68 {
								resr = "LXVIII"
							} else if res == 69 {
								resr = "LXIX"
							} else if res == 70 {
								resr = "LXX"
							} else if res == 71 {
								resr = "LXXI"
							} else if res == 72 {
								resr = "LXXII"
							} else if res == 73 {
								resr = "LXXIII"
							} else if res == 74 {
								resr = "LXXIV"
							} else if res == 75 {
								resr = "LXXV"
							} else if res == 76 {
								resr = "LXXVI"
							} else if res == 77 {
								resr = "LXXVII"
							} else if res == 78 {
								resr = "LXXVIII"
							} else if res == 79 {
								resr = "LXXIX"
							} else if res == 80 {
								resr = "LXXX"
							} else if res == 81 {
								resr = "LXXXI"
							} else if res == 82 {
								resr = "LXXXII"
							} else if res == 83 {
								resr = "LXXXIII"
							} else if res == 84 {
								resr = "LXXXIV"
							} else if res == 85 {
								resr = "LXXXV"
							} else if res == 86 {
								resr = "LXXXVI"
							} else if res == 87 {
								resr = "LXXXVII"
							} else if res == 88 {
								resr = "LXXXVIII"
							} else if res == 89 {
								resr = "LXXXIX"
							} else if res == 90 {
								resr = "XC"
							} else if res == 91 {
								resr = "XCI"
							} else if res == 92 {
								resr = "XCII"
							} else if res == 93 {
								resr = "XCIII"
							} else if res == 94 {
								resr = "XCIV"
							} else if res == 95 {
								resr = "XCV"
							} else if res == 96 {
								resr = "XCVI"
							} else if res == 97 {
								resr = "XCVII"
							} else if res == 98 {
								resr = "XCVIII"
							} else if res == 99 {
								resr = "XCIX"
							} else if res == 100 {
								resr = "C"
							}
							fmt.Println("Вот что получилось: ", a, "=", resr)

						} else if aa[1] == "-" {
							res := x - y
							if res > 0 {
								if res == 1 {
									resr = "I"
								} else if res == 2 {
									resr = "II"
								} else if res == 3 {
									resr = "III"
								} else if res == 4 {
									resr = "IV"
								} else if res == 5 {
									resr = "V"
								} else if res == 6 {
									resr = "VI"
								} else if res == 7 {
									resr = "VII"
								} else if res == 8 {
									resr = "VIII"
								} else if res == 9 {
									resr = "IX"
								} else if res == 10 {
									resr = "X"
								} else if res == 11 {
									resr = "XI"
								} else if res == 12 {
									resr = "XII"
								} else if res == 13 {
									resr = "XIII"
								} else if res == 14 {
									resr = "XIV"
								} else if res == 15 {
									resr = "XV"
								} else if res == 16 {
									resr = "XVI"
								} else if res == 17 {
									resr = "XVII"
								} else if res == 18 {
									resr = "XVIII"
								} else if res == 19 {
									resr = "XIX"
								} else if res == 20 {
									resr = "XX"
								} else if res == 21 {
									resr = "XXI"
								} else if res == 22 {
									resr = "XXII"
								} else if res == 23 {
									resr = "XXIII"
								} else if res == 24 {
									resr = "XXIV"
								} else if res == 25 {
									resr = "XXV"
								} else if res == 26 {
									resr = "XXVI"
								} else if res == 27 {
									resr = "XXVII"
								} else if res == 28 {
									resr = "XXVIII"
								} else if res == 29 {
									resr = "XXIX"
								} else if res == 30 {
									resr = "XXX"
								} else if res == 31 {
									resr = "XXXI"
								} else if res == 32 {
									resr = "XXXII"
								} else if res == 33 {
									resr = "XXXIII"
								} else if res == 34 {
									resr = "XXXIV"
								} else if res == 35 {
									resr = "XXXV"
								} else if res == 36 {
									resr = "XXXVI"
								} else if res == 37 {
									resr = "XXXVII"
								} else if res == 38 {
									resr = "XXXVIII"
								} else if res == 39 {
									resr = "XXXIX"
								} else if res == 40 {
									resr = "XL"
								} else if res == 41 {
									resr = "XLI"
								} else if res == 42 {
									resr = "XLII"
								} else if res == 43 {
									resr = "XLIII"
								} else if res == 44 {
									resr = "XLIV"
								} else if res == 45 {
									resr = "XLV"
								} else if res == 46 {
									resr = "XLVI"
								} else if res == 47 {
									resr = "XLVII"
								} else if res == 48 {
									resr = "XLVIII"
								} else if res == 49 {
									resr = "XLIX"
								} else if res == 50 {
									resr = "L"
								} else if res == 51 {
									resr = "LI"
								} else if res == 52 {
									resr = "LII"
								} else if res == 53 {
									resr = "LIII"
								} else if res == 54 {
									resr = "LIV"
								} else if res == 55 {
									resr = "LV"
								} else if res == 56 {
									resr = "LVI"
								} else if res == 57 {
									resr = "LVII"
								} else if res == 58 {
									resr = "LVIII"
								} else if res == 59 {
									resr = "LIX"
								} else if res == 60 {
									resr = "LX"
								} else if res == 61 {
									resr = "LXI"
								} else if res == 62 {
									resr = "LXII"
								} else if res == 63 {
									resr = "LXIII"
								} else if res == 64 {
									resr = "LXIV"
								} else if res == 65 {
									resr = "LXV"
								} else if res == 66 {
									resr = "LXVI"
								} else if res == 67 {
									resr = "LXVII"
								} else if res == 68 {
									resr = "LXVIII"
								} else if res == 69 {
									resr = "LXIX"
								} else if res == 70 {
									resr = "LXX"
								} else if res == 71 {
									resr = "LXXI"
								} else if res == 72 {
									resr = "LXXII"
								} else if res == 73 {
									resr = "LXXIII"
								} else if res == 74 {
									resr = "LXXIV"
								} else if res == 75 {
									resr = "LXXV"
								} else if res == 76 {
									resr = "LXXVI"
								} else if res == 77 {
									resr = "LXXVII"
								} else if res == 78 {
									resr = "LXXVIII"
								} else if res == 79 {
									resr = "LXXIX"
								} else if res == 80 {
									resr = "LXXX"
								} else if res == 81 {
									resr = "LXXXI"
								} else if res == 82 {
									resr = "LXXXII"
								} else if res == 83 {
									resr = "LXXXIII"
								} else if res == 84 {
									resr = "LXXXIV"
								} else if res == 85 {
									resr = "LXXXV"
								} else if res == 86 {
									resr = "LXXXVI"
								} else if res == 87 {
									resr = "LXXXVII"
								} else if res == 88 {
									resr = "LXXXVIII"
								} else if res == 89 {
									resr = "LXXXIX"
								} else if res == 90 {
									resr = "XC"
								} else if res == 91 {
									resr = "XCI"
								} else if res == 92 {
									resr = "XCII"
								} else if res == 93 {
									resr = "XCIII"
								} else if res == 94 {
									resr = "XCIV"
								} else if res == 95 {
									resr = "XCV"
								} else if res == 96 {
									resr = "XCVI"
								} else if res == 97 {
									resr = "XCVII"
								} else if res == 98 {
									resr = "XCVIII"
								} else if res == 99 {
									resr = "XCIX"
								} else if res == 100 {
									resr = "C"
								}
								fmt.Println("Вот что получилось: ", a, "=", resr)
							} else if res <= 0 {
								fmt.Println("К сожалению при работе с римскими цифрами результат не может быть отрицательным. Попробуйте использовать арабские цифры")
							}

						} else if aa[1] == "*" {
							res := x * y
							if res == 1 {
								resr = "I"
							} else if res == 2 {
								resr = "II"
							} else if res == 3 {
								resr = "III"
							} else if res == 4 {
								resr = "IV"
							} else if res == 5 {
								resr = "V"
							} else if res == 6 {
								resr = "VI"
							} else if res == 7 {
								resr = "VII"
							} else if res == 8 {
								resr = "VIII"
							} else if res == 9 {
								resr = "IX"
							} else if res == 10 {
								resr = "X"
							} else if res == 11 {
								resr = "XI"
							} else if res == 12 {
								resr = "XII"
							} else if res == 13 {
								resr = "XIII"
							} else if res == 14 {
								resr = "XIV"
							} else if res == 15 {
								resr = "XV"
							} else if res == 16 {
								resr = "XVI"
							} else if res == 17 {
								resr = "XVII"
							} else if res == 18 {
								resr = "XVIII"
							} else if res == 19 {
								resr = "XIX"
							} else if res == 20 {
								resr = "XX"
							} else if res == 21 {
								resr = "XXI"
							} else if res == 22 {
								resr = "XXII"
							} else if res == 23 {
								resr = "XXIII"
							} else if res == 24 {
								resr = "XXIV"
							} else if res == 25 {
								resr = "XXV"
							} else if res == 26 {
								resr = "XXVI"
							} else if res == 27 {
								resr = "XXVII"
							} else if res == 28 {
								resr = "XXVIII"
							} else if res == 29 {
								resr = "XXIX"
							} else if res == 30 {
								resr = "XXX"
							} else if res == 31 {
								resr = "XXXI"
							} else if res == 32 {
								resr = "XXXII"
							} else if res == 33 {
								resr = "XXXIII"
							} else if res == 34 {
								resr = "XXXIV"
							} else if res == 35 {
								resr = "XXXV"
							} else if res == 36 {
								resr = "XXXVI"
							} else if res == 37 {
								resr = "XXXVII"
							} else if res == 38 {
								resr = "XXXVIII"
							} else if res == 39 {
								resr = "XXXIX"
							} else if res == 40 {
								resr = "XL"
							} else if res == 41 {
								resr = "XLI"
							} else if res == 42 {
								resr = "XLII"
							} else if res == 43 {
								resr = "XLIII"
							} else if res == 44 {
								resr = "XLIV"
							} else if res == 45 {
								resr = "XLV"
							} else if res == 46 {
								resr = "XLVI"
							} else if res == 47 {
								resr = "XLVII"
							} else if res == 48 {
								resr = "XLVIII"
							} else if res == 49 {
								resr = "XLIX"
							} else if res == 50 {
								resr = "L"
							} else if res == 51 {
								resr = "LI"
							} else if res == 52 {
								resr = "LII"
							} else if res == 53 {
								resr = "LIII"
							} else if res == 54 {
								resr = "LIV"
							} else if res == 55 {
								resr = "LV"
							} else if res == 56 {
								resr = "LVI"
							} else if res == 57 {
								resr = "LVII"
							} else if res == 58 {
								resr = "LVIII"
							} else if res == 59 {
								resr = "LIX"
							} else if res == 60 {
								resr = "LX"
							} else if res == 61 {
								resr = "LXI"
							} else if res == 62 {
								resr = "LXII"
							} else if res == 63 {
								resr = "LXIII"
							} else if res == 64 {
								resr = "LXIV"
							} else if res == 65 {
								resr = "LXV"
							} else if res == 66 {
								resr = "LXVI"
							} else if res == 67 {
								resr = "LXVII"
							} else if res == 68 {
								resr = "LXVIII"
							} else if res == 69 {
								resr = "LXIX"
							} else if res == 70 {
								resr = "LXX"
							} else if res == 71 {
								resr = "LXXI"
							} else if res == 72 {
								resr = "LXXII"
							} else if res == 73 {
								resr = "LXXIII"
							} else if res == 74 {
								resr = "LXXIV"
							} else if res == 75 {
								resr = "LXXV"
							} else if res == 76 {
								resr = "LXXVI"
							} else if res == 77 {
								resr = "LXXVII"
							} else if res == 78 {
								resr = "LXXVIII"
							} else if res == 79 {
								resr = "LXXIX"
							} else if res == 80 {
								resr = "LXXX"
							} else if res == 81 {
								resr = "LXXXI"
							} else if res == 82 {
								resr = "LXXXII"
							} else if res == 83 {
								resr = "LXXXIII"
							} else if res == 84 {
								resr = "LXXXIV"
							} else if res == 85 {
								resr = "LXXXV"
							} else if res == 86 {
								resr = "LXXXVI"
							} else if res == 87 {
								resr = "LXXXVII"
							} else if res == 88 {
								resr = "LXXXVIII"
							} else if res == 89 {
								resr = "LXXXIX"
							} else if res == 90 {
								resr = "XC"
							} else if res == 91 {
								resr = "XCI"
							} else if res == 92 {
								resr = "XCII"
							} else if res == 93 {
								resr = "XCIII"
							} else if res == 94 {
								resr = "XCIV"
							} else if res == 95 {
								resr = "XCV"
							} else if res == 96 {
								resr = "XCVI"
							} else if res == 97 {
								resr = "XCVII"
							} else if res == 98 {
								resr = "XCVIII"
							} else if res == 99 {
								resr = "XCIX"
							} else if res == 100 {
								resr = "C"
							}
							fmt.Println("Вот что получилось: ", a, "=", resr)
						} else if aa[1] == "/" {
							res := x / y
							if res > 0 {
								if res == 1 {
									resr = "I"
								} else if res == 2 {
									resr = "II"
								} else if res == 3 {
									resr = "III"
								} else if res == 4 {
									resr = "IV"
								} else if res == 5 {
									resr = "V"
								} else if res == 6 {
									resr = "VI"
								} else if res == 7 {
									resr = "VII"
								} else if res == 8 {
									resr = "VIII"
								} else if res == 9 {
									resr = "IX"
								} else if res == 10 {
									resr = "X"
								} else if res == 11 {
									resr = "XI"
								} else if res == 12 {
									resr = "XII"
								} else if res == 13 {
									resr = "XIII"
								} else if res == 14 {
									resr = "XIV"
								} else if res == 15 {
									resr = "XV"
								} else if res == 16 {
									resr = "XVI"
								} else if res == 17 {
									resr = "XVII"
								} else if res == 18 {
									resr = "XVIII"
								} else if res == 19 {
									resr = "XIX"
								} else if res == 20 {
									resr = "XX"
								} else if res == 21 {
									resr = "XXI"
								} else if res == 22 {
									resr = "XXII"
								} else if res == 23 {
									resr = "XXIII"
								} else if res == 24 {
									resr = "XXIV"
								} else if res == 25 {
									resr = "XXV"
								} else if res == 26 {
									resr = "XXVI"
								} else if res == 27 {
									resr = "XXVII"
								} else if res == 28 {
									resr = "XXVIII"
								} else if res == 29 {
									resr = "XXIX"
								} else if res == 30 {
									resr = "XXX"
								} else if res == 31 {
									resr = "XXXI"
								} else if res == 32 {
									resr = "XXXII"
								} else if res == 33 {
									resr = "XXXIII"
								} else if res == 34 {
									resr = "XXXIV"
								} else if res == 35 {
									resr = "XXXV"
								} else if res == 36 {
									resr = "XXXVI"
								} else if res == 37 {
									resr = "XXXVII"
								} else if res == 38 {
									resr = "XXXVIII"
								} else if res == 39 {
									resr = "XXXIX"
								} else if res == 40 {
									resr = "XL"
								} else if res == 41 {
									resr = "XLI"
								} else if res == 42 {
									resr = "XLII"
								} else if res == 43 {
									resr = "XLIII"
								} else if res == 44 {
									resr = "XLIV"
								} else if res == 45 {
									resr = "XLV"
								} else if res == 46 {
									resr = "XLVI"
								} else if res == 47 {
									resr = "XLVII"
								} else if res == 48 {
									resr = "XLVIII"
								} else if res == 49 {
									resr = "XLIX"
								} else if res == 50 {
									resr = "L"
								} else if res == 51 {
									resr = "LI"
								} else if res == 52 {
									resr = "LII"
								} else if res == 53 {
									resr = "LIII"
								} else if res == 54 {
									resr = "LIV"
								} else if res == 55 {
									resr = "LV"
								} else if res == 56 {
									resr = "LVI"
								} else if res == 57 {
									resr = "LVII"
								} else if res == 58 {
									resr = "LVIII"
								} else if res == 59 {
									resr = "LIX"
								} else if res == 60 {
									resr = "LX"
								} else if res == 61 {
									resr = "LXI"
								} else if res == 62 {
									resr = "LXII"
								} else if res == 63 {
									resr = "LXIII"
								} else if res == 64 {
									resr = "LXIV"
								} else if res == 65 {
									resr = "LXV"
								} else if res == 66 {
									resr = "LXVI"
								} else if res == 67 {
									resr = "LXVII"
								} else if res == 68 {
									resr = "LXVIII"
								} else if res == 69 {
									resr = "LXIX"
								} else if res == 70 {
									resr = "LXX"
								} else if res == 71 {
									resr = "LXXI"
								} else if res == 72 {
									resr = "LXXII"
								} else if res == 73 {
									resr = "LXXIII"
								} else if res == 74 {
									resr = "LXXIV"
								} else if res == 75 {
									resr = "LXXV"
								} else if res == 76 {
									resr = "LXXVI"
								} else if res == 77 {
									resr = "LXXVII"
								} else if res == 78 {
									resr = "LXXVIII"
								} else if res == 79 {
									resr = "LXXIX"
								} else if res == 80 {
									resr = "LXXX"
								} else if res == 81 {
									resr = "LXXXI"
								} else if res == 82 {
									resr = "LXXXII"
								} else if res == 83 {
									resr = "LXXXIII"
								} else if res == 84 {
									resr = "LXXXIV"
								} else if res == 85 {
									resr = "LXXXV"
								} else if res == 86 {
									resr = "LXXXVI"
								} else if res == 87 {
									resr = "LXXXVII"
								} else if res == 88 {
									resr = "LXXXVIII"
								} else if res == 89 {
									resr = "LXXXIX"
								} else if res == 90 {
									resr = "XC"
								} else if res == 91 {
									resr = "XCI"
								} else if res == 92 {
									resr = "XCII"
								} else if res == 93 {
									resr = "XCIII"
								} else if res == 94 {
									resr = "XCIV"
								} else if res == 95 {
									resr = "XCV"
								} else if res == 96 {
									resr = "XCVI"
								} else if res == 97 {
									resr = "XCVII"
								} else if res == 98 {
									resr = "XCVIII"
								} else if res == 99 {
									resr = "XCIX"
								} else if res == 100 {
									resr = "C"
								}
								fmt.Println("Вот что получилось: ", a, "=", resr)
							} else if res <= 0 {
								fmt.Println("В римских цифрах нет отрицательных чисел и числа 0")
							}
						}
					}
				}
			}
		}
	}
}





