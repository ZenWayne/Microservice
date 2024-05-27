@setlocal enabledelayedexpansion

@echo off
for /f "tokens=*" %%a in ('type ".env"') do (
  set "%%a"=!%%a:~0,-1!
)

forge script script/NFT.s.sol --rpc-url %RPC_URL% --broadcast --verify -vvvv

endlocal