# SSH登录镜像模板使用说明 
成功启动该作业模板后，在个人的邮箱中会受到一封邮件，邮件中带有一个公共IP，此处记为public_ip，然后
使用ssh命令，登录该邮箱,具体登录步骤如下面说明：


## 登录 

ssh root@public  -p 1422(输入你指定的宿主机端口，默认提供的模板是1422)

## 输入密码
第一次登录会出现额外提示，提醒是否免密登录，输入yes，没有的话则是直接提示输入密码，该作业模板
密码如下：

**password**: 123456

## 自由配置
登录成功之后可以自由操作