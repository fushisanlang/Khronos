// index.js
// 获取应用实例

var a
var session
wx.request({
  //登录获取cookie
  url: 'https://daoserver.fushisanlang.cn/user/login', 
  data: {
    name: 'zhangyin',
    pass: '1'
  },
  header: {
    'content-type': 'application/json' // 默认值
  },
  success (res) {
    
    a=res.cookies
    session = a[0].split(";")[0]
    console.log(session)
   
  }
})
,

wx.request({


  //带cookie请求
  url: 'https://daoserver.fushisanlang.cn/role/get',
 
  header: {
    'content-type': 'application/json',
    'cookie': session
  },
  success (res) {
    console.log(res.data)
  }
})


