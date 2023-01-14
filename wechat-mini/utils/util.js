const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : `0${n}`
}

const setStorage = (key, value) => {
	try {
		wx.setStorageSync(key, value)
	} catch (e) { }
}

const getStorage = (key) => {
	try {
		return wx.getStorageSync(key)
	} catch (e) { }
	return null
}

module.exports = {
	setStorage,
	getStorage
}
