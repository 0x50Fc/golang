package top

import (
	"fmt"

	"github.com/hailongz/golang/db"
	"github.com/hailongz/golang/micro"
)

func (S *Service) RankSort(app micro.IContext, task *RankSortTask) (interface{}, error) {

	conn, prefix, err := app.GetDB("wd")

	if err != nil {
		return nil, err
	}

	prefix = fmt.Sprintf("%s%s_", prefix, task.Name)

	v := Top{}

	items := []*Top{}

	rm := []*Top{}

	fixedSet := []*Top{}
	fixedIndex := 0
	n := int32(0)

	scaner := db.NewScaner(&v)

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE fixed!=0 ORDER BY fixed ASC, sid DESC, id DESC")

		if err != nil {
			return err
		}

		defer rs.Close()

		for rs.Next() {
			err := scaner.Scan(rs)
			if err != nil {
				return err
			}

			item := Top{}
			item = v
			fixedSet = append(fixedSet, &item)
		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	err = func() error {

		rs, err := db.Query(conn, &v, prefix, " WHERE fixed=0 ORDER BY sid DESC, id DESC")

		if err != nil {
			return err
		}

		defer rs.Close()

		for rs.Next() {

			err := scaner.Scan(rs)

			if err != nil {
				return err
			}

			item := Top{}
			item = v

			n = n + 1

			if n > task.Limit {
				rm = append(rm, &item)
			} else {
				for fixedIndex < len(fixedSet) {
					p := fixedSet[fixedIndex]
					if p.Fixed <= n {
						p.Rank = p.Fixed
						items = append(items, p)
						fixedIndex = fixedIndex + 1
						n = n + 1
						if n > task.Limit {
							break
						}
					} else {
						break
					}
				}
				if n > task.Limit {
					rm = append(rm, &item)
				} else {
					item.Rank = n
					items = append(items, &item)
				}
			}

		}

		return nil
	}()

	if err != nil {
		return nil, err
	}

	err = db.Transaction(conn, func(conn db.Database) error {

		keys := map[string]bool{"rank": true}

		var err error = nil

		for _, item := range items {
			_, err = db.UpdateWithKeys(conn, item, prefix, keys)
			if err != nil {
				return err
			}
		}

		for fixedIndex < len(fixedSet) {
			p := fixedSet[fixedIndex]
			_, err = db.Delete(conn, p, prefix)
			if err != nil {
				return err
			}
			fixedIndex = fixedIndex + 1
		}

		for _, item := range rm {
			_, err = db.Delete(conn, item, prefix)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	{
		// 清除缓存

		redis, prefix, err := app.GetRedis("default")

		if err == nil {
			redis.Del(fmt.Sprintf("%s%s_rank", prefix, task.Name)).Result()
		}
	}

	return nil, nil
}
