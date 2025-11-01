@echo off
echo Generating Swagger documentation...

cd /d "%~dp0\.."

swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal

echo Swagger documentation generated successfully!
echo Available at: http://localhost:8004/swagger/index.html
pause

