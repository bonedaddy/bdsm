// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package datetime

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DatetimeABI is the input ABI used to generate the binding from.
const DatetimeABI = "[{\"inputs\":[],\"name\":\"DOW_FRI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOW_MON\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOW_SAT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOW_SUN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOW_THU\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOW_TUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOW_WED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OFFSET19700101\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_HOUR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_MINUTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"_daysFromDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_days\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_days\",\"type\":\"uint256\"}],\"name\":\"_daysToDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"}],\"name\":\"_getDaysInMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"daysInMonth\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"}],\"name\":\"_isLeapYear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"leapYear\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_now\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_nowDateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hour\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"second\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_days\",\"type\":\"uint256\"}],\"name\":\"addDays\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"addHours\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minutes\",\"type\":\"uint256\"}],\"name\":\"addMinutes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_months\",\"type\":\"uint256\"}],\"name\":\"addMonths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"name\":\"addSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_years\",\"type\":\"uint256\"}],\"name\":\"addYears\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"diffDays\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_days\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"diffHours\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"diffMinutes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_minutes\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"diffMonths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_months\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"diffSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"diffYears\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_years\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getDayOfWeek\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dayOfWeek\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getDaysInMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"daysInMonth\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getHour\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"hour\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getMinute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minute\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getMonth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getSecond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"second\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"getYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"isLeapYear\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"leapYear\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"isValidDate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hour\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"second\",\"type\":\"uint256\"}],\"name\":\"isValidDateTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"isWeekDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"weekDay\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"isWeekEnd\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"weekEnd\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_days\",\"type\":\"uint256\"}],\"name\":\"subDays\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_hours\",\"type\":\"uint256\"}],\"name\":\"subHours\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minutes\",\"type\":\"uint256\"}],\"name\":\"subMinutes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_months\",\"type\":\"uint256\"}],\"name\":\"subMonths\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seconds\",\"type\":\"uint256\"}],\"name\":\"subSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_years\",\"type\":\"uint256\"}],\"name\":\"subYears\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"timestampFromDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hour\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"second\",\"type\":\"uint256\"}],\"name\":\"timestampFromDateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"timestampToDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"timestampToDateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"year\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"month\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hour\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minute\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"second\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DatetimeBin is the compiled bytecode used for deploying new contracts.
var DatetimeBin = "0x608060405234801561001057600080fd5b506112c7806100206000396000f3fe608060405234801561001057600080fd5b506004361061030b5760003560e01c80638aa001fc1161019d578063c7b6fd6a116100e9578063de5101af116100a2578063f615ed541161007c578063f615ed54146108c1578063f9fd5250146108e4578063fa93f883146108ec578063ff2258cb146109095761030b565b8063de5101af1461087f578063e95564301461089c578063ea1c1690146108a45761030b565b8063c7b6fd6a14610804578063c7edf88c14610827578063c9d346221461082f578063cfbb9f3714610852578063d2b507431461085a578063d6582d0d146108625761030b565b80639e524caa11610156578063ad203bd411610130578063ad203bd41461079f578063b05eb08d146107c2578063b3bb8cd4146107df578063b8d16dbc146107e75761030b565b80639e524caa14610757578063a324ad241461077a578063a3f144ae146107975761030b565b80638aa001fc1461067c5780638bbf51b7146106995780638d4a2d39146106a157806390059aed146106c45780639220d426146106ff57806392d663131461073a5761030b565b80634355644d1161025c5780635e05bd6d116102155780637217523c116101ef5780637217523c1461060b57806374f0314f1461062e5780637be341091461063657806389a3a00d146106595761030b565b80635e05bd6d1461059057806362fb9697146105cb57806365c72840146105ee5761030b565b80634355644d146104df5780634371c46514610502578063442b8c791461051f578063444fda82146105425780634b321502146105655780634df86126146105885761030b565b80631f4f77b2116102c95780632af123b8116102a35780632af123b8146104415780633293d007146104645780633e239e1a1461049f5780633f9e0eb7146104bc5761030b565b80631f4f77b2146103f357806322f8a2b81461041c57806329441674146104395761030b565b80625015531461031057806302e98e0d1461034557806310848ddf14610368578063126702a01461038557806314b2d6dc1461038d5780631e0582e9146103ca575b600080fd5b6103336004803603604081101561032657600080fd5b508035906020013561092c565b60408051918252519081900360200190f35b6103336004803603604081101561035b57600080fd5b508035906020013561093f565b6103336004803603602081101561037e57600080fd5b503561094b565b61033361095c565b6103b6600480360360608110156103a357600080fd5b5080359060208101359060400135610961565b604080519115158252519081900360200190f35b610333600480360360608110156103e057600080fd5b5080359060208101359060400135610976565b6103336004803603606081101561040957600080fd5b5080359060208101359060400135610983565b6103336004803603602081101561043257600080fd5b5035610990565b61033361099b565b6103336004803603604081101561045757600080fd5b50803590602001356109a0565b6103b6600480360360c081101561047a57600080fd5b5080359060208101359060408101359060608101359060808101359060a001356109ac565b610333600480360360208110156104b557600080fd5b50356109c7565b610333600480360360408110156104d257600080fd5b50803590602001356109d2565b610333600480360360408110156104f557600080fd5b50803590602001356109de565b6103b66004803603602081101561051857600080fd5b50356109ea565b6103336004803603604081101561053557600080fd5b50803590602001356109f5565b6103336004803603604081101561055857600080fd5b5080359060200135610a01565b6103336004803603604081101561057b57600080fd5b5080359060200135610a0d565b610333610a19565b610333600480360360c08110156105a657600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610a1f565b610333600480360360408110156105e157600080fd5b5080359060200135610a39565b6103336004803603602081101561060457600080fd5b5035610a45565b6103336004803603604081101561062157600080fd5b5080359060200135610a50565b610333610a5c565b6103336004803603604081101561064c57600080fd5b5080359060200135610a63565b6103336004803603604081101561066f57600080fd5b5080359060200135610a6f565b6103336004803603602081101561069257600080fd5b5035610a7b565b610333610a86565b610333600480360360408110156106b757600080fd5b5080359060200135610a8b565b6106e1600480360360208110156106da57600080fd5b5035610a97565b60408051938452602084019290925282820152519081900360600190f35b610707610ab2565b604080519687526020870195909552858501939093526060850191909152608084015260a0830152519081900360c00190f35b6103336004803603602081101561075057600080fd5b5035610ad7565b6103336004803603604081101561076d57600080fd5b5080359060200135610ae2565b6103336004803603602081101561079057600080fd5b5035610aee565b610333610af9565b610333600480360360408110156107b557600080fd5b5080359060200135610b00565b6103b6600480360360208110156107d857600080fd5b5035610b0c565b610333610b17565b6103b6600480360360208110156107fd57600080fd5b5035610b1b565b6103336004803603604081101561081a57600080fd5b5080359060200135610b26565b610333610b32565b6103336004803603604081101561084557600080fd5b5080359060200135610b37565b610333610b43565b610333610b48565b6103b66004803603602081101561087857600080fd5b5035610b4d565b6106e16004803603602081101561089557600080fd5b5035610b58565b610333610b73565b610707600480360360208110156108ba57600080fd5b5035610b78565b610333600480360360408110156108d757600080fd5b5080359060200135610b9e565b610333610baa565b6103336004803603602081101561090257600080fd5b5035610baf565b6103336004803603604081101561091f57600080fd5b5080359060200135610bba565b60006109388383610bc6565b9392505050565b60006109388383610bda565b600061095682610bf7565b92915050565b600281565b600061096e848484610c19565b949350505050565b600061096e848484610c6f565b600061096e848484610ceb565b600061095682610d05565b600781565b60006109388383610d18565b60006109bc878787878787610d32565b979650505050505050565b600061095682610d72565b60006109388383610d80565b60006109388383610e06565b600061095682610e80565b60006109388383610e95565b60006109388383610ebc565b60006109388383610f18565b610e1081565b60006109bc878787878787610f2c565b9695505050505050565b60006109388383610f56565b600061095682610f6a565b60006109388383610f79565b6201518081565b60006109388383610f8e565b60006109388383610fe0565b600061095682610ff3565b600381565b60006109388383610ffa565b6000806000610aa58461100a565b9250925092509193909250565b600080600080600080610ac4426110a0565b949b939a50919850965094509092509050565b6000610956826110df565b600061093883836110f7565b60006109568261110a565b62253d8c81565b60006109388383611119565b600061095682611199565b4290565b6000610956826111be565b600061093883836111db565b600681565b600061093883836111f0565b600481565b600581565b60006109568261120b565b6000806000610b6684611220565b9196909550909350915050565b603c81565b600080600080600080610b8a876110a0565b949c939b5091995097509550909350915050565b60006109388383611231565b600181565b600061095682611241565b60006109388383611250565b610e10810282038281111561095657600080fd5b600081831115610be957600080fd5b603c8383035b049392505050565b60008080610c0a62015180855b0461100a565b509150915061096e8282610d80565b60006107b28410158015610c2d5750600083115b8015610c3a5750600c8311155b15610938576000610c4b8585610d80565b9050600083118015610c5d5750808311155b15610c6757600191505b509392505050565b60006107b2841015610c8057600080fd5b838383600062253d8c600460036064611324600c600d19890105890101050205600c80600d19870105600c02600287030361016f0281610cbc57fe5b0560046105b5600c600d1989010589016112c0010205617d4b8603010103039050809450505050509392505050565b600062015180610cfc858585610c6f565b02949350505050565b6007620151809091046003010660010190565b600081831115610d2757600080fd5b610e10838303610bef565b6000610d3f878787610c19565b15610a2f57601884108015610d545750603c83105b8015610d605750603c82105b15610a2f575060019695505050505050565b610e10620151809091060490565b60008160011480610d915750816003145b80610d9c5750816005145b80610da75750816007145b80610db25750816008145b80610dbd575081600a145b80610dc8575081600c145b15610dd55750601f610956565b81600214610de55750601e610956565b610dee83611199565b610df957601c610dfc565b601d5b60ff169392505050565b6000808080610e186201518087610c04565b600c918801600019810183810494909401965094509250900660010191506000610e428484610d80565b905080821115610e50578091505b62015180870662015180610e65868686610c6f565b0201945086851015610e7657600080fd5b5050505092915050565b60006006610e8d83610d05565b101592915050565b6000808080610ea76201518087610c04565b9187019450925090506000610e428484610d80565b6000808080610ece6201518087610c04565b918790039450925090506000610ee48484610d80565b905080821115610ef2578091505b62015180870662015180610f07868686610c6f565b0201945086851115610e7657600080fd5b610e10810282018281101561095657600080fd5b600081603c8402610e10860262015180610f478b8b8b610c6f565b02010101979650505050505050565b600081831115610f6557600080fd5b500390565b600061096e6201518083610c04565b62015180810282018281101561095657600080fd5b600081831115610f9d57600080fd5b600080610fad6201518086610c04565b5091509150600080610fc4620151808781610c0457fe5b50600c9586029590910201939093039190910395945050505050565b603c810282018281101561095657600080fd5b603c900690565b8181018281101561095657600080fd5b60008080836226496581018262023ab1600483020590506004600362023ab18302010590910390600062164b09610fa0600185010205905060046105b58202058303601f019250600061098f846050028161106157fe5b0590506000605061098f83020585039050600b820560301994909401606402929092018301996002600c90940290910392909201975095509350505050565b600080808080806110b46201518088610c04565b91999098919750610e10620151809092068281049750603c9290068281049650919091069350915050565b60006110ee6201518083610c04565b50909392505050565b603c810282038281111561095657600080fd5b6000610c676201518083610c04565b600080808061112b6201518087610c04565b91945092509050600c8084028301869003600019019081049350600c81066001019250600061115a8585610d80565b905080831115611168578092505b6201518088066201518061117d878787610c6f565b020195508786111561118e57600080fd5b505050505092915050565b6000600482061580156111ae57506064820615155b8061095657505061019090061590565b6000806111ce6201518084610c04565b5050905061093881611199565b62015180810282038281111561095657600080fd5b6000818311156111ff57600080fd5b62015180838303610bef565b6000600561121883610d05565b111592915050565b60008080610b666201518085610c04565b8082038281111561095657600080fd5b6000610e108206603c81610bef565b60008183111561125f57600080fd5b600061126e6201518085610c04565b505090506000611283620151808581610c0457fe5b50509190910394935050505056fea26469706673582212203110ab73d6389074523983b2b07d58f30436f3097b21fc5b22b8b297a319a8e964736f6c63430006000033"

