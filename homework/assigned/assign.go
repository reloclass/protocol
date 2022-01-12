package assigned

import (
	`strconv`
)

func (u *AssignReq) ParseAttachmentIds() (results []int64, err error) {
	results = make([]int64, len(u.AttachmentIds))
	for i, id := range u.AttachmentIds {
		var val int64
		if val, err = strconv.ParseInt(id, 10, 64); nil != err {
			return
		}

		results[i] = val
	}

	return
}
