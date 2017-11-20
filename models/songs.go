// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Song is an object representing the database table.
type Song struct {
	ID uint64 `boil:"id" json:"id" toml:"id" yaml:"id"`

	R *songR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L songL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SongColumns = struct {
	ID string
}{
	ID: "id",
}

// songR is where relationships are stored.
type songR struct {
	Variations VariationSlice
}

// songL is where Load methods for each relationship are stored.
type songL struct{}

var (
	songColumns               = []string{"id"}
	songColumnsWithoutDefault = []string{}
	songColumnsWithDefault    = []string{"id"}
	songPrimaryKeyColumns     = []string{"id"}
)

type (
	// SongSlice is an alias for a slice of pointers to Song.
	// This should generally be used opposed to []Song.
	SongSlice []*Song
	// SongHook is the signature for custom Song hook methods
	SongHook func(boil.Executor, *Song) error

	songQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	songType                 = reflect.TypeOf(&Song{})
	songMapping              = queries.MakeStructMapping(songType)
	songPrimaryKeyMapping, _ = queries.BindMapping(songType, songMapping, songPrimaryKeyColumns)
	songInsertCacheMut       sync.RWMutex
	songInsertCache          = make(map[string]insertCache)
	songUpdateCacheMut       sync.RWMutex
	songUpdateCache          = make(map[string]updateCache)
	songUpsertCacheMut       sync.RWMutex
	songUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var songBeforeInsertHooks []SongHook
var songBeforeUpdateHooks []SongHook
var songBeforeDeleteHooks []SongHook
var songBeforeUpsertHooks []SongHook

var songAfterInsertHooks []SongHook
var songAfterSelectHooks []SongHook
var songAfterUpdateHooks []SongHook
var songAfterDeleteHooks []SongHook
var songAfterUpsertHooks []SongHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Song) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range songBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Song) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range songBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Song) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range songBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Song) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range songBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Song) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range songAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Song) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range songAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Song) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range songAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Song) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range songAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Song) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range songAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSongHook registers your hook function for all future operations.
func AddSongHook(hookPoint boil.HookPoint, songHook SongHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		songBeforeInsertHooks = append(songBeforeInsertHooks, songHook)
	case boil.BeforeUpdateHook:
		songBeforeUpdateHooks = append(songBeforeUpdateHooks, songHook)
	case boil.BeforeDeleteHook:
		songBeforeDeleteHooks = append(songBeforeDeleteHooks, songHook)
	case boil.BeforeUpsertHook:
		songBeforeUpsertHooks = append(songBeforeUpsertHooks, songHook)
	case boil.AfterInsertHook:
		songAfterInsertHooks = append(songAfterInsertHooks, songHook)
	case boil.AfterSelectHook:
		songAfterSelectHooks = append(songAfterSelectHooks, songHook)
	case boil.AfterUpdateHook:
		songAfterUpdateHooks = append(songAfterUpdateHooks, songHook)
	case boil.AfterDeleteHook:
		songAfterDeleteHooks = append(songAfterDeleteHooks, songHook)
	case boil.AfterUpsertHook:
		songAfterUpsertHooks = append(songAfterUpsertHooks, songHook)
	}
}

// OneP returns a single song record from the query, and panics on error.
func (q songQuery) OneP() *Song {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single song record from the query.
func (q songQuery) One() (*Song, error) {
	o := &Song{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for songs")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Song records from the query, and panics on error.
func (q songQuery) AllP() SongSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Song records from the query.
func (q songQuery) All() (SongSlice, error) {
	var o []*Song

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Song slice")
	}

	if len(songAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Song records in the query, and panics on error.
func (q songQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Song records in the query.
func (q songQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count songs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q songQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q songQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if songs exists")
	}

	return count > 0, nil
}

// VariationsG retrieves all the variation's variations.
func (o *Song) VariationsG(mods ...qm.QueryMod) variationQuery {
	return o.Variations(boil.GetDB(), mods...)
}

// Variations retrieves all the variation's variations with an executor.
func (o *Song) Variations(exec boil.Executor, mods ...qm.QueryMod) variationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`variations`.`song_id`=?", o.ID),
	)

	query := Variations(exec, queryMods...)
	queries.SetFrom(query.Query, "`variations`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`variations`.*"})
	}

	return query
}

// LoadVariations allows an eager lookup of values, cached into the
// loaded structs of the objects.
func (songL) LoadVariations(e boil.Executor, singular bool, maybeSong interface{}) error {
	var slice []*Song
	var object *Song

	count := 1
	if singular {
		object = maybeSong.(*Song)
	} else {
		slice = *maybeSong.(*[]*Song)
		count = len(slice)
	}

	args := make([]interface{}, count)
	if singular {
		if object.R == nil {
			object.R = &songR{}
		}
		args[0] = object.ID
	} else {
		for i, obj := range slice {
			if obj.R == nil {
				obj.R = &songR{}
			}
			args[i] = obj.ID
		}
	}

	query := fmt.Sprintf(
		"select * from `variations` where `song_id` in (%s)",
		strmangle.Placeholders(dialect.IndexPlaceholders, count, 1, 1),
	)
	if boil.DebugMode {
		fmt.Fprintf(boil.DebugWriter, "%s\n%v\n", query, args)
	}

	results, err := e.Query(query, args...)
	if err != nil {
		return errors.Wrap(err, "failed to eager load variations")
	}
	defer results.Close()

	var resultSlice []*Variation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice variations")
	}

	if len(variationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Variations = resultSlice
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SongID.Uint64 {
				local.R.Variations = append(local.R.Variations, foreign)
				break
			}
		}
	}

	return nil
}

