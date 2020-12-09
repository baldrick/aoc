package com.gavinmanning;

import java.io.IOException;
import java.net.URISyntaxException;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;
import java.util.Map;
import java.util.Queue;
import java.util.stream.Collectors;

public class Utils {
    public static List<String> getInput(final String filename) throws URISyntaxException, IOException {
        final URL resource = Day2.class.getClassLoader().getResource(filename);
        if (null == resource) {
            throw new IOException("Failed to open " + filename);
        } else {
            return Files.readAllLines(Paths.get(resource.toURI()));
        }
    }

    public static String repeat(final String str, final int count) {
        final StringBuilder sb = new StringBuilder();
        for (int i = 0;  i < count; i++) {
            sb.append(str);
        }
        return sb.toString();
    }

    public static <K, V> String dumpMap(final Map<K, V> map) {
        final StringBuilder sb = new StringBuilder();
        for (Map.Entry<K, V> entry : map.entrySet()) {
            sb.append(entry.getKey()).append(" = ").append(entry.getValue()).append(System.lineSeparator());
        }
        return sb.toString();
    }

    public static <T> String dumpList(final List<T> list) {
        return list.stream().map(T::toString).collect(Collectors.joining(","));
    }

    public static <T> String dumpQueue(final Queue<T> queue) {
        return queue.stream().map(T::toString).collect(Collectors.joining(","));
    }
}
