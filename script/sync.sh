#!/bin/sh
# 三个参数 版本管理工具类型（git or svn） 本地目录，文件夹名称
if ! $1
then
	git clone $1 $0
fi
cd $0
git fetch -p