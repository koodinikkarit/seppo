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
	"gopkg.in/volatiletech/null.v6"
)

// Log is an object representing the database table.
type Log struct {
	ID          uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	LogType     int         `boil:"log_type" json:"log_type" toml:"log_type" yaml:"log_type"`
	Message     null.String `boil:"message" json:"message,omitempty" toml:"message" yaml:"message,omitempty"`
	MessageDate null.Time   `boil:"message_date" json:"message_date,omitempty" toml:"message_date" yaml:"message_date,omitempty"`

	R *logR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L logL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LogColumns = struct {
	ID          string
	LogType     string
	Message     string
	MessageDate string
}{
	ID:          "id",
	LogType:     "log_type",
	Message:     "message",
	MessageDate: "message_date",
}

// logR is where relationships are stored.
type logR struct {
}

// logL is where Load methods for each relationship are stored.
type logL struct{}

var (
	logColumns               = []string{"id", "log_type", "message", "message_date"}
	logColumnsWithoutDefault = []string{"log_type", "message", "message_date"}
	logColumnsWithDefault    = []string{"id"}
	logPrimaryKeyColumns     = []string{"id"}
)

type (
	// LogSlice is an alias for a slice of pointers to Log.
	// This should generally be used opposed to []Log.
	LogSlice []*Log
	// LogHook is the signature for custom Log hook methods
	LogHook func(boil.Executor, *Log) error

	logQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	logType                 = reflect.TypeOf(&Log{})
	logMapping              = queries.MakeStructMapping(logType)
	logPrimaryKeyMapping, _ = queries.BindMapping(logType, logMapping, logPrimaryKeyColumns)
	logInsertCacheMut       sync.RWMutex
	logInsertCache          = make(map[string]insertCache)
	logUpdateCacheMut       sync.RWMutex
	logUpdateCache          = make(map[string]updateCache)
	logUpsertCacheMut       sync.RWMutex
	logUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var logBeforeInsertHooks []LogHook
var logBeforeUpdateHooks []LogHook
var logBeforeDeleteHooks []LogHook
var logBeforeUpsertHooks []LogHook

var logAfterInsertHooks []LogHook
var logAfterSelectHooks []LogHook
var logAfterUpdateHooks []LogHook
var logAfterDeleteHooks []LogHook
var logAfterUpsertHooks []LogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Log) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range logBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Log) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range logBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Log) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range logBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Log) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range logBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Log) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range logAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Log) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range logAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Log) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range logAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Log) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range logAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Log) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range logAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLogHook registers your hook function for all future operations.
func AddLogHook(hookPoint boil.HookPoint, logHook LogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		logBeforeInsertHooks = append(logBeforeInsertHooks, logHook)
	case boil.BeforeUpdateHook:
		logBeforeUpdateHooks = append(logBeforeUpdateHooks, logHook)
	case boil.BeforeDeleteHook:
		logBeforeDeleteHooks = append(logBeforeDeleteHooks, logHook)
	case boil.BeforeUpsertHook:
		logBeforeUpsertHooks = append(logBeforeUpsertHooks, logHook)
	case boil.AfterInsertHook:
		logAfterInsertHooks = append(logAfterInsertHooks, logHook)
	case boil.AfterSelectHook:
		logAfterSelectHooks = append(logAfterSelectHooks, logHook)
	case boil.AfterUpdateHook:
		logAfterUpdateHooks = append(logAfterUpdateHooks, logHook)
	case boil.AfterDeleteHook:
		logAfterDeleteHooks = append(logAfterDeleteHooks, logHook)
	case boil.AfterUpsertHook:
		logAfterUpsertHooks = append(logAfterUpsertHooks, logHook)
	}
}

