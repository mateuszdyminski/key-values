package com.md.models;

import java.io.Serializable;
import java.util.Arrays;

/**
 * User: mdyminski
 */
public final class Id implements Serializable {

    private static final long serialVersionUID = 3030161700978779654L;

    public static final String SEPARATOR = "$";

    private int trs;

    private int t;

    private long e;

    private long tm;

    public Id() {
    }

    public static Id make(long timeInMicroseconds, long eNodeBId, int traceRecordingSessionId, int traceId) {
        Id id = new Id();
        id.setTm(timeInMicroseconds);
        id.setE(eNodeBId);
        id.setTrs(traceRecordingSessionId);
        id.setT(traceId);
        return id;
    }

    public int getTrs() {
        return trs;
    }

    public void setTrs(int trs) {
        this.trs = trs;
    }

    public int getT() {
        return t;
    }

    public void setT(int t) {
        this.t = t;
    }

    public long getE() {
        return e;
    }

    public void setE(long e) {
        this.e = e;
    }

    public long getTm() {
        return tm;
    }

    public void setTm(long tm) {
        this.tm = tm;
    }

    @Override
    public boolean equals(Object o) {
        if (o == null) {
            return false;
        }
        if (o == this) {
            return true;
        }
        if (!(o instanceof Id)) {
            return false;
        }
        Id other = (Id) o;
        return t == other.getT() && trs == other.getTrs()
                && e == other.getE() && tm == other.getTm();
    }

    @Override
    public int hashCode() {
        return Arrays.hashCode(new Object[]{trs, t, e, tm});
    }

    @Override
    public String toString() {
        return idAsString();
    }

    public String idAsString() {
        return tm + SEPARATOR + e + SEPARATOR + trs + SEPARATOR + t;
    }

    public String idAsString2() {
        return new StringBuilder().append(tm).append(SEPARATOR).append(e).append(SEPARATOR).append(trs).append(SEPARATOR).append(t).toString();
    }
}
