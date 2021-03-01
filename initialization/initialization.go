package initialization

import (
	"MP1/errorchecker"
	"MP1/tcp"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// InitializeNode parses the specified process id from the command line, and returns a corresponding node,
// as well as a list of potential nodes to send messages to.
func InitializeNode() (node tcp.Node, nodes []tcp.Node){
	// Read command line for id of process.
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Incorrect Usage. Do: go run process.go <node_id_number>\n")
		os.Exit(1)
	}

	// Parse process id from commandline.
	processId, err := strconv.Atoi(os.Args[1])
	errorchecker.CheckError(err)

	// Parse information from configuration file about specified process.
	node, nodes = ParseConfiguration(processId)
	return
}

// ParseConfiguration goes line by line through a node configuration file and extracts minimum delay, maximum delay,
// and node information, storing the node with the corresponding ID into a single variable.
func ParseConfiguration(id int) (node tcp.Node, nodes []tcp.Node) {
	// Open configuration file.
	file, err := os.Open("config/nodes.txt")
	if err != nil {
		errorchecker.CheckError(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// Parse minimum and maxmimum delay from configuration file.
	delay := scanner.Text()
	re := regexp.MustCompile(`\((.*?)\)`)
	submatchall := re.FindAllString(delay, -1)
	for i := range submatchall {
		submatchall[i] = strings.Trim(submatchall[i], "(")
		submatchall[i] = strings.Trim(submatchall[i], ")")
	}

	var minDelay, maxDelay int
	minDelay, err = strconv.Atoi(submatchall[0])
	errorchecker.CheckError(err)
	maxDelay, err = strconv.Atoi(submatchall[1])
	errorchecker.CheckError(err)

	// Parse node information and add them to a slice of nodes.
	for scanner.Scan() {
		nodeInfo := strings.Fields(scanner.Text())
		id, err := strconv.Atoi(nodeInfo[0])
		errorchecker.CheckError(err)
		nodes = append(nodes, tcp.Node{Id: id, Ip: nodeInfo[1], Port: nodeInfo[2], MinDelay: minDelay, MaxDelay: maxDelay})
	}

	node = nodes[id]

	return
}
