   git filter-branch -f --tree-filter 'rm -rf https/server/server.crt' HEAD
   git filter-branch -f --tree-filter 'rm -rf https/server/server.key' HEAD
   git filter-branch -f --tree-filter 'rm -rf https/server.key' HEAD
   git filter-branch -f --tree-filter 'rm -rf https/server.crt' HEAD
   git push origin --force
