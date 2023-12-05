
set fiels=config api_http auth giftpack getid api_ws

for %%i in (%fiels%) do (
	echo on
	@echo "start -->%%i"
	echo off
	cd apps/%%i/
	echo on
	start %%i.exe
	echo off
	cd ../../
)

pause