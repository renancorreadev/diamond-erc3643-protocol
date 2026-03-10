package indexer

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// TokenState holds indexed state for a single tokenId.
type TokenState struct {
	TokenID     *big.Int          `json:"tokenId"`
	TotalSupply *big.Int          `json:"totalSupply"`
	HolderCount uint64            `json:"holderCount"`
	Holders     map[string]*big.Int `json:"holders"` // address → balance
}

// TransferEvent is an indexed transfer record.
type TransferEvent struct {
	TxHash   common.Hash    `json:"txHash"`
	Block    uint64         `json:"block"`
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	TokenID  *big.Int       `json:"tokenId"`
	Amount   *big.Int       `json:"amount"`
	EventType string        `json:"eventType"` // "mint", "burn", "transfer", "forced_transfer"
}

// MintEvent is an indexed mint record.
type MintEvent struct {
	TxHash  common.Hash    `json:"txHash"`
	Block   uint64         `json:"block"`
	TokenID *big.Int       `json:"tokenId"`
	To      common.Address `json:"to"`
	Amount  *big.Int       `json:"amount"`
}

// State is the in-memory indexed state. Thread-safe via RWMutex.
type State struct {
	mu        sync.RWMutex
	Tokens    map[string]*TokenState  // tokenId string → state
	Events    []TransferEvent
	LastBlock uint64
}

// NewState creates an empty state.
func NewState() *State {
	return &State{
		Tokens: make(map[string]*TokenState),
	}
}

// GetOrCreateToken returns existing token state or creates a new one.
func (s *State) GetOrCreateToken(tokenID *big.Int) *TokenState {
	key := tokenID.String()
	s.mu.Lock()
	defer s.mu.Unlock()

	if t, ok := s.Tokens[key]; ok {
		return t
	}
	t := &TokenState{
		TokenID:     new(big.Int).Set(tokenID),
		TotalSupply: big.NewInt(0),
		Holders:     make(map[string]*big.Int),
	}
	s.Tokens[key] = t
	return t
}

// RecordMint updates state for a mint event.
func (s *State) RecordMint(tokenID *big.Int, to common.Address, amount *big.Int, txHash common.Hash, block uint64) {
	t := s.GetOrCreateToken(tokenID)
	s.mu.Lock()
	defer s.mu.Unlock()

	t.TotalSupply.Add(t.TotalSupply, amount)

	addr := to.Hex()
	if _, ok := t.Holders[addr]; !ok {
		t.Holders[addr] = big.NewInt(0)
		t.HolderCount++
	}
	t.Holders[addr].Add(t.Holders[addr], amount)

	s.Events = append(s.Events, TransferEvent{
		TxHash:    txHash,
		Block:     block,
		From:      common.Address{},
		To:        to,
		TokenID:   tokenID,
		Amount:    amount,
		EventType: "mint",
	})
	s.LastBlock = block
}

// RecordBurn updates state for a burn event.
func (s *State) RecordBurn(tokenID *big.Int, from common.Address, amount *big.Int, txHash common.Hash, block uint64) {
	t := s.GetOrCreateToken(tokenID)
	s.mu.Lock()
	defer s.mu.Unlock()

	t.TotalSupply.Sub(t.TotalSupply, amount)

	addr := from.Hex()
	if bal, ok := t.Holders[addr]; ok {
		bal.Sub(bal, amount)
		if bal.Sign() <= 0 {
			delete(t.Holders, addr)
			t.HolderCount--
		}
	}

	s.Events = append(s.Events, TransferEvent{
		TxHash:    txHash,
		Block:     block,
		From:      from,
		To:        common.Address{},
		TokenID:   tokenID,
		Amount:    amount,
		EventType: "burn",
	})
	s.LastBlock = block
}

// RecordTransfer updates state for a transfer event.
func (s *State) RecordTransfer(tokenID *big.Int, from, to common.Address, amount *big.Int, eventType string, txHash common.Hash, block uint64) {
	t := s.GetOrCreateToken(tokenID)
	s.mu.Lock()
	defer s.mu.Unlock()

	fromAddr := from.Hex()
	toAddr := to.Hex()

	// Debit from
	if bal, ok := t.Holders[fromAddr]; ok {
		bal.Sub(bal, amount)
		if bal.Sign() <= 0 {
			delete(t.Holders, fromAddr)
			t.HolderCount--
		}
	}

	// Credit to
	if _, ok := t.Holders[toAddr]; !ok {
		t.Holders[toAddr] = big.NewInt(0)
		t.HolderCount++
	}
	t.Holders[toAddr].Add(t.Holders[toAddr], amount)

	s.Events = append(s.Events, TransferEvent{
		TxHash:    txHash,
		Block:     block,
		From:      from,
		To:        to,
		TokenID:   tokenID,
		Amount:    amount,
		EventType: eventType,
	})
	s.LastBlock = block
}

// GetTokenState returns a snapshot of token state (read-locked).
func (s *State) GetTokenState(tokenID string) (*TokenState, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.Tokens[tokenID]
	return t, ok
}

// GetAllTokens returns all token states.
func (s *State) GetAllTokens() map[string]*TokenState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result := make(map[string]*TokenState, len(s.Tokens))
	for k, v := range s.Tokens {
		result[k] = v
	}
	return result
}

// GetEvents returns recent events (last N).
func (s *State) GetEvents(limit int) []TransferEvent {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if limit <= 0 || limit > len(s.Events) {
		limit = len(s.Events)
	}
	start := len(s.Events) - limit
	result := make([]TransferEvent, limit)
	copy(result, s.Events[start:])
	return result
}
