// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.enums.fbe

// Fast Binary Encoding com.chronoxor.enums final sender
@Suppress("MemberVisibilityCanBePrivate", "PropertyName")
open class FinalSender : com.chronoxor.fbe.Sender
{
    // Sender models accessors
    val EnumsModel: EnumsFinalModel

    constructor() : super(true)
    {
        EnumsModel = EnumsFinalModel(buffer)
    }

    constructor(buffer: com.chronoxor.fbe.Buffer) : super(buffer, true)
    {
        EnumsModel = EnumsFinalModel(buffer)
    }

    @Suppress("JoinDeclarationAndAssignment")
    fun send(obj: Any): Long
    {
        when (obj)
        {
            is com.chronoxor.enums.Enums -> if (obj.fbeType == EnumsModel.fbeType) return send(obj)
        }

        return 0
    }

    fun send(value: com.chronoxor.enums.Enums): Long
    {
        // Serialize the value into the FBE stream
        val serialized = EnumsModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.enums.Enums serialization failed!" }
        assert(EnumsModel.verify()) { "com.chronoxor.enums.Enums validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }

    // Send message handler
    override fun onSend(buffer: ByteArray, offset: Long, size: Long): Long { throw UnsupportedOperationException("com.chronoxor.enums.fbe.Sender.onSend() not implemented!") }
}
