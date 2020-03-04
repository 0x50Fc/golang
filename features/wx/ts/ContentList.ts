import { int32 } from "./lib/less";
import {Content} from "./Content";
import {Page} from "./Page";


export interface CountData {
    /**
     * 记录数量
     */
    total: int32
}


export interface ContentList {
    /**
     * 内容
     */
    items: Content[]

    /**
     * 分页
     */
    page?: Page
}