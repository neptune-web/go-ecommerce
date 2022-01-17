// Code generated by entc, DO NOT EDIT.

package model

import (
	"mall-go/app/sku/service/internal/data/ent/schema"
	"mall-go/app/sku/service/internal/data/model/brand"
	"mall-go/app/sku/service/internal/data/model/saleexplain"
	"mall-go/app/sku/service/internal/data/model/sku"
	"mall-go/app/sku/service/internal/data/model/speckey"
	"mall-go/app/sku/service/internal/data/model/specvalue"
	"mall-go/app/sku/service/internal/data/model/spu"
	"mall-go/app/sku/service/internal/data/model/spudetailimg"
	"mall-go/app/sku/service/internal/data/model/spuimg"
	"mall-go/app/sku/service/internal/data/model/tag"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	brandMixin := schema.Brand{}.Mixin()
	brandMixinFields0 := brandMixin[0].Fields()
	_ = brandMixinFields0
	brandFields := schema.Brand{}.Fields()
	_ = brandFields
	// brandDescCreateTime is the schema descriptor for create_time field.
	brandDescCreateTime := brandMixinFields0[0].Descriptor()
	// brand.DefaultCreateTime holds the default value on creation for the create_time field.
	brand.DefaultCreateTime = brandDescCreateTime.Default.(func() time.Time)
	// brandDescUpdateTime is the schema descriptor for update_time field.
	brandDescUpdateTime := brandMixinFields0[1].Descriptor()
	// brand.DefaultUpdateTime holds the default value on creation for the update_time field.
	brand.DefaultUpdateTime = brandDescUpdateTime.Default.(func() time.Time)
	// brand.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	brand.UpdateDefaultUpdateTime = brandDescUpdateTime.UpdateDefault.(func() time.Time)
	saleexplainMixin := schema.SaleExplain{}.Mixin()
	saleexplainMixinFields0 := saleexplainMixin[0].Fields()
	_ = saleexplainMixinFields0
	saleexplainFields := schema.SaleExplain{}.Fields()
	_ = saleexplainFields
	// saleexplainDescCreateTime is the schema descriptor for create_time field.
	saleexplainDescCreateTime := saleexplainMixinFields0[0].Descriptor()
	// saleexplain.DefaultCreateTime holds the default value on creation for the create_time field.
	saleexplain.DefaultCreateTime = saleexplainDescCreateTime.Default.(func() time.Time)
	// saleexplainDescUpdateTime is the schema descriptor for update_time field.
	saleexplainDescUpdateTime := saleexplainMixinFields0[1].Descriptor()
	// saleexplain.DefaultUpdateTime holds the default value on creation for the update_time field.
	saleexplain.DefaultUpdateTime = saleexplainDescUpdateTime.Default.(func() time.Time)
	// saleexplain.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	saleexplain.UpdateDefaultUpdateTime = saleexplainDescUpdateTime.UpdateDefault.(func() time.Time)
	skuMixin := schema.Sku{}.Mixin()
	skuMixinFields0 := skuMixin[0].Fields()
	_ = skuMixinFields0
	skuFields := schema.Sku{}.Fields()
	_ = skuFields
	// skuDescCreateTime is the schema descriptor for create_time field.
	skuDescCreateTime := skuMixinFields0[0].Descriptor()
	// sku.DefaultCreateTime holds the default value on creation for the create_time field.
	sku.DefaultCreateTime = skuDescCreateTime.Default.(func() time.Time)
	// skuDescUpdateTime is the schema descriptor for update_time field.
	skuDescUpdateTime := skuMixinFields0[1].Descriptor()
	// sku.DefaultUpdateTime holds the default value on creation for the update_time field.
	sku.DefaultUpdateTime = skuDescUpdateTime.Default.(func() time.Time)
	// sku.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	sku.UpdateDefaultUpdateTime = skuDescUpdateTime.UpdateDefault.(func() time.Time)
	speckeyMixin := schema.SpecKey{}.Mixin()
	speckeyMixinFields0 := speckeyMixin[0].Fields()
	_ = speckeyMixinFields0
	speckeyFields := schema.SpecKey{}.Fields()
	_ = speckeyFields
	// speckeyDescCreateTime is the schema descriptor for create_time field.
	speckeyDescCreateTime := speckeyMixinFields0[0].Descriptor()
	// speckey.DefaultCreateTime holds the default value on creation for the create_time field.
	speckey.DefaultCreateTime = speckeyDescCreateTime.Default.(func() time.Time)
	// speckeyDescUpdateTime is the schema descriptor for update_time field.
	speckeyDescUpdateTime := speckeyMixinFields0[1].Descriptor()
	// speckey.DefaultUpdateTime holds the default value on creation for the update_time field.
	speckey.DefaultUpdateTime = speckeyDescUpdateTime.Default.(func() time.Time)
	// speckey.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	speckey.UpdateDefaultUpdateTime = speckeyDescUpdateTime.UpdateDefault.(func() time.Time)
	specvalueMixin := schema.SpecValue{}.Mixin()
	specvalueMixinFields0 := specvalueMixin[0].Fields()
	_ = specvalueMixinFields0
	specvalueFields := schema.SpecValue{}.Fields()
	_ = specvalueFields
	// specvalueDescCreateTime is the schema descriptor for create_time field.
	specvalueDescCreateTime := specvalueMixinFields0[0].Descriptor()
	// specvalue.DefaultCreateTime holds the default value on creation for the create_time field.
	specvalue.DefaultCreateTime = specvalueDescCreateTime.Default.(func() time.Time)
	// specvalueDescUpdateTime is the schema descriptor for update_time field.
	specvalueDescUpdateTime := specvalueMixinFields0[1].Descriptor()
	// specvalue.DefaultUpdateTime holds the default value on creation for the update_time field.
	specvalue.DefaultUpdateTime = specvalueDescUpdateTime.Default.(func() time.Time)
	// specvalue.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	specvalue.UpdateDefaultUpdateTime = specvalueDescUpdateTime.UpdateDefault.(func() time.Time)
	spuMixin := schema.Spu{}.Mixin()
	spuMixinFields0 := spuMixin[0].Fields()
	_ = spuMixinFields0
	spuFields := schema.Spu{}.Fields()
	_ = spuFields
	// spuDescCreateTime is the schema descriptor for create_time field.
	spuDescCreateTime := spuMixinFields0[0].Descriptor()
	// spu.DefaultCreateTime holds the default value on creation for the create_time field.
	spu.DefaultCreateTime = spuDescCreateTime.Default.(func() time.Time)
	// spuDescUpdateTime is the schema descriptor for update_time field.
	spuDescUpdateTime := spuMixinFields0[1].Descriptor()
	// spu.DefaultUpdateTime holds the default value on creation for the update_time field.
	spu.DefaultUpdateTime = spuDescUpdateTime.Default.(func() time.Time)
	// spu.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	spu.UpdateDefaultUpdateTime = spuDescUpdateTime.UpdateDefault.(func() time.Time)
	spudetailimgMixin := schema.SpuDetailImg{}.Mixin()
	spudetailimgMixinFields0 := spudetailimgMixin[0].Fields()
	_ = spudetailimgMixinFields0
	spudetailimgFields := schema.SpuDetailImg{}.Fields()
	_ = spudetailimgFields
	// spudetailimgDescCreateTime is the schema descriptor for create_time field.
	spudetailimgDescCreateTime := spudetailimgMixinFields0[0].Descriptor()
	// spudetailimg.DefaultCreateTime holds the default value on creation for the create_time field.
	spudetailimg.DefaultCreateTime = spudetailimgDescCreateTime.Default.(func() time.Time)
	// spudetailimgDescUpdateTime is the schema descriptor for update_time field.
	spudetailimgDescUpdateTime := spudetailimgMixinFields0[1].Descriptor()
	// spudetailimg.DefaultUpdateTime holds the default value on creation for the update_time field.
	spudetailimg.DefaultUpdateTime = spudetailimgDescUpdateTime.Default.(func() time.Time)
	// spudetailimg.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	spudetailimg.UpdateDefaultUpdateTime = spudetailimgDescUpdateTime.UpdateDefault.(func() time.Time)
	spuimgMixin := schema.SpuImg{}.Mixin()
	spuimgMixinFields0 := spuimgMixin[0].Fields()
	_ = spuimgMixinFields0
	spuimgFields := schema.SpuImg{}.Fields()
	_ = spuimgFields
	// spuimgDescCreateTime is the schema descriptor for create_time field.
	spuimgDescCreateTime := spuimgMixinFields0[0].Descriptor()
	// spuimg.DefaultCreateTime holds the default value on creation for the create_time field.
	spuimg.DefaultCreateTime = spuimgDescCreateTime.Default.(func() time.Time)
	// spuimgDescUpdateTime is the schema descriptor for update_time field.
	spuimgDescUpdateTime := spuimgMixinFields0[1].Descriptor()
	// spuimg.DefaultUpdateTime holds the default value on creation for the update_time field.
	spuimg.DefaultUpdateTime = spuimgDescUpdateTime.Default.(func() time.Time)
	// spuimg.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	spuimg.UpdateDefaultUpdateTime = spuimgDescUpdateTime.UpdateDefault.(func() time.Time)
	tagMixin := schema.Tag{}.Mixin()
	tagMixinFields0 := tagMixin[0].Fields()
	_ = tagMixinFields0
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreateTime is the schema descriptor for create_time field.
	tagDescCreateTime := tagMixinFields0[0].Descriptor()
	// tag.DefaultCreateTime holds the default value on creation for the create_time field.
	tag.DefaultCreateTime = tagDescCreateTime.Default.(func() time.Time)
	// tagDescUpdateTime is the schema descriptor for update_time field.
	tagDescUpdateTime := tagMixinFields0[1].Descriptor()
	// tag.DefaultUpdateTime holds the default value on creation for the update_time field.
	tag.DefaultUpdateTime = tagDescUpdateTime.Default.(func() time.Time)
	// tag.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	tag.UpdateDefaultUpdateTime = tagDescUpdateTime.UpdateDefault.(func() time.Time)
}
