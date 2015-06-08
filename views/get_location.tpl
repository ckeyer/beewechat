<!DOCTYPE html>

<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

    	<title><?.PageTitle?></title>
</head>

<body>

<script type="text/javascript" src="http://res.wx.qq.com/open/js/jweixin-1.0.0.js"></script>
<script type="text/javascript">
  document.write("Opened<br>");    
  // 微信SDK 获取用户地址
 function wxjs_load(){
  document.write("load Start<br>");    
    wx.config({
    debug: true, // 开启调试模式,调用的所有api的返回值会在客户端alert出来，若要查看传入的参数，可以在pc端打开，参数信息会通过log打出，仅在pc端时才会打印。
    appId: '<? .AppID ?>', // 必填，公众号的唯一标识
    timestamp: <? .TimeStamp ?>, // 必填，生成签名的时间戳
    nonceStr: '<? .Noncestr ?>', // 必填，生成签名的随机串
    signature: '<? .Signature ?>',// 必填，签名，见附录1
    jsApiList: ['openLocation', 'getLocation'] // 必填，需要使用的JS接口列表，所有JS接口列表见附录2
});
document.write("load Over<br>");    
  };

  wx.ready(function(){
document.write("ok<br>");

wx.getLocation({
    success: function (res) {
        var latitude = res.latitude; // 纬度，浮点数，范围为90 ~ -90
        var longitude = res.longitude; // 经度，浮点数，范围为180 ~ -180。
        var speed = res.speed; // 速度，以米/每秒计
        var accuracy = res.accuracy; // 位置精度
        document.write(res.latitude+'<br>')
        document.write(res.longitude+'<br>')
        document.write(res.speed+'<br>')
        document.write(res.accuracy+'<br>')
    }
});

    // config信息验证后会执行ready方法，所有接口调用都必须在config接口获得结果之后，config是一个客户端的异步操作，所以如果需要在页面加载时就调用相关接口，则须把相关接口放在ready函数中调用来确保正确执行。对于用户触发时才调用的接口，则可以直接调用，不需要放在ready函数中。
});
  wx.error(function(res){

document.write("error <br>");
    // config信息验证失败会执行error函数，如签名过期导致验证失败，具体错误信息可以打开config的debug模式查看，也可以在返回的res参数中查看，对于SPA可以在这里更新签名。

});
   window.onload = wxjs_load;
</script>
  
</body>
</html>
