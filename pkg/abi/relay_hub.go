// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// RelayHubMetaData contains all meta data concerning the RelayHub contract.
var RelayHubMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIProofVerificationFunction\",\"name\":\"defaultVerificationFunction\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"BridgeRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"BridgeUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"ChainReseted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"validatorSet\",\"type\":\"address[]\"}],\"name\":\"ValidatorSetUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"blockProofs\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes\",\"name\":\"rawReceipt\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofSiblings\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proofPath\",\"type\":\"bytes\"}],\"name\":\"checkReceiptProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"checkValidators\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"blockStart\",\"type\":\"uint256\"}],\"name\":\"checkValidatorsAndQuorumReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"enableCrossChainBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"getLatestEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"contractIProofVerificationFunction\",\"name\":\"verificationFunction\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawRegisterBlock\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"epochLength\",\"type\":\"uint32\"}],\"name\":\"registerBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"resetChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status_\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"}],\"name\":\"unregisterBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"blockProofs\",\"type\":\"bytes[]\"}],\"name\":\"updateValidatorSetUsingEpochBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611c44380380611c4483398101604081905261002f91610078565b60018054336001600160a01b0319918216811783556000908152600260205260408120805460ff19169093179092558154166001600160a01b03929092169190911790556100a8565b60006020828403121561008a57600080fd5b81516001600160a01b03811681146100a157600080fd5b9392505050565b611b8d806100b76000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637518ee25116100665780637518ee251461013a5780637e4b28db1461015d578063973ebcd81461018e578063c80f0c20146101a1578063d2e1a496146101b457600080fd5b80630c46a0e1146100a3578063383d0ad9146100ec578063558a7297146101015780636f72020d146101145780637018b4f914610127575b600080fd5b6100cf6100b1366004611324565b6000908152600460205260409020600101546001600160a01b031690565b6040516001600160a01b0390911681526020015b60405180910390f35b6100ff6100fa366004611388565b6101c7565b005b6100ff61010f3660046113f6565b6105b7565b6100ff610122366004611324565b610613565b6100ff610135366004611324565b6106cc565b61014d61014836600461142f565b610754565b60405190151581526020016100e3565b61018061016b366004611324565b60009081526003602052604090206005015490565b6040519081526020016100e3565b61014d61019c3660046114c2565b610a3e565b6100ff6101af366004611591565b610c36565b6100ff6101c236600461161e565b610f8e565b6000838152600460205260408082208151608081019092528054829060ff1660028111156101f7576101f7611643565b600281111561020857610208611643565b8152815461010090046001600160a01b0390811660208301526001928301549081166040830152600160a01b900463ffffffff166060909101529091508151600281111561025857610258611643565b1480610276575060028151600281111561027457610274611643565b145b6102b85760405162461bcd60e51b815260206004820152600e60248201526d1b9bdd081c9959da5cdd195c995960921b60448201526064015b60405180910390fd5b60006102c78260200151610fff565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091529091506060826001600160a01b0316639c27dfe6888888600081811061032857610328611659565b905060200281019061033a919061166f565b88606001516040518563ffffffff1660e01b815260040161035e94939291906116de565b600060405180830381865afa15801561037b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103a391908101906117fb565b8092508193505050836060015163ffffffff1682604001516103c591906118ca565b6001600160401b03161561040d5760405162461bcd60e51b815260206004820152600f60248201526e6e6f742065706f636820626c6f636b60881b60448201526064016102af565b60006003600089815260200190815260200160002090506104518282876060015163ffffffff1686604001516104439190611906565b6001600160401b0316611029565b61045e8560200151610fff565b6001600160a01b03166342fcdfe38989898960600151306040518663ffffffff1660e01b815260040161049595949392919061192c565b600060405180830381865afa1580156104b2573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104da91908101906117fb565b505060028086526000898152600460205260409020865181548893839160ff191690600190849081111561051057610510611643565b0217905550602082015181546001600160a01b0391821661010002610100600160a81b03199091161782556040808401516001909301805460609095015163ffffffff16600160a01b026001600160c01b03199095169390921692909217929092179091555188907f3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e906105a59085906119ef565b60405180910390a25050505050505050565b3360009081526002602052604090205460ff16806105df57506001546001600160a01b031633145b6105e857600080fd5b6001600160a01b03919091166000908152600260205260409020805460ff1916911515919091179055565b3360009081526002602052604090205460ff168061063b57506001546001600160a01b031633145b61064457600080fd5b600081815260046020908152604080832080546001600160a81b031916815560010180546001600160c01b031916905560039091528120908181818161068a82826112ee565b50506000600485018190556005909401849055505060405183927f1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed925090a250565b3360009081526002602052604090205460ff16806106f457506001546001600160a01b031633145b6106fd57600080fd5b60008181526004602052604080822080546001600160a81b031916815560010180546001600160c01b03191690555182917f5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac8791a250565b60008481526004602090815260408083206001015460039092528220600160a01b90910463ffffffff16908261078a8386611a3c565b905081600401548110156107d45760405162461bcd60e51b815260206004820152601160248201527018985908195c1bd8da081cdd185c9d1959607a1b60448201526064016102af565b60008181526003830160205260408120548160086107f38360ff611a50565b901c6001600160401b0381111561080c5761080c61170f565b604051908082528060200260200182016040528015610835578160200160208202803683370190505b5090506000805b8a811015610a13578761084f828c611a50565b6108599190611a3c565b91506001821180156108aa57506002600388016000610879600186611a68565b8152602001908152602001600020546108929190611a3c565b8861089d838d611a50565b6108a79190611a7f565b11155b156108bd576108ba600183611a68565b91505b86600501548211156109045760405162461bcd60e51b815260206004820152601060248201526f3130b21032b837b1b410373ab6b132b960811b60448201526064016102af565b60006001888101828f8f8681811061091e5761091e611659565b90506020020160208101906109339190611a93565b6001600160a01b031660001b8152602001908152602001600020546109589190611a68565b9050610965816001611a50565b15806109885750600083815260028901602052604090206109869082611254565b155b156109935750610a01565b60008160ff166001901b90508085600884901c815181106109b6576109b6611659565b602002602001015116600014156109d557866109d181611ab0565b9750505b8085600884901c815181106109ec576109ec611659565b60200260200101818151179150818152505050505b80610a0b81611ab0565b91505061083c565b506003610a21846002611acb565b610a2b9190611a3c565b909310159b9a5050505050505050505050565b600089815260046020526040808220815160808101909252805483929190829060ff166002811115610a7257610a72611643565b6002811115610a8357610a83611643565b8152815461010090046001600160a01b0390811660208301526001909201549182166040820152600160a01b90910463ffffffff166060909101529050600281516002811115610ad557610ad5611643565b14610b0f5760405162461bcd60e51b815260206004820152600a6024820152696e6f742061637469766560b01b60448201526064016102af565b6000610b1e8260200151610fff565b6001600160a01b03166342fcdfe38d8d8d8660600151306040518663ffffffff1660e01b8152600401610b5595949392919061192c565b600060405180830381865afa158015610b72573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610b9a91908101906117fb565b509050610baa8260200151610fff565b6001600160a01b0316637f55ad248a8a84608001518b8b8b8b6040518863ffffffff1660e01b8152600401610be59796959493929190611aea565b602060405180830381865afa158015610c02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c269190611b3a565b9c9b505050505050505050505050565b3360009081526002602052604090205460ff1680610c5e57506001546001600160a01b031633145b610c6757600080fd5b6000868152600460205260408082208151608081019092528054829060ff166002811115610c9757610c97611643565b6002811115610ca857610ca8611643565b8152815461010090046001600160a01b0390811660208301526001909201549182166040820152600160a01b90910463ffffffff166060909101529050600081516002811115610cfa57610cfa611643565b1480610d185750600181516002811115610d1657610d16611643565b145b610d595760405162461bcd60e51b8152602060048201526012602482015271185b1c9958591e481c9959da5cdd195c995960721b60448201526064016102af565b600080610d6588610fff565b6001600160a01b0316639c27dfe68a8989886040518563ffffffff1660e01b8152600401610d9694939291906116de565b600060405180830381865afa158015610db3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ddb91908101906117fb565b915091508363ffffffff168260400151610df591906118ca565b6001600160401b031615610e3d5760405162461bcd60e51b815260206004820152600f60248201526e6e6f742065706f636820626c6f636b60881b60448201526064016102af565b600183526001600160a01b0388811660208086019190915290861660408086019190915263ffffffff86166060860181905260008c8152600390935291819020908401519091610e9591849184916104439190611906565b8463ffffffff168360400151610eab9190611906565b6001600160401b031660048201556002848190525060008a81526004602052604090208451815486929190829060ff19166001836002811115610ef057610ef0611643565b0217905550602082015181546001600160a01b0391821661010002610100600160a81b03199091161782556040808401516001909301805460609095015163ffffffff16600160a01b026001600160c01b031990951693831693909317939093179091559051908716908b907fb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc5090600090a350505050505050505050565b3360009081526002602052604090205460ff1680610fb657506001546001600160a01b031633145b610fbf57600080fd5b6000918252600460205260409091206001810180546001600160a01b0319166001600160a01b03909316929092179091558054600260ff19909116179055565b60006001600160a01b0382166110205750506000546001600160a01b031690565b5090565b919050565b600582015415611088576005820154611043906001611a50565b81146110885760405162461bcd60e51b81526020600482015260146024820152733130b21032b837b1b4103a3930b739b4ba34b7b760611b60448201526064016102af565b60008351116110ce5760405162461bcd60e51b8152602060048201526012602482015271189859081d985b1a59185d1bdc9cc81cd95d60721b60448201526064016102af565b600060086110db84611279565b6110e7911c6001611a50565b6001600160401b038111156110fe576110fe61170f565b604051908082528060200260200182016040528015611127578160200160208202803683370190505b50905060005b84518110156111cf57600085828151811061114a5761114a611659565b602090810291909101015190506111618582611283565b506001600160a01b0381166000908152600186810160205260408220546111889190611a68565b90508060ff166001901b84600883901c815181106111a8576111a8611659565b602002602001018181511791508181525050505080806111c790611ab0565b91505061112d565b5060008281526002840160205260408120905b825181101561122a578281815181106111fd576111fd611659565b6020908102919091018101516000838152918490526040909120558061122281611ab0565b9150506111e2565b50509251600082815260038401602052604090206001600160401b03909116905560059091015550565b600881901c600090815260208390526040902054600160ff83161b1615155b92915050565b6000611273825490565b6000611298836001600160a01b03841661129f565b9392505050565b60008181526001830160205260408120546112e657508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611273565b506000611273565b508054600082559060005260206000209081019061130c919061130f565b50565b5b808211156110205760008155600101611310565b60006020828403121561133657600080fd5b5035919050565b60008083601f84011261134f57600080fd5b5081356001600160401b0381111561136657600080fd5b6020830191508360208260051b850101111561138157600080fd5b9250929050565b60008060006040848603121561139d57600080fd5b8335925060208401356001600160401b038111156113ba57600080fd5b6113c68682870161133d565b9497909650939450505050565b6001600160a01b038116811461130c57600080fd5b801515811461130c57600080fd5b6000806040838503121561140957600080fd5b8235611414816113d3565b91506020830135611424816113e8565b809150509250929050565b6000806000806060858703121561144557600080fd5b8435935060208501356001600160401b0381111561146257600080fd5b61146e8782880161133d565b9598909750949560400135949350505050565b60008083601f84011261149357600080fd5b5081356001600160401b038111156114aa57600080fd5b60208301915083602082850101111561138157600080fd5b600080600080600080600080600060a08a8c0312156114e057600080fd5b8935985060208a01356001600160401b03808211156114fe57600080fd5b61150a8d838e0161133d565b909a50985060408c013591508082111561152357600080fd5b61152f8d838e01611481565b909850965060608c013591508082111561154857600080fd5b6115548d838e01611481565b909650945060808c013591508082111561156d57600080fd5b5061157a8c828d01611481565b915080935050809150509295985092959850929598565b60008060008060008060a087890312156115aa57600080fd5b8635955060208701356115bc816113d3565b945060408701356001600160401b038111156115d757600080fd5b6115e389828a01611481565b90955093505060608701356115f7816113d3565b9150608087013563ffffffff8116811461161057600080fd5b809150509295509295509295565b6000806040838503121561163157600080fd5b823591506020830135611424816113d3565b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000808335601e1984360301811261168657600080fd5b8301803591506001600160401b038211156116a057600080fd5b60200191503681900382131561138157600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b8481526060602082015260006116f86060830185876116b5565b905063ffffffff8316604083015295945050505050565b634e487b7160e01b600052604160045260246000fd5b60405160e081016001600160401b03811182821017156117475761174761170f565b60405290565b8051611024816113d3565b600082601f83011261176957600080fd5b815160206001600160401b03808311156117855761178561170f565b8260051b604051601f19603f830116810181811084821117156117aa576117aa61170f565b6040529384528581018301938381019250878511156117c857600080fd5b83870191505b848210156117f05781516117e1816113d3565b835291830191908301906117ce565b979650505050505050565b60008082840361010081121561181057600080fd5b60e081121561181e57600080fd5b50611827611725565b835181526020840151602082015260408401516001600160401b03808216821461185057600080fd5b8160408401526118626060870161174d565b60608401526080860151608084015260a086015160a084015260c086015160c084015282945060e086015192508083111561189c57600080fd5b50506118aa85828601611758565b9150509250929050565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b03808416806118e4576118e46118b4565b92169190910692915050565b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380841680611920576119206118b4565b92169190910492915050565b60006080820187835260206080818501528187835260a08501905060a08860051b86010192508860005b898110156119c457868503609f190183528135368c9003601e1901811261197c57600080fd5b8b0180356001600160401b0381111561199457600080fd5b8036038d13156119a357600080fd5b6119b087828885016116b5565b965050509183019190830190600101611956565b50505063ffffffff861660408501525090506001600160a01b03831660608301529695505050505050565b6020808252825182820181905260009190848201906040850190845b81811015611a305783516001600160a01b031683529284019291840191600101611a0b565b50909695505050505050565b600082611a4b57611a4b6118b4565b500490565b60008219821115611a6357611a636118f0565b500190565b600082821015611a7a57611a7a6118f0565b500390565b600082611a8e57611a8e6118b4565b500690565b600060208284031215611aa557600080fd5b8135611298816113d3565b6000600019821415611ac457611ac46118f0565b5060010190565b6000816000190483118215151615611ae557611ae56118f0565b500290565b608081526000611afe60808301898b6116b5565b8760208401528281036040840152611b178187896116b5565b90508281036060840152611b2c8185876116b5565b9a9950505050505050505050565b600060208284031215611b4c57600080fd5b8151611298816113e856fea264697066735822122015585927dbc2b73f6632f0cfabad9b2245b85746f7feb19f19354a2efd689e3e64736f6c634300080a0033",
}

// RelayHubABI is the input ABI used to generate the binding from.
// Deprecated: Use RelayHubMetaData.ABI instead.
var RelayHubABI = RelayHubMetaData.ABI

// RelayHubBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RelayHubMetaData.Bin instead.
var RelayHubBin = RelayHubMetaData.Bin

// DeployRelayHub deploys a new Ethereum contract, binding an instance of RelayHub to it.
func DeployRelayHub(auth *bind.TransactOpts, backend bind.ContractBackend, defaultVerificationFunction common.Address) (common.Address, *types.Transaction, *RelayHub, error) {
	parsed, err := RelayHubMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RelayHubBin), backend, defaultVerificationFunction)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RelayHub{RelayHubCaller: RelayHubCaller{contract: contract}, RelayHubTransactor: RelayHubTransactor{contract: contract}, RelayHubFilterer: RelayHubFilterer{contract: contract}}, nil
}

// RelayHub is an auto generated Go binding around an Ethereum contract.
type RelayHub struct {
	RelayHubCaller     // Read-only binding to the contract
	RelayHubTransactor // Write-only binding to the contract
	RelayHubFilterer   // Log filterer for contract events
}

// RelayHubCaller is an auto generated read-only Go binding around an Ethereum contract.
type RelayHubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RelayHubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RelayHubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RelayHubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RelayHubSession struct {
	Contract     *RelayHub         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RelayHubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RelayHubCallerSession struct {
	Contract *RelayHubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RelayHubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RelayHubTransactorSession struct {
	Contract     *RelayHubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RelayHubRaw is an auto generated low-level Go binding around an Ethereum contract.
type RelayHubRaw struct {
	Contract *RelayHub // Generic contract binding to access the raw methods on
}

// RelayHubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RelayHubCallerRaw struct {
	Contract *RelayHubCaller // Generic read-only contract binding to access the raw methods on
}

// RelayHubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RelayHubTransactorRaw struct {
	Contract *RelayHubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRelayHub creates a new instance of RelayHub, bound to a specific deployed contract.
func NewRelayHub(address common.Address, backend bind.ContractBackend) (*RelayHub, error) {
	contract, err := bindRelayHub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RelayHub{RelayHubCaller: RelayHubCaller{contract: contract}, RelayHubTransactor: RelayHubTransactor{contract: contract}, RelayHubFilterer: RelayHubFilterer{contract: contract}}, nil
}

// NewRelayHubCaller creates a new read-only instance of RelayHub, bound to a specific deployed contract.
func NewRelayHubCaller(address common.Address, caller bind.ContractCaller) (*RelayHubCaller, error) {
	contract, err := bindRelayHub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubCaller{contract: contract}, nil
}

// NewRelayHubTransactor creates a new write-only instance of RelayHub, bound to a specific deployed contract.
func NewRelayHubTransactor(address common.Address, transactor bind.ContractTransactor) (*RelayHubTransactor, error) {
	contract, err := bindRelayHub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RelayHubTransactor{contract: contract}, nil
}

// NewRelayHubFilterer creates a new log filterer instance of RelayHub, bound to a specific deployed contract.
func NewRelayHubFilterer(address common.Address, filterer bind.ContractFilterer) (*RelayHubFilterer, error) {
	contract, err := bindRelayHub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RelayHubFilterer{contract: contract}, nil
}

// bindRelayHub binds a generic wrapper to an already deployed contract.
func bindRelayHub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RelayHubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHub *RelayHubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHub.Contract.RelayHubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHub *RelayHubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHub.Contract.RelayHubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHub *RelayHubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHub.Contract.RelayHubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RelayHub *RelayHubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RelayHub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RelayHub *RelayHubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RelayHub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RelayHub *RelayHubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RelayHub.Contract.contract.Transact(opts, method, params...)
}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHub *RelayHubCaller) CheckReceiptProof(opts *bind.CallOpts, chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "checkReceiptProof", chainId, blockProofs, rawReceipt, proofSiblings, proofPath)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHub *RelayHubSession) CheckReceiptProof(chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	return _RelayHub.Contract.CheckReceiptProof(&_RelayHub.CallOpts, chainId, blockProofs, rawReceipt, proofSiblings, proofPath)
}

