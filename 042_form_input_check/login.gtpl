<html>
<head>
<title></title>
</head>
<body>
<form action="http://localhost:12138/login" method="post">
    使用者名称:<input type="text" name="username">
    中文名称:<input type="text" name="realname">
    密码:<input type="password" name="password">
    年龄:<input type="text" name="age">
    <input type="submit" value="登入">

    </br>
    下拉式功能表
    <select name="fruit">
    <option value="apple">apple<option>
    <option value="pear">pear<option>
    <option value="banana">banana<option>
    </select>

    </br>
    必选项
    <input type="radio" name="gender" value="M">男
    <input type="radio" name="gender" value="F">女

    </br>
    勾选
    <input type="checkbox" name="interest" value="interest0">兴趣0
    <input type="checkbox" name="interest" value="interest1">兴趣1
    <input type="checkbox" name="interest" value="interest2">兴趣2
</form>
</body>
</html>
