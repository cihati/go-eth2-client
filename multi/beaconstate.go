// Copyright © 2021 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multi

import (
	"context"

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/spec"
)

// BeaconState fetches a beacon state.
// N.B if the requested beacon state is not available this will return nil without an error.
func (s *Service) BeaconState(ctx context.Context, stateID string) (*spec.VersionedBeaconState, error) {
	res, err := s.doCall(ctx, func(ctx context.Context, client consensusclient.Service) (interface{}, error) {
		beaconState, err := client.(consensusclient.BeaconStateProvider).BeaconState(ctx, stateID)
		if err != nil {
			return nil, err
		}
		return beaconState, nil
	}, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*spec.VersionedBeaconState), nil
}
