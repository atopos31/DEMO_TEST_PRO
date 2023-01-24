const container = document.getElementsByClassName('container')[0];
const signIn = document.getElementById('sign-in');
const signUp = document.getElementById('sign-up');
const username2 = document.querySelector('#username2');
const email2 = document.querySelector('#email2');
const password2 = document.querySelector('#password2');
const email = document.querySelector('#email1');
const password = document.querySelector('#password1');
const captcha = document.querySelector('#captcha');


signUp.onclick = function () {
    container.classList.add('active');
}

signIn.onclick = function () {
    container.classList.remove('active');
}

//登录请求
document.querySelector('.signIn').addEventListener('click', function () {
    var data = {
        email: email.value,
        password: password.value
    };


    fetch('/login', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(function (response) {
            return response.json();
        })
        .then(function (data) {
            console.log(data)
            if (data.ok) {
                //通过query参数定位指定页面
                window.location.href = '/home?name='+data.username+'&key='+data.key;
            } else {
                alert('Incorrect email or password, please try again.');
            }
        })
        .catch(function (error) {
            console.error('Error:', error);
        });

});

//验证码发送请求
document.querySelector('.send-captcha').addEventListener('click', function () {
    // send a request to the server to send the captcha
    let formData = new FormData();
    formData.append("email", email2.value);
    fetch('/send-captcha', {
        method: 'POST',
        body: formData,
    })
        .then(function (response) {
            // handle the response from the server
            return response.text();
        })
        .then(function(data){
            alert(data)
        })
        .catch(function (error) {
            console.log(error);
        });
});

//注册请求
document.querySelector('.signUp').addEventListener('click', function () {
    var data = {
        username: username2.value,
        email: email2.value,
        password: password2.value,
        captcha: captcha.value
    };
    console.log(data)
    fetch('/register', {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            'Content-Type': 'application/json'
        }
    })
        .then(function (response) {
            // handle the response from the server
            return response.text();
        })
        .then(function(data){
            alert(data)
            console.log(data)
        })
        .catch(function (error) {
            console.log(error);
        });
});    