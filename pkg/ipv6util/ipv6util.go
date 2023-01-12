package ipv6util

import "net"

func GenerateEUI64Address(prefix net.IP, mac net.HardwareAddr) net.IP {
	ip := make([]byte, 16)
	copy(ip[0:8], prefix[0:8])

	copy(ip[8:11], mac[0:3])
	ip[8] ^= 0x02
	ip[11] = 0xff
	ip[12] = 0xfe
	copy(ip[13:16], mac[3:6])

	return ip
}
