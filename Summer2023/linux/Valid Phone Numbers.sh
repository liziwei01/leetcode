###
 # @Author: liziwei01
 # @Date: 2023-07-15 00:02:35
 # @LastEditors: liziwei01
 # @LastEditTime: 2023-07-15 01:02:20
 # @Description: file content
### 

# Read from the file file.txt and output all valid phone numbers to stdout.
grep '^([0-9]{3}-|\([0-9]{3}\) )[0-9]{3}-[0-9]{4}$' file.txt