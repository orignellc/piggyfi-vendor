// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package celo

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GlobalP2PMetaData contains all meta data concerning the GlobalP2P contract.
var GlobalP2PMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"}],\"name\":\"newWallet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"deployWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletLogic\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usd\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletLogic\",\"type\":\"address\"}],\"name\":\"updateWalletLogic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletLogic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"wallets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611b11806100206000396000f3fe608060405260043610620000bb5760003560e01c806382787c7f116200006d57806382787c7f146200020c5780638da5cb5b146200023a578063beacf2cb146200026a578063da76d5cd146200029a578063f2fde38b14620002c8578063f897a22b14620002f657620000c3565b8063038defd714620000c857806337ba532c146200010c578063485cc95514620001505780635c975abb146200017e578063715018a614620001ae57806372f179f814620001c857620000c3565b36620000c357005b600080fd5b348015620000d557600080fd5b50620000f46004803603810190620000ee919062000d9b565b62000326565b60405162000103919062000dea565b60405180910390f35b3480156200011957600080fd5b5062000138600480360381019062000132919062000f69565b62000346565b60405162000147919062000fcb565b60405180910390f35b3480156200015d57600080fd5b506200017c600480360381019062000176919062000fe8565b62000520565b005b3480156200018b57600080fd5b5062000196620006fa565b604051620001a5919062000dea565b60405180910390f35b348015620001bb57600080fd5b50620001c662000711565b005b348015620001d557600080fd5b50620001f46004803603810190620001ee919062000f69565b62000729565b60405162000203919062000fcb565b60405180910390f35b3480156200021957600080fd5b5062000238600480360381019062000232919062000d9b565b62000772565b005b3480156200024757600080fd5b5062000252620007ca565b60405162000261919062000fcb565b60405180910390f35b3480156200027757600080fd5b5062000282620007f4565b60405162000291919062000fcb565b60405180910390f35b348015620002a757600080fd5b50620002c66004803603810190620002c091906200106a565b6200081a565b005b348015620002d557600080fd5b50620002f46004803603810190620002ee919062000d9b565b620008be565b005b3480156200030357600080fd5b506200030e62000948565b6040516200031d919062000fcb565b60405180910390f35b609a6020528060005260406000206000915054906101000a900460ff1681565b6000620003526200096e565b6200035c620009f3565b602082511015620003a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200039b90620010fd565b60405180910390fd5b600030609860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620003d89062000d14565b620003e59291906200111f565b604051809103906000f08015801562000402573d6000803e3d6000fd5b50905080609984604051620004189190620011cf565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001609a60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff1683604051620004e39190620011cf565b60405180910390207f1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b60405160405180910390a380915050919050565b60008060019054906101000a900460ff16159050808015620005525750600160008054906101000a900460ff1660ff16105b80620005835750620005643062000a42565b158015620005825750600160008054906101000a900460ff1660ff16145b5b620005c5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620005bc906200125e565b60405180910390fd5b60016000806101000a81548160ff021916908360ff160217905550801562000603576001600060016101000a81548160ff0219169083151502179055505b6200060d62000a65565b6200061762000ac3565b82609760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081609860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508015620006f55760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051620006ec9190620012da565b60405180910390a15b505050565b6000606560009054906101000a900460ff16905090565b6200071b6200096e565b62000727600062000b21565b565b6099818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6200077c6200096e565b6200078662000be7565b80609760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b609760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b620008246200096e565b478111156200086a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620008619062001347565b60405180910390fd5b62000874620007ca565b73ffffffffffffffffffffffffffffffffffffffff166108fc479081150290604051600060405180830381858888f19350505050158015620008ba573d6000803e3d6000fd5b5050565b620008c86200096e565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036200093a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200093190620013df565b60405180910390fd5b620009458162000b21565b50565b609860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6200097862000c35565b73ffffffffffffffffffffffffffffffffffffffff1662000998620007ca565b73ffffffffffffffffffffffffffffffffffffffff1614620009f1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620009e89062001451565b60405180910390fd5b565b620009fd620006fa565b1562000a40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a3790620014c3565b60405180910390fd5b565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1662000ab7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000aae906200155b565b60405180910390fd5b62000ac162000c3d565b565b600060019054906101000a900460ff1662000b15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000b0c906200155b565b60405180910390fd5b62000b1f62000ca5565b565b6000603360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b62000bf1620006fa565b62000c33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000c2a90620015cd565b60405180910390fd5b565b600033905090565b600060019054906101000a900460ff1662000c8f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000c86906200155b565b60405180910390fd5b62000ca362000c9d62000c35565b62000b21565b565b600060019054906101000a900460ff1662000cf7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000cee906200155b565b60405180910390fd5b6000606560006101000a81548160ff021916908315150217905550565b6104ec80620015f083390190565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000d638262000d36565b9050919050565b62000d758162000d56565b811462000d8157600080fd5b50565b60008135905062000d958162000d6a565b92915050565b60006020828403121562000db45762000db362000d2c565b5b600062000dc48482850162000d84565b91505092915050565b60008115159050919050565b62000de48162000dcd565b82525050565b600060208201905062000e01600083018462000dd9565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000e5c8262000e11565b810181811067ffffffffffffffff8211171562000e7e5762000e7d62000e22565b5b80604052505050565b600062000e9362000d22565b905062000ea1828262000e51565b919050565b600067ffffffffffffffff82111562000ec45762000ec362000e22565b5b62000ecf8262000e11565b9050602081019050919050565b82818337600083830152505050565b600062000f0262000efc8462000ea6565b62000e87565b90508281526020810184848401111562000f215762000f2062000e0c565b5b62000f2e84828562000edc565b509392505050565b600082601f83011262000f4e5762000f4d62000e07565b5b813562000f6084826020860162000eeb565b91505092915050565b60006020828403121562000f825762000f8162000d2c565b5b600082013567ffffffffffffffff81111562000fa35762000fa262000d31565b5b62000fb18482850162000f36565b91505092915050565b62000fc58162000d56565b82525050565b600060208201905062000fe2600083018462000fba565b92915050565b6000806040838503121562001002576200100162000d2c565b5b6000620010128582860162000d84565b9250506020620010258582860162000d84565b9150509250929050565b6000819050919050565b62001044816200102f565b81146200105057600080fd5b50565b600081359050620010648162001039565b92915050565b60006020828403121562001083576200108262000d2c565b5b6000620010938482850162001053565b91505092915050565b600082825260208201905092915050565b7f47503a496e76616c696420555549440000000000000000000000000000000000600082015250565b6000620010e5600f836200109c565b9150620010f282620010ad565b602082019050919050565b600060208201905081810360008301526200111881620010d6565b9050919050565b600060408201905062001136600083018562000fba565b62001145602083018462000fba565b9392505050565b600081519050919050565b600081905092915050565b60005b838110156200118257808201518184015260208101905062001165565b8381111562001192576000848401525b50505050565b6000620011a5826200114c565b620011b1818562001157565b9350620011c381856020860162001162565b80840191505092915050565b6000620011dd828462001198565b915081905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b600062001246602e836200109c565b91506200125382620011e8565b604082019050919050565b60006020820190508181036000830152620012798162001237565b9050919050565b6000819050919050565b600060ff82169050919050565b6000819050919050565b6000620012c2620012bc620012b68462001280565b62001297565b6200128a565b9050919050565b620012d481620012a1565b82525050565b6000602082019050620012f16000830184620012c9565b92915050565b7f47503a20496e73756666696369656e742062616c616e63650000000000000000600082015250565b60006200132f6018836200109c565b91506200133c82620012f7565b602082019050919050565b60006020820190508181036000830152620013628162001320565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000620013c76026836200109c565b9150620013d48262001369565b604082019050919050565b60006020820190508181036000830152620013fa81620013b8565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000620014396020836200109c565b9150620014468262001401565b602082019050919050565b600060208201905081810360008301526200146c816200142a565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b6000620014ab6010836200109c565b9150620014b88262001473565b602082019050919050565b60006020820190508181036000830152620014de816200149c565b9050919050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b600062001543602b836200109c565b91506200155082620014e5565b604082019050919050565b60006020820190508181036000830152620015768162001534565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b6000620015b56014836200109c565b9150620015c2826200157d565b602082019050919050565b60006020820190508181036000830152620015e881620015a6565b905091905056fe608060405234801561001057600080fd5b506040516104ec3803806104ec8339818101604052810190610032919061011d565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061015d565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100ea826100bf565b9050919050565b6100fa816100df565b811461010557600080fd5b50565b600081519050610117816100f1565b92915050565b60008060408385031215610134576101336100ba565b5b600061014285828601610108565b925050602061015385828601610108565b9150509250929050565b6103808061016c6000396000f3fe60806040526004361061002d5760003560e01c8063a890bd15146100a3578063c55dae63146100ce57610099565b36610099573073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f476c794848a86924c99aa79862df9075793bc6e48fa9e77f85f196f12731b0bd3460405161008f9190610239565b60405180910390a3005b6100a16100f9565b005b3480156100af57600080fd5b506100b8610113565b6040516100c59190610295565b60405180910390f35b3480156100da57600080fd5b506100e3610137565b6040516100f091906102d1565b60405180910390f35b61010161015d565b61011161010c61015f565b6101fa565b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508073ffffffffffffffffffffffffffffffffffffffff1663beacf2cb6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101f4919061031d565b91505090565b3660008037600080366000845af43d6000803e806000811461021b573d6000f35b3d6000fd5b6000819050919050565b61023381610220565b82525050565b600060208201905061024e600083018461022a565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061027f82610254565b9050919050565b61028f81610274565b82525050565b60006020820190506102aa6000830184610286565b92915050565b60006102bb82610254565b9050919050565b6102cb816102b0565b82525050565b60006020820190506102e660008301846102c2565b92915050565b600080fd5b6102fa816102b0565b811461030557600080fd5b50565b600081519050610317816102f1565b92915050565b600060208284031215610333576103326102ec565b5b600061034184828501610308565b9150509291505056fea264697066735822122002e35faf8fd6fc7b0a7bbc9c7220a2e5dacc2ea3004d7350a50ba35f06b9894264736f6c634300080f0033a26469706673582212202d3acfe1b7fc5e0d6ad43c5814a413734bbee0e064fb79785bc976913f27d4a764736f6c634300080f0033",
}

