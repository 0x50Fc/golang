import { int32 } from './lib/less';


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
     * 总记录数
     */
    total: int32
}

