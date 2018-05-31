/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package manager

import (
	"github.com/Loopring/relay-cluster/dao"
	"github.com/Loopring/relay-lib/log"
	"github.com/Loopring/relay-lib/types"
)

type SubmitRingHandler struct {
	Event *types.SubmitRingMethodEvent
	Rds   *dao.RdsService
}

func (handler *SubmitRingHandler) HandleFailed() error {
	event := handler.Event
	rds := handler.Rds

	if event.Status != types.TX_STATUS_FAILED {
		return nil
	}

	model, err := rds.FindRingMined(event.TxHash.Hex())
	if err == nil {
		log.Debugf("order manager,handle submitRing method,tx %s has already exist", event.TxHash.Hex())
		return nil
	} else {
		log.Debugf("order manager,handle submitRing method,tx:%s status:%s inserted", event.TxHash.Hex(), types.StatusStr(event.Status))
		model.FromSubmitRingMethod(event)
		return rds.Add(model)
	}
}

// todo
func (handler *SubmitRingHandler) HandlePending() error {
	return nil
}

func (handler *SubmitRingHandler) HandleSuccess() error {
	return nil
}

type RingMinedHandler struct {
	Event *types.RingMinedEvent
	Rds   *dao.RdsService
}

func (handler *RingMinedHandler) HandleFailed() error {
	return nil
}

func (handler *RingMinedHandler) HandlePending() error {
	return nil
}

func (handler *RingMinedHandler) HandleSuccess() error {
	event := handler.Event
	rds := handler.Rds

	if event.Status != types.TX_STATUS_SUCCESS {
		return nil
	}

	if model, err := rds.FindRingMined(event.TxHash.Hex()); err == nil {
		log.Errorf("order manager,handle ringmined event,tx:%s ringhash:%s already exist", event.TxHash.Hex(), event.Ringhash.Hex())
		return nil
	} else {
		log.Debugf("order manager,handle ringmined event,tx:%s, ringhash:%s inserted", event.TxHash.Hex(), event.Ringhash.Hex())
		model.ConvertDown(event)
		return rds.Add(model)
	}
}