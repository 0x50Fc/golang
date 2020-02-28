import { int64, int32 } from "./lib/less";
import { Top } from './Top';

export interface Page {
    /**
     * 分页位置
     */
    p: int32
    /**
    * 单页记录数
    */
    n: int32
    /**
     * 总页数
     */
    count: int32
    /**
     * 总记录数
     */
    total: int32

    /**
    * 顶部ID
    */
    topId: int64
}


export interface QueryData {
    /**
     * Top
     */
    items: Top[]

    /**
     * 分页
     */
    page?: Page
}


export interface CountData {
    /**
     * 总记录数
     */
    total: int32
}

