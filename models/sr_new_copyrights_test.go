// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testSRNewCopyrights(t *testing.T) {
	t.Parallel()

	query := SRNewCopyrights(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSRNewCopyrightsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = srNewCopyright.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSRNewCopyrightsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SRNewCopyrights(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSRNewCopyrightsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SRNewCopyrightSlice{srNewCopyright}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSRNewCopyrightsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SRNewCopyrightExists(tx, srNewCopyright.ID)
	if err != nil {
		t.Errorf("Unable to check if SRNewCopyright exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SRNewCopyrightExistsG to return true, but got false.")
	}
}
func testSRNewCopyrightsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	srNewCopyrightFound, err := FindSRNewCopyright(tx, srNewCopyright.ID)
	if err != nil {
		t.Error(err)
	}

	if srNewCopyrightFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSRNewCopyrightsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SRNewCopyrights(tx).Bind(srNewCopyright); err != nil {
		t.Error(err)
	}
}

func testSRNewCopyrightsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := SRNewCopyrights(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSRNewCopyrightsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyrightOne := &SRNewCopyright{}
	srNewCopyrightTwo := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyrightOne, srNewCopyrightDBTypes, false, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}
	if err = randomize.Struct(seed, srNewCopyrightTwo, srNewCopyrightDBTypes, false, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyrightOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = srNewCopyrightTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SRNewCopyrights(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSRNewCopyrightsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	srNewCopyrightOne := &SRNewCopyright{}
	srNewCopyrightTwo := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyrightOne, srNewCopyrightDBTypes, false, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}
	if err = randomize.Struct(seed, srNewCopyrightTwo, srNewCopyrightDBTypes, false, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyrightOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = srNewCopyrightTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func srNewCopyrightBeforeInsertHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightAfterInsertHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightAfterSelectHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightBeforeUpdateHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightAfterUpdateHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightBeforeDeleteHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightAfterDeleteHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightBeforeUpsertHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func srNewCopyrightAfterUpsertHook(e boil.Executor, o *SRNewCopyright) error {
	*o = SRNewCopyright{}
	return nil
}

func testSRNewCopyrightsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &SRNewCopyright{}
	o := &SRNewCopyright{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, srNewCopyrightDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright object: %s", err)
	}

	AddSRNewCopyrightHook(boil.BeforeInsertHook, srNewCopyrightBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightBeforeInsertHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.AfterInsertHook, srNewCopyrightAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightAfterInsertHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.AfterSelectHook, srNewCopyrightAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightAfterSelectHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.BeforeUpdateHook, srNewCopyrightBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightBeforeUpdateHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.AfterUpdateHook, srNewCopyrightAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightAfterUpdateHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.BeforeDeleteHook, srNewCopyrightBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightBeforeDeleteHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.AfterDeleteHook, srNewCopyrightAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightAfterDeleteHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.BeforeUpsertHook, srNewCopyrightBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightBeforeUpsertHooks = []SRNewCopyrightHook{}

	AddSRNewCopyrightHook(boil.AfterUpsertHook, srNewCopyrightAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	srNewCopyrightAfterUpsertHooks = []SRNewCopyrightHook{}
}
func testSRNewCopyrightsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSRNewCopyrightsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx, srNewCopyrightColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSRNewCopyrightToOneSynchronizationRaportUsingSR(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SRNewCopyright
	var foreign SynchronizationRaport

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, srNewCopyrightDBTypes, false, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, synchronizationRaportDBTypes, false, synchronizationRaportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SynchronizationRaport struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SRID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SR(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SRNewCopyrightSlice{&local}
	if err = local.L.LoadSR(tx, false, (*[]*SRNewCopyright)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.SR == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SR = nil
	if err = local.L.LoadSR(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SR == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSRNewCopyrightToOneCopyrightUsingCopyright(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SRNewCopyright
	var foreign Copyright

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, srNewCopyrightDBTypes, false, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, copyrightDBTypes, false, copyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Copyright struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.CopyrightID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Copyright(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SRNewCopyrightSlice{&local}
	if err = local.L.LoadCopyright(tx, false, (*[]*SRNewCopyright)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Copyright == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Copyright = nil
	if err = local.L.LoadCopyright(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Copyright == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSRNewCopyrightToOneSetOpSynchronizationRaportUsingSR(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SRNewCopyright
	var b, c SynchronizationRaport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, srNewCopyrightDBTypes, false, strmangle.SetComplement(srNewCopyrightPrimaryKeyColumns, srNewCopyrightColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, synchronizationRaportDBTypes, false, strmangle.SetComplement(synchronizationRaportPrimaryKeyColumns, synchronizationRaportColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, synchronizationRaportDBTypes, false, strmangle.SetComplement(synchronizationRaportPrimaryKeyColumns, synchronizationRaportColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SynchronizationRaport{&b, &c} {
		err = a.SetSR(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SR != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SRSRNewCopyrights[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SRID != x.ID {
			t.Error("foreign key was wrong value", a.SRID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SRID))
		reflect.Indirect(reflect.ValueOf(&a.SRID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SRID != x.ID {
			t.Error("foreign key was wrong value", a.SRID, x.ID)
		}
	}
}
func testSRNewCopyrightToOneSetOpCopyrightUsingCopyright(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SRNewCopyright
	var b, c Copyright

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, srNewCopyrightDBTypes, false, strmangle.SetComplement(srNewCopyrightPrimaryKeyColumns, srNewCopyrightColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, copyrightDBTypes, false, strmangle.SetComplement(copyrightPrimaryKeyColumns, copyrightColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, copyrightDBTypes, false, strmangle.SetComplement(copyrightPrimaryKeyColumns, copyrightColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Copyright{&b, &c} {
		err = a.SetCopyright(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Copyright != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SRNewCopyrights[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CopyrightID != x.ID {
			t.Error("foreign key was wrong value", a.CopyrightID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CopyrightID))
		reflect.Indirect(reflect.ValueOf(&a.CopyrightID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CopyrightID != x.ID {
			t.Error("foreign key was wrong value", a.CopyrightID, x.ID)
		}
	}
}
func testSRNewCopyrightsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = srNewCopyright.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSRNewCopyrightsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SRNewCopyrightSlice{srNewCopyright}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSRNewCopyrightsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SRNewCopyrights(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	srNewCopyrightDBTypes = map[string]string{`CopyrightID`: `bigint`, `ID`: `bigint`, `SRID`: `bigint`}
	_                     = bytes.MinRead
)

func testSRNewCopyrightsUpdate(t *testing.T) {
	t.Parallel()

	if len(srNewCopyrightColumns) == len(srNewCopyrightPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	if err = srNewCopyright.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSRNewCopyrightsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(srNewCopyrightColumns) == len(srNewCopyrightPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	srNewCopyright := &SRNewCopyright{}
	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, srNewCopyright, srNewCopyrightDBTypes, true, srNewCopyrightPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(srNewCopyrightColumns, srNewCopyrightPrimaryKeyColumns) {
		fields = srNewCopyrightColumns
	} else {
		fields = strmangle.SetComplement(
			srNewCopyrightColumns,
			srNewCopyrightPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(srNewCopyright))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SRNewCopyrightSlice{srNewCopyright}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSRNewCopyrightsUpsert(t *testing.T) {
	t.Parallel()

	if len(srNewCopyrightColumns) == len(srNewCopyrightPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	srNewCopyright := SRNewCopyright{}
	if err = randomize.Struct(seed, &srNewCopyright, srNewCopyrightDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewCopyright.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert SRNewCopyright: %s", err)
	}

	count, err := SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &srNewCopyright, srNewCopyrightDBTypes, false, srNewCopyrightPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SRNewCopyright struct: %s", err)
	}

	if err = srNewCopyright.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert SRNewCopyright: %s", err)
	}

	count, err = SRNewCopyrights(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