// GlobalP2PABI is the input ABI used to generate the binding from.
// Deprecated: Use GlobalP2PMetaData.ABI instead.
var GlobalP2PABI = GlobalP2PMetaData.ABI

// GlobalP2PBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GlobalP2PMetaData.Bin instead.
var GlobalP2PBin = GlobalP2PMetaData.Bin

// DeployGlobalP2P deploys a new Ethereum contract, binding an instance of GlobalP2P to it.
func DeployGlobalP2P(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *GlobalP2P, error) {
	parsed, err := GlobalP2PMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GlobalP2PBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlobalP2P{GlobalP2PCaller: GlobalP2PCaller{contract: contract}, GlobalP2PTransactor: GlobalP2PTransactor{contract: contract}, GlobalP2PFilterer: GlobalP2PFilterer{contract: contract}}, nil
}

// GlobalP2P is an auto generated Go binding around an Ethereum contract.
type GlobalP2P struct {
	GlobalP2PCaller     // Read-only binding to the contract
	GlobalP2PTransactor // Write-only binding to the contract
	GlobalP2PFilterer   // Log filterer for contract events
}

// GlobalP2PCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlobalP2PCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalP2PTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlobalP2PTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalP2PFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlobalP2PFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlobalP2PSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlobalP2PSession struct {
	Contract     *GlobalP2P        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlobalP2PCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlobalP2PCallerSession struct {
	Contract *GlobalP2PCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// GlobalP2PTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlobalP2PTransactorSession struct {
	Contract     *GlobalP2PTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// GlobalP2PRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlobalP2PRaw struct {
	Contract *GlobalP2P // Generic contract binding to access the raw methods on
}

