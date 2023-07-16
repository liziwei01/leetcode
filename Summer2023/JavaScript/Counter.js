/*
 * @Author: liziwei01
 * @Date: 2023-07-15 01:52:34
 * @LastEditors: liziwei01
 * @LastEditTime: 2023-07-15 01:53:38
 * @Description: file content
 */
/**
 * @param {number} n
 * @return {Function} counter
 */
var createCounter = function(n) {
    var a = n
    return function() {
        return a++
    };
};


const counter = createCounter(10)
console.log(counter()) // 10
console.log(counter()) // 11
