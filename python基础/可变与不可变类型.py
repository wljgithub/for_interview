﻿#_*_ coding:utf-8 _*_
'''
1.python中的可变数据类型与不可变类型有哪些?
可变:列表、字典、
不可变：数字、字符串、元组(tuple)

2.可变类型遇不可变类型有什么区别？
对于不可变类型类说，当将变量的值赋予另一个变量时，修改其中一个，另一个并 不会 发生改变

'''
a=1
b=a
print('a的值为{} \t b的值为{}'.format(a,b))
print('修改了变量a后')
a=2
print('a的值为{} \t b的值为{}'.format(a,b))
print('\n')

'''
因为当执行a=1，b=a代码的时候，解释器其实在内存中干了两件事，
1.在内存中创建一个值为1的对象
2.分别将对象指向别名a，和别名b
所以当执行a=2,代码后，（因为对象 “1” 是整数型，为不可变类型，要改变对象的值只能重新创建一个新的对象） 解释器重新创建了一个值为2的对象，然后将变量a指向值为2的对象
因此修改了a的值，b并不会发生改变
'''

#但是对于可变类型来说，当将变量的值赋予给另一个变量时，修改其中一个，另一个将 会 发生改变
c=[1,2,3]
d=c
print('c的值为{} \t d的值为{}'.format(c,d))
print('修改了变量c后')
c.append(4)
print('c的值为{} \t d的值为{}'.format(c,d))

'''
列表是可变类型，修改其中一个另一个 将会 发生改变，
因为变量c、变量d在内存中都指向同一个对象，列表是可变类型，值发生改变时不需要重新创建一个新的对象
'''




