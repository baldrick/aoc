package com.gavinmanning;

import java.awt.*;

public class Day3 {
    public static void main(String[] args) {
        final int startPoint = 325489;
        final int[][] spiral = createSpiral(startPoint);
        //dumpSpiral(spiral);
        final Point one = find(spiral, 1);
        final Point start = find(spiral, startPoint);
        System.out.println("Path length = " + (Math.abs(one.getX() - start.getX()) + Math.abs(one.getY() - start.getY())));
        System.out.println("Additive spiral value > " + startPoint + " = " + findSpiralPoint(startPoint));
    }

    private static int[][] createSpiral(final int end) {
        final int edgeLength = (int) Math.ceil(Math.sqrt(end)) + 1;
        final int[][] spiral = new int[edgeLength][edgeLength];
        final int middle = edgeLength / 2;
        int x = middle;
        int y = middle;
        Direction direction = Direction.RIGHT;
        for (int i = 1;  i <= end;  i++) {
            //System.out.println("Setting " + x + "," + y + "=" + i);
            spiral[x][y] = i;
            x += direction.change.x;
            y += direction.change.y;
            if (0 == next(spiral, x, y, direction.next)) {
                direction = direction.next;
            }
        }
        return spiral;
    }

    private static int findSpiralPoint(final int end) {
        final int edgeLength = 12; //(int) Math.ceil(Math.sqrt(end)) + 1;
        final int[][] spiral = new int[edgeLength][edgeLength];
        final int middle = edgeLength / 2;
        int x = middle;
        int y = middle;
        spiral[x][y] = 1;
        Direction direction = Direction.RIGHT;
        x += direction.change.x;
        y += direction.change.y;
        direction = Direction.UP;
        for (int i = 1;  i <= end;  i++) {
            //System.out.println("Setting " + x + "," + y + "=" + i);
            final int cumulativeValue = addNeighbours(spiral, x, y);
            if (cumulativeValue > end) {
                dumpSpiral(spiral);
                return cumulativeValue;
            }
            spiral[x][y] = cumulativeValue;
            x += direction.change.x;
            y += direction.change.y;
            if (0 == next(spiral, x, y, direction.next)) {
                direction = direction.next;
            }
        }
        return 0;
    }

    private static int addNeighbours(final int[][] spiral, final int x, final int y) {
        return addDirection(spiral, x, y, Direction.UP.change)
            + addDirection(spiral, x, y, Direction.DOWN.change)
            + addDirection(spiral, x, y, Direction.LEFT.change)
            + addDirection(spiral, x, y, Direction.RIGHT.change)
            + addDirection(spiral, x, y, new Point(1, 1))
            + addDirection(spiral, x, y, new Point(1, -1))
            + addDirection(spiral, x, y, new Point(-1, 1))
            + addDirection(spiral, x, y, new Point(-1, -1));
    }

    private static int addDirection(final int[][] spiral, final int x, final int y, final Point direction) {
        return spiral[x + direction.x][y + direction.y];
    }

    private static int next(final int[][] spiral, final int x, final int y, final Direction direction) {
        return spiral[x + direction.change.x][y + direction.change.y];
    }

    private static void dumpSpiral(final int[][] spiral) {
        for (int j = 0;  j < spiral.length;  j++) {
            for (int i = 0;  i < spiral.length;  i++) {
                System.out.print(String.format("%8d ", spiral[i][j]));
            }
            System.out.println();
        }
    }

    private static Point find(final int[][] spiral, final int target) {
        for (int j = 0;  j < spiral.length;  j++) {
            for (int i = 0; i < spiral.length; i++) {
                if (spiral[i][j] == target) {
                    return new Point(i, j);
                }
            }
        }
        throw new RuntimeException("Spiral does not contain " + target);
    }

    enum Direction {
        RIGHT   (new Point(1,  0)),
        UP      (new Point(0, -1)),
        LEFT    (new Point(-1, 0)),
        DOWN    (new Point(0,  1));

        private final Point change;
        private Direction next;

        static {
            RIGHT.next = UP;
            UP.next = LEFT;
            LEFT.next = DOWN;
            DOWN.next = RIGHT;
        }

        Direction(final Point change) {
            this.change = change;
        }
    }
}
