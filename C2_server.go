package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"path"
	"strconv"
	"time"
)

//[key is string representation of ip address x.x.x.x]
//value is string representing name/location of device
var serverList = make(map[string]string)

var clientList = make(map[string]string)

//contents is string representation of MAC addresses
//load from database
var macList []string

const port int = 80 //TODO decide on custom port or make part of config
func main() {
	fmt.Printf("hello, world\n")

	//On load:
	//need to query database to update serverList, clientList and macList
	//Send status command to all clients and servers and display them
	//wait for command
}

//What exactly does this need to do
//List of functions
//Status
//add device
//update device
//prepare firmware
//compile firmware
//upload firmware to tftp server and send remote reset to arduino
//manage database connection
//Simulate arduino button presses
//send query string on network
//getters and setters for serverList and clientList from database

// returns a string slice with all status info for servers.
func getNetworkStatus(queryString string) []string {
	statusReports := make([]string, len(serverList)) //TODO determine status query string
	for IP, _ := range serverList {
		statusReports = append(statusReports, networkQuery(queryString, net.ParseIP(IP)))
	}
	return statusReports
}

//This encapsulates s into tcp packets, establishes a tcp session w/ ip,
//transmits the packets and recieves and returns the response as a string
//TODO decide on exact return value
func networkQuery(s string, ip net.IP) string {
	var response string = ""

	return response
}

func receiveFromClient(port int) {
	incoming, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		// handle error
		log.Println(err)
	}
	for {
		conn, err := incoming.Accept()
		if err != nil {
			// handle error
		}
		//go handleConnection(conn)
	}
}

//macType defines the second digit of the mac address as one of the 4
//possible values for internally managed mac addresses, 2, 6, A, E, where the
//value is directly passed in and checked in the function.
func addDevice(macType rune, ip1Digit int, ip2Digit int) {
	macAddr := createMacAddress(macType)

	//generate mac address and IP address
	//check if they are already in database
	//prompt to connect
}

func createMacAddress(macType rune) net.HardwareAddr {
	//testMAC, _ := net.ParseMAC(02:00:00:00:00)
	rand.Seed(time.Now().Unix())
	macDigits := make([]int, 10)
	macChars := make([]string, 10)

	for i := 0; i < 10; i++ { //generate 11 random integer digits
		if i == 1 { //this means that index 1 does not get set
			continue
		}
		//returns a psuedo random integer in the range [0.n) where n = 16 in this case
		macDigits[i] = rand.Intn(16)
	}

	switch macType {

	case '2':
		macDigits[1] = 2
	case '6':
		macDigits[1] = 6
	case 'A':
		macDigits[1] = 10
	case 'E':
		macDigits[1] = 15
	default:
		macDigits[1] = 2
	}

	//fmt.Println(macDigits)

	for i, digit := range macDigits {
		macChars[i] = fmt.Sprintf("%v", strconv.FormatInt(int64(digit), 16))
	}
	//fmt.Println(macChars)
	tempString := macChars[0] + macChars[1] + ":" + macChars[2] + macChars[3] + ":" + macChars[4] + macChars[5] + ":" + macChars[6] + macChars[7] + ":" + macChars[8] + macChars[9]
	//fmt.Println(tempString)
	macAddress, _ := net.ParseMAC(tempString)
	return macAddress

}

func pullFromDatabase() int{
	//implement database functions here
	return 0
}

func createFileHandle(path string) *os.File {
	f, err := os.Open("")
	if err != nil {
		log.Panicf("Failed opening file %s", path)
		log.Panicln(err)
	}
	return f
}

func closeFileHandle(f *os.File) {
	log.Printf("Closing file %s", f.Name())
	f.Close()
}

//have meta list of all query strings
//take config options passed in and generate source file
//return fullPath which is the path to the generated source file
func clientSourceGen(selfIP net.IP, rootPath string, deviceType string, queryStrings []string, IPAddresses []net.IP, C2_IP net.IP, C2_query net.IP, MAC net.HardwareAddr, port int, DNS net.IP, gateway net.IP, subnet net.IP) string {
	fullPath := path.Join(rootPath, serverList[net.IP.String(selfIP)], time.Now().Format(time.RFC822))
	inFile := createFileHandle(fullPath)
	defer closeFileHandle(inFile)
	outFile := createFileHandle(fullPath)
	defer closeFileHandle(outFile)
	switch deviceType {
	case "LED client":
		//needs array of 10 server ip addresses, array of 10 query strings,
		//C2 server ip, C2 server query, MAC, self IP,
		//port, dns, gateway, subnet,

	case "Other client":

		//needs array of n server ip addresses, array of n query strings,
		//C2 server ip, C2 server query, MAC, self IP,
		//port, dns, gateway, subnet,
	}
	return fullPath
}

//serverSourceGen takes config options passed in and generate source file.
//It returns fullPath which is the path to the generated source file
func serverSourceGen(selfIP net.IP, rootPath string, deviceType string, queryStrings []string, C2_IP net.IP, C2_query net.IP, MAC net.HardwareAddr, port int, DNS net.IP, gateway net.IP, subnet net.IP) string {
	fullPath := path.Join(rootPath, serverList[net.IP.String(selfIP)], time.Now().Format(time.RFC822))
	inFile := createFileHandle(fullPath)
	defer closeFileHandle(inFile)
	outFile := createFileHandle(fullPath)
	defer closeFileHandle(outFile)
	switch deviceType {
	//case "Other server":
	//needs array of query strings to respond to,
	//For other server, query string will be of aproximate form "'Port'-{ON:OFF}"
	//C2 server ip, C2 server query, MAC, self IP,
	//port, dns, gateway, subnet,

	case "Relay server":

		//needs array of query strings to respond to,
		//For relay server, query string will be of aproximate form "'Port'-{ON:OFF}"
		//C2 server ip, C2 server query, MAC, self IP,
		//port, dns, gateway, subnet,

	default:

	}
	return fullPath
}
