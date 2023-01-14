const EVENTS = {} // 事件集合

export const watch = (evenName, fn) => {
  if (typeof fn === 'function') {
    EVENTS[evenName] = fn
  }
}

export const dispatch = (evenName, ...arg) => {
  if (typeof EVENTS[evenName] === 'function') {
    EVENTS[evenName](...arg)
  }
} 

export const rmoveEvent = (evenName) => {
  delete EVENTS[evenName]
}