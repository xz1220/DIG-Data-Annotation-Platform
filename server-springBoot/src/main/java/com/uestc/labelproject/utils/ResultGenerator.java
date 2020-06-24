package com.uestc.labelproject.utils;


/**
 * @Auther:kiritoghy
 * @Date:19-7-23 下午6:00
 */
public class ResultGenerator {
    private static final int SUCCESS_CODE = 200;
    private static int FAIL_CODE = 500;
    private static final String DEFAULT_SUCCESS_MESSAGE = "SUCCESS";
    private static final String DEFAULT_FAIL_MESSAGE = "FAIL";

    /**
     * 生成默认成功信息
     * @return
     */
    public static Result genSuccessResult(){
        Result result = new Result();
        result.setCode(SUCCESS_CODE);
        result.setMessage(DEFAULT_SUCCESS_MESSAGE);
        return result;
    }

    /**
     * 生成指定成功信息
     * @param message
     * @return
     */
    public static Result genSuccessResult(String message){
        Result result = new Result();
        result.setCode(SUCCESS_CODE);
        result.setMessage(message);
        return result;
    }

    /**
     * 生成带有返回数据的成功信息
     * @param data
     * @return
     */
    public static Result genSuccessResult(Object data){
        Result result = new Result();
        result.setCode(SUCCESS_CODE);
        result.setMessage(DEFAULT_SUCCESS_MESSAGE);
        result.setData(data);
        return result;
    }

    /**
     * 生成默认失败信息
     * @return
     */
    public static Result genFailResult(){
        Result result = new Result();
        result.setCode(FAIL_CODE);
        result.setMessage(DEFAULT_FAIL_MESSAGE);
        return result;
    }

    /**
     * 生成指定失败信息
     * @param message
     * @return
     */
    public static Result genFailResult(String message){
        Result result = new Result();
        result.setCode(FAIL_CODE);
        result.setMessage(message);
        return result;
    }

    /**
     * 生成带有数据的失败信息
     * @param data
     * @return
     */
    public static Result genFailResult(Object data){
        Result result = new Result();
        result.setCode(FAIL_CODE);
        result.setMessage(DEFAULT_FAIL_MESSAGE);
        result.setData(data);
        return result;
    }
}
