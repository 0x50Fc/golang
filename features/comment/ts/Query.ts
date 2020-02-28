import { int32 } from "./lib/less";

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
}

export interface CountData {
    /**
     * 记录数量
     */
    total: int32
}


export interface QueryData {
    /**
     * 评论
     */
    items: Comment[]

    /**
     * 分页
     */
    page?: Page
}

