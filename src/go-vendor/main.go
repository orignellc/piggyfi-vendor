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
	uuid2 = "83608ac4-0f79-11ed-861d-0242ac120002"
	globalP2PAddress = "0xDAae917F4B3990bf132Da4AefDAbC880f192d19D"
	mockUSDAddress = "0x0d0702A42F7d62F4dBd43FfeCc409A273Ea01844"
	signerPrivateKey = "" //Use env to load this
)

// ┌─────────┬───────────────────┬──────────────────────────────────────────────┐
// │ (index) │     contract      │                   address                    │
// ├─────────┼───────────────────┼──────────────────────────────────────────────┤
// │    0    │  'WalletLogicV1'  │ '0xc07be69C1741532dA4012f884eCC7819Ba740901' │
// │    1    │ 'GlobalP2P Proxy' │ '0xDAae917F4B3990bf132Da4AefDAbC880f192d19D' │
// │    2    │      'Admin'      │ '0x56921Eba54DB2d70A32B5186153680d01bb6f0a2' │
// │    3    │ 'Implementation'  │ '0xaC85dC97b271Da20d9f703666DC879e0fA87263E' │
// │    4    │     'MockUSD'     │ '0x0d0702A42F7d62F4dBd43FfeCc409A273Ea01844' │
// └─────────┴───────────────────┴──────────────────────────────────────────────┘

//0x8456cb59

// @Todo Anonymous mappings
var EventsIndexers = make(map[string]struct{})

