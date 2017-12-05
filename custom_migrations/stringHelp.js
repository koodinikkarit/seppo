
function clearFront(text) {
	if (text.startsWith("\n")) {
		return clearFront(text.slice(1))
	}
	return text
}

function clearBack(text) {
	if (text.endsWith("\n")) {
		return clearBack(text.slice(0, text.length - 1));
	}
	return text;
}

module.exports = {
	clearFront,
	clearBack
};