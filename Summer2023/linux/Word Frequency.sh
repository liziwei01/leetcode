###
 # @Author: liziwei01
 # @Date: 2023-07-14 03:27:25
 # @LastEditors: liziwei01
 # @LastEditTime: 2023-07-14 03:28:27
 # @Description: file content
### 
# Read from the file words.txt and output the word frequency list to stdout.
# tr translate -squeeze ' ' to '\n' 
# unique count
# sort numerically and reversely
# print 2nd 1st element
cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -nr | awk '{print $2, $1}'