// DeployDatetime deploys a new Ethereum contract, binding an instance of Datetime to it.
func DeployDatetime(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Datetime, error) {
	parsed, err := abi.JSON(strings.NewReader(DatetimeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DatetimeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Datetime{DatetimeCaller: DatetimeCaller{contract: contract}, DatetimeTransactor: DatetimeTransactor{contract: contract}, DatetimeFilterer: DatetimeFilterer{contract: contract}}, nil
}

// Datetime is an auto generated Go binding around an Ethereum contract.
type Datetime struct {
	DatetimeCaller     // Read-only binding to the contract
	DatetimeTransactor // Write-only binding to the contract
	DatetimeFilterer   // Log filterer for contract events
}

// DatetimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DatetimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatetimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DatetimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatetimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DatetimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DatetimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DatetimeSession struct {
	Contract     *Datetime         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DatetimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DatetimeCallerSession struct {
	Contract *DatetimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DatetimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DatetimeTransactorSession struct {
	Contract     *DatetimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DatetimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DatetimeRaw struct {
	Contract *Datetime // Generic contract binding to access the raw methods on
}

// DatetimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DatetimeCallerRaw struct {
	Contract *DatetimeCaller // Generic read-only contract binding to access the raw methods on
}

// DatetimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DatetimeTransactorRaw struct {
	Contract *DatetimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDatetime creates a new instance of Datetime, bound to a specific deployed contract.
func NewDatetime(address common.Address, backend bind.ContractBackend) (*Datetime, error) {
	contract, err := bindDatetime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Datetime{DatetimeCaller: DatetimeCaller{contract: contract}, DatetimeTransactor: DatetimeTransactor{contract: contract}, DatetimeFilterer: DatetimeFilterer{contract: contract}}, nil
}

// NewDatetimeCaller creates a new read-only instance of Datetime, bound to a specific deployed contract.
func NewDatetimeCaller(address common.Address, caller bind.ContractCaller) (*DatetimeCaller, error) {
	contract, err := bindDatetime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DatetimeCaller{contract: contract}, nil
}

// NewDatetimeTransactor creates a new write-only instance of Datetime, bound to a specific deployed contract.
func NewDatetimeTransactor(address common.Address, transactor bind.ContractTransactor) (*DatetimeTransactor, error) {
	contract, err := bindDatetime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DatetimeTransactor{contract: contract}, nil
}

// NewDatetimeFilterer creates a new log filterer instance of Datetime, bound to a specific deployed contract.
func NewDatetimeFilterer(address common.Address, filterer bind.ContractFilterer) (*DatetimeFilterer, error) {
	contract, err := bindDatetime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DatetimeFilterer{contract: contract}, nil
}

// bindDatetime binds a generic wrapper to an already deployed contract.
func bindDatetime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DatetimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datetime *DatetimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datetime.Contract.DatetimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datetime *DatetimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datetime.Contract.DatetimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datetime *DatetimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datetime.Contract.DatetimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Datetime *DatetimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Datetime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Datetime *DatetimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Datetime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Datetime *DatetimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Datetime.Contract.contract.Transact(opts, method, params...)
}

// DOWFRI is a free data retrieval call binding the contract method 0xd2b50743.
//
// Solidity: function DOW_FRI() view returns(uint256)
func (_Datetime *DatetimeCaller) DOWFRI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "DOW_FRI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOWFRI is a free data retrieval call binding the contract method 0xd2b50743.
//
// Solidity: function DOW_FRI() view returns(uint256)
func (_Datetime *DatetimeSession) DOWFRI() (*big.Int, error) {
	return _Datetime.Contract.DOWFRI(&_Datetime.CallOpts)
}

// DOWFRI is a free data retrieval call binding the contract method 0xd2b50743.
//
// Solidity: function DOW_FRI() view returns(uint256)
func (_Datetime *DatetimeCallerSession) DOWFRI() (*big.Int, error) {
	return _Datetime.Contract.DOWFRI(&_Datetime.CallOpts)
}

// DOWMON is a free data retrieval call binding the contract method 0xf9fd5250.
//
// Solidity: function DOW_MON() view returns(uint256)
func (_Datetime *DatetimeCaller) DOWMON(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "DOW_MON")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOWMON is a free data retrieval call binding the contract method 0xf9fd5250.
//
// Solidity: function DOW_MON() view returns(uint256)
func (_Datetime *DatetimeSession) DOWMON() (*big.Int, error) {
	return _Datetime.Contract.DOWMON(&_Datetime.CallOpts)
}

// DOWMON is a free data retrieval call binding the contract method 0xf9fd5250.
//
// Solidity: function DOW_MON() view returns(uint256)
func (_Datetime *DatetimeCallerSession) DOWMON() (*big.Int, error) {
	return _Datetime.Contract.DOWMON(&_Datetime.CallOpts)
}

// DOWSAT is a free data retrieval call binding the contract method 0xc7edf88c.
//
// Solidity: function DOW_SAT() view returns(uint256)
func (_Datetime *DatetimeCaller) DOWSAT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "DOW_SAT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOWSAT is a free data retrieval call binding the contract method 0xc7edf88c.
//
// Solidity: function DOW_SAT() view returns(uint256)
func (_Datetime *DatetimeSession) DOWSAT() (*big.Int, error) {
	return _Datetime.Contract.DOWSAT(&_Datetime.CallOpts)
}

// DOWSAT is a free data retrieval call binding the contract method 0xc7edf88c.
//
// Solidity: function DOW_SAT() view returns(uint256)
func (_Datetime *DatetimeCallerSession) DOWSAT() (*big.Int, error) {
	return _Datetime.Contract.DOWSAT(&_Datetime.CallOpts)
}

// DOWSUN is a free data retrieval call binding the contract method 0x29441674.
//
// Solidity: function DOW_SUN() view returns(uint256)
func (_Datetime *DatetimeCaller) DOWSUN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "DOW_SUN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOWSUN is a free data retrieval call binding the contract method 0x29441674.
//
// Solidity: function DOW_SUN() view returns(uint256)
func (_Datetime *DatetimeSession) DOWSUN() (*big.Int, error) {
	return _Datetime.Contract.DOWSUN(&_Datetime.CallOpts)
}

// DOWSUN is a free data retrieval call binding the contract method 0x29441674.
//
// Solidity: function DOW_SUN() view returns(uint256)
func (_Datetime *DatetimeCallerSession) DOWSUN() (*big.Int, error) {
	return _Datetime.Contract.DOWSUN(&_Datetime.CallOpts)
}

// DOWTHU is a free data retrieval call binding the contract method 0xcfbb9f37.
//
// Solidity: function DOW_THU() view returns(uint256)
func (_Datetime *DatetimeCaller) DOWTHU(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "DOW_THU")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOWTHU is a free data retrieval call binding the contract method 0xcfbb9f37.
//
// Solidity: function DOW_THU() view returns(uint256)
func (_Datetime *DatetimeSession) DOWTHU() (*big.Int, error) {
	return _Datetime.Contract.DOWTHU(&_Datetime.CallOpts)
}

// DOWTHU is a free data retrieval call binding the contract method 0xcfbb9f37.
//
// Solidity: function DOW_THU() view returns(uint256)
func (_Datetime *DatetimeCallerSession) DOWTHU() (*big.Int, error) {
	return _Datetime.Contract.DOWTHU(&_Datetime.CallOpts)
}

// DOWTUE is a free data retrieval call binding the contract method 0x126702a0.
//
// Solidity: function DOW_TUE() view returns(uint256)
func (_Datetime *DatetimeCaller) DOWTUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "DOW_TUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOWTUE is a free data retrieval call binding the contract method 0x126702a0.
//
// Solidity: function DOW_TUE() view returns(uint256)
func (_Datetime *DatetimeSession) DOWTUE() (*big.Int, error) {
	return _Datetime.Contract.DOWTUE(&_Datetime.CallOpts)
}

// DOWTUE is a free data retrieval call binding the contract method 0x126702a0.
//
// Solidity: function DOW_TUE() view returns(uint256)
func (_Datetime *DatetimeCallerSession) DOWTUE() (*big.Int, error) {
	return _Datetime.Contract.DOWTUE(&_Datetime.CallOpts)
}

// DOWWED is a free data retrieval call binding the contract method 0x8bbf51b7.
//
// Solidity: function DOW_WED() view returns(uint256)
func (_Datetime *DatetimeCaller) DOWWED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "DOW_WED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DOWWED is a free data retrieval call binding the contract method 0x8bbf51b7.
//
// Solidity: function DOW_WED() view returns(uint256)
func (_Datetime *DatetimeSession) DOWWED() (*big.Int, error) {
	return _Datetime.Contract.DOWWED(&_Datetime.CallOpts)
}

// DOWWED is a free data retrieval call binding the contract method 0x8bbf51b7.
//
// Solidity: function DOW_WED() view returns(uint256)
func (_Datetime *DatetimeCallerSession) DOWWED() (*big.Int, error) {
	return _Datetime.Contract.DOWWED(&_Datetime.CallOpts)
}

// OFFSET19700101 is a free data retrieval call binding the contract method 0xa3f144ae.
//
// Solidity: function OFFSET19700101() view returns(int256)
func (_Datetime *DatetimeCaller) OFFSET19700101(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "OFFSET19700101")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OFFSET19700101 is a free data retrieval call binding the contract method 0xa3f144ae.
//
// Solidity: function OFFSET19700101() view returns(int256)
func (_Datetime *DatetimeSession) OFFSET19700101() (*big.Int, error) {
	return _Datetime.Contract.OFFSET19700101(&_Datetime.CallOpts)
}

// OFFSET19700101 is a free data retrieval call binding the contract method 0xa3f144ae.
//
// Solidity: function OFFSET19700101() view returns(int256)
func (_Datetime *DatetimeCallerSession) OFFSET19700101() (*big.Int, error) {
	return _Datetime.Contract.OFFSET19700101(&_Datetime.CallOpts)
}

// SECONDSPERDAY is a free data retrieval call binding the contract method 0x74f0314f.
//
// Solidity: function SECONDS_PER_DAY() view returns(uint256)
func (_Datetime *DatetimeCaller) SECONDSPERDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "SECONDS_PER_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERDAY is a free data retrieval call binding the contract method 0x74f0314f.
//
// Solidity: function SECONDS_PER_DAY() view returns(uint256)
func (_Datetime *DatetimeSession) SECONDSPERDAY() (*big.Int, error) {
	return _Datetime.Contract.SECONDSPERDAY(&_Datetime.CallOpts)
}

// SECONDSPERDAY is a free data retrieval call binding the contract method 0x74f0314f.
//
// Solidity: function SECONDS_PER_DAY() view returns(uint256)
func (_Datetime *DatetimeCallerSession) SECONDSPERDAY() (*big.Int, error) {
	return _Datetime.Contract.SECONDSPERDAY(&_Datetime.CallOpts)
}

// SECONDSPERHOUR is a free data retrieval call binding the contract method 0x4df86126.
//
// Solidity: function SECONDS_PER_HOUR() view returns(uint256)
func (_Datetime *DatetimeCaller) SECONDSPERHOUR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "SECONDS_PER_HOUR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERHOUR is a free data retrieval call binding the contract method 0x4df86126.
//
// Solidity: function SECONDS_PER_HOUR() view returns(uint256)
func (_Datetime *DatetimeSession) SECONDSPERHOUR() (*big.Int, error) {
	return _Datetime.Contract.SECONDSPERHOUR(&_Datetime.CallOpts)
}

// SECONDSPERHOUR is a free data retrieval call binding the contract method 0x4df86126.
//
// Solidity: function SECONDS_PER_HOUR() view returns(uint256)
func (_Datetime *DatetimeCallerSession) SECONDSPERHOUR() (*big.Int, error) {
	return _Datetime.Contract.SECONDSPERHOUR(&_Datetime.CallOpts)
}

// SECONDSPERMINUTE is a free data retrieval call binding the contract method 0xe9556430.
//
// Solidity: function SECONDS_PER_MINUTE() view returns(uint256)
func (_Datetime *DatetimeCaller) SECONDSPERMINUTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "SECONDS_PER_MINUTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERMINUTE is a free data retrieval call binding the contract method 0xe9556430.
//
// Solidity: function SECONDS_PER_MINUTE() view returns(uint256)
func (_Datetime *DatetimeSession) SECONDSPERMINUTE() (*big.Int, error) {
	return _Datetime.Contract.SECONDSPERMINUTE(&_Datetime.CallOpts)
}

// SECONDSPERMINUTE is a free data retrieval call binding the contract method 0xe9556430.
//
// Solidity: function SECONDS_PER_MINUTE() view returns(uint256)
func (_Datetime *DatetimeCallerSession) SECONDSPERMINUTE() (*big.Int, error) {
	return _Datetime.Contract.SECONDSPERMINUTE(&_Datetime.CallOpts)
}

// DaysFromDate is a free data retrieval call binding the contract method 0x1e0582e9.
//
// Solidity: function _daysFromDate(uint256 year, uint256 month, uint256 day) pure returns(uint256 _days)
func (_Datetime *DatetimeCaller) DaysFromDate(opts *bind.CallOpts, year *big.Int, month *big.Int, day *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "_daysFromDate", year, month, day)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DaysFromDate is a free data retrieval call binding the contract method 0x1e0582e9.
//
// Solidity: function _daysFromDate(uint256 year, uint256 month, uint256 day) pure returns(uint256 _days)
func (_Datetime *DatetimeSession) DaysFromDate(year *big.Int, month *big.Int, day *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DaysFromDate(&_Datetime.CallOpts, year, month, day)
}

// DaysFromDate is a free data retrieval call binding the contract method 0x1e0582e9.
//
// Solidity: function _daysFromDate(uint256 year, uint256 month, uint256 day) pure returns(uint256 _days)
func (_Datetime *DatetimeCallerSession) DaysFromDate(year *big.Int, month *big.Int, day *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DaysFromDate(&_Datetime.CallOpts, year, month, day)
}

// DaysToDate is a free data retrieval call binding the contract method 0x90059aed.
//
// Solidity: function _daysToDate(uint256 _days) pure returns(uint256 year, uint256 month, uint256 day)
func (_Datetime *DatetimeCaller) DaysToDate(opts *bind.CallOpts, _days *big.Int) (struct {
	Year  *big.Int
	Month *big.Int
	Day   *big.Int
}, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "_daysToDate", _days)

	outstruct := new(struct {
		Year  *big.Int
		Month *big.Int
		Day   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Year = out[0].(*big.Int)
	outstruct.Month = out[1].(*big.Int)
	outstruct.Day = out[2].(*big.Int)

	return *outstruct, err

}

// DaysToDate is a free data retrieval call binding the contract method 0x90059aed.
//
// Solidity: function _daysToDate(uint256 _days) pure returns(uint256 year, uint256 month, uint256 day)
func (_Datetime *DatetimeSession) DaysToDate(_days *big.Int) (struct {
	Year  *big.Int
	Month *big.Int
	Day   *big.Int
}, error) {
	return _Datetime.Contract.DaysToDate(&_Datetime.CallOpts, _days)
}

// DaysToDate is a free data retrieval call binding the contract method 0x90059aed.
//
// Solidity: function _daysToDate(uint256 _days) pure returns(uint256 year, uint256 month, uint256 day)
func (_Datetime *DatetimeCallerSession) DaysToDate(_days *big.Int) (struct {
	Year  *big.Int
	Month *big.Int
	Day   *big.Int
}, error) {
	return _Datetime.Contract.DaysToDate(&_Datetime.CallOpts, _days)
}

// GetDaysInMonth is a free data retrieval call binding the contract method 0x3f9e0eb7.
//
// Solidity: function _getDaysInMonth(uint256 year, uint256 month) pure returns(uint256 daysInMonth)
func (_Datetime *DatetimeCaller) GetDaysInMonth(opts *bind.CallOpts, year *big.Int, month *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "_getDaysInMonth", year, month)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDaysInMonth is a free data retrieval call binding the contract method 0x3f9e0eb7.
//
// Solidity: function _getDaysInMonth(uint256 year, uint256 month) pure returns(uint256 daysInMonth)
func (_Datetime *DatetimeSession) GetDaysInMonth(year *big.Int, month *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDaysInMonth(&_Datetime.CallOpts, year, month)
}

// GetDaysInMonth is a free data retrieval call binding the contract method 0x3f9e0eb7.
//
// Solidity: function _getDaysInMonth(uint256 year, uint256 month) pure returns(uint256 daysInMonth)
func (_Datetime *DatetimeCallerSession) GetDaysInMonth(year *big.Int, month *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDaysInMonth(&_Datetime.CallOpts, year, month)
}

// IsLeapYear2 is a free data retrieval call binding the contract method 0xb05eb08d.
//
// Solidity: function _isLeapYear(uint256 year) pure returns(bool leapYear)
func (_Datetime *DatetimeCaller) IsLeapYear2(opts *bind.CallOpts, year *big.Int) (bool, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "_isLeapYear", year)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeapYear2 is a free data retrieval call binding the contract method 0xb05eb08d.
//
// Solidity: function _isLeapYear(uint256 year) pure returns(bool leapYear)
func (_Datetime *DatetimeSession) IsLeapYear2(year *big.Int) (bool, error) {
	return _Datetime.Contract.IsLeapYear2(&_Datetime.CallOpts, year)
}

// IsLeapYear2 is a free data retrieval call binding the contract method 0xb05eb08d.
//
// Solidity: function _isLeapYear(uint256 year) pure returns(bool leapYear)
func (_Datetime *DatetimeCallerSession) IsLeapYear2(year *big.Int) (bool, error) {
	return _Datetime.Contract.IsLeapYear2(&_Datetime.CallOpts, year)
}

// Now is a free data retrieval call binding the contract method 0xb3bb8cd4.
//
// Solidity: function _now() view returns(uint256 timestamp)
func (_Datetime *DatetimeCaller) Now(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "_now")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Now is a free data retrieval call binding the contract method 0xb3bb8cd4.
//
// Solidity: function _now() view returns(uint256 timestamp)
func (_Datetime *DatetimeSession) Now() (*big.Int, error) {
	return _Datetime.Contract.Now(&_Datetime.CallOpts)
}

// Now is a free data retrieval call binding the contract method 0xb3bb8cd4.
//
// Solidity: function _now() view returns(uint256 timestamp)
func (_Datetime *DatetimeCallerSession) Now() (*big.Int, error) {
	return _Datetime.Contract.Now(&_Datetime.CallOpts)
}

// NowDateTime is a free data retrieval call binding the contract method 0x9220d426.
//
// Solidity: function _nowDateTime() view returns(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second)
func (_Datetime *DatetimeCaller) NowDateTime(opts *bind.CallOpts) (struct {
	Year   *big.Int
	Month  *big.Int
	Day    *big.Int
	Hour   *big.Int
	Minute *big.Int
	Second *big.Int
}, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "_nowDateTime")

	outstruct := new(struct {
		Year   *big.Int
		Month  *big.Int
		Day    *big.Int
		Hour   *big.Int
		Minute *big.Int
		Second *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Year = out[0].(*big.Int)
	outstruct.Month = out[1].(*big.Int)
	outstruct.Day = out[2].(*big.Int)
	outstruct.Hour = out[3].(*big.Int)
	outstruct.Minute = out[4].(*big.Int)
	outstruct.Second = out[5].(*big.Int)

	return *outstruct, err

}

// NowDateTime is a free data retrieval call binding the contract method 0x9220d426.
//
// Solidity: function _nowDateTime() view returns(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second)
func (_Datetime *DatetimeSession) NowDateTime() (struct {
	Year   *big.Int
	Month  *big.Int
	Day    *big.Int
	Hour   *big.Int
	Minute *big.Int
	Second *big.Int
}, error) {
	return _Datetime.Contract.NowDateTime(&_Datetime.CallOpts)
}

// NowDateTime is a free data retrieval call binding the contract method 0x9220d426.
//
// Solidity: function _nowDateTime() view returns(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second)
func (_Datetime *DatetimeCallerSession) NowDateTime() (struct {
	Year   *big.Int
	Month  *big.Int
	Day    *big.Int
	Hour   *big.Int
	Minute *big.Int
	Second *big.Int
}, error) {
	return _Datetime.Contract.NowDateTime(&_Datetime.CallOpts)
}

// AddDays is a free data retrieval call binding the contract method 0x7217523c.
//
// Solidity: function addDays(uint256 timestamp, uint256 _days) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) AddDays(opts *bind.CallOpts, timestamp *big.Int, _days *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "addDays", timestamp, _days)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddDays is a free data retrieval call binding the contract method 0x7217523c.
//
// Solidity: function addDays(uint256 timestamp, uint256 _days) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) AddDays(timestamp *big.Int, _days *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddDays(&_Datetime.CallOpts, timestamp, _days)
}

// AddDays is a free data retrieval call binding the contract method 0x7217523c.
//
// Solidity: function addDays(uint256 timestamp, uint256 _days) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) AddDays(timestamp *big.Int, _days *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddDays(&_Datetime.CallOpts, timestamp, _days)
}

// AddHours is a free data retrieval call binding the contract method 0x4b321502.
//
// Solidity: function addHours(uint256 timestamp, uint256 _hours) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) AddHours(opts *bind.CallOpts, timestamp *big.Int, _hours *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "addHours", timestamp, _hours)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddHours is a free data retrieval call binding the contract method 0x4b321502.
//
// Solidity: function addHours(uint256 timestamp, uint256 _hours) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) AddHours(timestamp *big.Int, _hours *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddHours(&_Datetime.CallOpts, timestamp, _hours)
}

// AddHours is a free data retrieval call binding the contract method 0x4b321502.
//
// Solidity: function addHours(uint256 timestamp, uint256 _hours) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) AddHours(timestamp *big.Int, _hours *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddHours(&_Datetime.CallOpts, timestamp, _hours)
}

