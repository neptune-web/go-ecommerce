// Code generated by entc, DO NOT EDIT.

package model

import (
	"mall-go/app/app/service/internal/data/ent/schema"
	"mall-go/app/app/service/internal/data/model/activity"
	"mall-go/app/app/service/internal/data/model/banner"
	"mall-go/app/app/service/internal/data/model/banneritem"
	"mall-go/app/app/service/internal/data/model/category"
	"mall-go/app/app/service/internal/data/model/charge"
	"mall-go/app/app/service/internal/data/model/gridcategory"
	"mall-go/app/app/service/internal/data/model/refund"
	"mall-go/app/app/service/internal/data/model/theme"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	activityMixin := schema.Activity{}.Mixin()
	activityMixinFields0 := activityMixin[0].Fields()
	_ = activityMixinFields0
	activityFields := schema.Activity{}.Fields()
	_ = activityFields
	// activityDescCreateTime is the schema descriptor for create_time field.
	activityDescCreateTime := activityMixinFields0[0].Descriptor()
	// activity.DefaultCreateTime holds the default value on creation for the create_time field.
	activity.DefaultCreateTime = activityDescCreateTime.Default.(func() time.Time)
	// activityDescUpdateTime is the schema descriptor for update_time field.
	activityDescUpdateTime := activityMixinFields0[1].Descriptor()
	// activity.DefaultUpdateTime holds the default value on creation for the update_time field.
	activity.DefaultUpdateTime = activityDescUpdateTime.Default.(func() time.Time)
	// activity.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	activity.UpdateDefaultUpdateTime = activityDescUpdateTime.UpdateDefault.(func() time.Time)
	bannerMixin := schema.Banner{}.Mixin()
	bannerMixinFields0 := bannerMixin[0].Fields()
	_ = bannerMixinFields0
	bannerFields := schema.Banner{}.Fields()
	_ = bannerFields
	// bannerDescCreateTime is the schema descriptor for create_time field.
	bannerDescCreateTime := bannerMixinFields0[0].Descriptor()
	// banner.DefaultCreateTime holds the default value on creation for the create_time field.
	banner.DefaultCreateTime = bannerDescCreateTime.Default.(func() time.Time)
	// bannerDescUpdateTime is the schema descriptor for update_time field.
	bannerDescUpdateTime := bannerMixinFields0[1].Descriptor()
	// banner.DefaultUpdateTime holds the default value on creation for the update_time field.
	banner.DefaultUpdateTime = bannerDescUpdateTime.Default.(func() time.Time)
	// banner.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	banner.UpdateDefaultUpdateTime = bannerDescUpdateTime.UpdateDefault.(func() time.Time)
	banneritemMixin := schema.BannerItem{}.Mixin()
	banneritemMixinFields0 := banneritemMixin[0].Fields()
	_ = banneritemMixinFields0
	banneritemFields := schema.BannerItem{}.Fields()
	_ = banneritemFields
	// banneritemDescCreateTime is the schema descriptor for create_time field.
	banneritemDescCreateTime := banneritemMixinFields0[0].Descriptor()
	// banneritem.DefaultCreateTime holds the default value on creation for the create_time field.
	banneritem.DefaultCreateTime = banneritemDescCreateTime.Default.(func() time.Time)
	// banneritemDescUpdateTime is the schema descriptor for update_time field.
	banneritemDescUpdateTime := banneritemMixinFields0[1].Descriptor()
	// banneritem.DefaultUpdateTime holds the default value on creation for the update_time field.
	banneritem.DefaultUpdateTime = banneritemDescUpdateTime.Default.(func() time.Time)
	// banneritem.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	banneritem.UpdateDefaultUpdateTime = banneritemDescUpdateTime.UpdateDefault.(func() time.Time)
	categoryMixin := schema.Category{}.Mixin()
	categoryMixinFields0 := categoryMixin[0].Fields()
	_ = categoryMixinFields0
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescCreateTime is the schema descriptor for create_time field.
	categoryDescCreateTime := categoryMixinFields0[0].Descriptor()
	// category.DefaultCreateTime holds the default value on creation for the create_time field.
	category.DefaultCreateTime = categoryDescCreateTime.Default.(func() time.Time)
	// categoryDescUpdateTime is the schema descriptor for update_time field.
	categoryDescUpdateTime := categoryMixinFields0[1].Descriptor()
	// category.DefaultUpdateTime holds the default value on creation for the update_time field.
	category.DefaultUpdateTime = categoryDescUpdateTime.Default.(func() time.Time)
	// category.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	category.UpdateDefaultUpdateTime = categoryDescUpdateTime.UpdateDefault.(func() time.Time)
	chargeMixin := schema.Charge{}.Mixin()
	chargeMixinFields0 := chargeMixin[0].Fields()
	_ = chargeMixinFields0
	chargeFields := schema.Charge{}.Fields()
	_ = chargeFields
	// chargeDescCreateTime is the schema descriptor for create_time field.
	chargeDescCreateTime := chargeMixinFields0[0].Descriptor()
	// charge.DefaultCreateTime holds the default value on creation for the create_time field.
	charge.DefaultCreateTime = chargeDescCreateTime.Default.(func() time.Time)
	// chargeDescUpdateTime is the schema descriptor for update_time field.
	chargeDescUpdateTime := chargeMixinFields0[1].Descriptor()
	// charge.DefaultUpdateTime holds the default value on creation for the update_time field.
	charge.DefaultUpdateTime = chargeDescUpdateTime.Default.(func() time.Time)
	// charge.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	charge.UpdateDefaultUpdateTime = chargeDescUpdateTime.UpdateDefault.(func() time.Time)
	gridcategoryMixin := schema.GridCategory{}.Mixin()
	gridcategoryMixinFields0 := gridcategoryMixin[0].Fields()
	_ = gridcategoryMixinFields0
	gridcategoryFields := schema.GridCategory{}.Fields()
	_ = gridcategoryFields
	// gridcategoryDescCreateTime is the schema descriptor for create_time field.
	gridcategoryDescCreateTime := gridcategoryMixinFields0[0].Descriptor()
	// gridcategory.DefaultCreateTime holds the default value on creation for the create_time field.
	gridcategory.DefaultCreateTime = gridcategoryDescCreateTime.Default.(func() time.Time)
	// gridcategoryDescUpdateTime is the schema descriptor for update_time field.
	gridcategoryDescUpdateTime := gridcategoryMixinFields0[1].Descriptor()
	// gridcategory.DefaultUpdateTime holds the default value on creation for the update_time field.
	gridcategory.DefaultUpdateTime = gridcategoryDescUpdateTime.Default.(func() time.Time)
	// gridcategory.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	gridcategory.UpdateDefaultUpdateTime = gridcategoryDescUpdateTime.UpdateDefault.(func() time.Time)
	refundMixin := schema.Refund{}.Mixin()
	refundMixinFields0 := refundMixin[0].Fields()
	_ = refundMixinFields0
	refundFields := schema.Refund{}.Fields()
	_ = refundFields
	// refundDescCreateTime is the schema descriptor for create_time field.
	refundDescCreateTime := refundMixinFields0[0].Descriptor()
	// refund.DefaultCreateTime holds the default value on creation for the create_time field.
	refund.DefaultCreateTime = refundDescCreateTime.Default.(func() time.Time)
	// refundDescUpdateTime is the schema descriptor for update_time field.
	refundDescUpdateTime := refundMixinFields0[1].Descriptor()
	// refund.DefaultUpdateTime holds the default value on creation for the update_time field.
	refund.DefaultUpdateTime = refundDescUpdateTime.Default.(func() time.Time)
	// refund.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	refund.UpdateDefaultUpdateTime = refundDescUpdateTime.UpdateDefault.(func() time.Time)
	themeMixin := schema.Theme{}.Mixin()
	themeMixinFields0 := themeMixin[0].Fields()
	_ = themeMixinFields0
	themeFields := schema.Theme{}.Fields()
	_ = themeFields
	// themeDescCreateTime is the schema descriptor for create_time field.
	themeDescCreateTime := themeMixinFields0[0].Descriptor()
	// theme.DefaultCreateTime holds the default value on creation for the create_time field.
	theme.DefaultCreateTime = themeDescCreateTime.Default.(func() time.Time)
	// themeDescUpdateTime is the schema descriptor for update_time field.
	themeDescUpdateTime := themeMixinFields0[1].Descriptor()
	// theme.DefaultUpdateTime holds the default value on creation for the update_time field.
	theme.DefaultUpdateTime = themeDescUpdateTime.Default.(func() time.Time)
	// theme.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	theme.UpdateDefaultUpdateTime = themeDescUpdateTime.UpdateDefault.(func() time.Time)
}
