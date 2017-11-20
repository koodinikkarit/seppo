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

func testEvents(t *testing.T) {
	t.Parallel()

	query := Events(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testEventsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = event.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Events(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEventsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EventSlice{event}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testEventsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := EventExists(tx, event.ID)
	if err != nil {
		t.Errorf("Unable to check if Event exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EventExistsG to return true, but got false.")
	}
}
func testEventsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	eventFound, err := FindEvent(tx, event.ID)
	if err != nil {
		t.Error(err)
	}

	if eventFound == nil {
		t.Error("want a record, got nil")
	}
}
func testEventsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Events(tx).Bind(event); err != nil {
		t.Error(err)
	}
}

func testEventsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Events(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEventsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	eventOne := &Event{}
	eventTwo := &Event{}
	if err = randomize.Struct(seed, eventOne, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}
	if err = randomize.Struct(seed, eventTwo, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = eventOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = eventTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Events(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEventsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	eventOne := &Event{}
	eventTwo := &Event{}
	if err = randomize.Struct(seed, eventOne, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}
	if err = randomize.Struct(seed, eventTwo, eventDBTypes, false, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = eventOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = eventTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func eventBeforeInsertHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterInsertHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterSelectHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventBeforeUpdateHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterUpdateHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventBeforeDeleteHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterDeleteHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventBeforeUpsertHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func eventAfterUpsertHook(e boil.Executor, o *Event) error {
	*o = Event{}
	return nil
}

func testEventsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Event{}
	o := &Event{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, eventDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Event object: %s", err)
	}

	AddEventHook(boil.BeforeInsertHook, eventBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	eventBeforeInsertHooks = []EventHook{}

	AddEventHook(boil.AfterInsertHook, eventAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	eventAfterInsertHooks = []EventHook{}

	AddEventHook(boil.AfterSelectHook, eventAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	eventAfterSelectHooks = []EventHook{}

	AddEventHook(boil.BeforeUpdateHook, eventBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	eventBeforeUpdateHooks = []EventHook{}

	AddEventHook(boil.AfterUpdateHook, eventAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	eventAfterUpdateHooks = []EventHook{}

	AddEventHook(boil.BeforeDeleteHook, eventBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	eventBeforeDeleteHooks = []EventHook{}

	AddEventHook(boil.AfterDeleteHook, eventAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	eventAfterDeleteHooks = []EventHook{}

	AddEventHook(boil.BeforeUpsertHook, eventBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	eventBeforeUpsertHooks = []EventHook{}

	AddEventHook(boil.AfterUpsertHook, eventAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	eventAfterUpsertHooks = []EventHook{}
}
func testEventsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx, eventColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEventToManyEventSchedules(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Event
	var b, c EventSchedule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, eventScheduleDBTypes, false, eventScheduleColumnsWithDefault...)
	randomize.Struct(seed, &c, eventScheduleDBTypes, false, eventScheduleColumnsWithDefault...)

	b.EventID = a.ID
	c.EventID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	eventSchedule, err := a.EventSchedules(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range eventSchedule {
		if v.EventID == b.EventID {
			bFound = true
		}
		if v.EventID == c.EventID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EventSlice{&a}
	if err = a.L.LoadEventSchedules(tx, false, (*[]*Event)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.EventSchedules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.EventSchedules = nil
	if err = a.L.LoadEventSchedules(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.EventSchedules); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", eventSchedule)
	}
}

func testEventToManyAddOpEventSchedules(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Event
	var b, c, d, e EventSchedule

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, eventDBTypes, false, strmangle.SetComplement(eventPrimaryKeyColumns, eventColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*EventSchedule{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, eventScheduleDBTypes, false, strmangle.SetComplement(eventSchedulePrimaryKeyColumns, eventScheduleColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*EventSchedule{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddEventSchedules(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.EventID {
			t.Error("foreign key was wrong value", a.ID, first.EventID)
		}
		if a.ID != second.EventID {
			t.Error("foreign key was wrong value", a.ID, second.EventID)
		}

		if first.R.Event != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Event != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.EventSchedules[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.EventSchedules[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.EventSchedules(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testEventsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = event.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testEventsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := EventSlice{event}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testEventsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Events(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	eventDBTypes = map[string]string{`CreatedAt`: `datetime`, `DeletedAt`: `datetime`, `End`: `datetime`, `ID`: `bigint`, `Name`: `varchar`, `Start`: `datetime`, `UpdatedAt`: `datetime`}
	_            = bytes.MinRead
)

func testEventsUpdate(t *testing.T) {
	t.Parallel()

	if len(eventColumns) == len(eventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if err = event.Update(tx); err != nil {
		t.Error(err)
	}
}

func testEventsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(eventColumns) == len(eventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	event := &Event{}
	if err = randomize.Struct(seed, event, eventDBTypes, true, eventColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, event, eventDBTypes, true, eventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(eventColumns, eventPrimaryKeyColumns) {
		fields = eventColumns
	} else {
		fields = strmangle.SetComplement(
			eventColumns,
			eventPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(event))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := EventSlice{event}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testEventsUpsert(t *testing.T) {
	t.Parallel()

	if len(eventColumns) == len(eventPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	event := Event{}
	if err = randomize.Struct(seed, &event, eventDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = event.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Event: %s", err)
	}

	count, err := Events(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &event, eventDBTypes, false, eventPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Event struct: %s", err)
	}

	if err = event.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Event: %s", err)
	}

	count, err = Events(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
