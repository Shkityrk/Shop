@echo off
setlocal

cd admin
docker build -t admin:latest .
cd ../auth
docker build -t auth:latest .
cd ../cart
docker build -t cart:latest .
cd ../product
docker build -t product:latest .

cd ..
docker tag admin:latest shkityrk/admin:latest
docker tag auth:latest shkityrk/auth:latest
docker tag cart:latest shkityrk/cart:latest
docker tag product:latest shkityrk/product:latest

docker push shkityrk/admin:latest
docker push shkityrk/auth:latest
docker push shkityrk/cart:latest
docker push shkityrk/product:latest




echo done <3
pause


docker build -t shkityrk/shop_admin ./admin
docker build -t shkityrk/shop_product ./product
docker build -t shkityrk/shop_auth ./auth
docker build -t shkityrk/shop_cart ./cart