G:
cd G:\kafka\2.13_3.5.1\bin\windows
start zookeeper-server-start.bat ..\..\config\zookeeper.properties
timeout /t 3 /nobreak
start kafka-server-start.bat ..\..\config\server.properties
H:
