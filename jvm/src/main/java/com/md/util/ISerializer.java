package com.md.util;

import java.nio.ByteBuffer;

public interface ISerializer<T> {
    byte[] toBytes(T toSerialize);
    T fromBytes(byte[] buffer);
}