// AddMinutes is a free data retrieval call binding the contract method 0x89a3a00d.
//
// Solidity: function addMinutes(uint256 timestamp, uint256 _minutes) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) AddMinutes(opts *bind.CallOpts, timestamp *big.Int, _minutes *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "addMinutes", timestamp, _minutes)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddMinutes is a free data retrieval call binding the contract method 0x89a3a00d.
//
// Solidity: function addMinutes(uint256 timestamp, uint256 _minutes) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) AddMinutes(timestamp *big.Int, _minutes *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddMinutes(&_Datetime.CallOpts, timestamp, _minutes)
}

// AddMinutes is a free data retrieval call binding the contract method 0x89a3a00d.
//
// Solidity: function addMinutes(uint256 timestamp, uint256 _minutes) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) AddMinutes(timestamp *big.Int, _minutes *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddMinutes(&_Datetime.CallOpts, timestamp, _minutes)
}

// AddMonths is a free data retrieval call binding the contract method 0x4355644d.
//
// Solidity: function addMonths(uint256 timestamp, uint256 _months) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) AddMonths(opts *bind.CallOpts, timestamp *big.Int, _months *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "addMonths", timestamp, _months)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddMonths is a free data retrieval call binding the contract method 0x4355644d.
//
// Solidity: function addMonths(uint256 timestamp, uint256 _months) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) AddMonths(timestamp *big.Int, _months *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddMonths(&_Datetime.CallOpts, timestamp, _months)
}

