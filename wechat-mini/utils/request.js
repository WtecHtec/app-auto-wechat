import { mode, BASE_URL } from "../config";
import { getStorage } from "./util";
import { MINIKET_KEY } from './storage-keys'
const Request = (url, data, method = 'POST') => {
	const minikey = getStorage(MINIKET_KEY) || '';
  return new Promise((resolve, reject) => {
    wx.request({
      url: `${BASE_URL[mode]}${url}`,
      data,
      header: {
        'content-type': 'application/json',
        'Authorization': `Bearer ${minikey}`
      },
			method,
      success (res) {
				if (res && res.data && typeof res.data === 'object') {
					return resolve(['', res.data]);
				}
			  resolve([`post ${url} response error `, null]);
      },
			fail() {
				resolve([`post ${url} error`, null]);
			}
    })
  })
}
export default Request