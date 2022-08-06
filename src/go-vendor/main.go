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
	"reflect"
	"regexp"
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

	gp2pInstance, err := globalP2PInstance(client)

	if err != nil {
		log.Fatalln("Error getting instance: ",err)
	}

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

	//decimals, err := cUSDInstance.Decimals(&bind.CallOpts{Context: context.Background()})
	//
	//if err != nil {
	//	log.Fatalln("Cannot get mock usd instance: ", err)
	//}

	agent := getAgentWalletAddress(gp2pInstance, uuid2)

	//toMint := convertUintToBigWei(1000, decimals)
	//contractABI, _ := celo.ParseMockUSDABI()
	//transactions, err := MintMockUSDToken(client, cUSDInstance, contractABI, agent, toMint)
	//if err != nil {
	//	log.Fatalln("An error occurred: ", err)
	//}
	//fmt.Println("Funded Agent Wallet: ",transactions.Hash().Hex())

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

	//Step: Approve an Order to Receiver address & Read Agent Spendable Balance
	agentWallet, err := AgentWalletInstance(client,agent)
	if err != nil {
		log.Fatalln("Error getting instance: ",err)
	}
	//contractABI, err := celo.ParseWalletLogicV1ABI()
	//if err != nil {
	//	log.Fatalln("Error getting logic ABI: ",err)
	//}
	//
	//
	//send := convertUintToBigWei(200, decimals)
	//fee := convertUintToBigWei(2, decimals)
	//to := common.HexToAddress("0xadA32d1905DB6FF74F08801ac4016E56D3dF4375")
	//quoteId := "20387149469590097920"
	//transactionHash, err := approveOrder(agentWallet, contractABI, client, agent.Hex(), send, fee, to, quoteId)
	//if err != nil {
	//	log.Fatalln("Cannot approve transaction: ", err)
	//}
	//
	//fmt.Println(transactionHash)

	//Step: Freeze agent balance
	//transactionHash, err := freezeBalance(agentWallet, contractABI, client, agent.Hex(), convertUintToBigWei(1000, decimals))
	//
	//fmt.Println("Froze balance: ",transactionHash)

	//Step: Withdrawal Available Liquidity
	//@dev Notice total and fee is equal allowable balance
	//total := convertUintToBigWeiFloat(790.02, decimals)
	//fee := convertUintToBigWeiFloat(7.98, decimals)
	//to := common.HexToAddress("0xadA32d1905DB6FF74F08801ac4016E56D3dF4375")
	//transactionHash, err := withdrawLiquidityBalance(agentWallet, contractABI, client, agent.Hex(), total, fee, to)
	//
	//fmt.Println("Withdrawal of liquidity balance: ", transactionHash)

	//Step: UnFreeze agent balance
	//transactionHash, err := unfreezeBalance(agentWallet, contractABI, client, agent.Hex(), convertUintToBigWei(1000, decimals))
	//
	//fmt.Println("Unfroze balance: ",transactionHash)

	spendableBal, err := agentWallet.SpendableBalance(&bind.CallOpts{Context: context.Background()})

	fmt.Println("Spendable Bal: ", spendableBal.String())
}

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// IsZeroAddress validate if it's a 0 address
func IsZeroAddress(iaddress interface{}) bool {
	var address common.Address
	switch v := iaddress.(type) {
	case string:
		address = common.HexToAddress(v)
	case common.Address:
		address = v
	default:
		return false
	}

	zeroAddressBytes := common.FromHex("0x0000000000000000000000000000000000000000")
	addressBytes := address.Bytes()
	return reflect.DeepEqual(addressBytes, zeroAddressBytes)
}

//if IsZeroAddress(agentWalletAddress) || !IsValidAddress(agentWalletInstance) {
//log.Fatalln("Invalid Agent Address")
//}