// AddMonths is a free data retrieval call binding the contract method 0x4355644d.
//
// Solidity: function addMonths(uint256 timestamp, uint256 _months) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) AddMonths(timestamp *big.Int, _months *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddMonths(&_Datetime.CallOpts, timestamp, _months)
}

// AddSeconds is a free data retrieval call binding the contract method 0x8d4a2d39.
//
// Solidity: function addSeconds(uint256 timestamp, uint256 _seconds) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) AddSeconds(opts *bind.CallOpts, timestamp *big.Int, _seconds *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "addSeconds", timestamp, _seconds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddSeconds is a free data retrieval call binding the contract method 0x8d4a2d39.
//
// Solidity: function addSeconds(uint256 timestamp, uint256 _seconds) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) AddSeconds(timestamp *big.Int, _seconds *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddSeconds(&_Datetime.CallOpts, timestamp, _seconds)
}

// AddSeconds is a free data retrieval call binding the contract method 0x8d4a2d39.
//
// Solidity: function addSeconds(uint256 timestamp, uint256 _seconds) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) AddSeconds(timestamp *big.Int, _seconds *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddSeconds(&_Datetime.CallOpts, timestamp, _seconds)
}

// AddYears is a free data retrieval call binding the contract method 0x442b8c79.
//
// Solidity: function addYears(uint256 timestamp, uint256 _years) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) AddYears(opts *bind.CallOpts, timestamp *big.Int, _years *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "addYears", timestamp, _years)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddYears is a free data retrieval call binding the contract method 0x442b8c79.
//
// Solidity: function addYears(uint256 timestamp, uint256 _years) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) AddYears(timestamp *big.Int, _years *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddYears(&_Datetime.CallOpts, timestamp, _years)
}

