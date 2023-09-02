// Copyright 2023 The concrete-geth Authors
//
// The concrete-geth library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The concrete library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the concrete library. If not, see <http://www.gnu.org/licenses/>.

package api

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Logger interface {
	Debug(msg string)
}

type BlockContext interface {
	GetHash(uint64) common.Hash
	GasLimit() uint64
	BlockNumber() uint64
	Timestamp() uint64
	Difficulty() *big.Int
	BaseFee() *big.Int
	Coinbase() common.Address
	Random() common.Hash
}

type CallContext interface {
	TxGasPrice() *big.Int
	TxOrigin() common.Address
	CallData() []byte
	CallDataSize() int
	Caller() common.Address
	CallValue() *big.Int
}

type Caller interface {
	CallStatic(addr common.Address, input []byte, gas uint64) ([]byte, uint64, error)
	Call(addr common.Address, input []byte, gas uint64, value *big.Int) ([]byte, uint64, error)
	CallDelegate(addr common.Address, input []byte, gas uint64) ([]byte, uint64, error)
	Create(input []byte, gas uint64, value *big.Int) (common.Address, uint64, error)
	Create2(input []byte, salt common.Hash, gas uint64, value *big.Int) (common.Address, uint64, error)
}
