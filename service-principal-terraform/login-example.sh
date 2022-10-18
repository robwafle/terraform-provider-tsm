# use your tenantId, your appId, and your pem to login to your subscription,
# Please save as login.sh (included in .gitignore) so you don't commit your sensitive data by accident.
az login \
--service-principal \
-u YOUR-APPID-HERE \
-p terraform-sp-cert.pem \
--tenant YOUR-TENANTID-HERE
