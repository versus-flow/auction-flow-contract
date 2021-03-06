package main

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bjartek/go-with-the-flow/gwtf"
	"github.com/onflow/cadence"
)

func fileAsImageData(path string) string {
	f, _ := os.Open("./" + path)

	defer f.Close()

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	contentType := http.DetectContentType(content)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	return "data:" + contentType + ";base64, " + encoded
}

func main() {

	flow := gwtf.NewGoWithTheFlowDevNet()
	//g.TransactionFromFile("buy/settle").SignProposeAndPayAs("versus").Argument(cadence.UInt64(2)).RunPrintEventsFull()

	//value := flow.ScriptFromFile("get_active_auction").AccountArgument("versus").RunReturns()
	//value := flow.ScriptFromFile("check_account").AccountArgument("buyer1").RunReturns()
	///	spew.Dump(value)

	//	flow.TransactionFromFile("buy/settle").SignProposeAndPayAs("versus").Argument(cadence.UInt64(1)).RunPrintEventsFull()

	//	result := flow.ScriptFromFile("get_active_auction").AccountArgument("versus").RunReturns()
	//spew.Dump(result)

	//accountArg := cadence.String("adf")
	//flow.TransactionFromFile("setup/transfer_flow").SignProposeAndPayAsService().UFix64Argument("1000.0").RawAccountArgument("bcfd4d4868a916e5").RunPrintEventsFull()

	//setup buyer1
	//flow.CreateAccountPrintEvents("buyer1")
	//transfer flow to versus account
	//flow.TransactionFromFile("setup/transfer_flow").SignProposeAndPayAsService().UFix64Argument("1000.0").AccountArgument("buyer1").RunPrintEventsFull()

	//run this, copy the address you get out and put it into ~/.flow-dev.json for the versus name
	//flow.CreateAccountPrintEvents("versus")
	//transfer flow to versus account
	//flow.TransactionFromFile("setup/transfer_flow").SignProposeAndPayAsService().UFix64Argument("1000.0").AccountArgument("versus").RunPrintEventsFull()

	//run this, copy the address you get out and put it into ~/.flow-dev.json for the artist name
	//flow.CreateAccountPrintEvents("artist")
	//flow.TransactionFromFile("setup/transfer_flow").SignProposeAndPayAsService().UFix64Argument("1000.0").AccountArgument("artist").RunPrintEventsFull()

	/*
		flow.TransactionFromFile("setup/versus").
			SignProposeAndPayAs("versus").
			UFix64Argument("0.15").    //cut percentage,
			UFix64Argument("86400.0"). //length
			UFix64Argument("300.0").   // bump on late bid
			RunPrintEventsFull()

	*/
	now := time.Now()
	t := now.Unix()
	timeString := strconv.FormatInt(t, 10) + ".0"

	image := fileAsImageData("bull.png")

	flow.TransactionFromFile("setup/drop").
		SignProposeAndPayAs("versus").
		AccountArgument("artist").     //marketplace location
		UFix64Argument("10.01").       //start price
		UFix64Argument(timeString).    //start time
		StringArgument("Kinger9999").  //artist name
		StringArgument("CryptoBull2"). //name of art
		StringArgument(image).         //imaage
		StringArgument("An Angry bull").
		Argument(cadence.NewUInt64(10)). //number of editions to use for the editioned auction
		UFix64Argument("5.0").           //min bid increment
		RunPrintEventsFull()
	/*


		flow.TransactionFromFile("setup/mint_art").
			SignProposeAndPayAs("versus").
			RawAccountArgument("0x42de7e7e48d17e2a").
			StringArgument("Kinger9999").    //artist name
			StringArgument("CryptoBull0").   //name of art
			StringArgument(image).           //imaage
			StringArgument("An angry bull"). //description
			RunPrintEventsFull()

			flow.TransactionFromFile("buy/bid").
				SignProposeAndPayAs("buyer1").
				AccountArgument("versus").
				Argument(cadence.UInt64(1)).  //id of drop
				Argument(cadence.UInt64(21)). //id of unique auction auction to bid on
				UFix64Argument("10.01").      //amount to bid
				RunPrintEventsFull()

	*/

}
