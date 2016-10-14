#!/bin/bash
# ---
echo "Cross compiling for Windows, Mac and Linux\n";
echo "-----------------------------------------";
#----

#---- Compile Windows
echo "[BEGIN] Compiling for Windows to bin/binary-win32\n";
env GOOS=windows go build -o bin/binary-win32; # Windows
echo "[DONE] Compiling for Windows to bin/binary-win32\n";
echo "-----------------------------------------";

# ---- Compile Mac
echo "[BEGIN] Compiling for Mac to bin/binary-darwin\n"
env GOOS=darwin go build -o bin/binary-darwin # darwin / Mac
echo "[DONE] Compiling for Mac to bin/binary-darwin\n"
echo "-----------------------------------------";

# ---- Start Linux
echo "[BEGIN] Compiling for Linux to bin/binary-linux\n"
env GOOS=linux go build -o bin/binary-linux # Linux
echo "[DONE] Compiling for Linux to bin/binary-linux\n"