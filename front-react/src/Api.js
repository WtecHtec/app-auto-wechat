import httpRequest from './requset'

/**
 * 检查是否登录
 * @returns
 */
export function check() {
  return httpRequest.post('/auth/check')
}

/**
 * 登录
 * @returns
 */
export function login(param) {
  return httpRequest.post('/login', param)
}

/**
 * 获取配置
 * @returns
 */
export function getAutoConfig() {
  return httpRequest.get('/auth/getautoconfig')
}

/**
 * 更新配置
 * @returns
 */
export function updateAutoConfig(param) {
  return httpRequest.post('/auth/updateautoconfig', param)
}





