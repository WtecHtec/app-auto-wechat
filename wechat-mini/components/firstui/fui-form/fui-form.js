import form from './fui-validator.js'
Component({
  properties: {
    //表单padding值
    padding: {
      type: String,
      value: "0"
    },
    //是否显示校验错误信息
    show: {
      type: Boolean,
      value: true
    },
    //是否禁用该表单内的所有组件,透明遮罩层
    disabled: {
      type: Boolean,
      value: false
    },
    //提示框top值 px
    top: {
      type: Number,
      optionalTypes: [String],
      value: 0
    },
    left: {
      type: Number,
      optionalTypes: [String],
      value: 24
    },
    right: {
      type: Number,
      optionalTypes: [String],
      value: 24
    },
    //错误提示框背景色
    background: {
      type: String,
      value: ''
    },
    //错误提示字体大小
    size: {
      type: Number,
      optionalTypes: [String],
      value: 28
    },
    //错误提示字体颜色
    color: {
      type: String,
      value: '#fff'
    },
    //错误提示框圆角值
    ridus: {
      type: Number,
      optionalTypes: [String],
      value: 16
    },
    //错误消息显示时间 ms
    duration: {
      type: Number,
      value: 2000
    }
  },
  data: {
    errorMsg: '',
    timer: null
  },
  lifetimes: {
    detached: function () {
      this.clearTimer()
    }
  },
  methods: {
    clearTimer() {
      clearTimeout(this.data.timer)
      this.data.timer = null;
    },
    /*
        @param model 表单数据对象
        @param rules 表单验证规则
        @param checkAll 校验所有元素
       */
    validator(model, rules, checkAll = false) {
      return new Promise((resolve, reject) => {
        try {
          let res = form.validator(model, rules, checkAll);
          if (!res.isPassed && this.data.show) {
            this.clearTimer()
            let errors = res.errorMsg;
            if (checkAll) {
              errors = errors[0].msg
            }
            this.setData({
              errorMsg: errors
            }, () => {
              this.data.timer = setTimeout(() => {
                this.setData({
                  errorMsg: ''
                })
              }, this.data.duration)
            })
          }
          resolve(res)
        } catch (e) {
          reject({
            isPassed: false,
            errorMsg: '校验出错，请检查传入的数据格式是否有误！'
          })
        }
      })
    }
  }
})