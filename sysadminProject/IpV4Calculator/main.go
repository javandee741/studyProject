package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Калькулятор подсетей IPv4")
	fmt.Println("Введите IP-адрес и маску (форматы: 192.168.1.1/24 или 192.168.1.1 255.255.255.0)")

	for {
		fmt.Print("\nВведите IP и маску: ")
		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) == "exit" {
			break
		}

		// Обработка ввода
		ip, ipNet, err := parseInput(input)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}

		// Вычисление информации о подсети
		network, broadcast, firstIP, lastIP, totalHosts := calculateSubnetInfo(ipNet)

		// Вывод результатов
		fmt.Println("\nРезультат:")
		fmt.Printf("IP-адрес:        %s\n", ip)
		fmt.Printf("Маска подсети:   %s\n", net.IP(ipNet.Mask))
		fmt.Printf("Префикс:        /%d\n", prefixLength(ipNet.Mask))
		fmt.Printf("Адрес сети:     %s\n", network)
		fmt.Printf("Широковещательный адрес: %s\n", broadcast)
		fmt.Printf("Доступные хосты: %s - %s\n", firstIP, lastIP)
		fmt.Printf("Всего хостов:    %d\n", totalHosts)
		fmt.Printf("Класс сети:      %s\n", getNetworkClass(ip))
		fmt.Printf("Тип адреса:      %s\n", getAddressType(ip))
	}
}

// Парсинг ввода (поддержка форматов: 192.168.1.1/24 и 192.168.1.1 255.255.255.0)
func parseInput(input string) (net.IP, *net.IPNet, error) {
	parts := strings.Fields(input)

	// Если ввод через пробел (IP и маска отдельно)
	if len(parts) == 2 {
		ipStr := parts[0]
		maskStr := parts[1]

		ip := net.ParseIP(ipStr)
		if ip == nil {
			return nil, nil, fmt.Errorf("неверный IP-адрес")
		}

		mask := net.ParseIP(maskStr)
		if mask == nil {
			return nil, nil, fmt.Errorf("неверная маска подсети")
		}

		// Преобразование маски в net.IPMask
		maskv4 := mask.To4()
		if maskv4 == nil {
			return nil, nil, fmt.Errorf("только IPv4 поддерживается")
		}

		maskLen := net.IPv4Mask(maskv4[0], maskv4[1], maskv4[2], maskv4[3])
		prefixSize, _ := maskLen.Size()

		// Создаем IPNet
		_, ipNet, err := net.ParseCIDR(fmt.Sprintf("%s/%d", ipStr, prefixSize))
		return ip, ipNet, err
	}

	// Если ввод в формате CIDR (192.168.1.1/24)
	ip, ipNet, err := net.ParseCIDR(input)
	if err != nil {
		return nil, nil, fmt.Errorf("неверный формат ввода. Используйте: 192.168.1.1/24 или 192.168.1.1 255.255.255.0")
	}

	return ip, ipNet, nil
}

// Вычисление информации о подсети
func calculateSubnetInfo(ipNet *net.IPNet) (network, broadcast, firstIP, lastIP net.IP, totalHosts uint32) {
	// Адрес сети
	network = ipNet.IP

	// Широковещательный адрес
	mask := ipNet.Mask
	broadcast = make(net.IP, len(network))
	copy(broadcast, network)
	for i := 0; i < len(mask); i++ {
		broadcast[i] |= ^mask[i]
	}

	// Первый и последний доступный IP
	firstIP = make(net.IP, len(network))
	copy(firstIP, network)
	firstIP[len(firstIP)-1]++

	lastIP = make(net.IP, len(broadcast))
	copy(lastIP, broadcast)
	lastIP[len(lastIP)-1]--

	// Количество хостов
	ones, bits := mask.Size()
	totalHosts = (1 << (bits - ones)) - 2

	return network, broadcast, firstIP, lastIP, totalHosts
}

// Определение длины префикса
func prefixLength(mask net.IPMask) int {
	ones, _ := mask.Size()
	return ones
}

// Определение класса сети
func getNetworkClass(ip net.IP) string {
	ip = ip.To4()
	if ip == nil {
		return "N/A"
	}

	firstOctet := ip[0]
	switch {
	case firstOctet < 128:
		return "A"
	case firstOctet < 192:
		return "B"
	case firstOctet < 224:
		return "C"
	case firstOctet < 240:
		return "D (multicast)"
	default:
		return "E (reserved)"
	}
}

// Определение типа адреса
func getAddressType(ip net.IP) string {
	ip = ip.To4()
	if ip == nil {
		return "N/A"
	}

	switch {
	case ip.IsLoopback():
		return "Loopback"
	case ip.IsPrivate():
		return "Private"
	case ip.IsMulticast():
		return "Multicast"
	case ip.IsUnspecified():
		return "Unspecified"
	case ip.Equal(net.IPv4bcast):
		return "Broadcast"
	default:
		return "Public"
	}
}
