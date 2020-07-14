有时候误提交了某个文件 , 想把这个文件删掉 , 用下面的方法

# 删除包括历史
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch 文件相对路径' --prune-empty --tag-name-filter cat -- --all
# 同步到远程
git push origin master --force


还有就是如果设置了忽略文件 , 一定在提交之前就设置好 ,提交后的同样会被git管理

2.单独设置

在项目父级目录下 找到.git文件（此文件为隐藏文件 需要 ls -a 命令查看）
打开.git下的config文件 在最下方增加
[user]
 
    name = 'xxxx'
    email = 'xxx@163.com'
 

或者直接在项目目录下运行命令：

git config  user.name xxx
 
git config  user.email xxx@163.com

