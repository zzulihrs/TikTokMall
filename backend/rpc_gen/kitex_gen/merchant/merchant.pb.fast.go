// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package merchant

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *Category) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Category[number], err)
}

func (x *Category) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Category) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Category) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MerchantProductSimpleInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MerchantProductSimpleInfo[number], err)
}

func (x *MerchantProductSimpleInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *MerchantProductSimpleInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MerchantProductSimpleInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MerchantProductSimpleInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Stock, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *MerchantProductSimpleInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *MerchantProductSimpleInfo) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.ImgUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MerchantProductDetailInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MerchantProductDetailInfo[number], err)
}

func (x *MerchantProductDetailInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *MerchantProductDetailInfo) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MerchantProductDetailInfo) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Description, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MerchantProductDetailInfo) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Stock, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *MerchantProductDetailInfo) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.Price, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *MerchantProductDetailInfo) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.ImgUrl, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MerchantProductDetailInfo) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	var v Category
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Category = append(x.Category, &v)
	return offset, nil
}

func (x *MerchantProductDetailInfo) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	var v string
	v, offset, err = fastpb.ReadString(buf, _type)
	if err != nil {
		return offset, err
	}
	x.SliderImgs = append(x.SliderImgs, v)
	return offset, err
}

func (x *GetMerchantReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMerchantReq[number], err)
}

func (x *GetMerchantReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetMerchantResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetMerchantResp[number], err)
}

func (x *GetMerchantResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetMerchantResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetMerchantResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Email, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AddProductReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AddProductReq[number], err)
}

func (x *AddProductReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MerchantId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AddProductReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v MerchantProductDetailInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Product = &v
	return offset, nil
}

func (x *AddProductResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *DeleteProductReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteProductReq[number], err)
}

func (x *DeleteProductReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MerchantId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteProductReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Pid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteProductResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *UpdateProductReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UpdateProductReq[number], err)
}

func (x *UpdateProductReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MerchantId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UpdateProductReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v MerchantProductDetailInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Product = &v
	return offset, nil
}

func (x *UpdateProductResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *SearchProductsReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SearchProductsReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.CategoryId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SearchProductsReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.MaxStock, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *SearchProductsReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.MinPrice, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *SearchProductsReq) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.MaxPrice, offset, err = fastpb.ReadFloat(buf, _type)
	return offset, err
}

func (x *SearchProductsReq) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.PageNo, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *SearchProductsReq) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.PageSize, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *SearchProductsReq) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.MerchantId, offset, err = fastpb.ReadInt64(buf, _type)
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
	var v MerchantProductSimpleInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Products = append(x.Products, &v)
	return offset, nil
}

func (x *ProductDetailReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductDetailReq[number], err)
}

func (x *ProductDetailReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.MerchantId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ProductDetailReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Pid, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ProductDetailResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ProductDetailResp[number], err)
}

func (x *ProductDetailResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v MerchantProductDetailInfo
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Product = &v
	return offset, nil
}

func (x *Category) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *Category) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *Category) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *Category) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *MerchantProductSimpleInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *MerchantProductSimpleInfo) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *MerchantProductSimpleInfo) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *MerchantProductSimpleInfo) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *MerchantProductSimpleInfo) fastWriteField4(buf []byte) (offset int) {
	if x.Stock == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.GetStock())
	return offset
}

func (x *MerchantProductSimpleInfo) fastWriteField5(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 5, x.GetPrice())
	return offset
}

func (x *MerchantProductSimpleInfo) fastWriteField6(buf []byte) (offset int) {
	if x.ImgUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetImgUrl())
	return offset
}

func (x *MerchantProductDetailInfo) FastWrite(buf []byte) (offset int) {
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
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetName())
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField3(buf []byte) (offset int) {
	if x.Description == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDescription())
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField4(buf []byte) (offset int) {
	if x.Stock == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 4, x.GetStock())
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField5(buf []byte) (offset int) {
	if x.Price == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 5, x.GetPrice())
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField6(buf []byte) (offset int) {
	if x.ImgUrl == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.GetImgUrl())
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField7(buf []byte) (offset int) {
	if x.Category == nil {
		return offset
	}
	for i := range x.GetCategory() {
		offset += fastpb.WriteMessage(buf[offset:], 7, x.GetCategory()[i])
	}
	return offset
}

func (x *MerchantProductDetailInfo) fastWriteField8(buf []byte) (offset int) {
	if len(x.SliderImgs) == 0 {
		return offset
	}
	for i := range x.GetSliderImgs() {
		offset += fastpb.WriteString(buf[offset:], 8, x.GetSliderImgs()[i])
	}
	return offset
}

func (x *GetMerchantReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetMerchantReq) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *GetMerchantResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *GetMerchantResp) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetId())
	return offset
}

func (x *GetMerchantResp) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUsername())
	return offset
}

func (x *GetMerchantResp) fastWriteField3(buf []byte) (offset int) {
	if x.Email == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetEmail())
	return offset
}

