import { int32, int64 } from "./lib/less";


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

export interface TopPage extends Page {

    /**
     * 顶部ID
     */
    topId: int64
}

export enum GroupBy {
    /**
     * 按 mid 分组
     */
    mid = 'mid',
    /**
     * 按 fuid 分组
     */
    fuid = 'fuid'
}