// CheckReceiptProof is a free data retrieval call binding the contract method 0x973ebcd8.
//
// Solidity: function checkReceiptProof(uint256 chainId, bytes[] blockProofs, bytes rawReceipt, bytes proofSiblings, bytes proofPath) view returns(bool)
func (_RelayHub *RelayHubCallerSession) CheckReceiptProof(chainId *big.Int, blockProofs [][]byte, rawReceipt []byte, proofSiblings []byte, proofPath []byte) (bool, error) {
	return _RelayHub.Contract.CheckReceiptProof(&_RelayHub.CallOpts, chainId, blockProofs, rawReceipt, proofSiblings, proofPath)
}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubCaller) CheckValidatorsAndQuorumReached(opts *bind.CallOpts, chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "checkValidatorsAndQuorumReached", chainId, checkValidators, blockStart)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubSession) CheckValidatorsAndQuorumReached(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHub.Contract.CheckValidatorsAndQuorumReached(&_RelayHub.CallOpts, chainId, checkValidators, blockStart)
}

// CheckValidatorsAndQuorumReached is a free data retrieval call binding the contract method 0x7518ee25.
//
// Solidity: function checkValidatorsAndQuorumReached(uint256 chainId, address[] checkValidators, uint256 blockStart) view returns(bool)
func (_RelayHub *RelayHubCallerSession) CheckValidatorsAndQuorumReached(chainId *big.Int, checkValidators []common.Address, blockStart *big.Int) (bool, error) {
	return _RelayHub.Contract.CheckValidatorsAndQuorumReached(&_RelayHub.CallOpts, chainId, checkValidators, blockStart)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHub *RelayHubCaller) GetBridgeAddress(opts *bind.CallOpts, chainId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "getBridgeAddress", chainId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHub *RelayHubSession) GetBridgeAddress(chainId *big.Int) (common.Address, error) {
	return _RelayHub.Contract.GetBridgeAddress(&_RelayHub.CallOpts, chainId)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0x0c46a0e1.
//
// Solidity: function getBridgeAddress(uint256 chainId) view returns(address)
func (_RelayHub *RelayHubCallerSession) GetBridgeAddress(chainId *big.Int) (common.Address, error) {
	return _RelayHub.Contract.GetBridgeAddress(&_RelayHub.CallOpts, chainId)
}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHub *RelayHubCaller) GetLatestEpochNumber(opts *bind.CallOpts, chainId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RelayHub.contract.Call(opts, &out, "getLatestEpochNumber", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHub *RelayHubSession) GetLatestEpochNumber(chainId *big.Int) (*big.Int, error) {
	return _RelayHub.Contract.GetLatestEpochNumber(&_RelayHub.CallOpts, chainId)
}

// GetLatestEpochNumber is a free data retrieval call binding the contract method 0x7e4b28db.
//
// Solidity: function getLatestEpochNumber(uint256 chainId) view returns(uint256)
func (_RelayHub *RelayHubCallerSession) GetLatestEpochNumber(chainId *big.Int) (*big.Int, error) {
	return _RelayHub.Contract.GetLatestEpochNumber(&_RelayHub.CallOpts, chainId)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHub *RelayHubTransactor) EnableCrossChainBridge(opts *bind.TransactOpts, chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "enableCrossChainBridge", chainId, bridgeAddress)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHub *RelayHubSession) EnableCrossChainBridge(chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHub.Contract.EnableCrossChainBridge(&_RelayHub.TransactOpts, chainId, bridgeAddress)
}