// AddVariationsG adds the given related objects to the existing relationships
// of the song, optionally inserting them as new records.
// Appends related to o.R.Variations.
// Sets related.R.Song appropriately.
// Uses the global database handle.
func (o *Song) AddVariationsG(insert bool, related ...*Variation) error {
	return o.AddVariations(boil.GetDB(), insert, related...)
}

// AddVariationsP adds the given related objects to the existing relationships
// of the song, optionally inserting them as new records.
// Appends related to o.R.Variations.
// Sets related.R.Song appropriately.
// Panics on error.
func (o *Song) AddVariationsP(exec boil.Executor, insert bool, related ...*Variation) {
	if err := o.AddVariations(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddVariationsGP adds the given related objects to the existing relationships
// of the song, optionally inserting them as new records.
// Appends related to o.R.Variations.
// Sets related.R.Song appropriately.
// Uses the global database handle and panics on error.
func (o *Song) AddVariationsGP(insert bool, related ...*Variation) {
	if err := o.AddVariations(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// AddVariations adds the given related objects to the existing relationships
// of the song, optionally inserting them as new records.
// Appends related to o.R.Variations.
// Sets related.R.Song appropriately.
func (o *Song) AddVariations(exec boil.Executor, insert bool, related ...*Variation) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SongID.Uint64 = o.ID
			rel.SongID.Valid = true
			if err = rel.Insert(exec); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `variations` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"song_id"}),
				strmangle.WhereClause("`", "`", 0, variationPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.SongID.Uint64 = o.ID
			rel.SongID.Valid = true
		}
	}

	if o.R == nil {
		o.R = &songR{
			Variations: related,
		}
	} else {
		o.R.Variations = append(o.R.Variations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &variationR{
				Song: o,
			}
		} else {
			rel.R.Song = o
		}
	}
	return nil
}

// SetVariationsG removes all previously related items of the
// song replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Song's Variations accordingly.
// Replaces o.R.Variations with related.
// Sets related.R.Song's Variations accordingly.
// Uses the global database handle.
func (o *Song) SetVariationsG(insert bool, related ...*Variation) error {
	return o.SetVariations(boil.GetDB(), insert, related...)
}

