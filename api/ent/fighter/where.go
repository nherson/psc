// Code generated by ent, DO NOT EDIT.

package fighter

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nherson/psc/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldUpdatedAt, v))
}

// UfcFighterID applies equality check predicate on the "ufc_fighter_id" field. It's identical to UfcFighterIDEQ.
func UfcFighterID(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldUfcFighterID, v))
}

// MmaID applies equality check predicate on the "mma_id" field. It's identical to MmaIDEQ.
func MmaID(v int) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldMmaID, v))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldLastName, v))
}

// NickName applies equality check predicate on the "nick_name" field. It's identical to NickNameEQ.
func NickName(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldNickName, v))
}

// FightinsiderID applies equality check predicate on the "fightinsider_id" field. It's identical to FightinsiderIDEQ.
func FightinsiderID(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldFightinsiderID, v))
}

// TapologyID applies equality check predicate on the "tapology_id" field. It's identical to TapologyIDEQ.
func TapologyID(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldTapologyID, v))
}

// Temporary applies equality check predicate on the "temporary" field. It's identical to TemporaryEQ.
func Temporary(v bool) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldTemporary, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldUpdatedAt, v))
}

// UfcFighterIDEQ applies the EQ predicate on the "ufc_fighter_id" field.
func UfcFighterIDEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldUfcFighterID, v))
}

// UfcFighterIDNEQ applies the NEQ predicate on the "ufc_fighter_id" field.
func UfcFighterIDNEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldUfcFighterID, v))
}

// UfcFighterIDIn applies the In predicate on the "ufc_fighter_id" field.
func UfcFighterIDIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldUfcFighterID, vs...))
}

// UfcFighterIDNotIn applies the NotIn predicate on the "ufc_fighter_id" field.
func UfcFighterIDNotIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldUfcFighterID, vs...))
}

// UfcFighterIDGT applies the GT predicate on the "ufc_fighter_id" field.
func UfcFighterIDGT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldUfcFighterID, v))
}

// UfcFighterIDGTE applies the GTE predicate on the "ufc_fighter_id" field.
func UfcFighterIDGTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldUfcFighterID, v))
}

// UfcFighterIDLT applies the LT predicate on the "ufc_fighter_id" field.
func UfcFighterIDLT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldUfcFighterID, v))
}

// UfcFighterIDLTE applies the LTE predicate on the "ufc_fighter_id" field.
func UfcFighterIDLTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldUfcFighterID, v))
}

// UfcFighterIDContains applies the Contains predicate on the "ufc_fighter_id" field.
func UfcFighterIDContains(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContains(FieldUfcFighterID, v))
}

// UfcFighterIDHasPrefix applies the HasPrefix predicate on the "ufc_fighter_id" field.
func UfcFighterIDHasPrefix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasPrefix(FieldUfcFighterID, v))
}

// UfcFighterIDHasSuffix applies the HasSuffix predicate on the "ufc_fighter_id" field.
func UfcFighterIDHasSuffix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasSuffix(FieldUfcFighterID, v))
}

// UfcFighterIDEqualFold applies the EqualFold predicate on the "ufc_fighter_id" field.
func UfcFighterIDEqualFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEqualFold(FieldUfcFighterID, v))
}

// UfcFighterIDContainsFold applies the ContainsFold predicate on the "ufc_fighter_id" field.
func UfcFighterIDContainsFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContainsFold(FieldUfcFighterID, v))
}

// MmaIDEQ applies the EQ predicate on the "mma_id" field.
func MmaIDEQ(v int) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldMmaID, v))
}

// MmaIDNEQ applies the NEQ predicate on the "mma_id" field.
func MmaIDNEQ(v int) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldMmaID, v))
}

// MmaIDIn applies the In predicate on the "mma_id" field.
func MmaIDIn(vs ...int) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldMmaID, vs...))
}

// MmaIDNotIn applies the NotIn predicate on the "mma_id" field.
func MmaIDNotIn(vs ...int) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldMmaID, vs...))
}

// MmaIDGT applies the GT predicate on the "mma_id" field.
func MmaIDGT(v int) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldMmaID, v))
}

// MmaIDGTE applies the GTE predicate on the "mma_id" field.
func MmaIDGTE(v int) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldMmaID, v))
}

// MmaIDLT applies the LT predicate on the "mma_id" field.
func MmaIDLT(v int) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldMmaID, v))
}

// MmaIDLTE applies the LTE predicate on the "mma_id" field.
func MmaIDLTE(v int) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldMmaID, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContainsFold(FieldLastName, v))
}

// NickNameEQ applies the EQ predicate on the "nick_name" field.
func NickNameEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldNickName, v))
}

