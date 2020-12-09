package cvm

import "math/big"

type accountObject struct {
	Address      common.Address              `json:"address,omitempty"`
	AddrHash     common.Hash                 `json:"addr_hash,omitempty"` // hash of ethereum address of the account
	ByteCode     []byte                      `json:"byte_code,omitempty"`
	Data         accountData                 `json:"data,omitempty"`
	CacheStorage map[common.Hash]common.Hash `json:"cache_storage,omitempty"` // 用于缓存存储的变量
}

type accountData struct {
	Nonce    uint64      `json:"nonce,omitempty"`
	Balance  *big.Int    `json:"balance,omitempty"`
	Root     common.Hash `json:"root,omitempty"` // merkle root of the storage trie
	CodeHash []byte      `json:"code_hash,omitempty"`
}

//AccountState 实现vm的StateDB的接口 用于进行测试
type AccountState struct {
	Accounts map[common.Address]*accountObject `json:"accounts,omitempty"`
}
