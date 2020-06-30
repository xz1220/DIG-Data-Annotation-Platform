package com.uestc.labelproject.utils;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.uestc.labelproject.entity.Data;
import com.uestc.labelproject.entity.Point;
import com.uestc.labelproject.entity.RleData;
import com.uestc.labelproject.entity.TempRleData;

import java.util.ArrayList;
import java.util.List;

/**
 * @Auther: kiritoghy
 * @Date: 19-8-29 下午5:05
 */
public class DataGeneratorUtil {

    /**
     * 判断是否在多边形内
     * @param point
     * @param data
     * @return
     */
    public static boolean isInPolygon(Point point, Data data){
        List<Point> pts = data.getPoint();
        int N = pts.size();
        boolean boundOrVertex = true;
        int intersectCount = 0;
        double precision = 2e-10;
        Point p1, p2;
        Point p = point;

        p1 = pts.get(0);
        for(int i = 1; i <= N; ++i){
            if(p.equals(p1)){
                return boundOrVertex;
            }

            p2 = pts.get(i % N);
            if(p.getX() < Math.min(p1.getX(), p2.getX()) || p.getX() > Math.max(p1.getX(), p2.getX())){
                p1 = p2;
                continue;
            }

            if(p.getX() > Math.min(p1.getX(), p2.getX()) && p.getX() < Math.max(p1.getX(), p2.getX())){
                if(p.getY() <= Math.max(p1.getY(), p2.getY())){
                    if(p1.getX() == p2.getX() && p.getY() >= Math.min(p1.getY(), p2.getY())){
                        return boundOrVertex;
                    }

                    if(p1.getY() == p2.getY()){
                        if(p1.getY() == p.getY()){
                            return boundOrVertex;
                        }else{
                            ++intersectCount;
                        }
                    }else{
                        double xinters = (p.getX() - p1.getX()) * (p2.getY() - p1.getY()) / (p2.getX() - p1.getX()) + p1.getY();//cross point of y
                        if(Math.abs(p.getY() - xinters) < precision){
                            return boundOrVertex;
                        }

                        if(p.getY() < xinters){
                            ++intersectCount;
                        }
                    }
                }
            }else{
                if(p.getX() == p2.getX() && p.getY() <= p2.getY()){
                    Point p3 = pts.get((i+1) % N);
                    if(p.getX() >= Math.min(p1.getX(), p3.getX()) && p.getX() <= Math.max(p1.getX(), p3.getX())){
                        ++intersectCount;
                    }else{
                        intersectCount += 2;
                    }
                }
            }
            p1 = p2;
        }
        if(intersectCount % 2 == 0){
            return false;
        } else { //奇数在多边形内
            return true;
        }
    }

    /**
     * 获取RLE数据
     * @param width
     * @param height
     * @param data
     * @return
     */
    public static RleData genRleData(Integer width, Integer height, Data data){
        boolean temp = false;
        int count = 0;
        List<Integer> rle = new ArrayList<>();
        RleData rleData = new RleData();
        for(int i = 0; i < height; ++i){
            for(int j = 0; j < width; ++j){
                Point point = new Point(j,i);
                if(DataGeneratorUtil.isInPolygon(point,data) == temp){
                    count++;
                }
                else{
                    rle.add(count);
                    count = 1;
                    temp = !temp;
                }
            }
        }
        rle.add(count);
        List<Integer> size = new ArrayList<>();
        size.add(width);
        size.add(height);
        rleData.setCounts(rle);
        rleData.setSize(size);
        return rleData;
    }

    /**
     * 获取bbox数据
     * @param data
     * @return
     * 获取box数据，不是很确定坐标系的确定
     */
    public static List<Double> getBbox(Data data){
        List<Double> bbox = new ArrayList<>();
        Double maxWidth = Double.NEGATIVE_INFINITY;
        Double minWidth = Double.POSITIVE_INFINITY;
        Double maxHeight = Double.NEGATIVE_INFINITY;
        Double minHeight = Double.POSITIVE_INFINITY;
        for(Point point: data.getPoint()){
            maxHeight = Math.max(maxHeight,point.getY());
            minHeight = Math.min(minHeight,point.getY());
            maxWidth = Math.max(maxWidth,point.getX());
            minWidth = Math.min(minWidth,point.getX());
        }
        /**
         * box数据
         */
        bbox.add(minWidth);
        bbox.add(maxHeight);
        bbox.add(maxWidth-minWidth);
        bbox.add(maxHeight-minHeight);
        return bbox;
    }

    /**
     * 将RLE数据转为字符串方便数据库保存
     * @param rleData
     * @param dataId
     * @return
     */
    public static TempRleData rleDataToString(RleData rleData, Long dataId){
        TempRleData tempRleData = new TempRleData();
        String rle = JSON.toJSONString(rleData);
        tempRleData.setData(rle);
        tempRleData.setDataId(dataId);
        return tempRleData;
    }

    /**
     * 生成polygon数据
     * @param data
     * @return
     * 生成list,从data中挨个读取值，在分割任务（taskid=3）返回的json数据中，处于segmentation字段
     */
    public static List<Double> genPolygonData(Data data){
        List<Double> seg = new ArrayList<>();
        for (Point point: data.getPoint()){
            seg.add(point.getX());
            seg.add(point.getY());
        }
        return seg;
    }

    /**
     * 将RLEString转为实体
     * @param rleString
     * @return
     */
    public static RleData stringToRleData(String rleString){
        return JSONObject.parseObject(rleString, RleData.class);
    }

    /**
     * 计算多边形面积
     * @param points
     * @return
     * 计算多边形面积
     */
    public static double CalculateArea(List<Point> points){
        double area = 0;
        for(int i = 0; i < points.size() - 1; ++i){
            area += points.get(i).getX() * points.get(i+1).getY() - points.get(i).getY() * points.get(i+1).getX();//x[i]*y[i+1]-y[i]*x[i+1];
        }
        area = 0.5 * Math.abs(area + points.get(points.size() - 1).getX() * points.get(0).getY() - points.get(points.size() - 1).getY()*points.get(0).getX());//x[n]*y[1]-y[n]*x[1]
        return area;
    }
}