/**
* @dev Usage
* contractABI, _ := celo.ParseGlobalP2PABI()
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

	//gp2pInstance, err := globalP2PInstance(client)
	//
	//if err != nil {
	//	log.Fatalln("Error getting instance: ",err)
	//}

	//Step 1: Deploy Wallet for Agent
	//contractABI, _ := celo.ParseGlobalP2PABI()
	//transactionHash, err := createAgentWallet(gp2pInstance, contractABI, client, uuid2)
	//
	//fmt.Println("Transaction hash to track agent wallet creation: ", transactionHash)

	//Step 2: Mint 1000 MockUSD token for Agent address
	//cUSDInstance, err := MockUSDInstance(client)
	//if err != nil {
	//	log.Fatalln("Error getting instance: ",err)
	//}
	//
	//decimals, err := cUSDInstance.Decimals(&bind.CallOpts{Context: context.Background()})
	//
	//if err != nil {
	//	log.Fatalln("Cannot get mock usd instance: ", err)
	//}
	//
	//agent := getAgentWalletAddress(gp2pInstance, uuid)
	//
	//toMint := convertUintToBigWei(1000, decimals)
	//contractABI, _ := celo.ParseMockUSDABI()
	//transactions, err := MintMockUSDToken(client, cUSDInstance, contractABI, agent, toMint)
	//if err != nil {
	//	log.Fatalln("An error occurred: ", err)
	//}
	//fmt.Println(transactions.Hash().Hex())

	//listenToEvent(client, []string{mockUSDAddress}, cUSDInstance.TryParseLog)


	//Step : Set new Logic Address
	//First Pause transaction, only when you pause can we update logic
	//contractABI, _ := celo.ParseGlobalP2PABI()
	//paused, err := gp2pInstance.Paused(&bind.CallOpts{Context: context.Background()})
	//if err != nil {
	//	log.Fatalln("Could not check paused: ", err)
	//}
	//
	//if !paused {
	//	transactionHash, err := PauseGP2P(gp2pInstance, contractABI, client)
	//	if err != nil {
	//		log.Fatalln("Could not pause gp2p, reason: ", err)
	//	}
	//	fmt.Println("Pause transaction hash: ", transactionHash)
	//}

	//Step : Set to the same address, since we don't have an upgrade
	//callOpt := &bind.CallOpts{Context: context.Background()}
	//logic, err := gp2pInstance.WalletLogic(callOpt)
	//
	//transactionHash, err := setNewLogicAddress(gp2pInstance, client, logic)
	//
	//if err != nil {
	//	log.Fatalln("Could not update logic: ", logic)
	//}
	//
	//fmt.Println(transactionHash)

	//Step : Unpause to perform other transactions
	//contractABI, _ := celo.ParseGlobalP2PABI()
	//
	//transactionHash, err := UnpauseGP2P(gp2pInstance, contractABI, client)
	//if err != nil {
	//	log.Fatalln("Could not unpause gp2p, reason: ", err)
	//}
	//fmt.Println("Pause transaction hash: ", transactionHash)

	//Final Step: Keeps the program running and listen for contract events
	//listenToEvent(client, []string{globalP2PAddress})
}

func convertUintToBigWei(amount int64, decimals uint8) *big.Int {
	sum := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(decimals)), nil)
	return new(big.Int).Mul(sum, big.NewInt(amount))
}


func listenToEvent(
	client *ethclient.Client,
	addresses []string,
	logParser func(log types.Log) (string, interface{}, bool, error)) {
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
			eventName, event, ok,  err := logParser(vLog)
			if err != nil {
				log.Fatalln("Error parsing event log: ", err)
			}

			fmt.Println("Event okay?: ", ok) // pointer to event log
			fmt.Println("Event name: ", eventName) // pointer to event log
			fmt.Println("Event Body: ", event) // pointer to event log

			return
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

	var (
		data []byte
		err error
	)

	if len(interfaceArgs) > 0 {
		data, err = abi.Pack(methodName, interfaceArgs...)
	}else{
		data, err = abi.Pack(methodName)
	}

	if err != nil {
		log.Fatalln("Cannot run abi pack: ",err)
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
	fmt.Println("Gas Limit: ", gasLimit)
	if err != nil {
		log.Fatalln("Could not estimate gas: ", err)
	}

	return gasLimit
}

func globalP2PInstance(client bind.ContractBackend) (*celo.GlobalP2P, error) {
	address := common.HexToAddress(globalP2PAddress)

	return celo.NewGlobalP2P(address, client)
}

func PauseGP2P(gp2pInstance *celo.GlobalP2P, contractABI *abi.ABI, client *ethclient.Client) (string, error) {
	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}

	cAddress := common.HexToAddress(globalP2PAddress)

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"pause", params)

	opt.GasLimit = gasLimit

	transaction , err := gp2pInstance.Pause(opt)

	if err != nil {
		log.Fatalln("Cannot pause GP2P: ", err)
	}

	fmt.Printf("%+v",transaction)
	fmt.Println()
	return transaction.Hash().Hex(), err
}

func UnpauseGP2P(gp2pInstance *celo.GlobalP2P, contractABI *abi.ABI, client *ethclient.Client) (string, error) {
	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}

	cAddress := common.HexToAddress(globalP2PAddress)


	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"unpause", params)

	opt.GasLimit = gasLimit

	transaction , err := gp2pInstance.Unpause(opt)

	if err != nil {
		log.Fatalln("Cannot unpause GP2P: ", err)
	}

	return transaction.Hash().Hex(), err
}

func MintMockUSDToken(client *ethclient.Client, cUsd *celo.MockUSD, contractABI *abi.ABI, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	opt := transactOpt(*client, big.NewInt(0))

	//var params []interface{}

	//paddedAddress := common.LeftPadBytes(receiver.Bytes(), 32)
	//
	//params = append(params, paddedAddress...)
	//
	//paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	//
	//params = append(params, paddedAmount...)
	//
	//cAddress := common.HexToAddress(globalP2PAddress)
	//
	//
	//gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"mint", params)
	//
	//opt.GasLimit = gasLimit

	transaction, err := cUsd.Mint(opt, receiver, amount)
	if err != nil {
		log.Fatalln("Cannot mint tokens for ", receiver.String()," ", err)
	}

	fmt.Printf("%+v",transaction)

	return transaction, err
}

func MockUSDInstance(client bind.ContractBackend) (*celo.MockUSD, error) {
	address := common.HexToAddress(mockUSDAddress)

	return celo.NewMockUSD(address, client)
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
func createAgentWallet(gp2pInstance *celo.GlobalP2P, contractABI *abi.ABI, client *ethclient.Client, userID string) (string, error) {
	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}
	params = append(params, userID)

	cAddress := common.HexToAddress(globalP2PAddress)

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"deployWallet", params)

	opt.GasLimit = gasLimit

	transaction , err := gp2pInstance.DeployWallet(opt, userID)

	if err != nil {
		log.Fatalln("Cannot deploy wallet: ", err)
	}

	fmt.Println(transaction.Hash())

	return transaction.Hash().Hex(), err
}

//@Todo check if wallet address given is null
func setNewLogicAddress(gp2pInstance *celo.GlobalP2P, client *ethclient.Client, logic common.Address) (string, error) {
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

	return transaction.Hash().Hex(), err
}

//@Todo check if wallet address given is null
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