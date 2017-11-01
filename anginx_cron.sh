#! /bin/bash
local_time=`date "+\/%d\\\\\\\\\\\/%b\\\\\\\\\\\/%Y\/"`
str="startdate $local_time"
sed -i "s/^.*startdate.*$/$str/" test_config.yml
