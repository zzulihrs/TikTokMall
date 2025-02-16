// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package product

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *ListProductsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListProductsReq[number], err)
}

func (x *ListProductsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Page, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListProductsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *ListProductsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CategoryName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Product[number], err)
}

func (x *Product) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Product) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Picture, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *Product) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Categories = append(x.Categories, v)
	return offset, err
}

func (x *Product) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.Stock, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Product) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.OwnerId, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *Product) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.OwnerName, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Product) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.SliderImgs = append(x.SliderImgs, v)
	return offset, err
}

func (x *ListProductsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ListProductsResp[number], err)
}

func (x *ListProductsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Product
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Products = append(x.Products, &v)
	return offset, nil
}

func (x *GetProductReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetProductReq[number], err)
}

func (x *GetProductReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadUint32(buf, _type)
	return offset, err
}

func (x *GetProductResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetProductResp[number], err)
}

func (x *GetProductResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Product
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Product = &v
	return offset, nil
}

func (x *SearchProductsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchProductsReq[number], err)
}

func (x *SearchProductsReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Query, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchProductsResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SearchProductsResp[number], err)
}

func (x *SearchProductsResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v Product
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Results = append(x.Results, &v)
	return offset, nil
}

func (x *ListProductsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ListProductsReq) fastWriteField1(buf []byte) (offset int) {
	if x.Page == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.GetPage())
	return offset
}

func (x *ListProductsReq) fastWriteField2(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 2, x.GetPageSize())
	return offset
}

func (x *ListProductsReq) fastWriteField3(buf []byte) (offset int) {
	if x.CategoryName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetCategoryName())
	return offset
}

func (x *Product) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	return offset
}

func (x *Product) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Product) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *Product) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *Product) fastWriteField4(buf []byte) (offset int) {
	if x.Picture == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetPicture())
	return offset
}

func (x *Product) fastWriteField5(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 5, x.GetPrice())
	return offset
}

func (x *Product) fastWriteField6(buf []byte) (offset int) {
	if len(x.Categories) == 0 {
		return offset
	}
	for i := range x.GetCategories() {
		offset += fastpb.WriteString(buf[offset:], 6, x.GetCategories()[i])
	}
	return offset
}

func (x *Product) fastWriteField7(buf []byte) (offset int) {
	if x.Stock == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 7, x.GetStock())
	return offset
}

func (x *Product) fastWriteField8(buf []byte) (offset int) {
	if x.OwnerId == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 8, x.GetOwnerId())
	return offset
}

func (x *Product) fastWriteField9(buf []byte) (offset int) {
	if x.OwnerName == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 9, x.GetOwnerName())
	return offset
}

func (x *Product) fastWriteField10(buf []byte) (offset int) {
	if len(x.SliderImgs) == 0 {
		return offset
	}
	for i := range x.GetSliderImgs() {
		offset += fastpb.WriteString(buf[offset:], 10, x.GetSliderImgs()[i])
	}
	return offset
}

func (x *ListProductsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ListProductsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Products == nil {
		return offset
	}
	for i := range x.GetProducts() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetProducts()[i])
	}
	return offset
}

func (x *GetProductReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetProductReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteUint32(buf[offset:], 1, x.GetId())
	return offset
}

func (x *GetProductResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetProductResp) fastWriteField1(buf []byte) (offset int) {
	if x.Product == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetProduct())
	return offset
}

func (x *SearchProductsReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SearchProductsReq) fastWriteField1(buf []byte) (offset int) {
	if x.Query == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetQuery())
	return offset
}

func (x *SearchProductsResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *SearchProductsResp) fastWriteField1(buf []byte) (offset int) {
	if x.Results == nil {
		return offset
	}
	for i := range x.GetResults() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetResults()[i])
	}
	return offset
}

func (x *ListProductsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ListProductsReq) sizeField1() (n int) {
	if x.Page == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.GetPage())
	return n
}

func (x *ListProductsReq) sizeField2() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt32(2, x.GetPageSize())
	return n
}

func (x *ListProductsReq) sizeField3() (n int) {
	if x.CategoryName == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetCategoryName())
	return n
}

func (x *Product) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	return n
}

func (x *Product) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetId())
	return n
}

func (x *Product) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *Product) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *Product) sizeField4() (n int) {
	if x.Picture == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetPicture())
	return n
}

func (x *Product) sizeField5() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeFloat(5, x.GetPrice())
	return n
}

func (x *Product) sizeField6() (n int) {
	if len(x.Categories) == 0 {
		return n
	}
	for i := range x.GetCategories() {
		n += fastpb.SizeString(6, x.GetCategories()[i])
	}
	return n
}

func (x *Product) sizeField7() (n int) {
	if x.Stock == 0 {
		return n
	}
	n += fastpb.SizeUint32(7, x.GetStock())
	return n
}

func (x *Product) sizeField8() (n int) {
	if x.OwnerId == 0 {
		return n
	}
	n += fastpb.SizeUint32(8, x.GetOwnerId())
	return n
}

func (x *Product) sizeField9() (n int) {
	if x.OwnerName == "" {
		return n
	}
	n += fastpb.SizeString(9, x.GetOwnerName())
	return n
}

func (x *Product) sizeField10() (n int) {
	if len(x.SliderImgs) == 0 {
		return n
	}
	for i := range x.GetSliderImgs() {
		n += fastpb.SizeString(10, x.GetSliderImgs()[i])
	}
	return n
}

func (x *ListProductsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ListProductsResp) sizeField1() (n int) {
	if x.Products == nil {
		return n
	}
	for i := range x.GetProducts() {
		n += fastpb.SizeMessage(1, x.GetProducts()[i])
	}
	return n
}

func (x *GetProductReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetProductReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeUint32(1, x.GetId())
	return n
}

func (x *GetProductResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetProductResp) sizeField1() (n int) {
	if x.Product == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetProduct())
	return n
}

func (x *SearchProductsReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SearchProductsReq) sizeField1() (n int) {
	if x.Query == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetQuery())
	return n
}

func (x *SearchProductsResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *SearchProductsResp) sizeField1() (n int) {
	if x.Results == nil {
		return n
	}
	for i := range x.GetResults() {
		n += fastpb.SizeMessage(1, x.GetResults()[i])
	}
	return n
}

var fieldIDToName_ListProductsReq = map[int32]string{
	1: "Page",
	2: "PageSize",
	3: "CategoryName",
}

var fieldIDToName_Product = map[int32]string{
	1:  "Id",
	2:  "Name",
	3:  "Description",
	4:  "Picture",
	5:  "Price",
	6:  "Categories",
	7:  "Stock",
	8:  "OwnerId",
	9:  "OwnerName",
	10: "SliderImgs",
}

var fieldIDToName_ListProductsResp = map[int32]string{
	1: "Products",
}

var fieldIDToName_GetProductReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_GetProductResp = map[int32]string{
	1: "Product",
}

var fieldIDToName_SearchProductsReq = map[int32]string{
	1: "Query",
}

var fieldIDToName_SearchProductsResp = map[int32]string{
	1: "Results",
}
