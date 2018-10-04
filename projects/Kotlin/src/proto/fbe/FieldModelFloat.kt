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

// Fast Binary Encoding Float field model class
class FieldModelFloat(buffer: Buffer, offset: Long) : FieldModel(buffer, offset)
{
    // Field size
    override val FBESize: Long = 4

    // Get the value
    fun get(defaults: Float = 0.0f): Float {
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return defaults

        return readFloat(FBEOffset)
    }

    // Set the value
    fun set(value: Float) {
        assert(_buffer.offset + FBEOffset + FBESize <= _buffer.size) { "Model is broken!" }
        if (_buffer.offset + FBEOffset + FBESize > _buffer.size)
            return

        write(FBEOffset, value)
    }
}
