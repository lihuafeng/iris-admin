/**
 email: /^[\w+-.]+@[a-z\d-]+(\.[a-z\d-]+)*\.[a-z]+$/,
 phone: /^[1,8][0-9]{10,11}$/,
 imageCaptcha: /^[a-zA-Z0-9]{6}$/,
 captcha: /^[a-zA-Z0-9]{6}$/
 */
jQuery( function() {
    $("#sendEmailCode").bind('click',function(){
        var countTime = 60
        var email = $("#email").val();
        var email_code = $("#email_code").val();
        var username = $("#username").val();
        if(email ==""){
            lightyear.notify('请输入绑定邮箱~', 'danger', 3000, 'mdi mdi-emoticon-happy', 'top', 'center');
            return false
        }
        if(!/^[\w+-.]+@[a-z\d-]+(\.[a-z\d-]+)*\.[a-z]+$/.test(email)){
            lightyear.notify('请输入正确的邮箱~', 'danger', 3000, 'mdi mdi-emoticon-happy', 'top', 'center');
            return false
        }

        $.ajax({
            url:"/admin/send_email_code",
            type: "POST",
            data: {
                username:username,
                email:email,
            },
            dataType: "json",
            async: false,  // 默认是true
            success:function(result){
                if(!result.code){
                    $(this).attr("disabled","true")
                    //倒计时
                    setTimeout(function countTimer() {
                        if ( countTime-- === 0 ) {
                            $("#sendEmailCode").html("获取验证码");
                            //启用按钮
                            $("#sendEmailCode").removeAttr("disabled")
                        }else{
                            $("#sendEmailCode").html(countTime + "s后重新发送");
                            setTimeout(countTimer, 1000);
                        }
                    })
                }else{
                    lightyear.notify('发送验证码失败，请重试！~', 'danger', 3000, 'mdi mdi-emoticon-happy', 'top', 'center');
                    return false
                }

            }});

    })
})