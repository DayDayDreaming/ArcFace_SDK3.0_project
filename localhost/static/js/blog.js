$(document).ready(function () {
    //注册
    $("#register-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            },
            // repassword: {
            //     required: true,
            //     rangelength: [5, 10],
            //     equalTo: "#register-password"
            // }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            },
            repassword: {
                required: "请确认密码",
                rangelength: "密码必须是5-10位",
                equalTo: "两次输入的密码必须相等"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/register";
            alert("urlStr:"+urlStr)
            $.ajax({
                url: urlStr,
                type: "post",
                dataType: "json",
                data: $(form).serialize(),
                success: function (data, status) {
                    alert("data:" + data.message)
                    if (data.Islogin == false){
                        window.location.href = "/error"
                    } else {
                        if (data.code == 1) {
                            console.log(data)
                            setTimeout(function () {
                                window.location.href = "/"
                            }, 1000)
                        }
                    }

                },
                err: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            })
        }
    });


    $("#login-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login"
            alert("urlStr:" + urlStr)
            $.ajax({
                url: urlStr,
                type: "post",
                dataType: "json",
                data: $(form).serialize(),
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.code == 1) {
                        console.log(data)
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    } else {
                        window.location.href = "/error"
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });


    $("#remove-form").validate({

        submitHandler: function (form) {
            var urlStr = "/remove"
            alert("urlStr:" + urlStr)
            $.ajax({
                url: urlStr,
                type: "post",
                dataType: "json",
                data: $(form).serialize(),
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.Islogin == false){
                        window.location.href = "/error"
                    } else {
                        if (data.code == 1) {
                            console.log(data)
                            setTimeout(function () {
                                window.location.href = "/"
                            }, 1000)
                        }
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });

    $("#add-form").validate({

        submitHandler: function (form) {
            var urlStr = "/add"
            alert("urlStr:" + urlStr)
            $.ajax({
                url: urlStr,
                type: "post",
                dataType: "json",
                data: $(form).serialize(),
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.Islogin == false){
                        window.location.href = "/error"
                    } else {
                        if (data.code == 1) {
                            console.log(data)
                            setTimeout(function () {
                                window.location.href = "/"
                            }, 1000)
                        }
                    }

                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });

    $("#delete-form").validate({

        submitHandler: function (form) {
            var urlStr = "/delete"
            alert("urlStr:" + urlStr)
            $.ajax({
                url: urlStr,
                type: "post",
                dataType: "json",
                data: $(form).serialize(),
                success: function (data, status) {
                    alert("data:" + data.message + ":" + status)
                    if (data.Islogin == false){
                        window.location.href = "/error"
                    } else{
                        if (data.code == 1) {
                            console.log(data)
                            setTimeout(function () {
                                window.location.href = "/"
                            }, 1000)
                        }
                    }

                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)

                }
            });
        }
    });

});