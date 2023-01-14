import autoUpdate from './utils/update'
App({
  onLaunch() {
		autoUpdate()
  },

  globalData: {
    userInfo: null
  }
})
