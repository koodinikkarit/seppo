package help

// testDb := getDb()

// 	variationVersions := []db.VariationVersion{}

// 	var variationsIds []uint32
// 	variationsIds = append(variationsIds, 2, 1)

// 	testDb.Where("variation_id in (?)", variationsIds).
// 		Group("variation_id desc").
// 		Find(&variationVersions)

// // variations := []db.Variation{}

// // testDb.Table("variations").
// // 	Preload("VariationVersions").
// // 	Find(&variations)

// // tx := testDb.Begin()

// // variations := []*db.Variation{}
// // variations = append(variations, &db.Variation{})
// // variations = append(variations, &db.Variation{})
// // variations = append(variations, &db.Variation{})

// sqlStr := "INSERT INTO `variations` (`song_id`,`language_id`,`variation_version_id`,`created_at`,`updated_at`,`deleted_at`) VALUES "
// vals := []interface{}{}

// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)
// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)
// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)
// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)
// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)
// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)
// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)
// sqlStr += "(?, ?, ?, ?, ?, ?), "
// vals = append(vals, nil, nil, nil, "2017-11-01 22:05:34", "2017-11-01 22:05:34", nil)

// sqlStr = sqlStr[0 : len(sqlStr)-2]
// //prepare the statement
// testDb.Exec(sqlStr, vals...)
