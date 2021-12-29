package coursetime

import (
	`strconv`
)

func (u *UpdateReq) ParseAssistantIds() (results []int64, err error) {
	results = make([]int64, len(u.AssistantIds))
	for i, id := range u.AssistantIds {
		var val int64
		if val, err = strconv.ParseInt(id, 10, 64); nil != err {
			return
		}

		results[i] = val
	}

	return
}
