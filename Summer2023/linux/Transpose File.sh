###
 # @Author: liziwei01
 # @Date: 2023-07-15 01:20:25
 # @LastEditors: liziwei01
 # @LastEditTime: 2023-07-15 01:33:45
 # @Description: file content
### 
#!/bin/bash

lines=`head -n1 file.txt | wc -w`

for i in `seq 1 $lines`
do
awk -va="$i" '{print $a}' file.txt | xargs
done