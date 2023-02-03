/**
 email: /^[\w+-.]+@[a-z\d-]+(\.[a-z\d-]+)*\.[a-z]+$/,
 phone: /^[1,8][0-9]{10,11}$/,
 imageCaptcha: /^[a-zA-Z0-9]{6}$/,
 captcha: /^[a-zA-Z0-9]{6}$/
 */
jQuery( function() {
    //发送验证码
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

    $("#change_avator").click(function(){
        $("#avator_file").click();
    });
    $("#avator_file").change(function(){

        //获取文件
        // var file = $("#avator_file").prop('files')['0'];
        // var file = document.getElementById("avator_file").files['0'];
        //创建读取文件的对象
        // var reader = new FileReader();
        //创建文件读取相关的变量
        // var imgFile;
        //为文件读取成功设置事件
        // reader.onload=function(e) {
        //     var e=window.event||e;
        //     imgFile = e.target.result;
        //     console.log(imgFile);
        // };
        // var formData = new FormData();
        // formData.append("upload", imgFile)
        // console.log(formData)
        //正式读取文件
        // reader.readAsDataURL(file);

        var maxSize = 2097152;
        var fileArray = document.getElementById("avator_file").files;
        var filesName = fileArray[0].name;
        var filesSize = fileArray[0].size;
        console.log(fileArray);
        console.log(filesName);
        console.log(filesSize);
        if (filesSize > maxSize) {
            alert('单张图片大小不能超过2Mb');
            $('#avator_file').val('');
            return;
        }
        var formData = new FormData();
        formData.append("uploadfile", fileArray['0'])
        $.ajax({
            url:"/admin/uploadImg",
            data:formData,
            type:"post",
            processData:false, //不序列化data
            contentType:false,
            // ContentType:"application/json",
            success:function (res) {
                if(res.code ===0){
                    $(".img-avatar").attr('src', res.fileUrl);
                    $("#creator_pic").val(res.fileUrl)
                    lightyear.notify('图片上传成功！~', 'success', 3000, 'mdi mdi-emoticon-happy', 'top', 'center');
                }
            }


        })
    })
})