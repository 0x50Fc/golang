package lookin

import (
	"github.com/hailongz/golang/dynamic"
	"github.com/hailongz/golang/micro"
)

func (S *Service) Code(app micro.IContext, task *CodeTask) (*CodeData, error) {

	maxLevel := int(dynamic.IntValue(dynamic.GetWithKeys(app.GetConfig(), []string{"lookin", "maxLevel"}), 3))

	var ids []int64

	if task.Fcode != nil {
		ids, _ = DeocdeString(dynamic.StringValue(task.Fuid, ""))
		if ids != nil {
			for i, id := range ids {
				if id == task.Uid {
					ids = ids[i+1:]
					break
				}
			}
		}
	} else if task.Fuid != nil {
		ids = []int64{dynamic.IntValue(task.Fuid, 0)}
	}

	if ids == nil {
		ids = []int64{task.Uid}
	} else {
		ids = append(ids, task.Uid)
	}

	n := len(ids)

	if n > maxLevel {
		ids = ids[n-maxLevel : 0]
	}

	return &CodeData{Code: EncodeToString(ids)}, nil
}
