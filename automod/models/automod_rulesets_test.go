// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testAutomodRulesets(t *testing.T) {
	t.Parallel()

	query := AutomodRulesets()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAutomodRulesetsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutomodRulesetsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := AutomodRulesets().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutomodRulesetsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AutomodRulesetSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAutomodRulesetsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AutomodRulesetExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if AutomodRuleset exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AutomodRulesetExists to return true, but got false.")
	}
}

func testAutomodRulesetsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	automodRulesetFound, err := FindAutomodRuleset(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if automodRulesetFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAutomodRulesetsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = AutomodRulesets().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAutomodRulesetsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := AutomodRulesets().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAutomodRulesetsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	automodRulesetOne := &AutomodRuleset{}
	automodRulesetTwo := &AutomodRuleset{}
	if err = randomize.Struct(seed, automodRulesetOne, automodRulesetDBTypes, false, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}
	if err = randomize.Struct(seed, automodRulesetTwo, automodRulesetDBTypes, false, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = automodRulesetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = automodRulesetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AutomodRulesets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAutomodRulesetsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	automodRulesetOne := &AutomodRuleset{}
	automodRulesetTwo := &AutomodRuleset{}
	if err = randomize.Struct(seed, automodRulesetOne, automodRulesetDBTypes, false, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}
	if err = randomize.Struct(seed, automodRulesetTwo, automodRulesetDBTypes, false, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = automodRulesetOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = automodRulesetTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testAutomodRulesetsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAutomodRulesetsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(automodRulesetColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAutomodRulesetToManyRulesetAutomodRules(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutomodRuleset
	var b, c AutomodRule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, automodRuleDBTypes, false, automodRuleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, automodRuleDBTypes, false, automodRuleColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RulesetID = a.ID
	c.RulesetID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	automodRule, err := a.RulesetAutomodRules().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range automodRule {
		if v.RulesetID == b.RulesetID {
			bFound = true
		}
		if v.RulesetID == c.RulesetID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AutomodRulesetSlice{&a}
	if err = a.L.LoadRulesetAutomodRules(ctx, tx, false, (*[]*AutomodRuleset)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RulesetAutomodRules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RulesetAutomodRules = nil
	if err = a.L.LoadRulesetAutomodRules(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RulesetAutomodRules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", automodRule)
	}
}

func testAutomodRulesetToManyRulesetAutomodRulesetConditions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutomodRuleset
	var b, c AutomodRulesetCondition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, automodRulesetConditionDBTypes, false, automodRulesetConditionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, automodRulesetConditionDBTypes, false, automodRulesetConditionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RulesetID = a.ID
	c.RulesetID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	automodRulesetCondition, err := a.RulesetAutomodRulesetConditions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range automodRulesetCondition {
		if v.RulesetID == b.RulesetID {
			bFound = true
		}
		if v.RulesetID == c.RulesetID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AutomodRulesetSlice{&a}
	if err = a.L.LoadRulesetAutomodRulesetConditions(ctx, tx, false, (*[]*AutomodRuleset)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RulesetAutomodRulesetConditions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RulesetAutomodRulesetConditions = nil
	if err = a.L.LoadRulesetAutomodRulesetConditions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RulesetAutomodRulesetConditions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", automodRulesetCondition)
	}
}

func testAutomodRulesetToManyAddOpRulesetAutomodRules(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutomodRuleset
	var b, c, d, e AutomodRule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, automodRulesetDBTypes, false, strmangle.SetComplement(automodRulesetPrimaryKeyColumns, automodRulesetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AutomodRule{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, automodRuleDBTypes, false, strmangle.SetComplement(automodRulePrimaryKeyColumns, automodRuleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AutomodRule{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRulesetAutomodRules(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RulesetID {
			t.Error("foreign key was wrong value", a.ID, first.RulesetID)
		}
		if a.ID != second.RulesetID {
			t.Error("foreign key was wrong value", a.ID, second.RulesetID)
		}

		if first.R.Ruleset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Ruleset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RulesetAutomodRules[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RulesetAutomodRules[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RulesetAutomodRules().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAutomodRulesetToManyAddOpRulesetAutomodRulesetConditions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a AutomodRuleset
	var b, c, d, e AutomodRulesetCondition

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, automodRulesetDBTypes, false, strmangle.SetComplement(automodRulesetPrimaryKeyColumns, automodRulesetColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AutomodRulesetCondition{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, automodRulesetConditionDBTypes, false, strmangle.SetComplement(automodRulesetConditionPrimaryKeyColumns, automodRulesetConditionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AutomodRulesetCondition{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRulesetAutomodRulesetConditions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RulesetID {
			t.Error("foreign key was wrong value", a.ID, first.RulesetID)
		}
		if a.ID != second.RulesetID {
			t.Error("foreign key was wrong value", a.ID, second.RulesetID)
		}

		if first.R.Ruleset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Ruleset != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RulesetAutomodRulesetConditions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RulesetAutomodRulesetConditions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RulesetAutomodRulesetConditions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAutomodRulesetsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAutomodRulesetsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AutomodRulesetSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAutomodRulesetsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := AutomodRulesets().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	automodRulesetDBTypes = map[string]string{`Enabled`: `boolean`, `GuildID`: `bigint`, `ID`: `bigint`, `Name`: `text`}
	_                     = bytes.MinRead
)

func testAutomodRulesetsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(automodRulesetPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(automodRulesetColumns) == len(automodRulesetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAutomodRulesetsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(automodRulesetColumns) == len(automodRulesetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &AutomodRuleset{}
	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, automodRulesetDBTypes, true, automodRulesetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(automodRulesetColumns, automodRulesetPrimaryKeyColumns) {
		fields = automodRulesetColumns
	} else {
		fields = strmangle.SetComplement(
			automodRulesetColumns,
			automodRulesetPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := AutomodRulesetSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAutomodRulesetsUpsert(t *testing.T) {
	t.Parallel()

	if len(automodRulesetColumns) == len(automodRulesetPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := AutomodRuleset{}
	if err = randomize.Struct(seed, &o, automodRulesetDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AutomodRuleset: %s", err)
	}

	count, err := AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, automodRulesetDBTypes, false, automodRulesetPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AutomodRuleset struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert AutomodRuleset: %s", err)
	}

	count, err = AutomodRulesets().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}