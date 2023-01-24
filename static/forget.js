const email = document.querySelector("#email");
const captcha = document.querySelector("#captcha");
const password = document.querySelector("#password");
// 返回登录页面
document.querySelector('.return').addEventListener('click', function (e) {
    e.preventDefault();
    window.location.href = "login";
});

//验证码请求
document.querySelector('.sendcaptcha').addEventListener('click', function (e) {
    e.preventDefault();
    //TODO: 实现发送重置密码邮件的操作
    // send a request to the server to send the captcha
    let formData = new FormData();
    formData.append("email", email.value);

    fetch('/forget', {
        method: 'POST',
        body: formData,
    })
        .then(function (response) {
            return response.text();
        })
        .then(function (data) {
            console.log(data.FormData)
            alert(data)
            //     document.querySelector('.forget-password-container').style.display = 'none';
            // document.querySelector('.forget-password-overlay').style.display = 'block';
        })
        .catch(function (error) {
            console.log(error);
        });

});

//重置密码请求
document.querySelector('.reset-password').addEventListener('click', function (e) {
    e.preventDefault();

    var data = {
        email: email.value,
        captcha: captcha.value,
        password: password.value,
    };

    console.log(data)
    fetch('/update', {
        method: 'POST',
        body: JSON.stringify(data),
    })
        .then(function (response) {
            return response.text();
        })
        .then(function (data) {
            alert(data)
            //     document.querySelector('.forget-password-container').style.display = 'none';
            // document.querySelector('.forget-password-overlay').style.display = 'block';
        })
        .catch(function (error) {
            console.log(error);
        });

});