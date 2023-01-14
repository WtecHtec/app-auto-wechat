import React from "react";
import { useState, useEffect } from "react";
import {
  Button,
  Input,
  QRCode,
  Spin,
  message,
 Form, Switch, Tooltip, Radio, Tag , Modal, Image 
} from 'antd';
import { InfoCircleOutlined, SyncOutlined, CloseCircleOutlined } from '@ant-design/icons';
import cookie from 'react-cookies'
import { check, login, getAutoConfig, updateAutoConfig } from './Api'
const { TextArea } = Input;
const Base_Url =  'http://127.0.0.1:4299/'

const PB_QR = `${Base_Url}public/pb-qrcode.png`
let ws = null
function Home() {
  let codeSin = ''

  const [settingForm] = Form.useForm();
  const [messageApi, contextHolder] = message.useMessage();

  const [loginType, setLoginType] = useState('code')
  // loading \ login \ logined
  const [pageStatus, setPageStatus] = useState('loading')

  const [settingShow, setSettingShow] = useState(false)
  const [settingConfig, setSettingConfig] = useState({ 
    enable: false,
    auto_reply: false, 
    auto_reply_group: false,
    auto_bot: 'nobot',
    auto_desc: '正在忙ing',
    tuling_api_key: '',
  })

  const [isModalOpen, setIsModalOpen ] = useState(false)
  const [isScan, setIsScan ] = useState(false)
  const [qrImg, setQrImg ] = useState("")
  
  useEffect(() => {
    try {
      if (!ws) {
        ws = new WebSocket("ws://127.0.0.1:4299/ws");
        ws.onopen = function (res) {
          console.log('onopen: ws 连接', res)
        };
        ws.onmessage = function (e) {
            const msg = JSON.parse(e.data);
            console.log('onmessage==ws', msg)
            switch(msg.type) {
              case 'wxcodeqr': 
                  if (msg.status === 0) {
                    setQrImg(`${Base_Url}${msg.content}`)
                    setIsModalOpen(true)
                  } else if (msg.status === 1) {
                    setIsScan(true)
                  } else if (msg.status === 2) {
                    setIsModalOpen(false)
                    // setSettingConfig.enable = true
                    setSettingConfig( {...setSettingConfig, enable: true} )
                  } else if (msg.status === -1) {
                    setIsModalOpen(false)
                    setIsScan(false)
                    showErrorMsg('微信登录失败')
                  } 
                break
              default:
                console.log('onmessage==ws', msg)
            }
        };
        ws.onerror = function () {
          console.error('onerror: ws 连接失败')
        };
      } 
    } catch (error) {
      console.error('WebSocket 初始化失败', error)
    }
  }, [])


  const sendMsg = (msg) => {
    console.log('msg===', msg)
    if (!ws) return
    const data = JSON.stringify(msg);
    ws.send(data);
  }

  useEffect(() => {
    const minikey = cookie.load('minikey')
    if (!minikey) {
      setPageStatus('login')
      return
    }
    check().then(() => {
      setPageStatus('logined')
    }).catch(()=> {
      console.log('登录失效')
    })
  }, [])

  useEffect(()=> {
    if (pageStatus ===  'logined') {
      getAutoConfig().then((res) => {
        console.log( 'getAutoConfig', res)
        const { data } = res;
        const info = {
          enable: data.Enable,
          auto_reply: data.auto_reply, 
          auto_reply_group: data.auto_reply_group,
          auto_bot: data.auto_bot,
          auto_desc: data.auto_desc,
          tuling_api_key: data.tuling_api_key,
        }

        // settingForm.resetFields()
        settingForm.setFieldsValue({
          ...info
        })
        setSettingConfig({
          ...info
        })
        setSettingShow(data.auto_reply)
      })
    }
  }, [pageStatus])



  const showErrorMsg = (content)=> {
    messageApi.open({
      content,
      type: 'error',
    });
  }
  const handleLogin = ()=> {
    if (codeSin) {
      login({
        code: codeSin
      }).then(res => {
        const { code, minikey, info} = res.data
        if (code === 200 && minikey) {
          cookie.save('minikey', minikey)
          cookie.save('minip', info.Id)
          setPageStatus("logined")
        } else {
          showErrorMsg('体验码\登录指令错误')
        }
        console.log('handleLogin===', res)
      }).catch(()=> {
        showErrorMsg('体验码\登录指令错误')
      })
    }
  } 

  const CodeLogin = ()=> {
    return <div className="App-instructions">
      <img className="App-logo" src={ PB_QR } />
      <Input.Group compact style={{ minWidth: '400px', }} >
          <Input style={{ width: '20%',}} placeholder="体验码\登录指令" onChange={ (e) => codeSin = e.target.value } />
          <Button type="primary" onClick={ ()=> handleLogin()}>登录</Button>
      </Input.Group>
      <span style={ { marginTop: '24px', color: '#333333', }}> *扫码关注公众号,发送"体验码"获取体验码\登录指令, 
        <Button type="link" onClick={ ()=> setLoginType('qrcode')}>扫码登录</Button></span>
    </div>
  }

  const QrLogin = () => {
    return <div className="App-instructions">
      <QRCode value="5555"  style={{margin: 'auto', }} status="expired" onRefresh={() => console.log('refresh')} />
      <span style={ { marginTop: '24px', color: '#333333', }}> *微信小程序扫码, <Button type="link" onClick={ ()=> setLoginType('code')}>切换登录方式</Button></span>
    </div>
  }
  
  const Login = () => {
    return <div> 
      { loginType === 'code' ? <CodeLogin></CodeLogin> : <QrLogin></QrLogin> }
    </div>
  }

  const onSettingFinish = (values) => {
    const newValues = { ...values }
    console.log(newValues)
    updateAutoConfig(newValues)
  }

  const onChange = (check) => {
    setSettingShow(check)
    settingForm.submit()
  }

  const handleWxLogin = () => {
      sendMsg({
        type: 'login',
        content: cookie.load('minip')
      })
      setIsModalOpen(true)
  }


  const Setting = () => {
    return  <div style={{padding: '24px 10% 0 0'}}>
        <span> *如设置没有效果,可尝试多次刷新登录 </span>
        <Form form={settingForm} labelCol={{ span: 4 }} initialValues={ settingConfig } layout="horizontal" onFinish={onSettingFinish}>
          <Form.Item label="登录状态"  name="enable">
            { settingConfig.enable 
              ? <>
                <Tag icon={<SyncOutlined spin />} color="processing">在线</Tag>
                <Button type="link" style={{color: '#f50'}} onClick={ ()=> { setIsScan(false), handleWxLogin() } }>强制刷新</Button>
              </>
              : <>
              <Tag icon={<CloseCircleOutlined />} color="error">离线</Tag>
              <Button type="link" onClick={ ()=> handleWxLogin()}>登录</Button></>
            }
          </Form.Item>
          <Form.Item name="auto_desc_qr" label="绑定微信小程序">
              <Image width={200} src={`${Base_Url}public/${cookie.load('minip')}.png`}></Image>
          </Form.Item>
          <Form.Item label="自动回复" name="auto_reply" valuePropName="checked">
            <Switch  onChange={ (checked)=>  onChange(checked)}/>
          </Form.Item>
          <div style={ { visibility: settingShow ? 'visible' : 'hidden'}}>
            <Form.Item label="群@自动回复" name="auto_reply_group" valuePropName="checked">
              <Switch  onChange= { ()=> settingForm.submit() } />
            </Form.Item>
            <Form.Item name="auto_desc" label={ <Tooltip title="低于机器人回复优先级,100字以内">
                <span>自动回复文案 <InfoCircleOutlined style={{ color: '#66666 !important',}} /></span>
              </Tooltip> }>
              <TextArea rows={4} maxLength={100} onBlur={ ()=> settingForm.submit() }/>
            </Form.Item>
            <Form.Item name="auto_bot" label="机器人">
              <Radio.Group onChange= { ()=> settingForm.submit() }>
                <Radio value="nobot"> 无 </Radio>
                <Radio value="tuling"> 图灵 </Radio>
                <Radio disabled value="chatgpt"> ChatGPT </Radio>
              </Radio.Group>
            </Form.Item>
            <Form.Item name="tuling_api_key" label={ <Tooltip title="http://www.turingapi.com/">
                <span>图灵机器人APP_KEY <InfoCircleOutlined style={{ color: '#66666 !important',}} /></span>
              </Tooltip> }>
            <Input placeholder="填写app_key,使用图灵机器人"  onBlur={ ()=> settingForm.submit() } />
            </Form.Item>
          </div>
        </Form>
        <Modal  open={isModalOpen} closable={false} footer={null} width="240px">
          <div className="wx-qr">
            <Image width={200} src={qrImg}/>
            { 
              isScan && <div  className="wx-qr-drawer wx-qr-con">
                <div> 扫码成功</div>
              </div>
            }
          </div>
        </Modal>
    </div>
  }


  const PageDom = () => {
    if (pageStatus === 'loading') {
      return <div className="App-instructions"> <Spin></Spin> </div>
    } else if (pageStatus === 'login') {
      return <Login></Login>
    } else {
      return <Setting></Setting>
    }
  }
  return  <> {contextHolder} <PageDom></PageDom> </>
}

export default Home