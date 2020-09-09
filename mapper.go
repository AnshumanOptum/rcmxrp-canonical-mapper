package mapper

import (
	"bytes"
	"log"
	"strconv"

	model "github.optum.com/AnshumanOptum/rcmxrp-canonical-mapper/model/generated"
	ebaasmapper "github.optum.com/ebaas/ebaas-mapper-common-golang/pkg/mapper"
)

type Rcmxrp_canonical_demo_mapper struct{}

func (m *Rcmxrp_canonical_demo_mapper) CanonicalName() string {
	return "rcmxrpCanonicalDemo"
}

func (m *Rcmxrp_canonical_demo_mapper) KeySchema() string {
	return model.NewRcmxrpCanonicalKey().Schema()
}

func (m *Employee_canonical_demo_mapper) ValueSchema() string {
	return model.NewRcmxrpCanonicalValue().Schema()
}

func (m *Rcmxrp_canonical_demo_mapper) MapKey(keyData map[string]interface{}) (bytes.Buffer, error) {
	result := model.RcmxrpCanonicalKey{}
	result.EnrolleeIndividualIdentifier = NewStringUnion(keyData["ADI9T3"].(string))

	var keyBuf bytes.Buffer
	err := result.Serialize(&keyBuf)
	return keyBuf, err
}

func (m *Rcmpxrp_canonical_demo_mapper) MapValue(keyData map[string]interface{}, rowMap ebaasmapper.RowMap) (bytes.Buffer, error) {
	var valueBuf bytes.Buffer

	result := model.RcmxrpCanonicalValue{
		EnrolleeIndividualIdentifier:                    &model.UnionNullString{},
		EnrolleeIndividualCrossReferenceListName:        &model.UnionNullString{},
		EnrolleeIndividualReferenceIdentifier:           &model.UnionNullInt{},
		EnrolleeIndividualReferenceCreateUserIdentifier: &model.UnionNullString{},
		EnrolleeIndividualReferenceCreateDate:           &model.UnionNullString{},
		EnrolleeIndividualReferenceCreateTime:           &model.UnionNullInt{},
	}

	if len(rowMap.GetRows("rcmxrp")) > 0 {
		rcmxrpRow, err := rowMap.GetRows("rcmxrp")[0].Value()

		if err != nil {
			return valueBuf, err
		}

		result.EnrolleeIndividualIdentifier = NewStringUnion(rcmxrpRow["ADI9T3"].(string))
		result.EnrolleeIndividualCrossReferenceListName = NewStringUnion(rcmxrpRow["ADZQHO"].(string))
		result.EnrolleeIndividualReferenceIdentifier = NewIntUnion(ConvertToInt(rcmxrpRow["ADTQC3"]))
		result.EnrolleeIndividualReferenceCreateUserIdentifier = NewStringUnion(rcmxrpRow["ADAKVN"].(string))
		result.EnrolleeIndividualReferenceCreateDate = NewStringUnion(rcmxrpRow["ADC2DT"].(string))
		result.EnrolleeIndividualReferenceCreateTime = NewStringUnion(rcmxrpRow["ADADTM"].(string))
	} else {
		return valueBuf, nil
	}

	err := result.Serialize(&valueBuf)
	return valueBuf, err
}

func ConvertToInt(value interface{}) (i int32) {
	switch v := value.(type) {
	case float64:
		i = int32(v)
	case int:
		i = int32(v)
	case int64:
		i = int32(v)
	case int32:
		i = int32(v)
	case string:
		i = convertToInt32(v)
	}

	return i
}

func convertToInt32(val string) (i int32) {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalln("Can't convert string to int. v = ", val)
	}
	return int32(intVal)
}

func convertToLong(value interface{}) (i int64) {
	switch v := value.(type) {
	case int64:
		i = int64(v)
	case string:
		i = convertToInt64(v)
	}
	return i
}

func convertToInt64(val string) (i int64) {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalln("Can't convert string to int. v = ", val)
	}
	return int64(intVal)
}

func NewStringUnion(s string) *model.UnionNullString {

	return &model.UnionNullString{
		String:    s,
		UnionType: model.UnionNullStringTypeEnumString,
	}
}

func NewIntUnion(i int32) *model.UnionNullInt {
	return &model.UnionNullInt{
		Int:       i,
		UnionType: model.UnionNullIntTypeEnumInt,
	}
}