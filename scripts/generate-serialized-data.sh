#!/usr/bin/env bash

OUTDIR=$(dirname "$0")/../java/testdata
[ -d "$OUTDIR" ] || mkdir -p "$OUTDIR"

rm -rf com/edutko/*.class
javac com/edutko/Main.java

rm -f "$OUTDIR"/*.ser
java com/edutko/Main "$OUTDIR"
rm -f com/edutko/*.class

keytool -genseckey -keystore "$OUTDIR"/jce.jks -storetype jceks -storepass "hunter2" \
    -alias aes256 -keyalg aes -keysize 256 -keypass "hunter2" 2>/dev/null
tail -c +33 "$OUTDIR"/jce.jks > "$OUTDIR"/SealedObjectForKeyProtector.ser
rm -f "$OUTDIR"/jce.jks
