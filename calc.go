package main

//используемые библиотеки

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	глобальная переменная в которой находится разделитель вводимой строки,

по совместительству наш будущий оператор операции
*/
var operator rune

/*
функция сплит проверяет, является-ли символ (r-rune) разделителем из списка. Если да - то переменной "operator" присваивается значение разделителя а функция сплит возвращает значение true
*/
func Split(r rune) bool {
	if r == '+' || r == '-' || r == '*' || r == '/' {
		operator = r
		return true
	}
	return false
}

// начало основной функции
func main() {
	//мапа для перевода римских чисел в арабские
	rome := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	// переменная out для записи целочисленного результата будущей операции
	var out int
	/* здесь программа принимает на вход строку пока не встретит символ новой строки, и разбивает её
	   на подстроки с помощью нашей функции split */
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	result := strings.FieldsFunc(input, Split)

	//Пеобразование римских чисел в арабские, если таковые имеются
	var rim bool
	var rim1 bool
	for key, value := range rome {
		if result[0] == key {
			rim = true
			result[0] = strconv.Itoa(value)
			result[1] = strings.TrimSuffix(result[1], "\n")
			for key, value := range rome {
				if result[1] == key {
					rim1 = true
					result[1] = strconv.Itoa(value)
				} else {
				}
			}
		}
	}

	//ограничение на две переменные в соответствии с ТЗ
	if len(result) != 2 {
		panic("нужно два числа (U.U) ")
	}

	/*дальше полученные подстроки преобразуются в целые числа без пробелов. Эти числа (num1 и num2) - наши переменные.
	Если числа не получаются, выдается ошибка */

	num1, err := strconv.Atoi(strings.TrimSpace(result[0]))
	if err != nil {
		panic("Какие-то неправильные данные :с ")
	}
	num2, err := strconv.Atoi(strings.TrimSpace(result[1]))
	if err != nil {
		panic("Какие-то неправильные данные :с ")
	}

	//ограничение на велечину переменных в соответствии с ТЗ
	if num1 > 10 || num2 > 10 {
		panic("слишком большие значения >_< ")
	}

	//применение операции к полученным переменным в зависимости от "оператора"
	switch operator {
	case '+':
		out = num1 + num2
	case '-':
		out = num1 - num2
	case '*':
		out = num1 * num2
	case '/':
		if num2 == 0 {
			fmt.Println("Давай давай подели на ноль ^-^ ")
			return
		}
		out = num1 / num2
	}

	if !rim {
		// вывода результата!
		fmt.Println(out)

	} else {
		if !rim1 {
			panic("оба числа должны быть римские")
		}
		//преобразование ответа из арабских в римские цифры
		if out <= 0 {
			panic("ответ меньше или равен нулю")
		}
		// здесь реализована структура содержащая в себе пары чисел
		romanNumerals := []struct {
			value   int
			numeral string
		}{
			{100, "C"},
			{90, "XC"},
			{50, "L"},
			{40, "XL"},
			{10, "X"},
			{9, "IX"},
			{5, "V"},
			{4, "IV"},
			{1, "I"},
		}
		/* дальше реализован цикл который проходит по парам чисел и подставляет римские числа в пустую строку, а арабские в переменную пока она не будет больше либо равна ответу операции*/
		resultat := ""
		for _, rn := range romanNumerals {
			for out >= rn.value {
				resultat += rn.numeral
				out -= rn.value
			}
		}
		// и выводит получившуюся строку
		fmt.Println(resultat)
	}
}

/*мой калькулятор не умеет работать с отрицательными числами в операторах, хотя я пытался
найти элегантное решение этой проблемы, в ТЗ об этом ничего не сказано, так что я решил
сильно с этим не заморачиваться*/
