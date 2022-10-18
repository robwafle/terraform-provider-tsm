openssl req \
-newkey rsa:2048 \
-new \
-nodes \
-x509 \
-days 365 \
-subj "/C=US/ST=Illinois/L=Chicago/O=Robert Wafle/OU=Prime/CN=wafle.com" \
-keyout key.pem \
-out cert.pem && \
cat key.pem > iac-sp-cert.pem && \
cat cert.pem >> terraform-sp-cert.pem