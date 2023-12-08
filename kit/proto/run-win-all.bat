@echo off
setlocal enabledelayedexpansion
REM 设置protoc.exe的路径
set PROTOC_PATH=C:\Users\AUSA\Documents\IdeaProject\tgf-example\kit\proto\protoc.exe
REM 设置.proto文件所在的目录
set PROTO_FILES_DIR=C:\Users\AUSA\Documents\IdeaProject\tgf-example
REM 设置生成文件的目标目录
set OUTPUT_CS_DIR=E:\unity\project\t2\Assets\Proto
set OUTPUT_GO_DIR=C:\Users\AUSA\Documents\IdeaProject\tgf-example\common
set PROTOC_GEN_GO=C:\Users\AUSA\Documents\IdeaProject\tgf-example\kit\proto\protoc-gen-go.exe
rem 查找目标文件夹及其子文件夹中所有以 .proto 结尾的文件
for /R %PROTO_FILES_DIR% %%f in (*.proto) do (
rem 将文件路径传递给 protoc
%PROTOC_PATH% --proto_path=C:\ --csharp_out=%OUTPUT_CS_DIR% --go_out=%OUTPUT_GO_DIR% --plugin=protoc-gen-go=%PROTOC_GEN_GO% %%f
echo %%f
)
endlocal