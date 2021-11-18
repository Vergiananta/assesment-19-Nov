package repository

import "strconv"

func Paginate(pages, pageSizes string) (int, int) {

	page, _ := strconv.Atoi(pages)
	if page == 0 {
		page = 1
	}

	pageSize, _ := strconv.Atoi(pageSizes)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	return offset, pageSize
}
