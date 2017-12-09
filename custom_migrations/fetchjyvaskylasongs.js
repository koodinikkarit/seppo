const request = require("request-promise");
const mysql = require("promise-mysql");
const batchInsert = require("./batchInsert");
const stringHelp = require("./stringHelp");

request("http://www.jklhelluntaisrk.fi/tykkiohjelma/songs_getUpdates.php?1499180362751&pw=1k2kDkaSkp&ts=1990-01-01+01%3A01%3A02").then((songs) => {
	const jsonSongs = JSON.parse(songs);
	let newVariations = [];
	let newKeyValues = [];
	let newAuthors = new Set();
	let newCopyrights = new Set();
	let variationAuthors = {};
	let variationCopyrights = {};
	let authorVariations = {};
	let copyrightVariations = {};
	jsonSongs.forEach((s, i) => {
		const name = s.name;
		const text = stringHelp.clearFront(stringHelp.clearBack(s.song.replace("-\r\n", "\n").
			replace("-k", "\n").
			replace(/[-]/g, "\n\n").
			replace(/\n\n\n\n\n/g, "\n\n").
			replace(/\n\n\n\n/g, "\n\n").
			replace(/\n\n\n/g, "\n\n").trim()));
		if (text.includes("-")) {
			console.log("contains -");
		}
		const year = parseInt(s.year);
		const author = s.lyrics_by;
		const copyright = s.copyright;
		if (author) {
			newAuthors.add(author);
			variationAuthors[i] = author;
			if (!authorVariations[author]) authorVariations[author] = new Set();
			authorVariations[author].add(i);
		}
		if (copyright && copyright != "null") {
			newCopyrights.add(copyright);
			variationCopyrights[i] = copyright;
			if (!copyrightVariations[i]) copyrightVariations[i] = new Set();
			copyrightVariations[i].add(i);
		}
		newVariations.push({
			name: name,
			text: text,
			year: !isNaN(year) ? year : null,
			author_id: null,
			copyright_id: null
		});
		let variationKeyValues = [];
		variationKeyValues.push({
			field_key: "added_at",
			field_value: s.added_at
		});
		if (s.added_by) {
			variationKeyValues.push({
				field_key: "added_by",
				field_value: s.added_by
			});
		}
		if (s.additional_info) {
			variationKeyValues.push({
				field_key: "additional_info",
				field_value: s.additional_info
			});
		}
		if (s.arrangement_by) {
			variationKeyValues.push({
				field_key: "arrangement_by",
				field_value: s.arrangement_by
			});
		}
		if (s.composed_by) {
			variationKeyValues.push({
				field_key: "composed_by",
				field_value: s.composed_by
			});
		}
		if (s.deleted) {
			variationKeyValues.push({
				field_key: "deleted",
				field_value: s.deleted
			});
		}
		if (s.modified) {
			variationKeyValues.push({
				field_key: "modified",
				field_value: s.modified
			});
		}
		if (s.orig_name) {
			variationKeyValues.push({
				field_key: "orig_name",
				field_value: s.orig_name
			});
		}
		if (s.songbook_id) {
			variationKeyValues.push({
				field_key: "songbook_id",
				field_value: s.songbook_id
			});
		}
		if (s.translated_by) {
			variationKeyValues.push({
				field_key: "translated_by",
				field_value: s.translated_by
			});
		}
		newKeyValues.push(variationKeyValues);
	});

	mysql.createConnection({
		host: "localhost",
		user: "jaska",
		password: "asdf321",
		database: "seppo"
	}).then(conn => {
		conn.beginTransaction();
		conn.query("insert into external_databases(name) values('Jyvaskyla')");
		Promise.all([
			conn.query("select * from authors"),
			conn.query("select * from copyrights"),
			conn.query("select * from variations v left join variation_versions vv on v.id = vv.variation_id"),
			conn.query("select * from variation_key_values")
		]).then(data => {
			const authors = data[0];
			const copyrights = data[1];
			const variations = data[2];
			const variationKeyValues = data[3];
			
			let filtteredAuthors = new Set();
			let filtteredCopyrights = new Set();
			let filtteredNewVariations = new Set();

			newAuthors.forEach(author => {
				var sameAuthor = authors.find(p => p.name === author);
				if (sameAuthor) {
					var authorVariation = authorVariations[author];
					if (authorVariation) {
						authorVariation.forEach(i => {
							newVariations[i].author_id = sameAuthor.id; 
						});
					}
				} else {
					filtteredAuthors.add(author);
				}
			});

			newCopyrights.forEach(copyright => {
				var sameCopyright = copyrights.find(p => p.name === copyright);
				if (sameCopyright) {
					let copyrightVariation = copyrightVariations[copyright];
					if (copyrightVariation) {
						copyrightVariations.forEach(i => {
							newCopyrights[i].copyright_id = sameCopyright.id;
						});
					}
				} else {
					filtteredCopyrights.add(copyright);
				}
			});
			let numberOfDublicates = 0;
			newVariations.forEach((variation, i) => {
				if (!variations.some(p => p.name === variation.name && p.text && variation.text)) {
					filtteredNewVariations.add(i);
				} else {
					numberOfDublicates++;
				}
			});

			console.log("Samojen laulujen määrä", numberOfDublicates);

			Promise.all([
				batchInsert(conn, "authors", Array.from(filtteredAuthors).map(p => ({
					name: p,
					created_at: new Date()
				}))),
				batchInsert(conn, "copyrights", Array.from(filtteredCopyrights).map(p => ({
					name: p,
					created_at: new Date()
				})))
			]).then(d => {
				const authorAffectedRows = d[0].affectedRows;
				const copyrightAffectedRows = d[1].affectedRows;
				return Promise.all([
					conn.query("select * from authors limit " + authorAffectedRows),
					conn.query("select * from copyrights limit " + copyrightAffectedRows)
				]);
			}, err => {
				console.log("err", err);
			}).then(data => {
				const addedAuthors = data[0];
				const addedCopyrights = data[0];

				addedAuthors.forEach(author => {
					const authorVariation = authorVariations[author.name];
					if (authorVariation) {
						authorVariation.forEach(p => {
							newVariations[p].author_id = author.id;
						});
					}
				});

				addedCopyrights.forEach(copyright => {
					const copyrightVariation = copyrightVariations[copyright.name];
					if (copyrightVariation) {
						copyrightVariation.forEach(p => {
							newVariations[p].copyright_id = copyright.id;
						});
					}
				});

				const newEws = newVariations.filter((p, i) => filtteredNewVariations.has(i)).map(e => ({
					year: e.year,
					author_id: e.author_id,
					copyright_id: e.copyright_id
				}));

				return batchInsert(conn, "variations", newEws);
			}, err => {
				console.log(err);
			}).then(r => {
				const affectedRows = r.affectedRows;
				return conn.query("select * from variations limit 10000 offset " + variations.length);
			}, err => {
				console.log("batch insert variations error");
			}).then(variations => {
				addNewKeyValues = [];
				let newVariationVersions = [];
				let newExternalVariations = [];
				let x = 0;			
				newVariations.forEach((newVariation, i) => {
					if (filtteredNewVariations.has(i)) {
						newExternalVariations.push({
							external_db_id: 1,
							variation_id: variations[x].id,
							external_id: jsonSongs[i].globalID
						});
						newVariationVersions.push({
							variation_id: variations[x].id,
							name: newVariation.name,
							text: newVariation.text,
							version: 1,
							created_at: new Date()
						});
						newKeyValues[i].forEach(keyValue => addNewKeyValues.push({
							variation_id: variations[x].id,
							field_key: keyValue.field_key,
							field_value: keyValue.field_value
						}));
						x++;
					}
				});
				return Promise.all([
					batchInsert(conn, "variation_key_values", addNewKeyValues),
					batchInsert(conn, "variation_versions", newVariationVersions),
					batchInsert(conn, "external_variations", newExternalVariations)
				]);
			}).then(() => {
				conn.commit().then(() => {
					console.log("success");
					conn.end();
				}, 
				(err) => {
					console.log("err", err);
					conn.end();
				});
			});
		});
	});
});