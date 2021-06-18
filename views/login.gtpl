<html>
<head>
<title></title>
<script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.4.4/jquery.min.js"></script>
<script type="text/javascript" src="https://raw.github.com/GRINPublishing/GTPL/master/lib/gtpl.js"></script>
</head>
<body>
<form name="login" method="post" action="/login">
    使用者名稱:<input type="text" id="name" name="Username">
    密碼:<input type="password" id="password" name="Password">
    <input type="submit" id="submit-btn" name="submit" value="登入">
</form>
<form action="/register" method="get">
    <input type="submit" value="regist account" />
</form>
</body>
<script>
    $(document).ready(function(){
        $("#submit-btn").click(function(e)
        {
            $.ajax({
                url : "http://192.168.56.105:8080/login",
                type: "POST",
                data : { "Username": $('#name').val(), "Password": $('#password').val()},
                success: function(data){
                    console.log('AJAX SUCCESS, data : '+data); 
                },
                error: function(errMsg){ 
                    console.log('AJAX FAILED, message : '+errMsg);
                }
            });
            e.preventDefault();
        });
    });
</script>
</html>

