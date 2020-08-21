package main

import (
	"github.com/onflow/cadence"
	"github.com/versus-flow/go-flow-tooling/tooling"
)

const nonFungibleToken = "NonFungibleToken"
const demoToken = "DemoToken"
const art = "Art"
const versus = "Versus"
const auction = "Auction"

const marketplace = "Marketplace"
const artist = "Artist"
const buyer1 = "Buyer1"
const buyer2 = "Buyer2"

func ufix(input string) cadence.UFix64 {
	amount, err := cadence.NewUFix64(input)
	if err != nil {
		panic(err)
	}
	return amount
}

func createArt(flow *tooling.FlowConfig, edition uint64, maxEdition uint64) {
	flow.SendTransactionWithArguments("setup/mint_art", art,
		flow.FindAddress(artist), //artist that owns the art
		//description
		cadence.NewUInt64(edition),    //edition
		cadence.NewUInt64(maxEdition)) //maxEdition
}

func bid(flow *tooling.FlowConfig, account string, auctionID int, amount string) {

	flow.CreateAccount(account)
	flow.SendTransactionWithArguments("setup/actor", account,
		ufix("100.0")) //tokens to mint

	flow.SendTransactionWithArguments("buy/bid", account,
		flow.FindAddress(marketplace),
		cadence.UInt64(1),         //id of drop
		cadence.UInt64(auctionID), //id of auction to bid on
		ufix(amount))              //amount to bid
}

// TODO create a script to check if account exist. or just use go flowConfig for that?
//TODO; Add sleep if started with storyteller mode?
//fmt.Println("Press the Enter Key to continue!")
//fmt.Scanln() // wai
func main() {
	flow := tooling.NewFlowConfigLocalhostWithGas(2000)

	flow.DeployContract(nonFungibleToken)
	flow.DeployContract(demoToken)
	flow.DeployContract(art)
	flow.DeployContract(auction)
	flow.DeployContract(versus)

	//We create the accounts and set up the stakeholdres in our scenario

	//Marketplace will own a marketplace and get a cut for each sale, this account does not own any NFT
	flow.CreateAccount(marketplace)
	flow.SendTransactionWithArguments("setup/actor", marketplace, ufix("0.0"))
	flow.SendTransactionWithArguments("setup/versus", marketplace,
		ufix("0.15"),      //cut percentage,
		cadence.UInt64(5), //drop length
		cadence.UInt64(5)) //minimumBlockRemainingAfterBidOrTie

	//The artist owns NFTs and sells in the marketplace
	flow.CreateAccount(artist)
	flow.SendTransactionWithArguments("setup/actor", artist, ufix("0.0"))
	flow.SendTransactionWithArguments("setup/actor", artist, ufix("0.0"))

	flow.SendTransactionWithArguments("setup/drop", marketplace,
		flow.FindAddress(artist),      //marketplace locaion
		ufix("10.0"),                  //start price
		cadence.NewUInt64(13),         //start block
		cadence.NewString("John Doe"), //artist name
		cadence.NewString("Name"),     //name of art
		cadence.NewString("https://cdn.discordapp.com/attachments/744365120268009472/744964330663051364/image0.png"), //url
		cadence.NewString("This is the description"),
		cadence.NewUInt64(10), //number of editions to use for the editioned auction
		ufix("5.0"))           //minimum bid increment

	bid(flow, buyer1, 1, "10.0")
	flow.SendTransactionWithArguments("tick", marketplace, cadence.NewUInt64(1))
	flow.SendTransactionWithArguments("tick", marketplace, cadence.NewUInt64(1))
	flow.SendTransactionWithArguments("tick", marketplace, cadence.NewUInt64(1))
	flow.SendTransactionWithArguments("tick", marketplace, cadence.NewUInt64(1))
	flow.SendTransactionWithArguments("tick", marketplace, cadence.NewUInt64(1))
	flow.SendTransactionWithArguments("tick", marketplace, cadence.NewUInt64(1))

	//bid(flow, buyer2, 2, "30.0")

	flow.SendTransactionWithArguments("buy/settle", marketplace, cadence.UInt64(1))
	flow.RunScript("check_account", flow.FindAddress(marketplace), cadence.NewString("marketplace"))
	flow.RunScript("check_account", flow.FindAddress(buyer1), cadence.NewString("buyer1"))
	flow.RunScript("check_account", flow.FindAddress(buyer2), cadence.NewString("buyer2"))
	flow.RunScript("check_account", flow.FindAddress(artist), cadence.NewString("artist"))

	/*
		//We try to settle the account but the acution has not ended yet
		flow.SendTransactionWithArguments("buy/settle", marketplace, cadence.UInt64(1))

		//now the auction has ended and we can settle

		//check the status of all the accounts involved in this scenario
		flow.RunScript("check_account", flow.FindAddress(marketplace), cadence.NewString("marketplace"))
		flow.RunScript("check_account", flow.FindAddress(artist), cadence.NewString("artist"))
		flow.RunScript("check_account", flow.FindAddress(buyer1), cadence.NewString("buyer1"))
		flow.RunScript("check_account", flow.FindAddress(buyer2), cadence.NewString("buyer2"))

	*/
}