// @Todo Check agent address passed is Valid
func freezeBalance(
	agentWalletInstance *celo.WalletLogicV1,
	contractABI *abi.ABI,
	client *ethclient.Client,
	agentWalletAddress string,
	amount *big.Int,
	) (string, error) {

	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}
	params = append(params, amount)

	cAddress := common.HexToAddress(agentWalletAddress)

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"freeze", params)

	opt.GasLimit = gasLimit

	transaction, err := agentWalletInstance.Freeze(opt, amount)

	if err != nil {
		log.Fatalln("Error freezing order: ", err)
	}

	return transaction.Hash().Hex(), err
}

func unfreezeBalance(
	agentWalletInstance *celo.WalletLogicV1,
	contractABI *abi.ABI,
	client *ethclient.Client,
	agentWalletAddress string,
	amount *big.Int,
) (string, error) {

	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}
	params = append(params, amount)

	cAddress := common.HexToAddress(agentWalletAddress)

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"unfreeze", params)

	opt.GasLimit = gasLimit

	transaction, err := agentWalletInstance.Unfreeze(opt, amount)

	if err != nil {
		log.Fatalln("Error unfreezing order: ", err)
	}

	return transaction.Hash().Hex(), err
}

func withdrawLiquidityBalance(
	agentWalletInstance *celo.WalletLogicV1,
	contractABI *abi.ABI,
	client *ethclient.Client,
	agentWalletAddress string,
	amount *big.Int,
	fee *big.Int,
	to common.Address,
) (string, error) {

	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}
	params = append(params, amount)
	params = append(params, fee)
	params = append(params, to)

	cAddress := common.HexToAddress(agentWalletAddress)

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"withdrawLiquidity", params)

	opt.GasLimit = gasLimit

	transaction, err := agentWalletInstance.WithdrawLiquidity(opt, amount, fee, to)

	if err != nil {
		log.Fatalln("Error withdrawing liquidity balance: ", err)
	}

	return transaction.Hash().Hex(), err
}

func approveOrder(
	agentWalletInstance *celo.WalletLogicV1,
	contractABI *abi.ABI,
	client *ethclient.Client,
	agentWalletAddress string,
	amount *big.Int,
	fee *big.Int,
	receiver common.Address,
	quoteId	string,
	) (string, error) {

	opt := transactOpt(*client, big.NewInt(0))

	var params []interface{}
	params = append(params, amount)
	params = append(params, fee)
	params = append(params, receiver)
	params = append(params, quoteId)

	cAddress := common.HexToAddress(agentWalletAddress)

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"execTransaction", params)

	opt.GasLimit = gasLimit

	transaction, err := agentWalletInstance.ExecTransaction(opt, amount, fee, receiver, quoteId)

	if err != nil {
		log.Fatalln("Error approving order: ", err)
	}

	return transaction.Hash().Hex(), err
}

func convertUintToBigWei(amount int64, decimals uint8) *big.Int {
	sum := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(decimals)), nil)
	return new(big.Int).Mul(sum, big.NewInt(amount))
}

//@Todo Fix to handle Float properly
func convertUintToBigWeiFloat(amount float64, decimals uint8) *big.Int {
	sum := new(big.Int).Exp(big.NewInt(10), new(big.Int).SetUint64(uint64(decimals)), nil)
	return new(big.Int).Mul(sum, big.NewInt(int64(amount)))
}

func AgentWalletInstance(client bind.ContractBackend, wallet common.Address) (*celo.WalletLogicV1, error) {
	agentWallet , err := celo.NewWalletLogicV1(wallet, client)

	if err != nil {
		log.Fatalln("")
	}

	return agentWallet, err
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
	fmt.Println("Max Transaction Gas: ", gasLimit)
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

	var params []interface{}

	params = append(params,receiver)

	params = append(params, amount)

	cAddress := common.HexToAddress(mockUSDAddress)

	gasLimit := estimateGasUsage(*client, cAddress, *contractABI,"mint", params)

	opt.GasLimit = gasLimit

	transaction, err := cUsd.Mint(opt, receiver, amount)
	if err != nil {
		log.Fatalln("Cannot mint tokens for ", receiver.String()," ", err)
	}

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