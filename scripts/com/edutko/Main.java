package com.edutko;

import java.io.FileOutputStream;
import java.io.ObjectOutputStream;
import java.io.Serializable;
import java.util.ArrayList;

public class Main {
    protected enum Status {
        SNAFU,
        TARFU,
        FUBAR
    }

    protected enum Prefix {
        KILO(1000),
        KIBI(1024),
        MEGA(1000000),
        MEBI(1048576),
        GIGA(1000000000),
        GIBI(1073741824);

        private long multiple;

        Prefix(int multiple) {
            this.multiple = multiple;
        }
    }

    private static class Bar implements Serializable {
        private int value;

        protected Bar(int value) {
            this.value = value;
        }
    }

    private static class Foo implements Serializable {
        private byte b;
        private char c;
        private double d;
        private float f;
        private int i;
        private long l;
        private short s;
        private boolean bool;
        private String o;
        private byte[] a;

        private Status status;

        private Prefix prefix;

        private Bar[] bars;

        protected Foo(byte b, boolean bool, char c, double d, float f, int i, long l, short s, String o, byte[] a, Status status, Prefix prefix, Bar[] bars) {
            this.b = b;
            this.bool = bool;
            this.c = c;
            this.d = d;
            this.f = f;
            this.i = i;
            this.l = l;
            this.s = s;
            this.o = o;
            this.a = a;
            this.status = status;
            this.prefix = prefix;
            this.bars = bars;
        }
    }

    private static void serializeToFile(Object obj, String path, String name) {
        try {
            FileOutputStream o = new FileOutputStream(path + "/" + name + ".ser");
            ObjectOutputStream out = new ObjectOutputStream(o);
            out.writeObject(obj);
            out.flush();
        } catch (Exception ex) {
            System.err.println(ex.getMessage());
        }
    }

    public static void main(String[] args) {
        Foo foo1 = new Foo((byte) 0x7f, true, 'e', 3.14, 2.718f, 7, 5000000000L, (short) 32767, "hello", new byte[]{0x11, 0x22, 0x33}, Status.SNAFU, Prefix.KILO, new Bar[]{new Bar(0x1111), new Bar(0x2222)});
        Foo foo2 = new Foo((byte) 0x00, true, 'e', 3.14, 2.718f, 8, 6000000000L, (short) 32766, "hola", new byte[]{0x44, 0x55, 0x66}, Status.TARFU, Prefix.MEGA, new Bar[]{new Bar(0x3333), new Bar(0x4444)});
        Foo foo3 = new Foo((byte) 0x55, true, 'e', 3.14, 2.718f, 9, 7000000000L, (short) 32765, "aloha", new byte[]{0x77, 0x77, 0x77}, Status.FUBAR, Prefix.GIGA, new Bar[]{new Bar(0x5555)});
        serializeToFile(foo1, args[0], "object");
        serializeToFile(new Foo[]{foo1, foo2, foo3}, args[0], "objects");

        ArrayList<Integer> al = new ArrayList<>();
        al.add(0xaaaaaa);
        al.add(0xbbbbbb);
        al.add(0xcccccc);
        al.add(0xdddddd);
        serializeToFile(al, args[0], "ArrayList");

        serializeToFile(Boolean.valueOf(true), args[0], "Boolean");
        serializeToFile(Byte.valueOf((byte)255), args[0], "Byte");
        serializeToFile(Character.valueOf('a'), args[0], "Character");
        serializeToFile(Double.valueOf(2.718), args[0], "Double");
        serializeToFile(Float.valueOf(3.14f), args[0], "Float");
        serializeToFile(Integer.valueOf(1234567890), args[0], "Integer");
        serializeToFile(Long.valueOf(9876543210L), args[0], "Long");
        serializeToFile(Short.valueOf((short)32767), args[0], "Short");

        serializeToFile("0123456789abcdef".repeat(4096), args[0], "long-string");

        serializeToFile(456789, args[0], "int");
        serializeToFile("hi", args[0], "string");
        serializeToFile(Status.FUBAR, args[0], "enum");
        serializeToFile(new byte[]{0x11, 0x22, 0x33, 0x44}, args[0], "bytes");
        serializeToFile(new String[]{"abc", "def", "ghi"}, args[0], "strings");
    }
}