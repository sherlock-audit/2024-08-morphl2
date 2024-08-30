// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"morph-l2/bindings/solc"
)

const MorphStandardERC20StorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1023_storage\"},{\"astId\":1003,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1004,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"53\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_name\",\"offset\":0,\"slot\":\"54\",\"type\":\"t_string_storage\"},{\"astId\":1007,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"55\",\"type\":\"t_string_storage\"},{\"astId\":1008,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"56\",\"type\":\"t_array(t_uint256)1020_storage\"},{\"astId\":1009,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_hashedName\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bytes32\"},{\"astId\":1010,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_hashedVersion\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_bytes32\"},{\"astId\":1011,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_name\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_string_storage\"},{\"astId\":1012,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_version\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_string_storage\"},{\"astId\":1013,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_array(t_uint256)1021_storage\"},{\"astId\":1014,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_nonces\",\"offset\":0,\"slot\":\"153\",\"type\":\"t_mapping(t_address,t_struct(Counter)1024_storage)\"},{\"astId\":1015,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"_PERMIT_TYPEHASH_DEPRECATED_SLOT\",\"offset\":0,\"slot\":\"154\",\"type\":\"t_bytes32\"},{\"astId\":1016,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"155\",\"type\":\"t_array(t_uint256)1022_storage\"},{\"astId\":1017,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"gateway\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1018,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"counterpart\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_address\"},{\"astId\":1019,\"contract\":\"contracts/libraries/token/MorphStandardERC20.sol:MorphStandardERC20\",\"label\":\"decimals_\",\"offset\":20,\"slot\":\"205\",\"type\":\"t_uint8\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1020_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[45]\",\"numberOfBytes\":\"1440\"},\"t_array(t_uint256)1021_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[48]\",\"numberOfBytes\":\"1536\"},\"t_array(t_uint256)1022_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1023_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_struct(Counter)1024_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e struct CountersUpgradeable.Counter)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_struct(Counter)1024_storage\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_struct(Counter)1024_storage\":{\"encoding\":\"inplace\",\"label\":\"struct CountersUpgradeable.Counter\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var MorphStandardERC20StorageLayout = new(solc.StorageLayout)

