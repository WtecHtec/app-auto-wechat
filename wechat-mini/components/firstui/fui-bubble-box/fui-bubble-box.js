// 本文件由FirstUI授权予车永钊（手机号：18276    45 30     19，身份证尾号：2  270 1  0）专用，请尊重知识产权，勿私下传播，违者追究法律责任。
Component({
  options: {
    multipleSlots: true
  },
  properties: {
    show: {
      type: Boolean,
      value: false
    },
    items: {
      type: Array,
      value: []
    },
    width: {
      type: Number,
      optionalTypes: [String],
      value: 300
    },
		padding: {
				type: Array,
				value: ['26rpx', '32rpx']
		},
    position: {
      type: String,
      value: 'fixed'
    },
    left: {
      type: Number,
      optionalTypes: [String],
      value: 0
    },
    right: {
      type: Number,
      optionalTypes: [String],
      value: 8
    },
    top: {
      type: Number,
      optionalTypes: [String],
      value: 0
    },
    bottom: {
      type: Number,
      optionalTypes: [String],
      value: 0
    },
    direction: {
      type: String,
      value: 'bottom'
    },
    zIndex: {
      type: Number,
      optionalTypes: [String],
      value: 1001
    },
    background: {
      type: String,
      value: '#fff'
    },
    size: {
      type: Number,
      optionalTypes: [String],
      value: 32
    },
    color: {
      type: String,
      value: '#181818'
    },
    fontWeight: {
      type: Number,
      optionalTypes: [String],
      value: 400
    },
    showLine: {
      type: Boolean,
      value: true
    },
    lineColor: {
      type: String,
      value: '#eee'
    },
    triangle: {
      type: Object,
      value: {}
    },
    isMask: {
      type: Boolean,
      value: true
    },
    maskBackground: {
      type: String,
      value: 'transparent'
    },
    maskClosable: {
      type: Boolean,
      value: true
    }
  },
  methods: {
    handleClose() {
      if (!this.data.maskClosable) return;
      this.triggerEvent('close', {});
    },
    handleClick(e) {
      let index = Number(e.currentTarget.dataset.index)
      this.triggerEvent('click', {
        index: index
      });
    },
    stop() {
      return false;
    }
  }
})