package globals

import "net"

var SocketMap = make(map[string] net.Conn)


func SetSocketMap(device_id string, conn net.Conn)  {

	SocketMap[device_id] = conn

}

func GetSocketMap(device_id string) net.Conn {

	return SocketMap[device_id]

}

func GetSocketMapAll() map[string] net.Conn {

	return SocketMap

}