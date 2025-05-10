# 安装 ssl 模块
echo "Installing ssl..."
sudo apt-get install openssl
sudo apt-get install libssl-dev

# # 创建根密钥，生成证书签名请求 (CSR)，创建根证书
openssl genrsa -out /etc/ssl/private/root.key 2048
openssl req -new -key /etc/ssl/private/root.key -out /etc/ssl/certs/root.csr
openssl x509 -req -in /etc/ssl/certs/root.csr -out /etc/ssl/certs/root.crt -signkey /etc/ssl/private/root.key -CAcreateserial -days 3650

# 生成服务器密钥，生成服务器证书签名请求 (CSR)，创建服务器证书扩展文件
openssl genrsa -out /etc/ssl/private/server.key 2048
openssl req -new -key /etc/ssl/private/server.key -out /etc/ssl/certs/server.csr
sudo nano v3.ext
# 内容如下
# authorityKeyIdentifier=keyid,issuer
# basicConstraints=CA:FALSE
# keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
# subjectAltName = @alt_names
# [alt_names]
# IP.1 = xxxxxxxxx # 你的云服务器地址

# # 使用根证书为服务器证书签名
openssl x509 -req -in /etc/ssl/certs/server.csr -CA /etc/ssl/certs/root.crt -CAkey /etc/ssl/private/root.key -CAcreateserial -out /etc/ssl/certs/server.crt -days 500 -sha256 -extfile v3.ext


# 打开Apache2配置文件
sudo nano /etc/apache2/sites-enabled/000-default.conf
# 添加如下内容
# <VirtualHost *:443>
#     ServerAdmin webmaster@localhost
#     DocumentRoot /var/www/html

#     SSLEngine on
#     SSLProxyEngine on

#     # 替换为您的自签名证书路径
#     SSLCertificateFile /etc/ssl/certs/server.crt
#     SSLCertificateKeyFile /etc/ssl/private/server.key

#     # 如果有中间证书，添加以下行
#     # SSLCertificateChainFile /path/to/your_intermediate.crt

#     ErrorLog ${APACHE_LOG_DIR}/error.log
#     CustomLog ${APACHE_LOG_DIR}/access.log combined

#     # 以下配置可选，用于启用 HTTP 到 HTTPS 重定向，也可以把这段添加到80端口那儿
#     <IfModule mod_rewrite.c>
#         RewriteEngine On
#         RewriteCond %{HTTPS} off
#         RewriteRule ^/?(.*) https://%{SERVER_NAME}%{REQUEST_URI} [R=301,L]
#     </IfModule>
# </VirtualHost>

# 启用ssl模块，启用ssl站点，重启服务
sudo a2enmod ssl
sudo a2ensite 000-default.conf
sudo systemctl restart apache2