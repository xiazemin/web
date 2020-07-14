 #!/bin/bash
 file="lowprice.conf"
 curl -O http://10.96.83.51:8999/download/$file
 tar -zxvf  $file -C /

  #tar -czvf lowprice.conf /home/conf /home.json /home/tus.json

 #scp lowprice.conf  xia@10.96.83.51:/home/aladdin/download

 #http://10.96.83.51:8999/download/
 #bash <(curl http://10.96.83.51:8999/download/low.sh)