// GlobalP2PCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlobalP2PCallerRaw struct {
	Contract *GlobalP2PCaller // Generic read-only contract binding to access the raw methods on
}

// GlobalP2PTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlobalP2PTransactorRaw struct {
	Contract *GlobalP2PTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlobalP2P creates a new instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2P(address common.Address, backend bind.ContractBackend) (*GlobalP2P, error) {
	contract, err := bindGlobalP2P(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlobalP2P{GlobalP2PCaller: GlobalP2PCaller{contract: contract}, GlobalP2PTransactor: GlobalP2PTransactor{contract: contract}, GlobalP2PFilterer: GlobalP2PFilterer{contract: contract}}, nil
}

// NewGlobalP2PCaller creates a new read-only instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2PCaller(address common.Address, caller bind.ContractCaller) (*GlobalP2PCaller, error) {
	contract, err := bindGlobalP2P(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PCaller{contract: contract}, nil
}

// NewGlobalP2PTransactor creates a new write-only instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2PTransactor(address common.Address, transactor bind.ContractTransactor) (*GlobalP2PTransactor, error) {
	contract, err := bindGlobalP2P(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PTransactor{contract: contract}, nil
}

// NewGlobalP2PFilterer creates a new log filterer instance of GlobalP2P, bound to a specific deployed contract.
func NewGlobalP2PFilterer(address common.Address, filterer bind.ContractFilterer) (*GlobalP2PFilterer, error) {
	contract, err := bindGlobalP2P(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PFilterer{contract: contract}, nil
}

// bindGlobalP2P binds a generic wrapper to an already deployed contract.
func bindGlobalP2P(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalP2PABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseGlobalP2PABI parses the ABI
func ParseGlobalP2PABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(GlobalP2PABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalP2P *GlobalP2PRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalP2P.Contract.GlobalP2PCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalP2P *GlobalP2PRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.Contract.GlobalP2PTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalP2P *GlobalP2PRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalP2P.Contract.GlobalP2PTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlobalP2P *GlobalP2PCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlobalP2P.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlobalP2P *GlobalP2PTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlobalP2P *GlobalP2PTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlobalP2P.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GlobalP2P *GlobalP2PCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlobalP2P.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GlobalP2P *GlobalP2PSession) Owner() (common.Address, error) {
	return _GlobalP2P.Contract.Owner(&_GlobalP2P.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) Owner() (common.Address, error) {
	return _GlobalP2P.Contract.Owner(&_GlobalP2P.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GlobalP2P *GlobalP2PCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GlobalP2P.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GlobalP2P *GlobalP2PSession) Paused() (bool, error) {
	return _GlobalP2P.Contract.Paused(&_GlobalP2P.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_GlobalP2P *GlobalP2PCallerSession) Paused() (bool, error) {
	return _GlobalP2P.Contract.Paused(&_GlobalP2P.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(bool)
func (_GlobalP2P *GlobalP2PCaller) Registry(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GlobalP2P.contract.Call(opts, &out, "registry", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(bool)
func (_GlobalP2P *GlobalP2PSession) Registry(arg0 common.Address) (bool, error) {
	return _GlobalP2P.Contract.Registry(&_GlobalP2P.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x038defd7.
//
// Solidity: function registry(address ) view returns(bool)
func (_GlobalP2P *GlobalP2PCallerSession) Registry(arg0 common.Address) (bool, error) {
	return _GlobalP2P.Contract.Registry(&_GlobalP2P.CallOpts, arg0)
}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_GlobalP2P *GlobalP2PCaller) UsdToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlobalP2P.contract.Call(opts, &out, "usdToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_GlobalP2P *GlobalP2PSession) UsdToken() (common.Address, error) {
	return _GlobalP2P.Contract.UsdToken(&_GlobalP2P.CallOpts)
}

// UsdToken is a free data retrieval call binding the contract method 0xf897a22b.
//
// Solidity: function usdToken() view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) UsdToken() (common.Address, error) {
	return _GlobalP2P.Contract.UsdToken(&_GlobalP2P.CallOpts)
}

// WalletLogic is a free data retrieval call binding the contract method 0xbeacf2cb.
//
// Solidity: function walletLogic() view returns(address)
func (_GlobalP2P *GlobalP2PCaller) WalletLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlobalP2P.contract.Call(opts, &out, "walletLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WalletLogic is a free data retrieval call binding the contract method 0xbeacf2cb.
//
// Solidity: function walletLogic() view returns(address)
func (_GlobalP2P *GlobalP2PSession) WalletLogic() (common.Address, error) {
	return _GlobalP2P.Contract.WalletLogic(&_GlobalP2P.CallOpts)
}

// WalletLogic is a free data retrieval call binding the contract method 0xbeacf2cb.
//
// Solidity: function walletLogic() view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) WalletLogic() (common.Address, error) {
	return _GlobalP2P.Contract.WalletLogic(&_GlobalP2P.CallOpts)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(address)
func (_GlobalP2P *GlobalP2PCaller) Wallets(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _GlobalP2P.contract.Call(opts, &out, "wallets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(address)
func (_GlobalP2P *GlobalP2PSession) Wallets(arg0 string) (common.Address, error) {
	return _GlobalP2P.Contract.Wallets(&_GlobalP2P.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x72f179f8.
//
// Solidity: function wallets(string ) view returns(address)
func (_GlobalP2P *GlobalP2PCallerSession) Wallets(arg0 string) (common.Address, error) {
	return _GlobalP2P.Contract.Wallets(&_GlobalP2P.CallOpts, arg0)
}

// DeployWallet is a paid mutator transaction binding the contract method 0x37ba532c.
//
// Solidity: function deployWallet(string uuid) returns(address)
func (_GlobalP2P *GlobalP2PTransactor) DeployWallet(opts *bind.TransactOpts, uuid string) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "deployWallet", uuid)
}

// DeployWallet is a paid mutator transaction binding the contract method 0x37ba532c.
//
// Solidity: function deployWallet(string uuid) returns(address)
func (_GlobalP2P *GlobalP2PSession) DeployWallet(uuid string) (*types.Transaction, error) {
	return _GlobalP2P.Contract.DeployWallet(&_GlobalP2P.TransactOpts, uuid)
}

// DeployWallet is a paid mutator transaction binding the contract method 0x37ba532c.
//
// Solidity: function deployWallet(string uuid) returns(address)
func (_GlobalP2P *GlobalP2PTransactorSession) DeployWallet(uuid string) (*types.Transaction, error) {
	return _GlobalP2P.Contract.DeployWallet(&_GlobalP2P.TransactOpts, uuid)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _walletLogic, address _usd) returns()
func (_GlobalP2P *GlobalP2PTransactor) Initialize(opts *bind.TransactOpts, _walletLogic common.Address, _usd common.Address) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "initialize", _walletLogic, _usd)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _walletLogic, address _usd) returns()
func (_GlobalP2P *GlobalP2PSession) Initialize(_walletLogic common.Address, _usd common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.Initialize(&_GlobalP2P.TransactOpts, _walletLogic, _usd)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _walletLogic, address _usd) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) Initialize(_walletLogic common.Address, _usd common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.Initialize(&_GlobalP2P.TransactOpts, _walletLogic, _usd)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GlobalP2P *GlobalP2PTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GlobalP2P *GlobalP2PSession) RenounceOwnership() (*types.Transaction, error) {
	return _GlobalP2P.Contract.RenounceOwnership(&_GlobalP2P.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GlobalP2P *GlobalP2PTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GlobalP2P.Contract.RenounceOwnership(&_GlobalP2P.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GlobalP2P *GlobalP2PTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GlobalP2P *GlobalP2PSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.TransferOwnership(&_GlobalP2P.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.TransferOwnership(&_GlobalP2P.TransactOpts, newOwner)
}

// UpdateWalletLogic is a paid mutator transaction binding the contract method 0x82787c7f.
//
// Solidity: function updateWalletLogic(address _walletLogic) returns()
func (_GlobalP2P *GlobalP2PTransactor) UpdateWalletLogic(opts *bind.TransactOpts, _walletLogic common.Address) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "updateWalletLogic", _walletLogic)
}

// UpdateWalletLogic is a paid mutator transaction binding the contract method 0x82787c7f.
//
// Solidity: function updateWalletLogic(address _walletLogic) returns()
func (_GlobalP2P *GlobalP2PSession) UpdateWalletLogic(_walletLogic common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.UpdateWalletLogic(&_GlobalP2P.TransactOpts, _walletLogic)
}

// UpdateWalletLogic is a paid mutator transaction binding the contract method 0x82787c7f.
//
// Solidity: function updateWalletLogic(address _walletLogic) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) UpdateWalletLogic(_walletLogic common.Address) (*types.Transaction, error) {
	return _GlobalP2P.Contract.UpdateWalletLogic(&_GlobalP2P.TransactOpts, _walletLogic)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PTransactor) WithdrawBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.contract.Transact(opts, "withdrawBalance", amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PSession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.Contract.WithdrawBalance(&_GlobalP2P.TransactOpts, amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns()
func (_GlobalP2P *GlobalP2PTransactorSession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _GlobalP2P.Contract.WithdrawBalance(&_GlobalP2P.TransactOpts, amount)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_GlobalP2P *GlobalP2PFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _GlobalP2P.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "Initialized":
		event, err = _GlobalP2P.ParseInitialized(log)
	case "OwnershipTransferred":
		event, err = _GlobalP2P.ParseOwnershipTransferred(log)
	case "Paused":
		event, err = _GlobalP2P.ParsePaused(log)
	case "Unpaused":
		event, err = _GlobalP2P.ParseUnpaused(log)
	case "NewWallet":
		event, err = _GlobalP2P.ParseNewWallet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlobalP2P *GlobalP2PTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlobalP2P.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlobalP2P *GlobalP2PSession) Receive() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Receive(&_GlobalP2P.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlobalP2P *GlobalP2PTransactorSession) Receive() (*types.Transaction, error) {
	return _GlobalP2P.Contract.Receive(&_GlobalP2P.TransactOpts)
}

// GlobalP2PInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the GlobalP2P contract.
type GlobalP2PInitializedIterator struct {
	Event *GlobalP2PInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlobalP2PInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlobalP2PInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlobalP2PInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PInitialized represents a Initialized event raised by the GlobalP2P contract.
type GlobalP2PInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GlobalP2P *GlobalP2PFilterer) FilterInitialized(opts *bind.FilterOpts) (*GlobalP2PInitializedIterator, error) {

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &GlobalP2PInitializedIterator{contract: _GlobalP2P.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GlobalP2P *GlobalP2PFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *GlobalP2PInitialized) (event.Subscription, error) {

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PInitialized)
				if err := _GlobalP2P.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_GlobalP2P *GlobalP2PFilterer) ParseInitialized(log types.Log) (*GlobalP2PInitialized, error) {
	event := new(GlobalP2PInitialized)
	if err := _GlobalP2P.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlobalP2POwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GlobalP2P contract.
type GlobalP2POwnershipTransferredIterator struct {
	Event *GlobalP2POwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlobalP2POwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2POwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlobalP2POwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlobalP2POwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2POwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2POwnershipTransferred represents a OwnershipTransferred event raised by the GlobalP2P contract.
type GlobalP2POwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GlobalP2P *GlobalP2PFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GlobalP2POwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GlobalP2POwnershipTransferredIterator{contract: _GlobalP2P.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GlobalP2P *GlobalP2PFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GlobalP2POwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2POwnershipTransferred)
				if err := _GlobalP2P.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GlobalP2P *GlobalP2PFilterer) ParseOwnershipTransferred(log types.Log) (*GlobalP2POwnershipTransferred, error) {
	event := new(GlobalP2POwnershipTransferred)
	if err := _GlobalP2P.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlobalP2PPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the GlobalP2P contract.
type GlobalP2PPausedIterator struct {
	Event *GlobalP2PPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlobalP2PPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlobalP2PPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlobalP2PPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PPaused represents a Paused event raised by the GlobalP2P contract.
type GlobalP2PPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GlobalP2P *GlobalP2PFilterer) FilterPaused(opts *bind.FilterOpts) (*GlobalP2PPausedIterator, error) {

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &GlobalP2PPausedIterator{contract: _GlobalP2P.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GlobalP2P *GlobalP2PFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *GlobalP2PPaused) (event.Subscription, error) {

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PPaused)
				if err := _GlobalP2P.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_GlobalP2P *GlobalP2PFilterer) ParsePaused(log types.Log) (*GlobalP2PPaused, error) {
	event := new(GlobalP2PPaused)
	if err := _GlobalP2P.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlobalP2PUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the GlobalP2P contract.
type GlobalP2PUnpausedIterator struct {
	Event *GlobalP2PUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlobalP2PUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlobalP2PUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlobalP2PUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PUnpaused represents a Unpaused event raised by the GlobalP2P contract.
type GlobalP2PUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GlobalP2P *GlobalP2PFilterer) FilterUnpaused(opts *bind.FilterOpts) (*GlobalP2PUnpausedIterator, error) {

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &GlobalP2PUnpausedIterator{contract: _GlobalP2P.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GlobalP2P *GlobalP2PFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *GlobalP2PUnpaused) (event.Subscription, error) {

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PUnpaused)
				if err := _GlobalP2P.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_GlobalP2P *GlobalP2PFilterer) ParseUnpaused(log types.Log) (*GlobalP2PUnpaused, error) {
	event := new(GlobalP2PUnpaused)
	if err := _GlobalP2P.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlobalP2PNewWalletIterator is returned from FilterNewWallet and is used to iterate over the raw logs and unpacked data for NewWallet events raised by the GlobalP2P contract.
type GlobalP2PNewWalletIterator struct {
	Event *GlobalP2PNewWallet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlobalP2PNewWalletIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlobalP2PNewWallet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlobalP2PNewWallet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlobalP2PNewWalletIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlobalP2PNewWalletIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlobalP2PNewWallet represents a NewWallet event raised by the GlobalP2P contract.
type GlobalP2PNewWallet struct {
	Uuid   common.Hash
	Wallet common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewWallet is a free log retrieval operation binding the contract event 0x1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b.
//
// Solidity: event newWallet(string indexed uuid, address indexed wallet)
func (_GlobalP2P *GlobalP2PFilterer) FilterNewWallet(opts *bind.FilterOpts, uuid []string, wallet []common.Address) (*GlobalP2PNewWalletIterator, error) {

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _GlobalP2P.contract.FilterLogs(opts, "newWallet", uuidRule, walletRule)
	if err != nil {
		return nil, err
	}
	return &GlobalP2PNewWalletIterator{contract: _GlobalP2P.contract, event: "newWallet", logs: logs, sub: sub}, nil
}

// WatchNewWallet is a free log subscription operation binding the contract event 0x1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b.
//
// Solidity: event newWallet(string indexed uuid, address indexed wallet)
func (_GlobalP2P *GlobalP2PFilterer) WatchNewWallet(opts *bind.WatchOpts, sink chan<- *GlobalP2PNewWallet, uuid []string, wallet []common.Address) (event.Subscription, error) {

	var uuidRule []interface{}
	for _, uuidItem := range uuid {
		uuidRule = append(uuidRule, uuidItem)
	}
	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _GlobalP2P.contract.WatchLogs(opts, "newWallet", uuidRule, walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlobalP2PNewWallet)
				if err := _GlobalP2P.contract.UnpackLog(event, "newWallet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewWallet is a log parse operation binding the contract event 0x1faae9711aa9dec64b6f886e5b953957f99597b26b899a71373d474b7928fd9b.
//
// Solidity: event newWallet(string indexed uuid, address indexed wallet)
func (_GlobalP2P *GlobalP2PFilterer) ParseNewWallet(log types.Log) (*GlobalP2PNewWallet, error) {
	event := new(GlobalP2PNewWallet)
	if err := _GlobalP2P.contract.UnpackLog(event, "newWallet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