func (x *AddProductReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AddProductReq) fastWriteField1(buf []byte) (offset int) {
	if x.MerchantId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetMerchantId())
	return offset
}

func (x *AddProductReq) fastWriteField2(buf []byte) (offset int) {
	if x.Product == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetProduct())
	return offset
}

func (x *AddProductResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *DeleteProductReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteProductReq) fastWriteField1(buf []byte) (offset int) {
	if x.MerchantId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetMerchantId())
	return offset
}

func (x *DeleteProductReq) fastWriteField2(buf []byte) (offset int) {
	if x.Pid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetPid())
	return offset
}

func (x *DeleteProductResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *UpdateProductReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *UpdateProductReq) fastWriteField1(buf []byte) (offset int) {
	if x.MerchantId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetMerchantId())
	return offset
}

func (x *UpdateProductReq) fastWriteField2(buf []byte) (offset int) {
	if x.Product == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetProduct())
	return offset
}

func (x *UpdateProductResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *SearchProductsReq) FastWrite(buf []byte) (offset int) {
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
	return offset
}

func (x *SearchProductsReq) fastWriteField1(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetName())
	return offset
}

func (x *SearchProductsReq) fastWriteField2(buf []byte) (offset int) {
	if x.CategoryId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetCategoryId())
	return offset
}

func (x *SearchProductsReq) fastWriteField3(buf []byte) (offset int) {
	if x.MaxStock == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 3, x.GetMaxStock())
	return offset
}

func (x *SearchProductsReq) fastWriteField4(buf []byte) (offset int) {
	if x.MinPrice == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 4, x.GetMinPrice())
	return offset
}

func (x *SearchProductsReq) fastWriteField5(buf []byte) (offset int) {
	if x.MaxPrice == 0 {
		return offset
	}
	offset += fastpb.WriteFloat(buf[offset:], 5, x.GetMaxPrice())
	return offset
}

func (x *SearchProductsReq) fastWriteField6(buf []byte) (offset int) {
	if x.PageNo == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 6, x.GetPageNo())
	return offset
}

func (x *SearchProductsReq) fastWriteField7(buf []byte) (offset int) {
	if x.PageSize == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 7, x.GetPageSize())
	return offset
}

func (x *SearchProductsReq) fastWriteField8(buf []byte) (offset int) {
	if x.MerchantId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 8, x.GetMerchantId())
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
	if x.Products == nil {
		return offset
	}
	for i := range x.GetProducts() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetProducts()[i])
	}
	return offset
}

func (x *ProductDetailReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ProductDetailReq) fastWriteField1(buf []byte) (offset int) {
	if x.MerchantId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetMerchantId())
	return offset
}

func (x *ProductDetailReq) fastWriteField2(buf []byte) (offset int) {
	if x.Pid == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetPid())
	return offset
}

func (x *ProductDetailResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *ProductDetailResp) fastWriteField1(buf []byte) (offset int) {
	if x.Product == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetProduct())
	return offset
}

func (x *Category) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *Category) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *Category) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *Category) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *MerchantProductSimpleInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *MerchantProductSimpleInfo) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *MerchantProductSimpleInfo) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *MerchantProductSimpleInfo) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *MerchantProductSimpleInfo) sizeField4() (n int) {
	if x.Stock == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.GetStock())
	return n
}

func (x *MerchantProductSimpleInfo) sizeField5() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeFloat(5, x.GetPrice())
	return n
}

func (x *MerchantProductSimpleInfo) sizeField6() (n int) {
	if x.ImgUrl == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetImgUrl())
	return n
}

func (x *MerchantProductDetailInfo) Size() (n int) {
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
	return n
}

func (x *MerchantProductDetailInfo) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *MerchantProductDetailInfo) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetName())
	return n
}

func (x *MerchantProductDetailInfo) sizeField3() (n int) {
	if x.Description == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDescription())
	return n
}

func (x *MerchantProductDetailInfo) sizeField4() (n int) {
	if x.Stock == 0 {
		return n
	}
	n += fastpb.SizeInt32(4, x.GetStock())
	return n
}

func (x *MerchantProductDetailInfo) sizeField5() (n int) {
	if x.Price == 0 {
		return n
	}
	n += fastpb.SizeFloat(5, x.GetPrice())
	return n
}

func (x *MerchantProductDetailInfo) sizeField6() (n int) {
	if x.ImgUrl == "" {
		return n
	}
	n += fastpb.SizeString(6, x.GetImgUrl())
	return n
}

func (x *MerchantProductDetailInfo) sizeField7() (n int) {
	if x.Category == nil {
		return n
	}
	for i := range x.GetCategory() {
		n += fastpb.SizeMessage(7, x.GetCategory()[i])
	}
	return n
}

func (x *MerchantProductDetailInfo) sizeField8() (n int) {
	if len(x.SliderImgs) == 0 {
		return n
	}
	for i := range x.GetSliderImgs() {
		n += fastpb.SizeString(8, x.GetSliderImgs()[i])
	}
	return n
}

