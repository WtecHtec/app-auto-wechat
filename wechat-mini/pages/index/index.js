import {postCheckLogin, getLoginCode, postLogin,
   getUserInfo, postUserInfo, postUpdateUserInfo,
   checkScanCode, scanCode } from '../servers/login'
import { setStorage } from '../../utils/util';
import { USERINFO_KEY, MINIKET_KEY } from '../../utils/storage-keys'
import { PAGE_STATUS, RE_OPT } from '../../utils/config'
let cache = ''
Page({
  data: {
    pageStatus: PAGE_STATUS.loading,
    from: 'init',
    noticeId: '',
    fromData: {},
  },
	onLoad(options) {
    options.from && (this.data.from = options.from);
    options.noticeId && (this.data.noticeId = options.noticeId)
    this.optNums = 0
		this.checkLoin()
  },
	async checkLoin() {
		this.setData({ pageStatus: PAGE_STATUS.loading })
		const [err, res] = await postCheckLogin();
		if (!err && res) {
      if (res.code === 200) {
        this._postUserInfo();
        return
      }
      const [e, info]  = await getUserInfo();
      if (!e && info) {
        const code = await getLoginCode();
				if (!code) {
					// 进入兜底
					this.setData({ pageStatus: PAGE_STATUS.error })
					return
				}
			  await this.handLogin(code, this.data.noticeId, info.nickName);
      } else {
        // 拒绝授权，进入兜底页
				this.setData({ pageStatus: this.optNums >= RE_OPT ? PAGE_STATUS.reopt : PAGE_STATUS.noright })
      }
		} else {
			this.setData({ pageStatus: PAGE_STATUS.nonetwork })
    }
	},
	async handLogin(code, noticeId, nickName) {
		const [err, res] = await postLogin(code, noticeId, nickName);
		if (!err && res && res.code === 200) {
      setStorage(MINIKET_KEY, res.minikey)
      this._postUserInfo();
		} else {
      this.setData({ pageStatus: PAGE_STATUS.error })
    }
	},
	bindRight() {
		this.optNums += 1 
		this.checkLoin()
  },
  async _postUserInfo() {
    const [err, res] = await postUserInfo();
    if (!err && res && Object.keys(res).length) {
      this.setData({ 
        pageStatus: PAGE_STATUS.normal,
        fromData: res 
       })
    } else {
      // 进入兜底
      this.setData({ pageStatus: PAGE_STATUS.error })
    }
  },
  bindReJoin() {
    // 清除缓存
    setStorage(USERINFO_KEY, null)
    setStorage(MINIKET_KEY, null)
    wx.redirectTo({ url: `/pages/index/index` });
  },

  fabClick() {
    wx.scanCode({
      onlyFromCamera: true,
      async success (res) {
        console.log(res)
        const { result } = res
        if (result) {
          const [err,] = await checkScanCode(result)
          if (!err) {
            wx.showModal({
              content: '是否登录?',
              success (res) {
                if (res.confirm) {
                  console.log('用户点击确定')
                  scanCode(result)
                } else if (res.cancel) {
                  console.log('用户点击取消')
                }
              }
            })
          } else {
            wx.showToast({
              title: '确认信息失败',
              icon: 'none',
              duration: 2000
            })
          }
        }
        
      }
    })
  },

  bubble() {
    wx.showToast({
      title: '前往PC端设置',
      icon: 'none',
      duration: 2000
    })
  },
  bindFromDataChange(e) {
    console.log(e)
    const { prop } = e.currentTarget.dataset
    const { value } = e.detail
    if (cache === prop) return
    cache = prop
    const upObj = {}
    upObj[`fromData.${prop}`] = value
    this.setData({
      ...upObj
    }, ()=> {
      postUpdateUserInfo(this.data.fromData)
    })
  }

})
