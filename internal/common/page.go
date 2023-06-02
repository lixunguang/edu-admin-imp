// package common some explain here.
package common

import (
	"math"
)

// Page some explain herer...
type Page struct {
	CurrentPage int64   //当前页码
	PageSize    int64   //每页条数
	TotalPage   int64   //总条数
	Nums        []int64 //分页序数
	PageCount   int64   //总页数
}

func (_ Page) CreatePageInfo(currentPage, pageSize, total int64) Page {
	if currentPage < 1 {
		currentPage = 1
	}
	if pageSize < 1 {
		pageSize = 1
	}
	this := Page{}
	this.CurrentPage = currentPage
	this.PageSize = pageSize
	this.TotalPage = total
	this.PageCount = int64(math.Ceil((float64(total)) / float64(pageSize)))
	this.setNum()
	return this
}

//设置序号
func (this *Page) setNum() {
	this.Nums = []int64{}
	if this.PageCount == 0 {
		return
	}
	var begin int64 = 1
	var end int64 = 5
	if this.PageCount <= 5 {
		end = this.PageCount
	} else {
		begin = this.CurrentPage - 2
		end = this.CurrentPage + 2
		if begin <= 0 {
			begin = 1
			end = 5
		}
		if end >= this.PageCount {
			end = this.PageCount
			begin = end - 4
		}
	}
	for i := begin; i <= end; i++ {
		this.Nums = append(this.Nums, i)
	}
}