// EnableCrossChainBridge is a paid mutator transaction binding the contract method 0xd2e1a496.
//
// Solidity: function enableCrossChainBridge(uint256 chainId, address bridgeAddress) returns()
func (_RelayHub *RelayHubTransactorSession) EnableCrossChainBridge(chainId *big.Int, bridgeAddress common.Address) (*types.Transaction, error) {
	return _RelayHub.Contract.EnableCrossChainBridge(&_RelayHub.TransactOpts, chainId, bridgeAddress)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns()
func (_RelayHub *RelayHubTransactor) RegisterBridge(opts *bind.TransactOpts, chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "registerBridge", chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns()
func (_RelayHub *RelayHubSession) RegisterBridge(chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHub.Contract.RegisterBridge(&_RelayHub.TransactOpts, chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// RegisterBridge is a paid mutator transaction binding the contract method 0xc80f0c20.
//
// Solidity: function registerBridge(uint256 chainId, address verificationFunction, bytes rawRegisterBlock, address bridgeAddress, uint32 epochLength) returns()
func (_RelayHub *RelayHubTransactorSession) RegisterBridge(chainId *big.Int, verificationFunction common.Address, rawRegisterBlock []byte, bridgeAddress common.Address, epochLength uint32) (*types.Transaction, error) {
	return _RelayHub.Contract.RegisterBridge(&_RelayHub.TransactOpts, chainId, verificationFunction, rawRegisterBlock, bridgeAddress, epochLength)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactor) ResetChain(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "resetChain", chainId)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHub *RelayHubSession) ResetChain(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.ResetChain(&_RelayHub.TransactOpts, chainId)
}

// ResetChain is a paid mutator transaction binding the contract method 0x6f72020d.
//
// Solidity: function resetChain(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactorSession) ResetChain(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.ResetChain(&_RelayHub.TransactOpts, chainId)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHub *RelayHubTransactor) SetOperator(opts *bind.TransactOpts, operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "setOperator", operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHub *RelayHubSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHub.Contract.SetOperator(&_RelayHub.TransactOpts, operator_, status_)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator_, bool status_) returns()
func (_RelayHub *RelayHubTransactorSession) SetOperator(operator_ common.Address, status_ bool) (*types.Transaction, error) {
	return _RelayHub.Contract.SetOperator(&_RelayHub.TransactOpts, operator_, status_)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactor) UnregisterBridge(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "unregisterBridge", chainId)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHub *RelayHubSession) UnregisterBridge(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.UnregisterBridge(&_RelayHub.TransactOpts, chainId)
}

// UnregisterBridge is a paid mutator transaction binding the contract method 0x7018b4f9.
//
// Solidity: function unregisterBridge(uint256 chainId) returns()
func (_RelayHub *RelayHubTransactorSession) UnregisterBridge(chainId *big.Int) (*types.Transaction, error) {
	return _RelayHub.Contract.UnregisterBridge(&_RelayHub.TransactOpts, chainId)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHub *RelayHubTransactor) UpdateValidatorSetUsingEpochBlocks(opts *bind.TransactOpts, chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHub.contract.Transact(opts, "updateValidatorSetUsingEpochBlocks", chainId, blockProofs)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHub *RelayHubSession) UpdateValidatorSetUsingEpochBlocks(chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHub.Contract.UpdateValidatorSetUsingEpochBlocks(&_RelayHub.TransactOpts, chainId, blockProofs)
}

// UpdateValidatorSetUsingEpochBlocks is a paid mutator transaction binding the contract method 0x383d0ad9.
//
// Solidity: function updateValidatorSetUsingEpochBlocks(uint256 chainId, bytes[] blockProofs) returns()
func (_RelayHub *RelayHubTransactorSession) UpdateValidatorSetUsingEpochBlocks(chainId *big.Int, blockProofs [][]byte) (*types.Transaction, error) {
	return _RelayHub.Contract.UpdateValidatorSetUsingEpochBlocks(&_RelayHub.TransactOpts, chainId, blockProofs)
}

// RelayHubBridgeRegisteredIterator is returned from FilterBridgeRegistered and is used to iterate over the raw logs and unpacked data for BridgeRegistered events raised by the RelayHub contract.
type RelayHubBridgeRegisteredIterator struct {
	Event *RelayHubBridgeRegistered // Event containing the contract specifics and raw log

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
func (it *RelayHubBridgeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBridgeRegistered)
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
		it.Event = new(RelayHubBridgeRegistered)
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
func (it *RelayHubBridgeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBridgeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBridgeRegistered represents a BridgeRegistered event raised by the RelayHub contract.
type RelayHubBridgeRegistered struct {
	ChainId       *big.Int
	BridgeAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRegistered is a free log retrieval operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHub *RelayHubFilterer) FilterBridgeRegistered(opts *bind.FilterOpts, chainId []*big.Int, bridgeAddress []common.Address) (*RelayHubBridgeRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var bridgeAddressRule []interface{}
	for _, bridgeAddressItem := range bridgeAddress {
		bridgeAddressRule = append(bridgeAddressRule, bridgeAddressItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "BridgeRegistered", chainIdRule, bridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBridgeRegisteredIterator{contract: _RelayHub.contract, event: "BridgeRegistered", logs: logs, sub: sub}, nil
}

// WatchBridgeRegistered is a free log subscription operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHub *RelayHubFilterer) WatchBridgeRegistered(opts *bind.WatchOpts, sink chan<- *RelayHubBridgeRegistered, chainId []*big.Int, bridgeAddress []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var bridgeAddressRule []interface{}
	for _, bridgeAddressItem := range bridgeAddress {
		bridgeAddressRule = append(bridgeAddressRule, bridgeAddressItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "BridgeRegistered", chainIdRule, bridgeAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBridgeRegistered)
				if err := _RelayHub.contract.UnpackLog(event, "BridgeRegistered", log); err != nil {
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

// ParseBridgeRegistered is a log parse operation binding the contract event 0xb72d37f67f55309a7342120ad674525186207bb1835dd860c16280472ea6fc50.
//
// Solidity: event BridgeRegistered(uint256 indexed chainId, address indexed bridgeAddress)
func (_RelayHub *RelayHubFilterer) ParseBridgeRegistered(log types.Log) (*RelayHubBridgeRegistered, error) {
	event := new(RelayHubBridgeRegistered)
	if err := _RelayHub.contract.UnpackLog(event, "BridgeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubBridgeUnregisteredIterator is returned from FilterBridgeUnregistered and is used to iterate over the raw logs and unpacked data for BridgeUnregistered events raised by the RelayHub contract.
type RelayHubBridgeUnregisteredIterator struct {
	Event *RelayHubBridgeUnregistered // Event containing the contract specifics and raw log

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
func (it *RelayHubBridgeUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubBridgeUnregistered)
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
		it.Event = new(RelayHubBridgeUnregistered)
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
func (it *RelayHubBridgeUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubBridgeUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubBridgeUnregistered represents a BridgeUnregistered event raised by the RelayHub contract.
type RelayHubBridgeUnregistered struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBridgeUnregistered is a free log retrieval operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) FilterBridgeUnregistered(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubBridgeUnregisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "BridgeUnregistered", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubBridgeUnregisteredIterator{contract: _RelayHub.contract, event: "BridgeUnregistered", logs: logs, sub: sub}, nil
}

// WatchBridgeUnregistered is a free log subscription operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) WatchBridgeUnregistered(opts *bind.WatchOpts, sink chan<- *RelayHubBridgeUnregistered, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "BridgeUnregistered", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubBridgeUnregistered)
				if err := _RelayHub.contract.UnpackLog(event, "BridgeUnregistered", log); err != nil {
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

// ParseBridgeUnregistered is a log parse operation binding the contract event 0x5ed0f717c1d62eb20db7f8f67a7f08a656c882be0834c2975e4991b546faac87.
//
// Solidity: event BridgeUnregistered(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) ParseBridgeUnregistered(log types.Log) (*RelayHubBridgeUnregistered, error) {
	event := new(RelayHubBridgeUnregistered)
	if err := _RelayHub.contract.UnpackLog(event, "BridgeUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubChainResetedIterator is returned from FilterChainReseted and is used to iterate over the raw logs and unpacked data for ChainReseted events raised by the RelayHub contract.
type RelayHubChainResetedIterator struct {
	Event *RelayHubChainReseted // Event containing the contract specifics and raw log

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
func (it *RelayHubChainResetedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubChainReseted)
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
		it.Event = new(RelayHubChainReseted)
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
func (it *RelayHubChainResetedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubChainResetedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubChainReseted represents a ChainReseted event raised by the RelayHub contract.
type RelayHubChainReseted struct {
	ChainId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChainReseted is a free log retrieval operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) FilterChainReseted(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubChainResetedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "ChainReseted", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubChainResetedIterator{contract: _RelayHub.contract, event: "ChainReseted", logs: logs, sub: sub}, nil
}

// WatchChainReseted is a free log subscription operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) WatchChainReseted(opts *bind.WatchOpts, sink chan<- *RelayHubChainReseted, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "ChainReseted", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubChainReseted)
				if err := _RelayHub.contract.UnpackLog(event, "ChainReseted", log); err != nil {
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

// ParseChainReseted is a log parse operation binding the contract event 0x1b3e6308933a8c83e369ce09eb5dc180027a75240e78b85b9832061d9f1716ed.
//
// Solidity: event ChainReseted(uint256 indexed chainId)
func (_RelayHub *RelayHubFilterer) ParseChainReseted(log types.Log) (*RelayHubChainReseted, error) {
	event := new(RelayHubChainReseted)
	if err := _RelayHub.contract.UnpackLog(event, "ChainReseted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RelayHubValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the RelayHub contract.
type RelayHubValidatorSetUpdatedIterator struct {
	Event *RelayHubValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *RelayHubValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RelayHubValidatorSetUpdated)
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
		it.Event = new(RelayHubValidatorSetUpdated)
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
func (it *RelayHubValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RelayHubValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RelayHubValidatorSetUpdated represents a ValidatorSetUpdated event raised by the RelayHub contract.
type RelayHubValidatorSetUpdated struct {
	ChainId      *big.Int
	ValidatorSet []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHub *RelayHubFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, chainId []*big.Int) (*RelayHubValidatorSetUpdatedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.FilterLogs(opts, "ValidatorSetUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &RelayHubValidatorSetUpdatedIterator{contract: _RelayHub.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHub *RelayHubFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *RelayHubValidatorSetUpdated, chainId []*big.Int) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _RelayHub.contract.WatchLogs(opts, "ValidatorSetUpdated", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RelayHubValidatorSetUpdated)
				if err := _RelayHub.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed chainId, address[] validatorSet)
func (_RelayHub *RelayHubFilterer) ParseValidatorSetUpdated(log types.Log) (*RelayHubValidatorSetUpdated, error) {
	event := new(RelayHubValidatorSetUpdated)
	if err := _RelayHub.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
