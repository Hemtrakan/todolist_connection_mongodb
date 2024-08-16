package utils

import (
	"fmt"
	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func PaginateQuery(orderBy *string, pageSkip, pageLimit int) (str string) {
	if orderBy == nil {
		order := "1"
		orderBy = &order
	}

	str = fmt.Sprintf("\nORDER BY %v\n OFFSET (%v-1)*%v ROWS FETCH NEXT %v ROWS ONLY", *orderBy, pageSkip, pageLimit, pageLimit)
	return
}
