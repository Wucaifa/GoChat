# 传输明文加密算法
使用AES CFB加密算法
data: 要加密的明文数据；
key: AES 加密密钥，长度必须为 16/24/32 字节（分别对应 AES-128/192/256）；
iv: 初始化向量（Initialization Vector），长度必须是 aes.BlockSize（=16 字节）；
返回：Base64 编码后的密文字符串和错误信息。