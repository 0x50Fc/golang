"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const less = __importStar(require("../Less"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
let url = require("url");
function getType(fd) {
    let v = "String";
    switch (fd.type) {
        case less.FieldType.INT32:
            v = "int";
            break;
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            v = "float";
            break;
        case less.FieldType.INT64:
            v = "long";
            break;
        case less.FieldType.BOOLEAN:
            v = "boolean";
            break;
        case less.FieldType.OBJECT:
        case less.FieldType.ENUM:
            if (fd.typeSymbol === undefined) {
                v = "any";
            }
            else {
                v = `[${fd.typeSymbol}](Object.md#${fd.typeSymbol})`;
            }
            break;
    }
    if (fd.isArray) {
        v = v + '[]';
    }
    return v;
}
function mkdirs(dir) {
    if (fs.existsSync(dir)) {
        return;
    }
    mkdirs(path.dirname(dir));
    fs.mkdirSync(dir);
}
function walk(basePath, outfile) {
    var objectSet = {};
    var enumKeys = [];
    var objectKeys = [];
    var code = [];
    var lesss = [];
    less.walk(basePath, (v) => {
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
    lesss.sort((a, b) => {
        return a.name.localeCompare(b.name);
    });
    for (let v of lesss) {
        code.push(`### ${v.name}.json ${v.request.method ? v.request.method : 'POST'}\r\n\r\n`);
        if (v.request.title) {
            code.push(`*${v.request.title || ''}\r\n`);
        }
        code.push(`| 参数 | 类型 | 必填 | 说明 ｜`);
    }
}
exports.walk = walk;
