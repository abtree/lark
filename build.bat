
set fiels=api_http auth config giftpack getid api_ws

for %%i in (%fiels%) do (
	echo on
	@echo "build -->%%i"
	echo off
	cd apps/%%i/
	echo on
	go build -gcflags "-N -l"
	echo off
	cd ../../
)

pause