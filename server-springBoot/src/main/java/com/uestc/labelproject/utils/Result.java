package com.uestc.labelproject.utils;

import java.io.Serializable;

/**
 * @Auther:kiritoghy
 * @Desc:用于返回对应消息
 * @Date:19-7-23 下午2:05
 */
public class Result implements Serializable {
    int code;
    String message;
    Object data;

    public Result() {
    }

    public Result(int code, String message) {
        this.code = code;
        this.message = message;
    }

    public int getCode() {
        return code;
    }

    public void setCode(int code) {
        this.code = code;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }

    public Object getData() {
        return data;
    }

    public void setData(Object data) {
        this.data = data;
    }

    @Override
    public String toString() {
        return "Result{" +
                "code=" + code +
                ", message='" + message + '\'' +
                ", data=" + data +
                '}';
    }
}