// SetVariationsP removes all previously related items of the
// song replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Song's Variations accordingly.
// Replaces o.R.Variations with related.
// Sets related.R.Song's Variations accordingly.
// Panics on error.
func (o *Song) SetVariationsP(exec boil.Executor, insert bool, related ...*Variation) {
	if err := o.SetVariations(exec, insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVariationsGP removes all previously related items of the
// song replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Song's Variations accordingly.
// Replaces o.R.Variations with related.
// Sets related.R.Song's Variations accordingly.
// Uses the global database handle and panics on error.
func (o *Song) SetVariationsGP(insert bool, related ...*Variation) {
	if err := o.SetVariations(boil.GetDB(), insert, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// SetVariations removes all previously related items of the
// song replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Song's Variations accordingly.
// Replaces o.R.Variations with related.
// Sets related.R.Song's Variations accordingly.
func (o *Song) SetVariations(exec boil.Executor, insert bool, related ...*Variation) error {
	query := "update `variations` set `song_id` = null where `song_id` = ?"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Variations {
			rel.SongID.Valid = false
			if rel.R == nil {
				continue
			}

			rel.R.Song = nil
		}

		o.R.Variations = nil
	}
	return o.AddVariations(exec, insert, related...)
}

// RemoveVariationsG relationships from objects passed in.
// Removes related items from R.Variations (uses pointer comparison, removal does not keep order)
// Sets related.R.Song.
// Uses the global database handle.
func (o *Song) RemoveVariationsG(related ...*Variation) error {
	return o.RemoveVariations(boil.GetDB(), related...)
}

// RemoveVariationsP relationships from objects passed in.
// Removes related items from R.Variations (uses pointer comparison, removal does not keep order)
// Sets related.R.Song.
// Panics on error.
func (o *Song) RemoveVariationsP(exec boil.Executor, related ...*Variation) {
	if err := o.RemoveVariations(exec, related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveVariationsGP relationships from objects passed in.
// Removes related items from R.Variations (uses pointer comparison, removal does not keep order)
// Sets related.R.Song.
// Uses the global database handle and panics on error.
func (o *Song) RemoveVariationsGP(related ...*Variation) {
	if err := o.RemoveVariations(boil.GetDB(), related...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// RemoveVariations relationships from objects passed in.
// Removes related items from R.Variations (uses pointer comparison, removal does not keep order)
// Sets related.R.Song.
func (o *Song) RemoveVariations(exec boil.Executor, related ...*Variation) error {
	var err error
	for _, rel := range related {
		rel.SongID.Valid = false
		if rel.R != nil {
			rel.R.Song = nil
		}
		if err = rel.Update(exec, "song_id"); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Variations {
			if rel != ri {
				continue
			}

			ln := len(o.R.Variations)
			if ln > 1 && i < ln-1 {
				o.R.Variations[i] = o.R.Variations[ln-1]
			}
			o.R.Variations = o.R.Variations[:ln-1]
			break
		}
	}

	return nil
}

// SongsG retrieves all records.
func SongsG(mods ...qm.QueryMod) songQuery {
	return Songs(boil.GetDB(), mods...)
}

// Songs retrieves all the records using an executor.
func Songs(exec boil.Executor, mods ...qm.QueryMod) songQuery {
	mods = append(mods, qm.From("`songs`"))
	return songQuery{NewQuery(exec, mods...)}
}

// FindSongG retrieves a single record by ID.
func FindSongG(id uint64, selectCols ...string) (*Song, error) {
	return FindSong(boil.GetDB(), id, selectCols...)
}

// FindSongGP retrieves a single record by ID, and panics on error.
func FindSongGP(id uint64, selectCols ...string) *Song {
	retobj, err := FindSong(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindSong retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSong(exec boil.Executor, id uint64, selectCols ...string) (*Song, error) {
	songObj := &Song{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `songs` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(songObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from songs")
	}

	return songObj, nil
}

// FindSongP retrieves a single record by ID with an executor, and panics on error.
func FindSongP(exec boil.Executor, id uint64, selectCols ...string) *Song {
	retobj, err := FindSong(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Song) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Song) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Song) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Song) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no songs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(songColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	songInsertCacheMut.RLock()
	cache, cached := songInsertCache[key]
	songInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			songColumns,
			songColumnsWithDefault,
			songColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(songType, songMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(songType, songMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `songs` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `songs` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `songs` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, songPrimaryKeyColumns))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into songs")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == songMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for songs")
	}

CacheNoHooks:
	if !cached {
		songInsertCacheMut.Lock()
		songInsertCache[key] = cache
		songInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Song record. See Update for
// whitelist behavior description.
func (o *Song) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Song record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Song) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Song, and panics on error.
// See Update for whitelist behavior description.
func (o *Song) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Song.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Song) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	songUpdateCacheMut.RLock()
	cache, cached := songUpdateCache[key]
	songUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			songColumns,
			songPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update songs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `songs` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, songPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(songType, songMapping, append(wl, songPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update songs row")
	}

	if !cached {
		songUpdateCacheMut.Lock()
		songUpdateCache[key] = cache
		songUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q songQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q songQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for songs")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SongSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o SongSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o SongSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SongSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), songPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `songs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, songPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in song slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Song) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Song) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Song) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Song) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no songs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(songColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	songUpsertCacheMut.RLock()
	cache, cached := songUpsertCache[key]
	songUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			songColumns,
			songColumnsWithDefault,
			songColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			songColumns,
			songPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert songs, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "songs", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `songs` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(songType, songMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(songType, songMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for songs")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == songMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for songs")
	}

CacheNoHooks:
	if !cached {
		songUpsertCacheMut.Lock()
		songUpsertCache[key] = cache
		songUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Song record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Song) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Song record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Song) DeleteG() error {
	if o == nil {
		return errors.New("models: no Song provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Song record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Song) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Song record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Song) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Song provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), songPrimaryKeyMapping)
	sql := "DELETE FROM `songs` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from songs")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q songQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q songQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no songQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from songs")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o SongSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o SongSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Song slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o SongSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SongSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Song slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(songBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), songPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `songs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, songPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from song slice")
	}

	if len(songAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Song) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Song) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Song) ReloadG() error {
	if o == nil {
		return errors.New("models: no Song provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Song) Reload(exec boil.Executor) error {
	ret, err := FindSong(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *SongSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *SongSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SongSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty SongSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SongSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	songs := SongSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), songPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `songs`.* FROM `songs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, songPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&songs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SongSlice")
	}

	*o = songs

	return nil
}

// SongExists checks if the Song row exists.
func SongExists(exec boil.Executor, id uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `songs` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if songs exists")
	}

	return exists, nil
}

// SongExistsG checks if the Song row exists.
func SongExistsG(id uint64) (bool, error) {
	return SongExists(boil.GetDB(), id)
}

// SongExistsGP checks if the Song row exists. Panics on error.
func SongExistsGP(id uint64) bool {
	e, err := SongExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// SongExistsP checks if the Song row exists. Panics on error.
func SongExistsP(exec boil.Executor, id uint64) bool {
	e, err := SongExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}