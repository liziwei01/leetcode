'''
Author: liziwei01
Date: 2023-08-27 00:04:50
LastEditors: liziwei01
LastEditTime: 2023-08-27 00:04:54
Description: file content
'''
total_weight = 0

for a in range(1, 10):  # 百位上的数字从 1 到 9
    for b in range(0, 10):  # 十位上的数字从 0 到 9
        for c in range(0, 10):  # 个位上的数字从 0 到 9
            odd_sum = a + c
            even_sum = b
            weight = odd_sum * even_sum
            total_weight += weight

print(total_weight)
