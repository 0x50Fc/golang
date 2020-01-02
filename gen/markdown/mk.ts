import * as less from "../Less"
import * as fs from "fs"
import * as path from "path"

let url = require("url");

function getType(fd: less.LessField): string {
    let v = "String"
    switch (fd.type) {
        case less.FieldType.INT32:
            v = "int"
            break;
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            v = "float"
            break;
        case less.FieldType.INT64:
            v = "long"
            break;
        case less.FieldType.BOOLEAN:
            v = "boolean"
            break
        case less.FieldType.OBJECT:
        case less.FieldType.ENUM:
            if (fd.typeSymbol === undefined) {
                v = "any"
            } else {
                v = `[${fd.typeSymbol}](Object.md#${fd.typeSymbol})`
            }
            break;
    }

    if (fd.isArray) {
        v = v + '[]';
    }

    return v
}

interface ObjectSet {
    [name: string]: less.LessObject | less.LessEnum
}

function mkdirs(dir: string) {
    if (fs.existsSync(dir)) {
        return
    }
    mkdirs(path.dirname(dir));
    fs.mkdirSync(dir)
}

interface PackageSet {
    [name: string]: less.Less[]
}

export function walk(basePath: string, outfile: string): void {

    var objectSet: ObjectSet = {}
    var enumKeys: string[] = []
    var objectKeys: string[] = []
    var code: string[] = [];
    var lesss: less.Less[] = [];

    less.walk(basePath, (v: less.Less): void => {

        for (let object of v.enums) {

            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                enumKeys.push(object.name);
            }

        }

        for (let object of v.objects) {

            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                objectKeys.push(object.name);
            }

        }

        lesss.push(v);

    });

    enumKeys.sort();
    objectKeys.sort();
    lesss.sort((a: less.Less, b: less.Less): number => {
        return a.name.localeCompare(b.name);
    })

    for (let v of lesss) {

        code.push(`### ${v.name}.json ${v.request.method ? v.request.method : 'POST'}\r\n\r\n`)
        if (v.request.title) {
            code.push(`*${v.request.title || ''}\r\n`);
        }

        code.push(`| 参数 | 类型 | 必填 | 说明 ｜`)
    }


}
