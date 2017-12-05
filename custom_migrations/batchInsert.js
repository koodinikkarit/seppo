const mysql = require("promise-mysql");

module.exports = (conn, tableName, items) => {
	return new Promise((resolve, reject) => {
		if (items.length === 0) {
			resolve({
				affectedRows: 0
			});
			return
		}
		console.log("Inserting to table", tableName, items.length);
		const fields = "(" + Object.keys(items[0]).join(",") + ")";
		const questonMarks = "(" + Object.keys(items[0]).map(p => "?").join(",") + ")";
		let sqlStr = `INSERT INTO ${tableName} ${fields} VALUES `;
		let valStrs = [];
		let vals = [];
		items.forEach(p => {
			valStrs.push(questonMarks);
			Object.keys(p).forEach(e => vals.push(p[e]));
		});
		sqlStr += valStrs.join(",");
		const sql = mysql.format(sqlStr, vals);
		resolve(conn.query(sql));
	});
}