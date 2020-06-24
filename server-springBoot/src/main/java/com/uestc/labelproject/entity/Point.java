package com.uestc.labelproject.entity;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-29 下午5:10
 */
public class Point {
    private double x;
    private double y;
    private int order;

    public Point() {
    }

    public Point(double x, double y) {
        this.x = x;
        this.y = y;
    }

    public double getX() {
        return x;
    }

    public void setX(double x) {
        this.x = x;
    }

    public double getY() {
        return y;
    }

    public void setY(double y) {
        this.y = y;
    }

    public int getOrder() {
        return order;
    }

    public void setOrder(int order) {
        this.order = order;
    }

    @Override
    public String toString() {
        return "Point{" +
                "x=" + x +
                ", y=" + y +
                ", order=" + order +
                '}';
    }
}
