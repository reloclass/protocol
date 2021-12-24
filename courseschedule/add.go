package courseschedule

import (
	`strconv`
)

func (a *AddReq) ParseAssistantIds() (results []int64, err error) {
	results = make([]int64, len(a.AssistantIds))
	for i, id := range a.AssistantIds {
		var val int64
		if val, err = strconv.ParseInt(id, 10, 64); nil != err {
			return
		}

		results[i] = val
	}

	return
}