// AddYears is a free data retrieval call binding the contract method 0x442b8c79.
//
// Solidity: function addYears(uint256 timestamp, uint256 _years) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) AddYears(timestamp *big.Int, _years *big.Int) (*big.Int, error) {
	return _Datetime.Contract.AddYears(&_Datetime.CallOpts, timestamp, _years)
}

// DiffDays is a free data retrieval call binding the contract method 0xc9d34622.
//
// Solidity: function diffDays(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _days)
func (_Datetime *DatetimeCaller) DiffDays(opts *bind.CallOpts, fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "diffDays", fromTimestamp, toTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiffDays is a free data retrieval call binding the contract method 0xc9d34622.
//
// Solidity: function diffDays(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _days)
func (_Datetime *DatetimeSession) DiffDays(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffDays(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffDays is a free data retrieval call binding the contract method 0xc9d34622.
//
// Solidity: function diffDays(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _days)
func (_Datetime *DatetimeCallerSession) DiffDays(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffDays(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffHours is a free data retrieval call binding the contract method 0x2af123b8.
//
// Solidity: function diffHours(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _hours)
func (_Datetime *DatetimeCaller) DiffHours(opts *bind.CallOpts, fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "diffHours", fromTimestamp, toTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiffHours is a free data retrieval call binding the contract method 0x2af123b8.
//
// Solidity: function diffHours(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _hours)
func (_Datetime *DatetimeSession) DiffHours(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffHours(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffHours is a free data retrieval call binding the contract method 0x2af123b8.
//
// Solidity: function diffHours(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _hours)
func (_Datetime *DatetimeCallerSession) DiffHours(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffHours(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffMinutes is a free data retrieval call binding the contract method 0x02e98e0d.
//
// Solidity: function diffMinutes(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _minutes)
func (_Datetime *DatetimeCaller) DiffMinutes(opts *bind.CallOpts, fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "diffMinutes", fromTimestamp, toTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiffMinutes is a free data retrieval call binding the contract method 0x02e98e0d.
//
// Solidity: function diffMinutes(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _minutes)
func (_Datetime *DatetimeSession) DiffMinutes(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffMinutes(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffMinutes is a free data retrieval call binding the contract method 0x02e98e0d.
//
// Solidity: function diffMinutes(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _minutes)
func (_Datetime *DatetimeCallerSession) DiffMinutes(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffMinutes(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffMonths is a free data retrieval call binding the contract method 0x7be34109.
//
// Solidity: function diffMonths(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _months)
func (_Datetime *DatetimeCaller) DiffMonths(opts *bind.CallOpts, fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "diffMonths", fromTimestamp, toTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiffMonths is a free data retrieval call binding the contract method 0x7be34109.
//
// Solidity: function diffMonths(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _months)
func (_Datetime *DatetimeSession) DiffMonths(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffMonths(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffMonths is a free data retrieval call binding the contract method 0x7be34109.
//
// Solidity: function diffMonths(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _months)
func (_Datetime *DatetimeCallerSession) DiffMonths(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffMonths(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffSeconds is a free data retrieval call binding the contract method 0x62fb9697.
//
// Solidity: function diffSeconds(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _seconds)
func (_Datetime *DatetimeCaller) DiffSeconds(opts *bind.CallOpts, fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "diffSeconds", fromTimestamp, toTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiffSeconds is a free data retrieval call binding the contract method 0x62fb9697.
//
// Solidity: function diffSeconds(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _seconds)
func (_Datetime *DatetimeSession) DiffSeconds(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffSeconds(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffSeconds is a free data retrieval call binding the contract method 0x62fb9697.
//
// Solidity: function diffSeconds(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _seconds)
func (_Datetime *DatetimeCallerSession) DiffSeconds(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffSeconds(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffYears is a free data retrieval call binding the contract method 0xff2258cb.
//
// Solidity: function diffYears(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _years)
func (_Datetime *DatetimeCaller) DiffYears(opts *bind.CallOpts, fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "diffYears", fromTimestamp, toTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DiffYears is a free data retrieval call binding the contract method 0xff2258cb.
//
// Solidity: function diffYears(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _years)
func (_Datetime *DatetimeSession) DiffYears(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffYears(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// DiffYears is a free data retrieval call binding the contract method 0xff2258cb.
//
// Solidity: function diffYears(uint256 fromTimestamp, uint256 toTimestamp) pure returns(uint256 _years)
func (_Datetime *DatetimeCallerSession) DiffYears(fromTimestamp *big.Int, toTimestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.DiffYears(&_Datetime.CallOpts, fromTimestamp, toTimestamp)
}

// GetDay is a free data retrieval call binding the contract method 0x65c72840.
//
// Solidity: function getDay(uint256 timestamp) pure returns(uint256 day)
func (_Datetime *DatetimeCaller) GetDay(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getDay", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDay is a free data retrieval call binding the contract method 0x65c72840.
//
// Solidity: function getDay(uint256 timestamp) pure returns(uint256 day)
func (_Datetime *DatetimeSession) GetDay(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDay(&_Datetime.CallOpts, timestamp)
}

// GetDay is a free data retrieval call binding the contract method 0x65c72840.
//
// Solidity: function getDay(uint256 timestamp) pure returns(uint256 day)
func (_Datetime *DatetimeCallerSession) GetDay(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDay(&_Datetime.CallOpts, timestamp)
}

// GetDayOfWeek is a free data retrieval call binding the contract method 0x22f8a2b8.
//
// Solidity: function getDayOfWeek(uint256 timestamp) pure returns(uint256 dayOfWeek)
func (_Datetime *DatetimeCaller) GetDayOfWeek(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getDayOfWeek", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDayOfWeek is a free data retrieval call binding the contract method 0x22f8a2b8.
//
// Solidity: function getDayOfWeek(uint256 timestamp) pure returns(uint256 dayOfWeek)
func (_Datetime *DatetimeSession) GetDayOfWeek(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDayOfWeek(&_Datetime.CallOpts, timestamp)
}

// GetDayOfWeek is a free data retrieval call binding the contract method 0x22f8a2b8.
//
// Solidity: function getDayOfWeek(uint256 timestamp) pure returns(uint256 dayOfWeek)
func (_Datetime *DatetimeCallerSession) GetDayOfWeek(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDayOfWeek(&_Datetime.CallOpts, timestamp)
}

// GetDaysInMonth2 is a free data retrieval call binding the contract method 0x10848ddf.
//
// Solidity: function getDaysInMonth(uint256 timestamp) pure returns(uint256 daysInMonth)
func (_Datetime *DatetimeCaller) GetDaysInMonth2(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getDaysInMonth", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDaysInMonth2 is a free data retrieval call binding the contract method 0x10848ddf.
//
// Solidity: function getDaysInMonth(uint256 timestamp) pure returns(uint256 daysInMonth)
func (_Datetime *DatetimeSession) GetDaysInMonth2(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDaysInMonth2(&_Datetime.CallOpts, timestamp)
}

// GetDaysInMonth2 is a free data retrieval call binding the contract method 0x10848ddf.
//
// Solidity: function getDaysInMonth(uint256 timestamp) pure returns(uint256 daysInMonth)
func (_Datetime *DatetimeCallerSession) GetDaysInMonth2(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetDaysInMonth2(&_Datetime.CallOpts, timestamp)
}

// GetHour is a free data retrieval call binding the contract method 0x3e239e1a.
//
// Solidity: function getHour(uint256 timestamp) pure returns(uint256 hour)
func (_Datetime *DatetimeCaller) GetHour(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getHour", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHour is a free data retrieval call binding the contract method 0x3e239e1a.
//
// Solidity: function getHour(uint256 timestamp) pure returns(uint256 hour)
func (_Datetime *DatetimeSession) GetHour(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetHour(&_Datetime.CallOpts, timestamp)
}

// GetHour is a free data retrieval call binding the contract method 0x3e239e1a.
//
// Solidity: function getHour(uint256 timestamp) pure returns(uint256 hour)
func (_Datetime *DatetimeCallerSession) GetHour(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetHour(&_Datetime.CallOpts, timestamp)
}

// GetMinute is a free data retrieval call binding the contract method 0xfa93f883.
//
// Solidity: function getMinute(uint256 timestamp) pure returns(uint256 minute)
func (_Datetime *DatetimeCaller) GetMinute(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getMinute", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinute is a free data retrieval call binding the contract method 0xfa93f883.
//
// Solidity: function getMinute(uint256 timestamp) pure returns(uint256 minute)
func (_Datetime *DatetimeSession) GetMinute(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetMinute(&_Datetime.CallOpts, timestamp)
}

// GetMinute is a free data retrieval call binding the contract method 0xfa93f883.
//
// Solidity: function getMinute(uint256 timestamp) pure returns(uint256 minute)
func (_Datetime *DatetimeCallerSession) GetMinute(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetMinute(&_Datetime.CallOpts, timestamp)
}

// GetMonth is a free data retrieval call binding the contract method 0xa324ad24.
//
// Solidity: function getMonth(uint256 timestamp) pure returns(uint256 month)
func (_Datetime *DatetimeCaller) GetMonth(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getMonth", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMonth is a free data retrieval call binding the contract method 0xa324ad24.
//
// Solidity: function getMonth(uint256 timestamp) pure returns(uint256 month)
func (_Datetime *DatetimeSession) GetMonth(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetMonth(&_Datetime.CallOpts, timestamp)
}

// GetMonth is a free data retrieval call binding the contract method 0xa324ad24.
//
// Solidity: function getMonth(uint256 timestamp) pure returns(uint256 month)
func (_Datetime *DatetimeCallerSession) GetMonth(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetMonth(&_Datetime.CallOpts, timestamp)
}

// GetSecond is a free data retrieval call binding the contract method 0x8aa001fc.
//
// Solidity: function getSecond(uint256 timestamp) pure returns(uint256 second)
func (_Datetime *DatetimeCaller) GetSecond(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getSecond", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSecond is a free data retrieval call binding the contract method 0x8aa001fc.
//
// Solidity: function getSecond(uint256 timestamp) pure returns(uint256 second)
func (_Datetime *DatetimeSession) GetSecond(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetSecond(&_Datetime.CallOpts, timestamp)
}

// GetSecond is a free data retrieval call binding the contract method 0x8aa001fc.
//
// Solidity: function getSecond(uint256 timestamp) pure returns(uint256 second)
func (_Datetime *DatetimeCallerSession) GetSecond(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetSecond(&_Datetime.CallOpts, timestamp)
}

// GetYear is a free data retrieval call binding the contract method 0x92d66313.
//
// Solidity: function getYear(uint256 timestamp) pure returns(uint256 year)
func (_Datetime *DatetimeCaller) GetYear(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "getYear", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetYear is a free data retrieval call binding the contract method 0x92d66313.
//
// Solidity: function getYear(uint256 timestamp) pure returns(uint256 year)
func (_Datetime *DatetimeSession) GetYear(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetYear(&_Datetime.CallOpts, timestamp)
}

// GetYear is a free data retrieval call binding the contract method 0x92d66313.
//
// Solidity: function getYear(uint256 timestamp) pure returns(uint256 year)
func (_Datetime *DatetimeCallerSession) GetYear(timestamp *big.Int) (*big.Int, error) {
	return _Datetime.Contract.GetYear(&_Datetime.CallOpts, timestamp)
}

// IsLeapYear is a free data retrieval call binding the contract method 0xb8d16dbc.
//
// Solidity: function isLeapYear(uint256 timestamp) pure returns(bool leapYear)
func (_Datetime *DatetimeCaller) IsLeapYear(opts *bind.CallOpts, timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "isLeapYear", timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeapYear is a free data retrieval call binding the contract method 0xb8d16dbc.
//
// Solidity: function isLeapYear(uint256 timestamp) pure returns(bool leapYear)
func (_Datetime *DatetimeSession) IsLeapYear(timestamp *big.Int) (bool, error) {
	return _Datetime.Contract.IsLeapYear(&_Datetime.CallOpts, timestamp)
}

// IsLeapYear is a free data retrieval call binding the contract method 0xb8d16dbc.
//
// Solidity: function isLeapYear(uint256 timestamp) pure returns(bool leapYear)
func (_Datetime *DatetimeCallerSession) IsLeapYear(timestamp *big.Int) (bool, error) {
	return _Datetime.Contract.IsLeapYear(&_Datetime.CallOpts, timestamp)
}

// IsValidDate is a free data retrieval call binding the contract method 0x14b2d6dc.
//
// Solidity: function isValidDate(uint256 year, uint256 month, uint256 day) pure returns(bool valid)
func (_Datetime *DatetimeCaller) IsValidDate(opts *bind.CallOpts, year *big.Int, month *big.Int, day *big.Int) (bool, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "isValidDate", year, month, day)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidDate is a free data retrieval call binding the contract method 0x14b2d6dc.
//
// Solidity: function isValidDate(uint256 year, uint256 month, uint256 day) pure returns(bool valid)
func (_Datetime *DatetimeSession) IsValidDate(year *big.Int, month *big.Int, day *big.Int) (bool, error) {
	return _Datetime.Contract.IsValidDate(&_Datetime.CallOpts, year, month, day)
}

// IsValidDate is a free data retrieval call binding the contract method 0x14b2d6dc.
//
// Solidity: function isValidDate(uint256 year, uint256 month, uint256 day) pure returns(bool valid)
func (_Datetime *DatetimeCallerSession) IsValidDate(year *big.Int, month *big.Int, day *big.Int) (bool, error) {
	return _Datetime.Contract.IsValidDate(&_Datetime.CallOpts, year, month, day)
}

// IsValidDateTime is a free data retrieval call binding the contract method 0x3293d007.
//
// Solidity: function isValidDateTime(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second) pure returns(bool valid)
func (_Datetime *DatetimeCaller) IsValidDateTime(opts *bind.CallOpts, year *big.Int, month *big.Int, day *big.Int, hour *big.Int, minute *big.Int, second *big.Int) (bool, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "isValidDateTime", year, month, day, hour, minute, second)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidDateTime is a free data retrieval call binding the contract method 0x3293d007.
//
// Solidity: function isValidDateTime(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second) pure returns(bool valid)
func (_Datetime *DatetimeSession) IsValidDateTime(year *big.Int, month *big.Int, day *big.Int, hour *big.Int, minute *big.Int, second *big.Int) (bool, error) {
	return _Datetime.Contract.IsValidDateTime(&_Datetime.CallOpts, year, month, day, hour, minute, second)
}

// IsValidDateTime is a free data retrieval call binding the contract method 0x3293d007.
//
// Solidity: function isValidDateTime(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second) pure returns(bool valid)
func (_Datetime *DatetimeCallerSession) IsValidDateTime(year *big.Int, month *big.Int, day *big.Int, hour *big.Int, minute *big.Int, second *big.Int) (bool, error) {
	return _Datetime.Contract.IsValidDateTime(&_Datetime.CallOpts, year, month, day, hour, minute, second)
}

// IsWeekDay is a free data retrieval call binding the contract method 0xd6582d0d.
//
// Solidity: function isWeekDay(uint256 timestamp) pure returns(bool weekDay)
func (_Datetime *DatetimeCaller) IsWeekDay(opts *bind.CallOpts, timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "isWeekDay", timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWeekDay is a free data retrieval call binding the contract method 0xd6582d0d.
//
// Solidity: function isWeekDay(uint256 timestamp) pure returns(bool weekDay)
func (_Datetime *DatetimeSession) IsWeekDay(timestamp *big.Int) (bool, error) {
	return _Datetime.Contract.IsWeekDay(&_Datetime.CallOpts, timestamp)
}

// IsWeekDay is a free data retrieval call binding the contract method 0xd6582d0d.
//
// Solidity: function isWeekDay(uint256 timestamp) pure returns(bool weekDay)
func (_Datetime *DatetimeCallerSession) IsWeekDay(timestamp *big.Int) (bool, error) {
	return _Datetime.Contract.IsWeekDay(&_Datetime.CallOpts, timestamp)
}

// IsWeekEnd is a free data retrieval call binding the contract method 0x4371c465.
//
// Solidity: function isWeekEnd(uint256 timestamp) pure returns(bool weekEnd)
func (_Datetime *DatetimeCaller) IsWeekEnd(opts *bind.CallOpts, timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "isWeekEnd", timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWeekEnd is a free data retrieval call binding the contract method 0x4371c465.
//
// Solidity: function isWeekEnd(uint256 timestamp) pure returns(bool weekEnd)
func (_Datetime *DatetimeSession) IsWeekEnd(timestamp *big.Int) (bool, error) {
	return _Datetime.Contract.IsWeekEnd(&_Datetime.CallOpts, timestamp)
}

// IsWeekEnd is a free data retrieval call binding the contract method 0x4371c465.
//
// Solidity: function isWeekEnd(uint256 timestamp) pure returns(bool weekEnd)
func (_Datetime *DatetimeCallerSession) IsWeekEnd(timestamp *big.Int) (bool, error) {
	return _Datetime.Contract.IsWeekEnd(&_Datetime.CallOpts, timestamp)
}

// SubDays is a free data retrieval call binding the contract method 0xc7b6fd6a.
//
// Solidity: function subDays(uint256 timestamp, uint256 _days) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) SubDays(opts *bind.CallOpts, timestamp *big.Int, _days *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "subDays", timestamp, _days)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubDays is a free data retrieval call binding the contract method 0xc7b6fd6a.
//
// Solidity: function subDays(uint256 timestamp, uint256 _days) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) SubDays(timestamp *big.Int, _days *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubDays(&_Datetime.CallOpts, timestamp, _days)
}

// SubDays is a free data retrieval call binding the contract method 0xc7b6fd6a.
//
// Solidity: function subDays(uint256 timestamp, uint256 _days) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) SubDays(timestamp *big.Int, _days *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubDays(&_Datetime.CallOpts, timestamp, _days)
}

// SubHours is a free data retrieval call binding the contract method 0x00501553.
//
// Solidity: function subHours(uint256 timestamp, uint256 _hours) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) SubHours(opts *bind.CallOpts, timestamp *big.Int, _hours *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "subHours", timestamp, _hours)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubHours is a free data retrieval call binding the contract method 0x00501553.
//
// Solidity: function subHours(uint256 timestamp, uint256 _hours) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) SubHours(timestamp *big.Int, _hours *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubHours(&_Datetime.CallOpts, timestamp, _hours)
}

// SubHours is a free data retrieval call binding the contract method 0x00501553.
//
// Solidity: function subHours(uint256 timestamp, uint256 _hours) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) SubHours(timestamp *big.Int, _hours *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubHours(&_Datetime.CallOpts, timestamp, _hours)
}

// SubMinutes is a free data retrieval call binding the contract method 0x9e524caa.
//
// Solidity: function subMinutes(uint256 timestamp, uint256 _minutes) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) SubMinutes(opts *bind.CallOpts, timestamp *big.Int, _minutes *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "subMinutes", timestamp, _minutes)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubMinutes is a free data retrieval call binding the contract method 0x9e524caa.
//
// Solidity: function subMinutes(uint256 timestamp, uint256 _minutes) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) SubMinutes(timestamp *big.Int, _minutes *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubMinutes(&_Datetime.CallOpts, timestamp, _minutes)
}

// SubMinutes is a free data retrieval call binding the contract method 0x9e524caa.
//
// Solidity: function subMinutes(uint256 timestamp, uint256 _minutes) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) SubMinutes(timestamp *big.Int, _minutes *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubMinutes(&_Datetime.CallOpts, timestamp, _minutes)
}

// SubMonths is a free data retrieval call binding the contract method 0xad203bd4.
//
// Solidity: function subMonths(uint256 timestamp, uint256 _months) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) SubMonths(opts *bind.CallOpts, timestamp *big.Int, _months *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "subMonths", timestamp, _months)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubMonths is a free data retrieval call binding the contract method 0xad203bd4.
//
// Solidity: function subMonths(uint256 timestamp, uint256 _months) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) SubMonths(timestamp *big.Int, _months *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubMonths(&_Datetime.CallOpts, timestamp, _months)
}

// SubMonths is a free data retrieval call binding the contract method 0xad203bd4.
//
// Solidity: function subMonths(uint256 timestamp, uint256 _months) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) SubMonths(timestamp *big.Int, _months *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubMonths(&_Datetime.CallOpts, timestamp, _months)
}

// SubSeconds is a free data retrieval call binding the contract method 0xf615ed54.
//
// Solidity: function subSeconds(uint256 timestamp, uint256 _seconds) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) SubSeconds(opts *bind.CallOpts, timestamp *big.Int, _seconds *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "subSeconds", timestamp, _seconds)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubSeconds is a free data retrieval call binding the contract method 0xf615ed54.
//
// Solidity: function subSeconds(uint256 timestamp, uint256 _seconds) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) SubSeconds(timestamp *big.Int, _seconds *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubSeconds(&_Datetime.CallOpts, timestamp, _seconds)
}

// SubSeconds is a free data retrieval call binding the contract method 0xf615ed54.
//
// Solidity: function subSeconds(uint256 timestamp, uint256 _seconds) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) SubSeconds(timestamp *big.Int, _seconds *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubSeconds(&_Datetime.CallOpts, timestamp, _seconds)
}

// SubYears is a free data retrieval call binding the contract method 0x444fda82.
//
// Solidity: function subYears(uint256 timestamp, uint256 _years) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCaller) SubYears(opts *bind.CallOpts, timestamp *big.Int, _years *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "subYears", timestamp, _years)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubYears is a free data retrieval call binding the contract method 0x444fda82.
//
// Solidity: function subYears(uint256 timestamp, uint256 _years) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeSession) SubYears(timestamp *big.Int, _years *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubYears(&_Datetime.CallOpts, timestamp, _years)
}

// SubYears is a free data retrieval call binding the contract method 0x444fda82.
//
// Solidity: function subYears(uint256 timestamp, uint256 _years) pure returns(uint256 newTimestamp)
func (_Datetime *DatetimeCallerSession) SubYears(timestamp *big.Int, _years *big.Int) (*big.Int, error) {
	return _Datetime.Contract.SubYears(&_Datetime.CallOpts, timestamp, _years)
}

// TimestampFromDate is a free data retrieval call binding the contract method 0x1f4f77b2.
//
// Solidity: function timestampFromDate(uint256 year, uint256 month, uint256 day) pure returns(uint256 timestamp)
func (_Datetime *DatetimeCaller) TimestampFromDate(opts *bind.CallOpts, year *big.Int, month *big.Int, day *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "timestampFromDate", year, month, day)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampFromDate is a free data retrieval call binding the contract method 0x1f4f77b2.
//
// Solidity: function timestampFromDate(uint256 year, uint256 month, uint256 day) pure returns(uint256 timestamp)
func (_Datetime *DatetimeSession) TimestampFromDate(year *big.Int, month *big.Int, day *big.Int) (*big.Int, error) {
	return _Datetime.Contract.TimestampFromDate(&_Datetime.CallOpts, year, month, day)
}

// TimestampFromDate is a free data retrieval call binding the contract method 0x1f4f77b2.
//
// Solidity: function timestampFromDate(uint256 year, uint256 month, uint256 day) pure returns(uint256 timestamp)
func (_Datetime *DatetimeCallerSession) TimestampFromDate(year *big.Int, month *big.Int, day *big.Int) (*big.Int, error) {
	return _Datetime.Contract.TimestampFromDate(&_Datetime.CallOpts, year, month, day)
}

// TimestampFromDateTime is a free data retrieval call binding the contract method 0x5e05bd6d.
//
// Solidity: function timestampFromDateTime(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second) pure returns(uint256 timestamp)
func (_Datetime *DatetimeCaller) TimestampFromDateTime(opts *bind.CallOpts, year *big.Int, month *big.Int, day *big.Int, hour *big.Int, minute *big.Int, second *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "timestampFromDateTime", year, month, day, hour, minute, second)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimestampFromDateTime is a free data retrieval call binding the contract method 0x5e05bd6d.
//
// Solidity: function timestampFromDateTime(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second) pure returns(uint256 timestamp)
func (_Datetime *DatetimeSession) TimestampFromDateTime(year *big.Int, month *big.Int, day *big.Int, hour *big.Int, minute *big.Int, second *big.Int) (*big.Int, error) {
	return _Datetime.Contract.TimestampFromDateTime(&_Datetime.CallOpts, year, month, day, hour, minute, second)
}

// TimestampFromDateTime is a free data retrieval call binding the contract method 0x5e05bd6d.
//
// Solidity: function timestampFromDateTime(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second) pure returns(uint256 timestamp)
func (_Datetime *DatetimeCallerSession) TimestampFromDateTime(year *big.Int, month *big.Int, day *big.Int, hour *big.Int, minute *big.Int, second *big.Int) (*big.Int, error) {
	return _Datetime.Contract.TimestampFromDateTime(&_Datetime.CallOpts, year, month, day, hour, minute, second)
}

// TimestampToDate is a free data retrieval call binding the contract method 0xde5101af.
//
// Solidity: function timestampToDate(uint256 timestamp) pure returns(uint256 year, uint256 month, uint256 day)
func (_Datetime *DatetimeCaller) TimestampToDate(opts *bind.CallOpts, timestamp *big.Int) (struct {
	Year  *big.Int
	Month *big.Int
	Day   *big.Int
}, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "timestampToDate", timestamp)

	outstruct := new(struct {
		Year  *big.Int
		Month *big.Int
		Day   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Year = out[0].(*big.Int)
	outstruct.Month = out[1].(*big.Int)
	outstruct.Day = out[2].(*big.Int)

	return *outstruct, err

}

// TimestampToDate is a free data retrieval call binding the contract method 0xde5101af.
//
// Solidity: function timestampToDate(uint256 timestamp) pure returns(uint256 year, uint256 month, uint256 day)
func (_Datetime *DatetimeSession) TimestampToDate(timestamp *big.Int) (struct {
	Year  *big.Int
	Month *big.Int
	Day   *big.Int
}, error) {
	return _Datetime.Contract.TimestampToDate(&_Datetime.CallOpts, timestamp)
}

// TimestampToDate is a free data retrieval call binding the contract method 0xde5101af.
//
// Solidity: function timestampToDate(uint256 timestamp) pure returns(uint256 year, uint256 month, uint256 day)
func (_Datetime *DatetimeCallerSession) TimestampToDate(timestamp *big.Int) (struct {
	Year  *big.Int
	Month *big.Int
	Day   *big.Int
}, error) {
	return _Datetime.Contract.TimestampToDate(&_Datetime.CallOpts, timestamp)
}

// TimestampToDateTime is a free data retrieval call binding the contract method 0xea1c1690.
//
// Solidity: function timestampToDateTime(uint256 timestamp) pure returns(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second)
func (_Datetime *DatetimeCaller) TimestampToDateTime(opts *bind.CallOpts, timestamp *big.Int) (struct {
	Year   *big.Int
	Month  *big.Int
	Day    *big.Int
	Hour   *big.Int
	Minute *big.Int
	Second *big.Int
}, error) {
	var out []interface{}
	err := _Datetime.contract.Call(opts, &out, "timestampToDateTime", timestamp)

	outstruct := new(struct {
		Year   *big.Int
		Month  *big.Int
		Day    *big.Int
		Hour   *big.Int
		Minute *big.Int
		Second *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Year = out[0].(*big.Int)
	outstruct.Month = out[1].(*big.Int)
	outstruct.Day = out[2].(*big.Int)
	outstruct.Hour = out[3].(*big.Int)
	outstruct.Minute = out[4].(*big.Int)
	outstruct.Second = out[5].(*big.Int)

	return *outstruct, err

}

// TimestampToDateTime is a free data retrieval call binding the contract method 0xea1c1690.
//
// Solidity: function timestampToDateTime(uint256 timestamp) pure returns(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second)
func (_Datetime *DatetimeSession) TimestampToDateTime(timestamp *big.Int) (struct {
	Year   *big.Int
	Month  *big.Int
	Day    *big.Int
	Hour   *big.Int
	Minute *big.Int
	Second *big.Int
}, error) {
	return _Datetime.Contract.TimestampToDateTime(&_Datetime.CallOpts, timestamp)
}

// TimestampToDateTime is a free data retrieval call binding the contract method 0xea1c1690.
//
// Solidity: function timestampToDateTime(uint256 timestamp) pure returns(uint256 year, uint256 month, uint256 day, uint256 hour, uint256 minute, uint256 second)
func (_Datetime *DatetimeCallerSession) TimestampToDateTime(timestamp *big.Int) (struct {
	Year   *big.Int
	Month  *big.Int
	Day    *big.Int
	Hour   *big.Int
	Minute *big.Int
	Second *big.Int
}, error) {
	return _Datetime.Contract.TimestampToDateTime(&_Datetime.CallOpts, timestamp)
}
