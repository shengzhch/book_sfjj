package common

/*
开地址哈希表 ：在开地址哈希表中解决冲突的方法是探查这个表，一直找到可以放置元素的槽位.
假设进行均匀散列，我们能够在一个开地址哈希表中探查的槽位个数是 1/(1-a)

h(k,i)= x ; 是到目前为止探查的次数

线性探查

h(k,i) = (h'(k)+i) mod m  -----> h'(x) 为辅助哈希函数


双散列

h(k,i) = (h'(k)+ih"(k)) mod m
*/