// 1. 引入 axios
import axios from 'axios'
import cookie from 'react-cookies'
import React from 'react';
import { Modal } from 'antd';
function info() {
    Modal.info({
      title: '系统提示',
      content: (
        <div>
          <p>登陆失效,请重新登陆！</p>
        </div>
      ),
      onOk() {
          cookie.remove('minikey')
          window.location.href = '/'
      },
    });
}
// 2. 创建axios对象，配置默认配置
const httpRequest = axios.create({
    // baseURL: 'https://sr7.top/wx/', // api的base_url
    baseURL: 'http://127.0.0.1:4299/', // api的base_url
    timeout: 15000 // 请求超时时间

})
axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
// 3. 创建 request拦截器
httpRequest.interceptors.request.use(config => { 
    config.headers.authorization = 'Bearer ' + cookie.load('minikey')
    return config
}, error => {
    Promise.reject(error)
})

// 4. 创建response 拦截器
httpRequest.interceptors.response.use(response => response, error => {
    if (error && error.response) {
        switch (error.response.status) {
            case 400:
                error.message = '请求错误'
                break
            case 401:
                if (error.request.responseURL.indexOf('/login') === -1) {
                  info()
                }
                error.message = '登陆失效'
                break
            case 402:
                error.message = '权限不足,请联系管理员'
                break

            case 404:
                error.message = `请求地址出错: ${error.response.config.url}`
                break

            case 500:
                error.message = error.response.data.errorMsg
                break

            default:
                error.message = '服务器内部错误'
        }
    }
    return Promise.reject(error)
})
// 5. 暴露出去
export default httpRequest
