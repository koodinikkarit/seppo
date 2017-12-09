const request = require("request");
const fs = require('fs');

request("http://www.jklhelluntaisrk.fi/tykkiohjelma/songs_getUpdates.php?1499180362751&pw=1k2kDkaSkp&ts=1990-01-01+01%3A01%3A02", (error, response, body) => {
	var jsonBody = JSON.parse(body);
	jsonBody.forEach(song => {
		try {
			fs.writeFileSync("./lauludata2/" + song.name + ".txt", song.song.replace("-k", "").replace("-", "\n\n"));
		} catch(e) {

		}
	});
})