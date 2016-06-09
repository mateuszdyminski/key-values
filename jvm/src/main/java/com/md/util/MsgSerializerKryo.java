package com.md.util;

import com.esotericsoftware.kryo.Kryo;
import com.esotericsoftware.kryo.Serializer;
import com.esotericsoftware.kryo.io.Input;
import com.esotericsoftware.kryo.io.Output;
import com.md.models.Message;

/**
 * User: mdyminski
 */
public class MsgSerializerKryo implements ISerializer<Message> {

    @Override
    public byte[] toBytes(Message toSerialize) {
        int size = 24 + toSerialize.getData().length;
        Output output = new Output(size, size);

        KRYO.get().writeObject(output, toSerialize);
        output.close();

        return output.toBytes();
    }

    @Override
    public Message fromBytes(byte[] bytes) {
        Input input = new Input(bytes);

        Message deserialized = KRYO.get().readObject(input, Message.class);
        input.close();

        return deserialized;
    }

    private static final ThreadLocal<Kryo> KRYO = new ThreadLocal<Kryo>() {
        @Override
        protected Kryo initialValue() {
            Kryo k = new Kryo();
            k.register(Message.class, new MsgSerializer());
            return k;
        }
    };


    private static class MsgSerializer extends Serializer<Message> {

        public void write (Kryo kryo, Output output, Message msg) {
            output.writeInt(msg.getTraceType());
            output.writeInt(msg.getDirection());
            output.writeInt(msg.getInterfaceName());
            output.writeInt(msg.getProtocol());
            output.writeInt(msg.getProtocolFormat());
            output.writeInt(msg.getProtocolMessageType());
            output.writeInt(msg.getData().length);
            output.writeBytes(msg.getData());
        }

        public Message read (Kryo kryo, Input input, Class<Message> type) {
            Message  msg = new Message();

            msg.setTraceType(input.readInt());
            msg.setDirection(input.readInt());
            msg.setInterfaceName(input.readInt());
            msg.setProtocol(input.readInt());
            msg.setProtocolFormat(input.readInt());
            msg.setProtocolMessageType(input.readInt());
            msg.setData(input.readBytes(input.readInt()));

            return msg;
        }
    }
}