// OneP returns a single log record from the query, and panics on error.
func (q logQuery) OneP() *Log {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single log record from the query.
func (q logQuery) One() (*Log, error) {
	o := &Log{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for logs")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Log records from the query, and panics on error.
func (q logQuery) AllP() LogSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Log records from the query.
func (q logQuery) All() (LogSlice, error) {
	var o []*Log

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Log slice")
	}

	if len(logAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Log records in the query, and panics on error.
func (q logQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Log records in the query.
func (q logQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count logs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q logQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q logQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if logs exists")
	}

	return count > 0, nil
}

// LogsG retrieves all records.
func LogsG(mods ...qm.QueryMod) logQuery {
	return Logs(boil.GetDB(), mods...)
}

// Logs retrieves all the records using an executor.
func Logs(exec boil.Executor, mods ...qm.QueryMod) logQuery {
	mods = append(mods, qm.From("`logs`"))
	return logQuery{NewQuery(exec, mods...)}
}

// FindLogG retrieves a single record by ID.
func FindLogG(id uint64, selectCols ...string) (*Log, error) {
	return FindLog(boil.GetDB(), id, selectCols...)
}

// FindLogGP retrieves a single record by ID, and panics on error.
func FindLogGP(id uint64, selectCols ...string) *Log {
	retobj, err := FindLog(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLog(exec boil.Executor, id uint64, selectCols ...string) (*Log, error) {
	logObj := &Log{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `logs` where `id`=?", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(logObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from logs")
	}

	return logObj, nil
}

// FindLogP retrieves a single record by ID with an executor, and panics on error.
func FindLogP(exec boil.Executor, id uint64, selectCols ...string) *Log {
	retobj, err := FindLog(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Log) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Log) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Log) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Log) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no logs provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(logColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	logInsertCacheMut.RLock()
	cache, cached := logInsertCache[key]
	logInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			logColumns,
			logColumnsWithDefault,
			logColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(logType, logMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(logType, logMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `logs` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `logs` () VALUES ()"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `logs` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, logPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into logs")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == logMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for logs")
	}

CacheNoHooks:
	if !cached {
		logInsertCacheMut.Lock()
		logInsertCache[key] = cache
		logInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Log record. See Update for
// whitelist behavior description.
func (o *Log) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Log record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Log) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Log, and panics on error.
// See Update for whitelist behavior description.
func (o *Log) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Log.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Log) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	logUpdateCacheMut.RLock()
	cache, cached := logUpdateCache[key]
	logUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			logColumns,
			logPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update logs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `logs` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, logPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(logType, logMapping, append(wl, logPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update logs row")
	}

	if !cached {
		logUpdateCacheMut.Lock()
		logUpdateCache[key] = cache
		logUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q logQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q logQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for logs")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o LogSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o LogSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o LogSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LogSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `logs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, logPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in log slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Log) UpsertG(updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Log) UpsertGP(updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Log) UpsertP(exec boil.Executor, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Log) Upsert(exec boil.Executor, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no logs provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(logColumnsWithDefault, o)

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

	logUpsertCacheMut.RLock()
	cache, cached := logUpsertCache[key]
	logUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			logColumns,
			logColumnsWithDefault,
			logColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			logColumns,
			logPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert logs, could not build update column list")
		}

		cache.query = queries.BuildUpsertQueryMySQL(dialect, "logs", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `logs` WHERE `id`=?",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
		)

		cache.valueMapping, err = queries.BindMapping(logType, logMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(logType, logMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for logs")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == logMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for logs")
	}

CacheNoHooks:
	if !cached {
		logUpsertCacheMut.Lock()
		logUpsertCache[key] = cache
		logUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Log record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Log) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Log record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Log) DeleteG() error {
	if o == nil {
		return errors.New("models: no Log provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Log record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Log) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Log record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Log) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Log provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), logPrimaryKeyMapping)
	sql := "DELETE FROM `logs` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from logs")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q logQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q logQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no logQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from logs")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o LogSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o LogSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Log slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o LogSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LogSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Log slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(logBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `logs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, logPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from log slice")
	}

	if len(logAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Log) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Log) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Log) ReloadG() error {
	if o == nil {
		return errors.New("models: no Log provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Log) Reload(exec boil.Executor) error {
	ret, err := FindLog(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LogSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *LogSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LogSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty LogSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LogSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	logs := LogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `logs`.* FROM `logs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, logPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&logs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LogSlice")
	}

	*o = logs

	return nil
}

// LogExists checks if the Log row exists.
func LogExists(exec boil.Executor, id uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `logs` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if logs exists")
	}

	return exists, nil
}

// LogExistsG checks if the Log row exists.
func LogExistsG(id uint64) (bool, error) {
	return LogExists(boil.GetDB(), id)
}

// LogExistsGP checks if the Log row exists. Panics on error.
func LogExistsGP(id uint64) bool {
	e, err := LogExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// LogExistsP checks if the Log row exists. Panics on error.
func LogExistsP(exec boil.Executor, id uint64) bool {
	e, err := LogExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}