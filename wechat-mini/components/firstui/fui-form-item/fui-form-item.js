Component({
  options: {
    multipleSlots: true
  },
  properties: {
    //padding值
    padding: {
      type: String,
      value: '28rpx 32rpx'
    },
    //margin-top 单位rpx
    marginTop: {
      type: String,
      optionalTypes: [Number],
      value: 0
    },
    //margin-bottom 单位rpx
    marginBottom: {
      type: String,
      optionalTypes: [Number],
      value: 0
    },
    //标签文本
    label: {
      type: String,
      value: ''
    },
    //标题字体大小 默认使用全局设置值
    labelSize: {
      type: String,
      optionalTypes: [Number],
      value: 32
    },
    labelColor: {
      type: String,
      value: '#333'
    },
    //label宽度 rpx 默认使用全局设置值
    labelWidth: {
      type: String,
      optionalTypes: [Number],
      value: 140
    },
    //默认使用全局设置值
    labelRight: {
      type: String,
      optionalTypes: [Number],
      value: 16
    },
    //是否显示必填的红色星号
    asterisk: {
      type: Boolean,
      value: false
    },
    asteriskColor: {
      type: String,
      value: '#FF2B2B'
    },
    background: {
      type: String,
      value: '#fff'
    },
    highlight: {
      type: Boolean,
      value: false
    },
    arrow: {
      type: Boolean,
      value: false
    },
    arrowColor: {
      type: String,
      value: '#B2B2B2'
    },
    bottomBorder: {
      type: Boolean,
      value: true
    },
    borderColor: {
      type: String,
      value: '#EEEEEE'
    },
    //下边框left值，单位rpx
    left: {
      type: String,
      optionalTypes: [Number],
      value: 32
    },
    //下边框right值，单位rpx
    right: {
      type: String,
      optionalTypes: [Number],
      value: 0
    },
    radius: {
      type: String,
      value: '0'
    },
    param: {
      type: String,
      optionalTypes: [Number],
      value: 0
    }

  },
  methods: {
    handleClick() {
      this.triggerEvent('click', {
        param: this.data.param
      });
    }
  }
})