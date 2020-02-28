import { int64 } from "./lib/less";
import { Relation } from "./Relation";

/**
 * 关系表基类
 * @type db
 */
export class Follow extends Relation {

    /**
    * 备注名
    * @length 255
    */
    title: string = ""

}
