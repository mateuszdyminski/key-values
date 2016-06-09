package com.md.models;

import java.io.Serializable;

/**
 * User: mdyminski
 */
public class Message implements Serializable {


    private static final long serialVersionUID = -7315952139250806790L;

    private byte[] data;
    private int traceType;
    private int direction;
    private int interfaceName;
    private int protocol;
    private int protocolFormat;
    private int protocolMessageType;

    public Message() {}

    public byte[] getData() {
        return data;
    }

    public void setData(byte[] data) {
        this.data = data.clone();
    }

    public int getTraceType() {
        return traceType;
    }

    public void setTraceType(int traceType) {
        this.traceType = traceType;
    }

    public int getDirection() {
        return direction;
    }

    public void setDirection(int direction) {
        this.direction = direction;
    }

    public int getInterfaceName() {
        return interfaceName;
    }

    public void setInterfaceName(int interfaceName) {
        this.interfaceName = interfaceName;
    }

    public int getProtocol() {
        return protocol;
    }

    public void setProtocol(int protocol) {
        this.protocol = protocol;
    }

    public int getProtocolFormat() {
        return protocolFormat;
    }

    public void setProtocolFormat(int protocolFormat) {
        this.protocolFormat = protocolFormat;
    }

    public int getProtocolMessageType() {
        return protocolMessageType;
    }

    public void setProtocolMessageType(int protocolMessageType) {
        this.protocolMessageType = protocolMessageType;
    }
}