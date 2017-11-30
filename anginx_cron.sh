#! /bin/bash
LANG="en_US"
sed_time=`date "+%d\\/%b\\/%Y"`
echo $sed_time
local_time=`date "+\/%d\\\\\\\\\\\/%b\\\\\\\\\\\/%Y\/"`
echo $local_time
str="startdate: $local_time"
sed -n "/$sed_time/p" /etc/nginx/logs/www.liuyangbao.com_443.access.log > temp.log
sed -i "s/^.*startdate.*$/$str/" /root/anginx/anginx_config.yml
/root/anginx/anginx-linux-amd64 -c /root/anginx/anginx_config.yml
rm -f /root/Anginx_*