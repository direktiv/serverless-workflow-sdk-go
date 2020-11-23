// Copyright 2020 The Serverless Workflow Specification Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"encoding/json"
	"fmt"
)

// UnmarshalJSON implementation for json Unmarshal function for the Workflow type
func (w *Workflow) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &w.WorkflowCommon); err != nil {
		return err
	}

	workflowMap := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &workflowMap); err != nil {
		return err
	}
	var rawStates []json.RawMessage
	if err := json.Unmarshal(workflowMap["states"], &rawStates); err != nil {
		return err
	}

	w.States = make([]State, len(rawStates))
	var mapState map[string]interface{}
	for i, rawState := range rawStates {
		if err := json.Unmarshal(rawState, &mapState); err != nil {
			return err
		}
		if _, ok := actionsModelMapping[mapState["type"].(string)]; !ok {
			return fmt.Errorf("state %s not supported", mapState["type"])
		}
		state := actionsModelMapping[mapState["type"].(string)](mapState)
		if err := json.Unmarshal(rawState, &state); err != nil {
			return err
		}
		w.States[i] = state
	}
	return nil
}

type basicEventbasedswitch struct {
	DataInputSchema  *string          `json:"dataInputSchema,omitempty"`
	DataOutputSchema *string          `json:"dataOutputSchema,omitempty"`
	Default          *Defaultdef      `json:"default,omitempty"`
	EventTimeout     *string          `json:"eventTimeout,omitempty"`
	Id               *string          `json:"id,omitempty"`
	Metadata         Metadata_1       `json:"metadata,omitempty"`
	Name             *string          `json:"name,omitempty"`
	OnErrors         []Error          `json:"onErrors,omitempty"`
	Start            *Start           `json:"start,omitempty"`
	StateDataFilter  *Statedatafilter `json:"stateDataFilter,omitempty"`
	Type             *string          `json:"type,omitempty"`
}

// UnmarshalJSON implementation for json Unmarshal function for the Eventbasedswitch type
func (j *Eventbasedswitch) UnmarshalJSON(data []byte) error {
	eventBasedSwitch := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &eventBasedSwitch)
	if err != nil {
		return err
	}
	var rawConditions []json.RawMessage
	err = json.Unmarshal(eventBasedSwitch["eventConditions"], &rawConditions)
	if err != nil {
		return err
	}

	j.EventConditions = make([]EventbasedswitchEventConditionsElem, len(rawConditions))
	var mapConditions map[string]interface{}
	for i, rawCondition := range rawConditions {
		err = json.Unmarshal(rawCondition, &mapConditions)
		if err != nil {
			return err
		}
		var condition EventbasedswitchEventConditionsElem
		if _, ok := mapConditions["end"]; ok {
			condition = &Enddeventcondition{}
		} else {
			condition = &Transitioneventcondition{}
		}
		err := json.Unmarshal(rawCondition, condition)
		if err != nil {
			return err
		}
		j.EventConditions[i] = condition
	}

	o := new(basicEventbasedswitch)
	err = json.Unmarshal(data, o)
	if err != nil {
		return err
	}

	j.DataInputSchema = o.DataInputSchema
	j.DataOutputSchema = o.DataOutputSchema
	j.Default = o.Default
	j.EventTimeout = o.EventTimeout
	j.Id = o.Id
	j.Metadata = o.Metadata
	j.Name = o.Name
	j.OnErrors = o.OnErrors
	j.Start = o.Start
	j.StateDataFilter = o.StateDataFilter
	j.Type = o.Type

	return nil
}

type basedDatabasedswitch struct {
	DataInputSchema  *string          `json:"dataInputSchema,omitempty"`
	DataOutputSchema *string          `json:"dataOutputSchema,omitempty"`
	Default          *Defaultdef      `json:"default,omitempty"`
	Id               *string          `json:"id,omitempty"`
	Metadata         Metadata_1       `json:"metadata,omitempty"`
	Name             *string          `json:"name,omitempty"`
	OnErrors         []Error          `json:"onErrors,omitempty"`
	Start            *Start           `json:"start,omitempty"`
	StateDataFilter  *Statedatafilter `json:"stateDataFilter,omitempty"`
	Type             *string          `json:"type,omitempty"`
}

// UnmarshalJSON implementation for json Unmarshal function for the Databasedswitch type
func (j *Databasedswitch) UnmarshalJSON(data []byte) error {
	dataBasedSwitch := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &dataBasedSwitch)
	if err != nil {
		return err
	}
	var rawConditions []json.RawMessage
	err = json.Unmarshal(dataBasedSwitch["dataConditions"], &rawConditions)
	if err != nil {
		return err
	}

	j.DataConditions = make([]DatabasedswitchDataConditionsElem, len(rawConditions))
	var mapConditions map[string]interface{}
	for i, rawCondition := range rawConditions {
		err = json.Unmarshal(rawCondition, &mapConditions)
		if err != nil {
			return err
		}
		var condition DatabasedswitchDataConditionsElem
		if _, ok := mapConditions["end"]; ok {
			condition = &Enddatacondition{}
		} else {
			condition = &Transitiondatacondition{}
		}
		err := json.Unmarshal(rawCondition, condition)
		if err != nil {
			return err
		}
		j.DataConditions[i] = condition
	}

	o := new(basedDatabasedswitch)
	err = json.Unmarshal(data, o)
	if err != nil {
		return err
	}

	j.DataInputSchema = o.DataInputSchema
	j.DataOutputSchema = o.DataOutputSchema
	j.Default = o.Default
	j.Id = o.Id
	j.Metadata = o.Metadata
	j.Name = o.Name
	j.OnErrors = o.OnErrors
	j.Start = o.Start
	j.StateDataFilter = o.StateDataFilter
	j.Type = o.Type

	return nil
}
