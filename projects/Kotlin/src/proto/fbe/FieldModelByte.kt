// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding

package fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

// Fast Binary Encoding Byte field model class
class FieldModelByte(buffer: Buffer, offset: Long) : FieldModel(buffer, offset)
{
    // Field size
    override val FBESize: Long = 1

    // Get the value
    fun get(defaults: Byte = 0.toByte()): Byte {
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return defaults

        return readByte(FBEOffset)
    }

    // Set the value
    fun set(value: Byte) {
        assert(_buffer.offset + FBEOffset + FBESize <= _buffer.size) { "Model is broken!" }
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return

        write(FBEOffset, value)
    }
}
