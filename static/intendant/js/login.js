/**
 * Created by iteny on 2016/3/3.
 */
// function asdf(e){
//     if (event.keyCode==13){
//         document.getElementById("loginsubmit").click();
//     }
// }
// //刷新验证码
// function refreshs(){
//     document.getElementById('code_img').src=verifycode+'?time='+Math.random();void(0);
// }
$(function(){
    $('.close').on('click', function(c){
		$('#username').val("");
        $('#password').val("");
	});

    $("input").removeAttr("disabled");
    $(':input').focus(function(){
        $(this).parents('.inputs').addClass('focus');
    }).blur(function(){
        $(this).parents('.inputs').removeClass('focus');
    });
    var formLogin = $('#form-login');
    formLogin.submit(function(e){
        // e.preventDefault();
        // alert("sdfasd");
        var username = $.trim($('#username').val()),
            password = $.trim($('#password').val()),
            userreg = /^[a-zA-Z][a-zA-Z0-9_]{4,15}$/;
            // verify = $.trim($('#verify').val());
        if (username.length === 0)
        {
            admin.error('用户名不能为空','#userli');
            return false;
        }
        else if (password.length === 0)
        {
            admin.error('密码不能为空','#passli');
            return false;
        }else if(!userreg.test(username))
        {
            admin.error('以字母开头，长度在5-15位之间，只能包含字符、数字和下划线。 ','#userli');
            return false;
        }else if(!userreg.test(password))
        {
            admin.error('以字母开头，长度在5-15位之间，只能包含字符、数字和下划线。 ','#passli');
            return false;
        }
        // else if (verify.length === 0)
        // {
        //     admin.error('验证码不能为空','#verifyli');
        //     return false;
        // }
        else
        {
            e.preventDefault();
            if(formLogin.attr('disabledSubmit')){
                admin.error('请勿重复登录','#loginsubmit');
                return false;
            }
            formLogin.attr('disabledSubmit',true);
            var param = formLogin.serialize();
            $.ajax({
                type: 'post',
                url: formLogin.attr('action'),
                data : param,
                dataType:"json",
                beforeSend: function(){
                    myload = layer.load(0,{time:3*1000});
                },
                success: function(msg){
                    if(msg.status === 1){
                        admin.success(msg.info,'#loginsubmit');
                        setTimeout(function(){window.location.href=redirect}, 3000);
                    }else{
                        layer.close(myload);
                        admin.error(msg.info,'#loginsubmit');
                        formLogin.attr('disabledSubmit','');
                        // refreshs();
                        // $('.yanzheng_img').eq(0).click();
                        // $('#verify').val('')
                    }
                },
                error: function(XMLHttpRequest, textStatus, errorThrown){
                    admin.error('网络连接异常！','#loginsubmit');
                    return false;
                }
            });

        }
    });

});