// NickNameNEQ applies the NEQ predicate on the "nick_name" field.
func NickNameNEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldNickName, v))
}

// NickNameIn applies the In predicate on the "nick_name" field.
func NickNameIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldNickName, vs...))
}

// NickNameNotIn applies the NotIn predicate on the "nick_name" field.
func NickNameNotIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldNickName, vs...))
}

// NickNameGT applies the GT predicate on the "nick_name" field.
func NickNameGT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldNickName, v))
}

// NickNameGTE applies the GTE predicate on the "nick_name" field.
func NickNameGTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldNickName, v))
}

// NickNameLT applies the LT predicate on the "nick_name" field.
func NickNameLT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldNickName, v))
}

// NickNameLTE applies the LTE predicate on the "nick_name" field.
func NickNameLTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldNickName, v))
}

// NickNameContains applies the Contains predicate on the "nick_name" field.
func NickNameContains(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContains(FieldNickName, v))
}

// NickNameHasPrefix applies the HasPrefix predicate on the "nick_name" field.
func NickNameHasPrefix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasPrefix(FieldNickName, v))
}

// NickNameHasSuffix applies the HasSuffix predicate on the "nick_name" field.
func NickNameHasSuffix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasSuffix(FieldNickName, v))
}

// NickNameEqualFold applies the EqualFold predicate on the "nick_name" field.
func NickNameEqualFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEqualFold(FieldNickName, v))
}

// NickNameContainsFold applies the ContainsFold predicate on the "nick_name" field.
func NickNameContainsFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContainsFold(FieldNickName, v))
}

// FightinsiderIDEQ applies the EQ predicate on the "fightinsider_id" field.
func FightinsiderIDEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldFightinsiderID, v))
}

// FightinsiderIDNEQ applies the NEQ predicate on the "fightinsider_id" field.
func FightinsiderIDNEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldFightinsiderID, v))
}

// FightinsiderIDIn applies the In predicate on the "fightinsider_id" field.
func FightinsiderIDIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldFightinsiderID, vs...))
}

// FightinsiderIDNotIn applies the NotIn predicate on the "fightinsider_id" field.
func FightinsiderIDNotIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldFightinsiderID, vs...))
}

// FightinsiderIDGT applies the GT predicate on the "fightinsider_id" field.
func FightinsiderIDGT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldFightinsiderID, v))
}

// FightinsiderIDGTE applies the GTE predicate on the "fightinsider_id" field.
func FightinsiderIDGTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldFightinsiderID, v))
}

// FightinsiderIDLT applies the LT predicate on the "fightinsider_id" field.
func FightinsiderIDLT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldFightinsiderID, v))
}

// FightinsiderIDLTE applies the LTE predicate on the "fightinsider_id" field.
func FightinsiderIDLTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldFightinsiderID, v))
}

// FightinsiderIDContains applies the Contains predicate on the "fightinsider_id" field.
func FightinsiderIDContains(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContains(FieldFightinsiderID, v))
}

// FightinsiderIDHasPrefix applies the HasPrefix predicate on the "fightinsider_id" field.
func FightinsiderIDHasPrefix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasPrefix(FieldFightinsiderID, v))
}

// FightinsiderIDHasSuffix applies the HasSuffix predicate on the "fightinsider_id" field.
func FightinsiderIDHasSuffix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasSuffix(FieldFightinsiderID, v))
}

// FightinsiderIDIsNil applies the IsNil predicate on the "fightinsider_id" field.
func FightinsiderIDIsNil() predicate.Fighter {
	return predicate.Fighter(sql.FieldIsNull(FieldFightinsiderID))
}

// FightinsiderIDNotNil applies the NotNil predicate on the "fightinsider_id" field.
func FightinsiderIDNotNil() predicate.Fighter {
	return predicate.Fighter(sql.FieldNotNull(FieldFightinsiderID))
}

// FightinsiderIDEqualFold applies the EqualFold predicate on the "fightinsider_id" field.
func FightinsiderIDEqualFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEqualFold(FieldFightinsiderID, v))
}

// FightinsiderIDContainsFold applies the ContainsFold predicate on the "fightinsider_id" field.
func FightinsiderIDContainsFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContainsFold(FieldFightinsiderID, v))
}

// TapologyIDEQ applies the EQ predicate on the "tapology_id" field.
func TapologyIDEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldTapologyID, v))
}

// TapologyIDNEQ applies the NEQ predicate on the "tapology_id" field.
func TapologyIDNEQ(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldTapologyID, v))
}

// TapologyIDIn applies the In predicate on the "tapology_id" field.
func TapologyIDIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldIn(FieldTapologyID, vs...))
}

