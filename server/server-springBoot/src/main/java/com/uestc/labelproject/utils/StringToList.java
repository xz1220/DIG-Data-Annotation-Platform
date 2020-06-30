package com.uestc.labelproject.utils;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

/**
 * @Auther: kiritoghy
 * @Date: 19-7-26 下午6:31
 */
public class StringToList {

    public static List<Long> longList(String str){
        return Arrays.stream(str.substring(1,str.length()-1).split(","))
                .map(s -> Long.parseLong(s.trim()))
                .collect(Collectors.toList());
    }
}
