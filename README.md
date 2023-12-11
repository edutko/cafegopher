# CAFEGOPHER

A Go library for interacting with serialized Java objects

## Introduction

This library grew out of a need to parse serialized Java in JCE keystores. (As far as I can tell,
the serialized Java objects in these keystores do not include a length prefix, so there is no way
skip over them without parsing them.) For now, this library is limited to prasing/unmarshalling.
Maybe someday I'll add serialization support as well, but that is not on the critical path for the
[more interesting work](https://github.com/edutko/what-is) I'm doing.

The name of the library is a reference to 0xCAFEBABE, the "magic bytes" of Java class files.

## References

The following documents and tools were immensely helpful when implementing Java deserialization:

* https://docs.oracle.com/javase/6/docs/platform/serialization/spec/protocol.html
* https://github.com/NickstaDB/SerializationDumper
