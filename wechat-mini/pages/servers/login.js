
import Request from '../../utils/request'
import { USERINFO_KEY } from '../../utils/storage-keys'
import { setStorage } from '../../utils/util'

/**
 * 获取微信登陆code
 * @returns 
 */
export function getLoginCode() {
	return new Promise(resolve => {
		wx.login({
			success (res) {
				if (res && res.code) {
					resolve(res.code)
				}
				resolve(null)
			},
			fail() {
				resolve(null)
			}
		})
	})
};

/**
 * 请求用户获取用户信息
 */
export function getUserInfo() {
  return new Promise(resolve => {
    wx.showModal({
      title: '温馨提示',
      content: '正在请求获取您的个人信息，以后续方便使用本程序！',
      success(res) {
        if (res.confirm) {
          wx.getUserProfile({
            desc: "获取你的昵称、头像、地区及性别",
            success: res => {
              if (res && res.userInfo) {
                setStorage(USERINFO_KEY, res.userInfo)
                resolve([null,  res.userInfo])
              } else {
                resolve(['获取失败',  null])
              }
            },
            fail: res => {
              //拒绝授权
              resolve(['拒绝授权',  null])
              return;
            }
          })
        } else if (res.cancel) {
          resolve(['您拒绝了请求',  null])
          return;
        }
      }
    })
  });
}

/**
 * 登录成功后页面跳转逻辑
 * @param {*} from 
 */
export function  NavToPage(from, noticeId) {
  if (from === 'init') {
    wx.redirectTo({ url: '/pages/home/home' });
  } else if (from === 'collect') {
    // 转发收集
    wx.redirectTo({ url: `/pkgDetail/pages/msg_collect/index?noticeId=${noticeId}` });
  }
}
/**
 * 请求登陆
 * @param {*} code 
 * @returns 
 */
export function postLogin(code, noticeId, nickName) {
	return Request('/login', { code, nickName, id: noticeId, login_type: 'wx_app'})
}

/**
 * 检查登录态
 * @returns 
 */
export function postCheckLogin() {
	return Request('/auth/check');
}

/**
 * 获取用户信息
 * @returns 
 */
 export function postUserInfo() {
	return Request('/auth/getautoconfig', '', 'GET');
}

/**
 * 完善用户信息
 * @returns 
 */
 export function postUpdateUserInfo(info) {
	return Request('/auth/updateautoconfig', { ...info });
}
