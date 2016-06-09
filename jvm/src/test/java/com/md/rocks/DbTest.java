package com.md.rocks;

import com.google.common.base.Stopwatch;
import com.google.common.collect.Lists;
import com.md.models.Id;
import com.md.models.Message;
import com.md.util.MsgSerializerKryo;
import org.junit.AfterClass;
import org.junit.BeforeClass;
import org.junit.Test;
import org.rocksdb.Options;
import org.rocksdb.RocksDB;
import org.rocksdb.RocksDBException;
import org.rocksdb.WriteOptions;

import java.io.IOException;
import java.nio.file.*;
import java.nio.file.attribute.BasicFileAttributes;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

public class DbTest {

    private static RocksDB db;
    private static Options dbOptions;

    private MsgSerializerKryo ser = new MsgSerializerKryo();
    public static final String DB_PATH = "/tmp/rocks-tests/";

    @BeforeClass
    public static void setup() throws IOException {
        removeOldDb();
        openDb();
    }

    @Test
    public void insertTest() throws RocksDBException {
        // given
        byte[] msg = ser.toBytes(prepareMsg());
        int insertsNum = 10000000;

        System.out.println("Msg length: " + msg.length);
        System.out.println("Id length: " + Id.make(System.nanoTime(), 12345, 678, 980).idAsString().getBytes().length);

        // when
        Stopwatch sw = Stopwatch.createStarted();
        for (int i = 0; i < insertsNum; i++) {
            Id id = Id.make(System.nanoTime(), 12345, 678, 980);
            db.put(id.idAsString().getBytes(), msg);
        }

        // then
        System.out.println("Rocks db inserted " + insertsNum + " messages in: " + sw.elapsed(TimeUnit.MILLISECONDS) + " ms");
    }

    @Test
    public void insertParallelTest() throws RocksDBException, InterruptedException {
        // given
        byte[] msg = ser.toBytes(prepareMsg());
        int insertsNum = 100000;
        int threadNum = 10;

        // when
        CountDownLatch latch = new CountDownLatch(threadNum);
        Stopwatch sw = Stopwatch.createStarted();

        for (int i = 0; i < threadNum; i++) {
            //New way:
            new Thread(() -> {
                for (int j = 0; j < insertsNum; j++) {
                    Id id = Id.make(System.nanoTime(), 12345, 678, 980);
                    try {
                        db.put(id.idAsString().getBytes(), msg);
                    } catch (RocksDBException e) {

                    }
                }
                latch.countDown();
            }
            ).start();
        }

        latch.await();

        // then
        System.out.println("Rocks db inserted " + insertsNum + " messages in: " + sw.elapsed(TimeUnit.MILLISECONDS) + " ms with threads number: " + threadNum);
    }

    @Test
    public void insertWithoutWalTest() throws RocksDBException {
        // given
        byte[] msg = ser.toBytes(prepareMsg());
        int insertsNum = 1000000;
        WriteOptions opt = new WriteOptions();
        opt.disableWAL();

        // when
        Stopwatch sw = Stopwatch.createStarted();
        for (int i = 0; i < insertsNum; i++) {
            Id id = Id.make(System.nanoTime(), 12345, 678, 980);
            db.put(opt, id.idAsString().getBytes(), msg);
        }

        // then
        System.out.println("Rocks db inserted " + insertsNum + " messages in: " + sw.elapsed(TimeUnit.MILLISECONDS) + " ms without WAL!");
    }

    @Test
    public void getTest() throws RocksDBException {
        // given
        Message msg = prepareMsg();
        List<byte[]> keys = Lists.newArrayList();
        int getsNum = 1000000;
        for (int i = 0; i < getsNum; i++) {
            byte[] key = Id.make(System.nanoTime(), 12345, 678, 980).idAsString().getBytes();
            keys.add(key);
            db.put(key, ser.toBytes(msg));
        }
        System.out.println("Insertion done!");

        // when
        Stopwatch sw = Stopwatch.createStarted();
        for (byte[] key : keys) {
            byte[] value = db.get(key);
        }

        // then
        System.out.println("Rocks db gets " + getsNum+ " messages in: " + sw.elapsed(TimeUnit.MILLISECONDS) + " ms");
    }

    private static void openDb() {
        // a static method that loads the RocksDB C++ library.
        RocksDB.loadLibrary();
        // the Options class contains a set of configurable DB options
        // that determines the behavior of a database.
        dbOptions = new Options().setCreateIfMissing(true);
        try {
            // a factory method that returns a RocksDB instance
            db = RocksDB.open(dbOptions, DB_PATH);
            // do something
        } catch (RocksDBException e) {
            // do some error handling
            throw new RuntimeException("Can't open db", e);
        }
    }

    private Message prepareMsg() {
        Message msg = new Message();
        msg.setData("00010060000008000000020000000100010000024002000000420004000000000035000E00001B00090100616263640000000068000100006B0004200080000028002100AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA".getBytes());
        msg.setTraceType(1);
        msg.setDirection(2);
        msg.setInterfaceName(3);
        msg.setProtocol(4);
        msg.setProtocolFormat(5);
        msg.setProtocolMessageType(6);

        return msg;
    }

    public static void removeOldDb() throws IOException {
        Path directory = Paths.get(DB_PATH);
        if (Files.exists(directory)) {
            Files.walkFileTree(directory, new SimpleFileVisitor<Path>() {
                @Override
                public FileVisitResult visitFile(Path file, BasicFileAttributes attrs) throws IOException {
                    Files.delete(file);
                    return FileVisitResult.CONTINUE;
                }

                @Override
                public FileVisitResult postVisitDirectory(Path dir, IOException exc) throws IOException {
                    Files.delete(dir);
                    return FileVisitResult.CONTINUE;
                }

            });
        }
    }

    @AfterClass
    public static void closeDb() {
        if (db != null) db.close();
        dbOptions.dispose();
    }
}
