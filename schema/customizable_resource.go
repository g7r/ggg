package schema

import "encoding/json"

type customJSON []interface{}

type deleteField struct{ dummy int }

var DeleteField deleteField

func (cjson *customJSON) Add(customJSON interface{}) {
	*cjson = append(*cjson, customJSON)
}

func marshalResource(r interface{}, cjson customJSON) ([]byte, error) {
	primaryJSON, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	var primaryJSONMap map[string]interface{}
	if err := json.Unmarshal(primaryJSON, &primaryJSONMap); err != nil {
		return nil, err
	}

	for _, secondary := range cjson {
		secondaryJSON, err := json.Marshal(secondary)
		if err != nil {
			return nil, err
		}

		var secondaryJSONMap map[string]interface{}
		if err := json.Unmarshal(secondaryJSON, &secondaryJSONMap); err != nil {
			return nil, err
		}

		for k, v := range secondaryJSONMap {
			if v == DeleteField {
				delete(primaryJSONMap, k)
			} else {
				primaryJSONMap[k] = v
			}
		}
	}

	finalJSON, err := json.Marshal(primaryJSONMap)
	if err != nil {
		return nil, err
	}

	return finalJSON, nil
}
