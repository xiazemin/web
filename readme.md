有时候误提交了某个文件 , 想把这个文件删掉 , 用下面的方法

# 删除包括历史
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch 文件相对路径' --prune-empty --tag-name-filter cat -- --all
# 同步到远程
git push origin master --force


还有就是如果设置了忽略文件 , 一定在提交之前就设置好 ,提交后的同样会被git管理
