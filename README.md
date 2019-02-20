# 图片合并工具

[GITHUB上的MERGI](!https://github.com/noelyahan/mergi)提供了go的图片合并功能，但是它的
cl只能两张两张的合并，对于要合并一个目录下多张图片的需求来说，运行起来非常麻烦，随手就以其library
为基础实现了一个合并目录下所有图片为一长条图的工具。

参数:
* -i 作为输入的路径名，不递归该路径下的子目录（其实要递归也很简单，使用filepath.Walk就可以了
* -o 输出的文件名，默认为out.png