func (x *GetMerchantReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetMerchantReq) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *GetMerchantResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *GetMerchantResp) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetId())
	return n
}

func (x *GetMerchantResp) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUsername())
	return n
}

func (x *GetMerchantResp) sizeField3() (n int) {
	if x.Email == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetEmail())
	return n
}

func (x *AddProductReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AddProductReq) sizeField1() (n int) {
	if x.MerchantId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetMerchantId())
	return n
}

func (x *AddProductReq) sizeField2() (n int) {
	if x.Product == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetProduct())
	return n
}

func (x *AddProductResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *DeleteProductReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteProductReq) sizeField1() (n int) {
	if x.MerchantId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetMerchantId())
	return n
}

func (x *DeleteProductReq) sizeField2() (n int) {
	if x.Pid == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetPid())
	return n
}

func (x *DeleteProductResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *UpdateProductReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *UpdateProductReq) sizeField1() (n int) {
	if x.MerchantId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetMerchantId())
	return n
}

func (x *UpdateProductReq) sizeField2() (n int) {
	if x.Product == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetProduct())
	return n
}

func (x *UpdateProductResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *SearchProductsReq) Size() (n int) {
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
	return n
}

func (x *SearchProductsReq) sizeField1() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetName())
	return n
}

func (x *SearchProductsReq) sizeField2() (n int) {
	if x.CategoryId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetCategoryId())
	return n
}

func (x *SearchProductsReq) sizeField3() (n int) {
	if x.MaxStock == 0 {
		return n
	}
	n += fastpb.SizeInt32(3, x.GetMaxStock())
	return n
}

func (x *SearchProductsReq) sizeField4() (n int) {
	if x.MinPrice == 0 {
		return n
	}
	n += fastpb.SizeFloat(4, x.GetMinPrice())
	return n
}

func (x *SearchProductsReq) sizeField5() (n int) {
	if x.MaxPrice == 0 {
		return n
	}
	n += fastpb.SizeFloat(5, x.GetMaxPrice())
	return n
}

func (x *SearchProductsReq) sizeField6() (n int) {
	if x.PageNo == 0 {
		return n
	}
	n += fastpb.SizeInt32(6, x.GetPageNo())
	return n
}

func (x *SearchProductsReq) sizeField7() (n int) {
	if x.PageSize == 0 {
		return n
	}
	n += fastpb.SizeInt32(7, x.GetPageSize())
	return n
}

func (x *SearchProductsReq) sizeField8() (n int) {
	if x.MerchantId == 0 {
		return n
	}
	n += fastpb.SizeInt64(8, x.GetMerchantId())
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
	if x.Products == nil {
		return n
	}
	for i := range x.GetProducts() {
		n += fastpb.SizeMessage(1, x.GetProducts()[i])
	}
	return n
}

func (x *ProductDetailReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ProductDetailReq) sizeField1() (n int) {
	if x.MerchantId == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetMerchantId())
	return n
}

func (x *ProductDetailReq) sizeField2() (n int) {
	if x.Pid == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetPid())
	return n
}

func (x *ProductDetailResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *ProductDetailResp) sizeField1() (n int) {
	if x.Product == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetProduct())
	return n
}

var fieldIDToName_Category = map[int32]string{
	1: "Id",
	2: "Name",
	3: "Description",
}

var fieldIDToName_MerchantProductSimpleInfo = map[int32]string{
	1: "Id",
	2: "Name",
	3: "Description",
	4: "Stock",
	5: "Price",
	6: "ImgUrl",
}

var fieldIDToName_MerchantProductDetailInfo = map[int32]string{
	1: "Id",
	2: "Name",
	3: "Description",
	4: "Stock",
	5: "Price",
	6: "ImgUrl",
	7: "Category",
	8: "SliderImgs",
}

var fieldIDToName_GetMerchantReq = map[int32]string{
	1: "Id",
}

var fieldIDToName_GetMerchantResp = map[int32]string{
	1: "Id",
	2: "Username",
	3: "Email",
}

var fieldIDToName_AddProductReq = map[int32]string{
	1: "MerchantId",
	2: "Product",
}

var fieldIDToName_AddProductResp = map[int32]string{}

var fieldIDToName_DeleteProductReq = map[int32]string{
	1: "MerchantId",
	2: "Pid",
}

var fieldIDToName_DeleteProductResp = map[int32]string{}

var fieldIDToName_UpdateProductReq = map[int32]string{
	1: "MerchantId",
	2: "Product",
}

var fieldIDToName_UpdateProductResp = map[int32]string{}

var fieldIDToName_SearchProductsReq = map[int32]string{
	1: "Name",
	2: "CategoryId",
	3: "MaxStock",
	4: "MinPrice",
	5: "MaxPrice",
	6: "PageNo",
	7: "PageSize",
	8: "MerchantId",
}

var fieldIDToName_SearchProductsResp = map[int32]string{
	1: "Products",
}

var fieldIDToName_ProductDetailReq = map[int32]string{
	1: "MerchantId",
	2: "Pid",
}

var fieldIDToName_ProductDetailResp = map[int32]string{
	1: "Product",
}