var MorphStandardERC20DeployedBin = "0x608060405234801561000f575f80fd5b506004361061016e575f3560e01c806370a08231116100d25780639dc29fac11610088578063c820f14611610063578063c820f14614610354578063d505accf14610367578063dd62ed3e1461037a575f80fd5b80639dc29fac1461031b578063a457c2d71461032e578063a9059cbb14610341575f80fd5b80637ecebe00116100b85780637ecebe00146102e557806384b0196e146102f857806395d89b4114610313575f80fd5b806370a0823114610290578063797594b0146102c5575f80fd5b8063313ce56711610127578063395093511161010d57806339509351146102555780634000aea01461026857806340c10f191461027b575f80fd5b8063313ce5671461021d5780633644e5151461024d575f80fd5b8063116191b611610157578063116191b6146101b357806318160ddd146101f857806323b872dd1461020a575f80fd5b806306fdde0314610172578063095ea7b314610190575b5f80fd5b61017a6103bf565b6040516101879190611c24565b60405180910390f35b6101a361019e366004611c65565b61044f565b6040519015158152602001610187565b60cc546101d39073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610187565b6035545b604051908152602001610187565b6101a3610218366004611c8d565b610468565b60cd5474010000000000000000000000000000000000000000900460ff1660405160ff9091168152602001610187565b6101fc61048b565b6101a3610263366004611c65565b610499565b6101a3610276366004611cc6565b6104e4565b61028e610289366004611c65565b61054d565b005b6101fc61029e366004611d46565b73ffffffffffffffffffffffffffffffffffffffff165f9081526033602052604090205490565b60cd546101d39073ffffffffffffffffffffffffffffffffffffffff1681565b6101fc6102f3366004611d46565b6105e1565b61030061060b565b6040516101879796959493929190611d5f565b61017a6106e2565b61028e610329366004611c65565b6106f1565b6101a361033c366004611c65565b61077c565b6101a361034f366004611c65565b61084c565b61028e610362366004611f03565b610859565b61028e610375366004611f93565b610a8a565b6101fc610388366004611ff8565b73ffffffffffffffffffffffffffffffffffffffff9182165f90815260346020908152604080832093909416825291909152205490565b6060603680546103ce90612029565b80601f01602080910402602001604051908101604052809291908181526020018280546103fa90612029565b80156104455780601f1061041c57610100808354040283529160200191610445565b820191905f5260205f20905b81548152906001019060200180831161042857829003601f168201915b5050505050905090565b5f3361045c818585610c46565b60019150505b92915050565b5f33610475858285610df9565b610480858585610ecf565b506001949350505050565b5f610494611143565b905090565b335f81815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061045c90829086906104df908790612074565b610c46565b5f6104ef858561084c565b5073ffffffffffffffffffffffffffffffffffffffff85163b1561048057610480858585858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061114c92505050565b60cc5473ffffffffffffffffffffffffffffffffffffffff1633146105d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4f6e6c792047617465776179000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6105dd82826111d7565b5050565b73ffffffffffffffffffffffffffffffffffffffff81165f90815260996020526040812054610462565b5f6060805f805f60606065545f801b1480156106275750606654155b61068d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f4549503731323a20556e696e697469616c697a6564000000000000000000000060448201526064016105ca565b6106956112ca565b61069d6112d9565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b6060603780546103ce90612029565b60cc5473ffffffffffffffffffffffffffffffffffffffff163314610772576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4f6e6c792047617465776179000000000000000000000000000000000000000060448201526064016105ca565b6105dd82826112e8565b335f81815260346020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561083f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084016105ca565b6104808286868403610c46565b5f3361045c818585610ecf565b5f54610100900460ff161580801561087757505f54600160ff909116105b806108905750303b15801561089057505f5460ff166001145b61091c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016105ca565b5f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790558015610978575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610981866114a9565b61098b8686611581565b60cd805460cc805473ffffffffffffffffffffffffffffffffffffffff8088167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925590851660ff88167401000000000000000000000000000000000000000002919091167fffffffffffffffffffffff000000000000000000000000000000000000000000909216919091171790558015610a82575f80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050565b83421115610af4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e6500000060448201526064016105ca565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9888888610b228c611621565b60408051602081019690965273ffffffffffffffffffffffffffffffffffffffff94851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f610b8982611655565b90505f610b988287878761169c565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610c2f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e6174757265000060448201526064016105ca565b610c3a8a8a8a610c46565b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610ce8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016105ca565b73ffffffffffffffffffffffffffffffffffffffff8216610d8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016105ca565b73ffffffffffffffffffffffffffffffffffffffff8381165f8181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381165f908152603460209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610ec95781811015610ebc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016105ca565b610ec98484848403610c46565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610f72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016105ca565b73ffffffffffffffffffffffffffffffffffffffff8216611015576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016105ca565b73ffffffffffffffffffffffffffffffffffffffff83165f90815260336020526040902054818110156110ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016105ca565b73ffffffffffffffffffffffffffffffffffffffff8085165f8181526033602052604080822086860390559286168082529083902080548601905591517fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906111369086815260200190565b60405180910390a3610ec9565b5f6104946116c2565b6040517fa4c0ed36000000000000000000000000000000000000000000000000000000008152839073ffffffffffffffffffffffffffffffffffffffff82169063a4c0ed36906111a4903390879087906004016120ac565b5f604051808303815f87803b1580156111bb575f80fd5b505af11580156111cd573d5f803e3d5ffd5b5050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8216611254576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016105ca565b8060355f8282546112659190612074565b909155505073ffffffffffffffffffffffffffffffffffffffff82165f818152603360209081526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6060606780546103ce90612029565b6060606880546103ce90612029565b73ffffffffffffffffffffffffffffffffffffffff821661138b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f730000000000000000000000000000000000000000000000000000000000000060648201526084016105ca565b73ffffffffffffffffffffffffffffffffffffffff82165f9081526033602052604090205481811015611440576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f636500000000000000000000000000000000000000000000000000000000000060648201526084016105ca565b73ffffffffffffffffffffffffffffffffffffffff83165f8181526033602090815260408083208686039055603580548790039055518581529192917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9101610dec565b505050565b5f54610100900460ff1661153f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105ca565b61157e816040518060400160405280600181526020017f3100000000000000000000000000000000000000000000000000000000000000815250611735565b50565b5f54610100900460ff16611617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105ca565b6105dd82826117f2565b73ffffffffffffffffffffffffffffffffffffffff81165f9081526099602052604090208054600181018255905b50919050565b5f610462611661611143565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f805f6116ab878787876118a1565b915091506116b881611989565b5095945050505050565b5f7f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6116ec611b3b565b6116f4611b93565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b5f54610100900460ff166117cb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105ca565b60676117d78382612134565b5060686117e48282612134565b50505f606581905560665550565b5f54610100900460ff16611888576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016105ca565b60366118948382612134565b5060376114a48282612134565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08311156118d657505f90506003611980565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015611927573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661197a575f60019250925050611980565b91505f90505b94509492505050565b5f81600481111561199c5761199c61224c565b036119a45750565b60018160048111156119b8576119b861224c565b03611a1f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f45434453413a20696e76616c6964207369676e6174757265000000000000000060448201526064016105ca565b6002816004811115611a3357611a3361224c565b03611a9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e6774680060448201526064016105ca565b6003816004811115611aae57611aae61224c565b0361157e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f756500000000000000000000000000000000000000000000000000000000000060648201526084016105ca565b5f80611b456112ca565b805190915015611b5c578051602090910120919050565b6065548015611b6b5792915050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4709250505090565b5f80611b9d6112d9565b805190915015611bb4578051602090910120919050565b6066548015611b6b5792915050565b5f81518084525f5b81811015611be757602081850181015186830182015201611bcb565b505f6020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081525f611c366020830184611bc3565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611c60575f80fd5b919050565b5f8060408385031215611c76575f80fd5b611c7f83611c3d565b946020939093013593505050565b5f805f60608486031215611c9f575f80fd5b611ca884611c3d565b9250611cb660208501611c3d565b9150604084013590509250925092565b5f805f8060608587031215611cd9575f80fd5b611ce285611c3d565b935060208501359250604085013567ffffffffffffffff80821115611d05575f80fd5b818701915087601f830112611d18575f80fd5b813581811115611d26575f80fd5b886020828501011115611d37575f80fd5b95989497505060200194505050565b5f60208284031215611d56575f80fd5b611c3682611c3d565b7fff00000000000000000000000000000000000000000000000000000000000000881681525f602060e06020840152611d9b60e084018a611bc3565b8381036040850152611dad818a611bc3565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c0860152855180825260208088019350909101905f5b81811015611e0d57835183529284019291840191600101611df1565b50909c9b505050505050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f82601f830112611e5b575f80fd5b813567ffffffffffffffff80821115611e7657611e76611e1f565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611ebc57611ebc611e1f565b81604052838152866020858801011115611ed4575f80fd5b836020870160208301375f602085830101528094505050505092915050565b803560ff81168114611c60575f80fd5b5f805f805f60a08688031215611f17575f80fd5b853567ffffffffffffffff80821115611f2e575f80fd5b611f3a89838a01611e4c565b96506020880135915080821115611f4f575f80fd5b50611f5c88828901611e4c565b945050611f6b60408701611ef3565b9250611f7960608701611c3d565b9150611f8760808701611c3d565b90509295509295909350565b5f805f805f805f60e0888a031215611fa9575f80fd5b611fb288611c3d565b9650611fc060208901611c3d565b95506040880135945060608801359350611fdc60808901611ef3565b925060a0880135915060c0880135905092959891949750929550565b5f8060408385031215612009575f80fd5b61201283611c3d565b915061202060208401611c3d565b90509250929050565b600181811c9082168061203d57607f821691505b60208210810361164f577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b80820180821115610462577f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201525f6120e06060830184611bc3565b95945050505050565b601f8211156114a457805f5260205f20601f840160051c8101602085101561210e5750805b601f840160051c820191505b8181101561212d575f815560010161211a565b5050505050565b815167ffffffffffffffff81111561214e5761214e611e1f565b6121628161215c8454612029565b846120e9565b602080601f8311600181146121b4575f841561217e5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555610a82565b5f858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015612200578886015182559484019460019091019084016121e1565b508582101561223c57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffdfea164736f6c6343000818000a"

func init() {
	if err := json.Unmarshal([]byte(MorphStandardERC20StorageLayoutJSON), MorphStandardERC20StorageLayout); err != nil {
		panic(err)
	}

	layouts["MorphStandardERC20"] = MorphStandardERC20StorageLayout
	deployedBytecodes["MorphStandardERC20"] = MorphStandardERC20DeployedBin
}
