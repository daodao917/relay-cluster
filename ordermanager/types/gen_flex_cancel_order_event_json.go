// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

// MarshalJSON marshals as JSON.
func (f FlexCancelOrderEvent) MarshalJSON() ([]byte, error) {
	type FlexCancelOrderEvent struct {
		Owner      common.Address `json:"owner"`
		OrderHash  common.Hash    `json:"order_hash"`
		CutoffTime int64          `json:"cutoff_time"`
		TokenS     common.Address `json:"token_s"`
		TokenB     common.Address `json:"token_b"`
		Type       FlexCancelType `json:"type"`
	}
	var enc FlexCancelOrderEvent
	enc.Owner = f.Owner
	enc.OrderHash = f.OrderHash
	enc.CutoffTime = f.CutoffTime
	enc.TokenS = f.TokenS
	enc.TokenB = f.TokenB
	enc.Type = f.Type
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (f *FlexCancelOrderEvent) UnmarshalJSON(input []byte) error {
	type FlexCancelOrderEvent struct {
		Owner      *common.Address `json:"owner"`
		OrderHash  *common.Hash    `json:"order_hash"`
		CutoffTime *int64          `json:"cutoff_time"`
		TokenS     *common.Address `json:"token_s"`
		TokenB     *common.Address `json:"token_b"`
		Type       *FlexCancelType `json:"type"`
	}
	var dec FlexCancelOrderEvent
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Owner != nil {
		f.Owner = *dec.Owner
	}
	if dec.OrderHash != nil {
		f.OrderHash = *dec.OrderHash
	}
	if dec.CutoffTime != nil {
		f.CutoffTime = *dec.CutoffTime
	}
	if dec.TokenS != nil {
		f.TokenS = *dec.TokenS
	}
	if dec.TokenB != nil {
		f.TokenB = *dec.TokenB
	}
	if dec.Type != nil {
		f.Type = *dec.Type
	}
	return nil
}
