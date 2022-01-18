// Code generated by entc, DO NOT EDIT.

package userinfo

import (
	"mall-go/app/user/service/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// DeleteTime applies equality check predicate on the "delete_time" field. It's identical to DeleteTimeEQ.
func DeleteTime(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// WxProfile applies equality check predicate on the "wx_profile" field. It's identical to WxProfileEQ.
func WxProfile(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWxProfile), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// DeleteTimeEQ applies the EQ predicate on the "delete_time" field.
func DeleteTimeEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeNEQ applies the NEQ predicate on the "delete_time" field.
func DeleteTimeNEQ(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIn applies the In predicate on the "delete_time" field.
func DeleteTimeIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeNotIn applies the NotIn predicate on the "delete_time" field.
func DeleteTimeNotIn(vs ...time.Time) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeleteTime), v...))
	})
}

// DeleteTimeGT applies the GT predicate on the "delete_time" field.
func DeleteTimeGT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeGTE applies the GTE predicate on the "delete_time" field.
func DeleteTimeGTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLT applies the LT predicate on the "delete_time" field.
func DeleteTimeLT(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeLTE applies the LTE predicate on the "delete_time" field.
func DeleteTimeLTE(v time.Time) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeleteTime), v))
	})
}

// DeleteTimeIsNil applies the IsNil predicate on the "delete_time" field.
func DeleteTimeIsNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeleteTime)))
	})
}

// DeleteTimeNotNil applies the NotNil predicate on the "delete_time" field.
func DeleteTimeNotNil() predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeleteTime)))
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNickname), v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNickname), v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNickname), v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNickname), v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNickname), v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNickname), v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNickname), v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNickname), v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNickname), v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNickname), v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNickname), v))
	})
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNickname), v))
	})
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNickname), v))
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	})
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	})
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	})
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAvatar), v))
	})
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAvatar), v))
	})
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAvatar), v...))
	})
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAvatar), v...))
	})
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAvatar), v))
	})
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAvatar), v))
	})
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAvatar), v))
	})
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAvatar), v))
	})
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAvatar), v))
	})
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAvatar), v))
	})
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAvatar), v))
	})
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAvatar), v))
	})
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAvatar), v))
	})
}

// WxProfileEQ applies the EQ predicate on the "wx_profile" field.
func WxProfileEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWxProfile), v))
	})
}

// WxProfileNEQ applies the NEQ predicate on the "wx_profile" field.
func WxProfileNEQ(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWxProfile), v))
	})
}

// WxProfileIn applies the In predicate on the "wx_profile" field.
func WxProfileIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWxProfile), v...))
	})
}

// WxProfileNotIn applies the NotIn predicate on the "wx_profile" field.
func WxProfileNotIn(vs ...string) predicate.UserInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWxProfile), v...))
	})
}

// WxProfileGT applies the GT predicate on the "wx_profile" field.
func WxProfileGT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWxProfile), v))
	})
}

// WxProfileGTE applies the GTE predicate on the "wx_profile" field.
func WxProfileGTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWxProfile), v))
	})
}

// WxProfileLT applies the LT predicate on the "wx_profile" field.
func WxProfileLT(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWxProfile), v))
	})
}

// WxProfileLTE applies the LTE predicate on the "wx_profile" field.
func WxProfileLTE(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWxProfile), v))
	})
}

// WxProfileContains applies the Contains predicate on the "wx_profile" field.
func WxProfileContains(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWxProfile), v))
	})
}

// WxProfileHasPrefix applies the HasPrefix predicate on the "wx_profile" field.
func WxProfileHasPrefix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWxProfile), v))
	})
}

// WxProfileHasSuffix applies the HasSuffix predicate on the "wx_profile" field.
func WxProfileHasSuffix(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWxProfile), v))
	})
}

// WxProfileEqualFold applies the EqualFold predicate on the "wx_profile" field.
func WxProfileEqualFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWxProfile), v))
	})
}

// WxProfileContainsFold applies the ContainsFold predicate on the "wx_profile" field.
func WxProfileContainsFold(v string) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWxProfile), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserInfo) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserInfo) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserInfo) predicate.UserInfo {
	return predicate.UserInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}