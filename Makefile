.PHONY: compile
compile:
	./scripts/solc_compile.sh \
		contracts/ERC20Factory.sol \
		bin/erc20_factory
	./scripts/solc_compile.sh \
		contracts/RuntimeFactory.sol \
			bin/runtime_factory
	./scripts/solc_compile.sh \
		contracts/EthErc20PaymentChannel.sol \
		bin/paymentchannel

.PHONY: abigen
abigen:
	./scripts/abigen.sh \
		bin/paymentchannel/PaymentChannels.bin \
		bin/paymentchannel/PaymentChannels.abi \
		paymentchannel \
		bindings/paymentchannel/bindings.go
	./scripts/abigen.sh \
		bin/runtime_factory/RuntimeFactory.bin \
		bin/runtime_factory/RuntimeFactory.abi \
		runtimefactory \
		bindings/runtimefactory/bindings.go
	./scripts/abigen.sh \
		bin/erc20_factory/ERC20Factory.bin \
		bin/erc20_factory/ERC20Factory.abi \
		erc20factory \
		bindings/erc20factory/bindings.go
	./scripts/solc_compile.sh \
		contracts/utils/time/BokkyPooBahsDateTimeContract.sol \
		bin/datetime \
		compilers/solc-v0.6.0
	./scripts/abigen.sh \
		bin/datetime/BokkyPooBahsDateTimeContract.bin \
		bin/datetime/BokkyPooBahsDateTimeContract.abi \
		datetime \
		bindings/datetime/bindings.go \
		--alias="getDaysInMonth=getDaysInMonth2,_isLeapYear=IsLeapYear2"

.PHONY: interfaces
interfaces:
	./scripts/abi2solidity.sh \
		bin/datetime/BokkyPooBahsDateTimeContract.abi \
		contracts/interfaces/DateTimeInterface.sol \
		DateTimeInterface \
		"pragma solidity 0.7.3;"

.PHONY: all
all: compile abigen interfaces