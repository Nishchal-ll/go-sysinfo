#!/usr/bin/env bash
set -e

PACKAGE_NAME="go-sysinfo"
VERSION="0.1.0"
ARCH="amd64"
BUILD_DIR="dist/deb/${PACKAGE_NAME}_${VERSION}_${ARCH}"

echo "🔨 [1/4] Cleaning old build directories..."
rm -rf dist/deb
mkdir -p "${BUILD_DIR}/usr/local/bin"
mkdir -p "${BUILD_DIR}/DEBIAN"

echo "⚙️  [2/4] Compiling Go binary for Linux ${ARCH}..."
CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -ldflags="-s -w" -o "${BUILD_DIR}/usr/local/bin/sysinfo" main.go
chmod 755 "${BUILD_DIR}/usr/local/bin/sysinfo"

echo "📝 [3/4] Creating DEBIAN/control metadata..."
cat <<EOF > "${BUILD_DIR}/DEBIAN/control"
Package: ${PACKAGE_NAME}
Version: ${VERSION}
Section: utils
Priority: optional
Architecture: ${ARCH}
Maintainer: Nishchal <github.com/Nishchal-ll>
Description: Fast, lightweight Linux system information CLI written in Go from scratch.
 Direct inspection of /proc and kernel syscalls with zero dependencies.
EOF

echo "📦 [4/4] Building .deb package with dpkg-deb..."
dpkg-deb --root-owner-group --build "${BUILD_DIR}" "dist/${PACKAGE_NAME}_${VERSION}_${ARCH}.deb"

echo ""
echo "🎉 SUCCESS! Your Debian/Ubuntu package is ready at:"
echo "   👉 dist/${PACKAGE_NAME}_${VERSION}_${ARCH}.deb"
