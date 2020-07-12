# v3.0版本虹软人脸识别sdk，python version：3.7

packaged.py 为 封装 使用摄像头进行人脸注册 和 人脸比对识别

demo.py 为 简单的示例

test.py 为尝试调用packaged.py    经测试可以正常运行

asserts文件夹是注册图片

arcface是对lib的封装

main.py不用管

不懂可以看官方文档doc 注释 tips.md


# 更新日志

- 4.16更新：
添加健身房，游泳池，餐厅人脸验证注册模块、签退模块，代码基本一致；

详见pool/canteen/gym check/delete文件

如需使用，请安装pymysql

并对代码文件进行了整理删除，删除了main.py

assert文件夹仅供自己调试使用，目前本包已经没有程序直接使用该文件夹


同时packaged因实际需要改为cmd指令输入形式，指令规则如下：

python test.py文件路径 需要存储的照片路径

因此需要按照环境要求修改


- 4.22更新：

仅更新packaged包，请自行酌情修改原代码。

添加 当检测人脸数量不为1时，不关闭摄像头，直到检测到人脸。可以考虑定时关闭
