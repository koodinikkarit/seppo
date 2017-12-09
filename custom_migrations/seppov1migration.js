const Promise = require("bluebird");
const mysql = require("promise-mysql");
const batchInsert = require("./batchInsert");
const stringHelp = require("./stringHelp");

Promise.all([
	mysql.createConnection({
		host: "localhost",
		user: "jaska",
		password: "asdf321",
		database: "seppoold"
	}),
	mysql.createConnection({
		host: "localhost",
		user: "jaska",
		password: "asdf321",
		database: "seppo"
	})
]).then((conns) => {
	const conn = conns[0];
	const conn2 = conns[1];
	Promise.all([
		conn.query("select * from ew_database_links"),
		conn.query("select * from ew_databases"),
		conn.query("select * from languages"),
		conn.query("select * from schedule_variations"),
		conn.query("select * from schedules"),
		conn.query("select * from song_database_tags"),
		conn.query("select * from song_database_variations"),
		conn.query("select * from song_databases"),
		conn.query("select * from songs"),
		conn.query("select * from tag_variations"),
		conn.query("select * from tags"),
		conn.query("select * from variation_ew_song_data"),
		conn.query("select * from variations v left join variation_texts vt on v.id = vt.variation_id")
	]).then(data => {		
		conn2.beginTransaction();

		var filtteredVariations = [];

		data[12].forEach(variation => {
			if (!filtteredVariations.some(p => {
				if (p.name === variation.name &&
					p.text === variation.text) {
						return true;
				}
			})) {
				filtteredVariations.push(variation);
			}
		});
		const nDublicates = data[12].length - filtteredVariations.length;

		console.log("number of variation duplicates", nDublicates);

		var linkFromOldVariationToNewVariation = {};
		var linkFromOldLanguageToNewlanguage = {};
		var linkFromOldTagToNewTag = {};
		var linkFromOldSongDatabaseToNewSongDatabase = {};
		var linkFromOldEwDatabaseToNewEwDatabase = {};

		var newLanguages = [];
		data[2].forEach(language => {
			newLanguages.push({
				name: language.name,
				created_at: new Date()
			});
		});
		batchInsert(conn2, "languages", newLanguages).then((r) => {
			return conn2.query("select * from languages limit " + r.affectedRows);
		}).then((languages) => {
			languages.forEach((language, i) => {
				const oldLanguage = data[2][i];
				linkFromOldLanguageToNewlanguage[oldLanguage.id] = language.id;
			});
			var newVariations = [];
			filtteredVariations.forEach(variation => {
				newVariations.push({
					language_id: variation.language_id ? linkFromOldLanguageToNewlanguage[variation.language_id] : null,
					created_at: new Date()
				});
			});
			return batchInsert(conn2, "variations", newVariations);
		}).then((r) => {
			return conn2.query("select * from variations limit " + r.affectedRows);
		}).then(variations => {
			var newVariationVersions = [];
			variations.forEach((variation, i) => {
				const oldVariation = data[12][i];
				linkFromOldVariationToNewVariation[oldVariation.id] = variation.id
				newVariationVersions.push({
					variation_id: variation.id,
					name: oldVariation.name,
					text: oldVariation.text ? stringHelp.clearFront(stringHelp.clearBack(
						oldVariation.text.replace(/[-]/g, "")).
						replace(/\n\n\n\n\n/g, "\n\n").
						replace(/\n\n\n\n/g, "\n\n").
						replace(/\n\n\n/g, "\n\n").trim()
					) : "",
					version: oldVariation.version,
					created_at: new Date()
				});
			});
			return batchInsert(conn2, "variation_versions", newVariationVersions);
		}).then(() => {
			newTags = [];
			data[10].forEach(tag => {
				newTags.push({
					name: tag.name,
					created_at: new Date()
				});
			});
			return batchInsert(conn2, "tags", newTags);
		}).then((r) => {
			return conn2.query("select * from tags limit " + r.affectedRows);
		}).then((tags) => {
			tags.forEach((tag, i) => {
				const oldTag = data[10][i];
				linkFromOldTagToNewTag[oldTag.id] = tag.id;
			});
			var newSongDatabases = [];
			data[7].forEach(songDatabase => {
				newSongDatabases.push({
					name: songDatabase.name,
					created_at: new Date()
				});
			});
			return batchInsert(conn2, "song_databases", newSongDatabases);
		}).then((r) => {
			return conn2.query("select * from song_databases limit " + r.affectedRows);
		}).then((songDatabases) => {
			songDatabases.forEach((songDatabase, i) => {
				const oldSongDatabase = data[7][i];
				linkFromOldSongDatabaseToNewSongDatabase[oldSongDatabase.id] = songDatabase.id;
			});
			var newEwDatabases = [];
			data[1].forEach(ewDatabase => {
				newEwDatabases.push({
					name: ewDatabase.name,
					song_database_id: linkFromOldSongDatabaseToNewSongDatabase[ewDatabase.song_database_id],
					ew_database_key: ewDatabase.key,
					created_at: new Date()
				})
			});
			return batchInsert(conn2, "ew_databases", newEwDatabases);
		}).then((r) => {
			return conn2.query("select * from ew_databases limit " + r.affectedRows);
		}).then((ewDatabases) => {
			ewDatabases.forEach((ewDatabase, i) => {
				const oldEwDatabase = data[1][i];
				linkFromOldEwDatabaseToNewEwDatabase[oldEwDatabase.id] = ewDatabase.id;
			});
			songDatabaseVariationExists = [];
			tagVariationExists = [];
			return Promise.all([
				batchInsert(conn2, "song_database_variations", data[6].filter(h => {
					const sameFound = songDatabaseVariationExists.some(p => p.song_database_id === h.song_database_id && p.variation_id === h.variation_id);
					if (!sameFound) songDatabaseVariationExists.push({
						song_database_id: h.song_database_id,
						variation_id: h.variation_id
					});
					return !sameFound && 
						h.variation_id && 
						linkFromOldVariationToNewVariation[h.variation_id] &&
						linkFromOldSongDatabaseToNewSongDatabase[h.song_database_id];
				}).map(songDatabaseVariation => ({
					song_database_id: linkFromOldSongDatabaseToNewSongDatabase[songDatabaseVariation.song_database_id],
					variation_id: linkFromOldVariationToNewVariation[songDatabaseVariation.variation_id],
					created_at: new Date()
				}))),
				batchInsert(conn2, "tag_variations", data[9].filter(h => {
					const sameFound = tagVariationExists.some(p => p.tag_id === h.tag_id && p.variation_id === h.variation_id);
					if (!sameFound) tagVariationExists.push({
						tag_id: h.tag_id,
						variation_id: h.variation_id
					});
					
					return !sameFound &&
						h.variation_id &&
						linkFromOldVariationToNewVariation[h.variation_id] &&
						linkFromOldTagToNewTag[h.tag_id];
				}).map(tagVariation => ({
					tag_id: linkFromOldTagToNewTag[tagVariation.tag_id],
					variation_id: linkFromOldVariationToNewVariation[tagVariation.variation_id],
					created_at: new Date()
				})))
			]);
		}).then(() => {
			conn2.commit().then(() => {
				console.log("success");
				conn.end();
				conn2.end();
			}, 
			(err) => {
				console.log("err", err);
				conn.end();
				conn2.end();
			});
		})
	});	
}, err => {
	console.log("error lol", err);
});