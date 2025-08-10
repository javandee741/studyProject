package main

import (
	"fmt"
	"strconv"
)

// Функция для перевода DEC в BIN (дополнительный код) с поддержкой N-бит
func decToBinTwosComplement(n int, bits int) (string, error) {
	if bits <= 0 {
		return "", fmt.Errorf("количество бит должно быть > 0")
	}

	max := 1<<(bits-1) - 1
	min := -1 << (bits - 1)

	if n > max || n < min {
		return "", fmt.Errorf("число %d выходит за диапазон %d-битного числа [%d, %d]", n, bits, min, max)
	}

	// Маска для ограничения бит
	mask := (1 << bits) - 1
	unsigned := uint64(n) & uint64(mask)
	format := fmt.Sprintf("%%0%db", bits) // Форматирование с ведущими нулями

	return fmt.Sprintf(format, unsigned), nil
}

// Функция для перевода BIN (доп. код) в DEC с поддержкой N-бит
func binToDecTwosComplement(binStr string) (int, error) {
	bits := len(binStr)
	if bits == 0 {
		return 0, fmt.Errorf("пустая строка")
	}

	// Проверка на валидность
	for _, ch := range binStr {
		if ch != '0' && ch != '1' {
			return 0, fmt.Errorf("неверный двоичный формат")
		}
	}

	val, err := strconv.ParseUint(binStr, 2, 64)
	if err != nil {
		return 0, err
	}

	// Если старший бит = 1 (отрицательное)
	if val&(1<<(bits-1)) != 0 {
		val = val - (1 << bits) // Эквивалентно: инвертируем и добавляем 1
	}

	return int(val), nil
}

func main() {
	// Примеры для разных битностей
	testCases := []struct {
		dec  int
		bits int
	}{
		{42, 8}, // 8 бит
		{-42, 8},
		{127, 8},
		{-128, 8},
		{1023, 10}, // 10 бит
		{-512, 10},
		{32767, 16}, // 16 бит
		{-32768, 16},
		{514, 11},
		{120436541, 32},
	}

	fmt.Println("DEC → BIN (доп. код):")
	for _, tc := range testCases {
		bin, err := decToBinTwosComplement(tc.dec, tc.bits)
		if err != nil {
			fmt.Printf("Ошибка для %d (%d бит): %v\n", tc.dec, tc.bits, err)
		} else {
			fmt.Printf("DEC %6d (%2d бит) → BIN %s\n", tc.dec, tc.bits, bin)
		}
	}

	// Примеры BIN → DEC
	binaries := []string{
		"00101010",   // 42 (8 бит)
		"11010110",   // -42 (8 бит)
		"1111111111", // -1 (10 бит)
		"0111111111", // 511 (10 бит)
		"1000000000", // -512 (10 бит)
		"1000110011",
	}

	fmt.Println("\nBIN → DEC:")
	for _, bin := range binaries {
		dec, err := binToDecTwosComplement(bin)
		if err != nil {
			fmt.Printf("Ошибка для %s: %v\n", bin, err)
		} else {
			fmt.Printf("BIN %s → DEC %6d (%d бит)\n", bin, dec, len(bin))
		}
	}
}
