<script type="text/javascript" src="http://code.jquery.com/jquery-1.7.1.min.js"></script>
<html>
<head>
<title></title>
</head>
<body>
<form name="login" action="/login" method="post">
    使用者名稱:<input type="text" name="Username">
    密碼:<input type="password" name="Password">
    <input type="submit" value="登入">
</form>
<form action="/register" method="get">
    <input type="submit" value="regist account" />
</form>
</body>
<script>
    $(document).ready(function(){
        $('[name="login"]').click(function(e)
        {
            var MyForm = JSON.stringify($('[name="Username"]').serializeJSON());
            console.log(MyForm);
            $.ajax({
                url : "http://localhost:8080/login",
                type: "POST",
                data : MyForm,

            });
            e.preventDefault(); //STOP default action

        });
    });
</script>
</html>