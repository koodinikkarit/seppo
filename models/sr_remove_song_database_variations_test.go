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

func testSRRemoveSongDatabaseVariations(t *testing.T) {
	t.Parallel()

	query := SRRemoveSongDatabaseVariations(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testSRRemoveSongDatabaseVariationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = srRemoveSongDatabaseVariation.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSRRemoveSongDatabaseVariationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SRRemoveSongDatabaseVariations(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSRRemoveSongDatabaseVariationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SRRemoveSongDatabaseVariationSlice{srRemoveSongDatabaseVariation}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testSRRemoveSongDatabaseVariationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SRRemoveSongDatabaseVariationExists(tx, srRemoveSongDatabaseVariation.ID)
	if err != nil {
		t.Errorf("Unable to check if SRRemoveSongDatabaseVariation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SRRemoveSongDatabaseVariationExistsG to return true, but got false.")
	}
}
func testSRRemoveSongDatabaseVariationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	srRemoveSongDatabaseVariationFound, err := FindSRRemoveSongDatabaseVariation(tx, srRemoveSongDatabaseVariation.ID)
	if err != nil {
		t.Error(err)
	}

	if srRemoveSongDatabaseVariationFound == nil {
		t.Error("want a record, got nil")
	}
}
func testSRRemoveSongDatabaseVariationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SRRemoveSongDatabaseVariations(tx).Bind(srRemoveSongDatabaseVariation); err != nil {
		t.Error(err)
	}
}

func testSRRemoveSongDatabaseVariationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := SRRemoveSongDatabaseVariations(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSRRemoveSongDatabaseVariationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariationOne := &SRRemoveSongDatabaseVariation{}
	srRemoveSongDatabaseVariationTwo := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariationOne, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariationTwo, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = srRemoveSongDatabaseVariationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SRRemoveSongDatabaseVariations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSRRemoveSongDatabaseVariationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	srRemoveSongDatabaseVariationOne := &SRRemoveSongDatabaseVariation{}
	srRemoveSongDatabaseVariationTwo := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariationOne, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariationTwo, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = srRemoveSongDatabaseVariationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func srRemoveSongDatabaseVariationBeforeInsertHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationAfterInsertHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationAfterSelectHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationBeforeUpdateHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationAfterUpdateHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationBeforeDeleteHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationAfterDeleteHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationBeforeUpsertHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func srRemoveSongDatabaseVariationAfterUpsertHook(e boil.Executor, o *SRRemoveSongDatabaseVariation) error {
	*o = SRRemoveSongDatabaseVariation{}
	return nil
}

func testSRRemoveSongDatabaseVariationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &SRRemoveSongDatabaseVariation{}
	o := &SRRemoveSongDatabaseVariation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, srRemoveSongDatabaseVariationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation object: %s", err)
	}

	AddSRRemoveSongDatabaseVariationHook(boil.BeforeInsertHook, srRemoveSongDatabaseVariationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationBeforeInsertHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.AfterInsertHook, srRemoveSongDatabaseVariationAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationAfterInsertHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.AfterSelectHook, srRemoveSongDatabaseVariationAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationAfterSelectHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.BeforeUpdateHook, srRemoveSongDatabaseVariationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationBeforeUpdateHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.AfterUpdateHook, srRemoveSongDatabaseVariationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationAfterUpdateHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.BeforeDeleteHook, srRemoveSongDatabaseVariationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationBeforeDeleteHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.AfterDeleteHook, srRemoveSongDatabaseVariationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationAfterDeleteHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.BeforeUpsertHook, srRemoveSongDatabaseVariationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationBeforeUpsertHooks = []SRRemoveSongDatabaseVariationHook{}

	AddSRRemoveSongDatabaseVariationHook(boil.AfterUpsertHook, srRemoveSongDatabaseVariationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	srRemoveSongDatabaseVariationAfterUpsertHooks = []SRRemoveSongDatabaseVariationHook{}
}
func testSRRemoveSongDatabaseVariationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSRRemoveSongDatabaseVariationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx, srRemoveSongDatabaseVariationColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSRRemoveSongDatabaseVariationToOneSynchronizationRaportUsingSR(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SRRemoveSongDatabaseVariation
	var foreign SynchronizationRaport

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
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

	slice := SRRemoveSongDatabaseVariationSlice{&local}
	if err = local.L.LoadSR(tx, false, (*[]*SRRemoveSongDatabaseVariation)(&slice)); err != nil {
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

func testSRRemoveSongDatabaseVariationToOneVariationUsingVariation(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SRRemoveSongDatabaseVariation
	var foreign Variation

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, variationDBTypes, false, variationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Variation struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.VariationID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Variation(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SRRemoveSongDatabaseVariationSlice{&local}
	if err = local.L.LoadVariation(tx, false, (*[]*SRRemoveSongDatabaseVariation)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Variation == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Variation = nil
	if err = local.L.LoadVariation(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Variation == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSRRemoveSongDatabaseVariationToOneSongDatabaseUsingSongDatabase(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local SRRemoveSongDatabaseVariation
	var foreign SongDatabase

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, songDatabaseDBTypes, false, songDatabaseColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SongDatabase struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SongDatabaseID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SongDatabase(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SRRemoveSongDatabaseVariationSlice{&local}
	if err = local.L.LoadSongDatabase(tx, false, (*[]*SRRemoveSongDatabaseVariation)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.SongDatabase == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SongDatabase = nil
	if err = local.L.LoadSongDatabase(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SongDatabase == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSRRemoveSongDatabaseVariationToOneSetOpSynchronizationRaportUsingSR(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SRRemoveSongDatabaseVariation
	var b, c SynchronizationRaport

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, srRemoveSongDatabaseVariationDBTypes, false, strmangle.SetComplement(srRemoveSongDatabaseVariationPrimaryKeyColumns, srRemoveSongDatabaseVariationColumnsWithoutDefault)...); err != nil {
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

		if x.R.SRSRRemoveSongDatabaseVariations[0] != &a {
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
func testSRRemoveSongDatabaseVariationToOneSetOpVariationUsingVariation(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SRRemoveSongDatabaseVariation
	var b, c Variation

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, srRemoveSongDatabaseVariationDBTypes, false, strmangle.SetComplement(srRemoveSongDatabaseVariationPrimaryKeyColumns, srRemoveSongDatabaseVariationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, variationDBTypes, false, strmangle.SetComplement(variationPrimaryKeyColumns, variationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, variationDBTypes, false, strmangle.SetComplement(variationPrimaryKeyColumns, variationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Variation{&b, &c} {
		err = a.SetVariation(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Variation != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SRRemoveSongDatabaseVariations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.VariationID != x.ID {
			t.Error("foreign key was wrong value", a.VariationID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.VariationID))
		reflect.Indirect(reflect.ValueOf(&a.VariationID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.VariationID != x.ID {
			t.Error("foreign key was wrong value", a.VariationID, x.ID)
		}
	}
}
func testSRRemoveSongDatabaseVariationToOneSetOpSongDatabaseUsingSongDatabase(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a SRRemoveSongDatabaseVariation
	var b, c SongDatabase

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, srRemoveSongDatabaseVariationDBTypes, false, strmangle.SetComplement(srRemoveSongDatabaseVariationPrimaryKeyColumns, srRemoveSongDatabaseVariationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, songDatabaseDBTypes, false, strmangle.SetComplement(songDatabasePrimaryKeyColumns, songDatabaseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, songDatabaseDBTypes, false, strmangle.SetComplement(songDatabasePrimaryKeyColumns, songDatabaseColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SongDatabase{&b, &c} {
		err = a.SetSongDatabase(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SongDatabase != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SRRemoveSongDatabaseVariations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SongDatabaseID != x.ID {
			t.Error("foreign key was wrong value", a.SongDatabaseID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SongDatabaseID))
		reflect.Indirect(reflect.ValueOf(&a.SongDatabaseID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SongDatabaseID != x.ID {
			t.Error("foreign key was wrong value", a.SongDatabaseID, x.ID)
		}
	}
}
func testSRRemoveSongDatabaseVariationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = srRemoveSongDatabaseVariation.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSRRemoveSongDatabaseVariationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SRRemoveSongDatabaseVariationSlice{srRemoveSongDatabaseVariation}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testSRRemoveSongDatabaseVariationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SRRemoveSongDatabaseVariations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	srRemoveSongDatabaseVariationDBTypes = map[string]string{`ID`: `bigint`, `SRID`: `bigint`, `SongDatabaseID`: `bigint`, `VariationID`: `bigint`}
	_                                    = bytes.MinRead
)

func testSRRemoveSongDatabaseVariationsUpdate(t *testing.T) {
	t.Parallel()

	if len(srRemoveSongDatabaseVariationColumns) == len(srRemoveSongDatabaseVariationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	if err = srRemoveSongDatabaseVariation.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSRRemoveSongDatabaseVariationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(srRemoveSongDatabaseVariationColumns) == len(srRemoveSongDatabaseVariationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	srRemoveSongDatabaseVariation := &SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true, srRemoveSongDatabaseVariationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(srRemoveSongDatabaseVariationColumns, srRemoveSongDatabaseVariationPrimaryKeyColumns) {
		fields = srRemoveSongDatabaseVariationColumns
	} else {
		fields = strmangle.SetComplement(
			srRemoveSongDatabaseVariationColumns,
			srRemoveSongDatabaseVariationPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(srRemoveSongDatabaseVariation))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SRRemoveSongDatabaseVariationSlice{srRemoveSongDatabaseVariation}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testSRRemoveSongDatabaseVariationsUpsert(t *testing.T) {
	t.Parallel()

	if len(srRemoveSongDatabaseVariationColumns) == len(srRemoveSongDatabaseVariationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	srRemoveSongDatabaseVariation := SRRemoveSongDatabaseVariation{}
	if err = randomize.Struct(seed, &srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = srRemoveSongDatabaseVariation.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert SRRemoveSongDatabaseVariation: %s", err)
	}

	count, err := SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &srRemoveSongDatabaseVariation, srRemoveSongDatabaseVariationDBTypes, false, srRemoveSongDatabaseVariationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SRRemoveSongDatabaseVariation struct: %s", err)
	}

	if err = srRemoveSongDatabaseVariation.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert SRRemoveSongDatabaseVariation: %s", err)
	}

	count, err = SRRemoveSongDatabaseVariations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
