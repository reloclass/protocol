package student

import (
	`strconv`
)

func (a *AddReq) ParseStudentIds() (results []int64, err error) {
	results = make([]int64, len(a.StudentIds))
	for i, id := range a.StudentIds {
		var val int64
		if val, err = strconv.ParseInt(id, 10, 64); nil != err {
			return
		}

		results[i] = val
	}

	return
}