// TapologyIDNotIn applies the NotIn predicate on the "tapology_id" field.
func TapologyIDNotIn(vs ...string) predicate.Fighter {
	return predicate.Fighter(sql.FieldNotIn(FieldTapologyID, vs...))
}

// TapologyIDGT applies the GT predicate on the "tapology_id" field.
func TapologyIDGT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGT(FieldTapologyID, v))
}

// TapologyIDGTE applies the GTE predicate on the "tapology_id" field.
func TapologyIDGTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldGTE(FieldTapologyID, v))
}

// TapologyIDLT applies the LT predicate on the "tapology_id" field.
func TapologyIDLT(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLT(FieldTapologyID, v))
}

// TapologyIDLTE applies the LTE predicate on the "tapology_id" field.
func TapologyIDLTE(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldLTE(FieldTapologyID, v))
}

// TapologyIDContains applies the Contains predicate on the "tapology_id" field.
func TapologyIDContains(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContains(FieldTapologyID, v))
}

// TapologyIDHasPrefix applies the HasPrefix predicate on the "tapology_id" field.
func TapologyIDHasPrefix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasPrefix(FieldTapologyID, v))
}

// TapologyIDHasSuffix applies the HasSuffix predicate on the "tapology_id" field.
func TapologyIDHasSuffix(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldHasSuffix(FieldTapologyID, v))
}

// TapologyIDIsNil applies the IsNil predicate on the "tapology_id" field.
func TapologyIDIsNil() predicate.Fighter {
	return predicate.Fighter(sql.FieldIsNull(FieldTapologyID))
}

// TapologyIDNotNil applies the NotNil predicate on the "tapology_id" field.
func TapologyIDNotNil() predicate.Fighter {
	return predicate.Fighter(sql.FieldNotNull(FieldTapologyID))
}

// TapologyIDEqualFold applies the EqualFold predicate on the "tapology_id" field.
func TapologyIDEqualFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldEqualFold(FieldTapologyID, v))
}

// TapologyIDContainsFold applies the ContainsFold predicate on the "tapology_id" field.
func TapologyIDContainsFold(v string) predicate.Fighter {
	return predicate.Fighter(sql.FieldContainsFold(FieldTapologyID, v))
}

// TemporaryEQ applies the EQ predicate on the "temporary" field.
func TemporaryEQ(v bool) predicate.Fighter {
	return predicate.Fighter(sql.FieldEQ(FieldTemporary, v))
}

// TemporaryNEQ applies the NEQ predicate on the "temporary" field.
func TemporaryNEQ(v bool) predicate.Fighter {
	return predicate.Fighter(sql.FieldNEQ(FieldTemporary, v))
}

// HasFights applies the HasEdge predicate on the "fights" edge.
func HasFights() predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, FightsTable, FightsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFightsWith applies the HasEdge predicate on the "fights" edge with a given conditions (other predicates).
func HasFightsWith(preds ...predicate.Fight) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := newFightsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpcomingFights applies the HasEdge predicate on the "upcoming_fights" edge.
func HasUpcomingFights() predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UpcomingFightsTable, UpcomingFightsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpcomingFightsWith applies the HasEdge predicate on the "upcoming_fights" edge with a given conditions (other predicates).
func HasUpcomingFightsWith(preds ...predicate.UpcomingFight) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := newUpcomingFightsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFighterAliases applies the HasEdge predicate on the "fighter_aliases" edge.
func HasFighterAliases() predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FighterAliasesTable, FighterAliasesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFighterAliasesWith applies the HasEdge predicate on the "fighter_aliases" edge with a given conditions (other predicates).
func HasFighterAliasesWith(preds ...predicate.FighterAlias) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := newFighterAliasesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFighterResults applies the HasEdge predicate on the "fighter_results" edge.
func HasFighterResults() predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, FighterResultsTable, FighterResultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFighterResultsWith applies the HasEdge predicate on the "fighter_results" edge with a given conditions (other predicates).
func HasFighterResultsWith(preds ...predicate.FighterResults) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := newFighterResultsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpcomingFighterOdds applies the HasEdge predicate on the "upcoming_fighter_odds" edge.
func HasUpcomingFighterOdds() predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UpcomingFighterOddsTable, UpcomingFighterOddsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpcomingFighterOddsWith applies the HasEdge predicate on the "upcoming_fighter_odds" edge with a given conditions (other predicates).
func HasUpcomingFighterOddsWith(preds ...predicate.UpcomingFighterOdds) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		step := newUpcomingFighterOddsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Fighter) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Fighter) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
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
func Not(p predicate.Fighter) predicate.Fighter {
	return predicate.Fighter(func(s *sql.Selector) {
		p(s.Not())
	})
}
