package main

import (
	"context"
	standardcryto "crypto"
	"crypto/ecdsa"
	"fmt"
	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/crypto"
	"github.com/celo-org/celo-blockchain/ethclient"
	"github.com/slim12kg/piggyfi-vendor/src/go-vendor/contracts/celo"
	"log"
	"math/big"
)

const (
	rpcUrl =  "wss://alfajores-forno.celo-testnet.org/ws"	//"https://alfajores-forno.celo-testnet.org"
	uuid = "83608880-0f79-11ed-861d-0242ac120002"
	globalP2PAddress = "0xbCCa5a93FC1C4DF01f94Ca3ce86EC99f667159B3"
	signerPrivateKey = ""
)

//┌─────────┬───────────────────┬──────────────────────────────────────────────┐
//│ (index) │     contract      │                   address                    │
//├─────────┼───────────────────┼──────────────────────────────────────────────┤
//│    0    │  'WalletLogicV1'  │ '0xFbd81B4d8EECD060A929F9574851fB968B4Ad667' │
//│    1    │ 'GlobalP2P Proxy' │ '0xbCCa5a93FC1C4DF01f94Ca3ce86EC99f667159B3' │
//│    2    │     'MockUSD'     │ '0xd2B98D367b2B49373239352f06059D1D7D91D517' │
//└─────────┴───────────────────┴──────────────────────────────────────────────┘

/**
* @dev Usage
* transactionHash := createAgentWallet(gp2pInstance, client)
* fmt.Println("Transaction hash to track agent wallet creation: ", transactionHash)
* agent := getAgentWalletAddress(gp2pInstance, uuid)
* fmt.Println(agent)
*/
func main() {
	client, err := ethclient.Dial(rpcUrl)
	defer client.Close()
	if err != nil {
		log.Fatalln("Could not connect to client: ", err)
	}

	gp2pInstance, err := globalP2PInstance(client)

	if err != nil {
		log.Fatalln("Error getting instance: ",err)
	}

	// Set to the same address, since we don't have an upgrade
	callOpt := &bind.CallOpts{Context: context.Background()}
	logic, err := gp2pInstance.WalletLogic(callOpt)

	trnx, err := setNewLogicAddress(gp2pInstance, client, logic)

	if err != nil {
		log.Fatalln("Could not update logic: ", logic)
	}

	listenToEvent(client, []string{globalP2PAddress})

	fmt.Printf("%+v",trnx)
}

func listenToEvent(client *ethclient.Client, addresses []string) {
	var contractAddresses []common.Address

	for _,address := range addresses{
		contractAddresses = append(contractAddresses, common.HexToAddress(address))
	}

	query := ethereum.FilterQuery{
		Addresses: contractAddresses,
	}

	logs := make(chan types.Log)

	subscription, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal("Client subscription error: ",err)
	}

	for {
		select {
		case err := <-subscription.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("Contract event log: ", vLog) // pointer to event log
		}
	}
}

func estimateGasUsage(
	client ethclient.Client,
	contractAddress common.Address,
	abi abi.ABI,
	methodName string,
	methodArgs []interface{}) uint64 {
	var interfaceArgs []interface{}

	interfaceArgs = append(interfaceArgs, methodArgs...)

	data, err := abi.Pack(methodName, interfaceArgs...)

	if err != nil {
		log.Fatalln(err)
	}

	_, _, publicKeyECDSA,  _ := getSigner()

	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	callMsg := ethereum.CallMsg{
		From:     from,
		To:       &contractAddress,
		Gas:      0,
		GasPrice: big.NewInt(0),
		Value:    big.NewInt(0),
		Data:     data,
	}

	gasLimit, err := client.EstimateGas(context.Background(), callMsg)

	if err != nil {
		log.Fatalln("Could not estimate gas: ", err)
	}

	return gasLimit
}

func globalP2PInstance(client bind.ContractBackend) (*celo.GlobalP2P, error) {
	address := common.HexToAddress(globalP2PAddress)

	return celo.NewGlobalP2P(address, client)
}

func getWalletLogicAddress(gp2pInstance *celo.GlobalP2P) {
	opts := &bind.CallOpts{Context: context.Background()}

	walletAddress, err := gp2pInstance.WalletLogic(opts)
	if err != nil {
		log.Fatalln("Cannot get wallet logic: ",err)
	}

	fmt.Println("Logic wallet address: ",walletAddress.String())
}

func getCUSDAddress(gp2pInstance *celo.GlobalP2P) {
	opts := &bind.CallOpts{Context: context.Background()}

	walletAddress, err := gp2pInstance.UsdToken(opts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("cUSD(Mock Token) wallet address: ",walletAddress.String())
}

/**
* This is a state changing method, so developer will wait for block to be mined
* and monitor on te frontend using the transaction hash returned
 */
func createAgentWallet(gp2pInstance *celo.GlobalP2P, client *ethclient.Client) (string, error) {
	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}
	params = append(params, uuid)

	cAddress := common.HexToAddress(globalP2PAddress)
	contractABI, err := celo.ParseGlobalP2PABI()
	if err != nil {
		log.Fatalln(err)
	}

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"withdrawBalance", params)

	opt.GasLimit = gasLimit

	transaction , err := gp2pInstance.DeployWallet(opt, uuid)

	if err != nil {
		log.Fatalln("Cannot deploy wallet: ", err)
	}

	fmt.Println(transaction.Hash().Hex())

	return transaction.Hash().Hex(), err
}

func setNewLogicAddress(gp2pInstance *celo.GlobalP2P, client *ethclient.Client, logic common.Address) (*types.Transaction, error) {
	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}
	params = append(params, logic)

	cAddress := common.HexToAddress(globalP2PAddress)
	contractABI, err := celo.ParseGlobalP2PABI()
	if err != nil {
		log.Fatalln("Cannot parse contract ABI: ", err)
	}

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"updateWalletLogic", params)

	opt.GasLimit = gasLimit

	transaction ,err := gp2pInstance.UpdateWalletLogic(opt, logic)

	if err != nil {
		log.Fatalln(err)
	}

	return transaction, err
}

func getAgentWalletAddress(gp2pInstance *celo.GlobalP2P, agentUUID string) common.Address {
	opts := &bind.CallOpts{Context: context.Background()}

	agentAddress, err := gp2pInstance.Wallets(opts, agentUUID)

	if err != nil {
		log.Fatalln("Cannot get wallet: ", err)
	}

	return agentAddress
}

func transactOpt(client ethclient.Client, value *big.Int) *bind.TransactOpts {
	signer, _, publicKeyECDSA,  _ := getSigner()

	from := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal("Cannot get pending nonce: ", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Cannot suggest gas price: ", err)
	}

	networkID, err := client.NetworkID(context.Background())

	if err != nil {
		log.Fatalln("Cannot get network ID: ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(signer, networkID)

	if err != nil {
		log.Fatalln("Error with NewKeyedTransactorWithChainID ",err)
	}

	auth.Nonce = new(big.Int).SetUint64(nonce)
	auth.Value = value     // in wei
	auth.GasLimit = uint64(800000) // set default Gas Limit to be overridden by call to GasEstimate eth_estimate
	auth.GasPrice = gasPrice

	return auth
}

func getSigner() (*ecdsa.PrivateKey, standardcryto.PublicKey, *ecdsa.PublicKey, error) {
	//In a balanced world, this should be kept safely in .env or deeper
	privateKey, err := crypto.HexToECDSA(signerPrivateKey)
	if err != nil {
		log.Fatal("Invalid private key: ", err)
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	return privateKey, publicKey, publicKeyECDSA, err
}