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

func testSRNewAuthors(t *testing.T) {
	t.Parallel()

	query := SRNewAuthors(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSRNewAuthorsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = srNewAuthor.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSRNewAuthorsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SRNewAuthors(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSRNewAuthorsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SRNewAuthorSlice{srNewAuthor}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSRNewAuthorsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SRNewAuthorExists(tx, srNewAuthor.ID)
	if err != nil {
		t.Errorf("Unable to check if SRNewAuthor exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SRNewAuthorExistsG to return true, but got false.")
	}
}
func testSRNewAuthorsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	srNewAuthorFound, err := FindSRNewAuthor(tx, srNewAuthor.ID)
	if err != nil {
		t.Error(err)
	}

	if srNewAuthorFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSRNewAuthorsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SRNewAuthors(tx).Bind(srNewAuthor); err != nil {
		t.Error(err)
	}
}

func testSRNewAuthorsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := SRNewAuthors(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSRNewAuthorsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthorOne := &SRNewAuthor{}
	srNewAuthorTwo := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthorOne, srNewAuthorDBTypes, false, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}
	if err = randomize.Struct(seed, srNewAuthorTwo, srNewAuthorDBTypes, false, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthorOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = srNewAuthorTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SRNewAuthors(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSRNewAuthorsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	srNewAuthorOne := &SRNewAuthor{}
	srNewAuthorTwo := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthorOne, srNewAuthorDBTypes, false, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}
	if err = randomize.Struct(seed, srNewAuthorTwo, srNewAuthorDBTypes, false, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthorOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = srNewAuthorTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func srNewAuthorBeforeInsertHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorAfterInsertHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorAfterSelectHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorBeforeUpdateHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorAfterUpdateHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorBeforeDeleteHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorAfterDeleteHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorBeforeUpsertHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func srNewAuthorAfterUpsertHook(e boil.Executor, o *SRNewAuthor) error {
	*o = SRNewAuthor{}
	return nil
}

func testSRNewAuthorsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &SRNewAuthor{}
	o := &SRNewAuthor{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, srNewAuthorDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor object: %s", err)
	}

	AddSRNewAuthorHook(boil.BeforeInsertHook, srNewAuthorBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	srNewAuthorBeforeInsertHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.AfterInsertHook, srNewAuthorAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	srNewAuthorAfterInsertHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.AfterSelectHook, srNewAuthorAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	srNewAuthorAfterSelectHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.BeforeUpdateHook, srNewAuthorBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	srNewAuthorBeforeUpdateHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.AfterUpdateHook, srNewAuthorAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	srNewAuthorAfterUpdateHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.BeforeDeleteHook, srNewAuthorBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	srNewAuthorBeforeDeleteHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.AfterDeleteHook, srNewAuthorAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	srNewAuthorAfterDeleteHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.BeforeUpsertHook, srNewAuthorBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	srNewAuthorBeforeUpsertHooks = []SRNewAuthorHook{}

	AddSRNewAuthorHook(boil.AfterUpsertHook, srNewAuthorAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	srNewAuthorAfterUpsertHooks = []SRNewAuthorHook{}
}
func testSRNewAuthorsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSRNewAuthorsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx, srNewAuthorColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSRNewAuthorToOneSynchronizationRaportUsingSR(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SRNewAuthor
	var foreign SynchronizationRaport

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, srNewAuthorDBTypes, false, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
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

	slice := SRNewAuthorSlice{&local}
	if err = local.L.LoadSR(tx, false, (*[]*SRNewAuthor)(&slice)); err != nil {
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

func testSRNewAuthorToOneAuthorUsingAuthor(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SRNewAuthor
	var foreign Author

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, srNewAuthorDBTypes, false, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, authorDBTypes, false, authorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Author struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.AuthorID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Author(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SRNewAuthorSlice{&local}
	if err = local.L.LoadAuthor(tx, false, (*[]*SRNewAuthor)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Author == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Author = nil
	if err = local.L.LoadAuthor(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Author == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSRNewAuthorToOneSetOpSynchronizationRaportUsingSR(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SRNewAuthor
	var b, c SynchronizationRaport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, srNewAuthorDBTypes, false, strmangle.SetComplement(srNewAuthorPrimaryKeyColumns, srNewAuthorColumnsWithoutDefault)...); err != nil {
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

		if x.R.SRSRNewAuthors[0] != &a {
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
func testSRNewAuthorToOneSetOpAuthorUsingAuthor(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SRNewAuthor
	var b, c Author

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, srNewAuthorDBTypes, false, strmangle.SetComplement(srNewAuthorPrimaryKeyColumns, srNewAuthorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, authorDBTypes, false, strmangle.SetComplement(authorPrimaryKeyColumns, authorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authorDBTypes, false, strmangle.SetComplement(authorPrimaryKeyColumns, authorColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Author{&b, &c} {
		err = a.SetAuthor(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Author != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SRNewAuthors[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AuthorID != x.ID {
			t.Error("foreign key was wrong value", a.AuthorID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AuthorID))
		reflect.Indirect(reflect.ValueOf(&a.AuthorID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AuthorID != x.ID {
			t.Error("foreign key was wrong value", a.AuthorID, x.ID)
		}
	}
}
func testSRNewAuthorsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = srNewAuthor.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSRNewAuthorsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SRNewAuthorSlice{srNewAuthor}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSRNewAuthorsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SRNewAuthors(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	srNewAuthorDBTypes = map[string]string{`AuthorID`: `bigint`, `ID`: `bigint`, `SRID`: `bigint`}
	_                  = bytes.MinRead
)

func testSRNewAuthorsUpdate(t *testing.T) {
	t.Parallel()

	if len(srNewAuthorColumns) == len(srNewAuthorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	if err = srNewAuthor.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSRNewAuthorsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(srNewAuthorColumns) == len(srNewAuthorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	srNewAuthor := &SRNewAuthor{}
	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, srNewAuthor, srNewAuthorDBTypes, true, srNewAuthorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(srNewAuthorColumns, srNewAuthorPrimaryKeyColumns) {
		fields = srNewAuthorColumns
	} else {
		fields = strmangle.SetComplement(
			srNewAuthorColumns,
			srNewAuthorPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(srNewAuthor))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SRNewAuthorSlice{srNewAuthor}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSRNewAuthorsUpsert(t *testing.T) {
	t.Parallel()

	if len(srNewAuthorColumns) == len(srNewAuthorPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	srNewAuthor := SRNewAuthor{}
	if err = randomize.Struct(seed, &srNewAuthor, srNewAuthorDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srNewAuthor.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert SRNewAuthor: %s", err)
	}

	count, err := SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &srNewAuthor, srNewAuthorDBTypes, false, srNewAuthorPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SRNewAuthor struct: %s", err)
	}

	if err = srNewAuthor.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert SRNewAuthor: %s", err)
	}

	count, err = SRNewAuthors(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}