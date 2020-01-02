"use strict";
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (Object.hasOwnProperty.call(mod, k)) result[k] = mod[k];
    result["default"] = mod;
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
const less = __importStar(require("./Less"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
let url = require("url");
function getType(fd, prefix) {
    let v = "string";
    switch (fd.type) {
        case less.FieldType.INT32:
            v = "int32";
            break;
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            v = "number";
            break;
        case less.FieldType.INT64:
            v = "int64";
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
                v = prefix + fd.typeSymbol;
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
function walk(basePath, outDir) {
    if (!fs.existsSync(outDir)) {
        fs.mkdirSync(outDir);
    }
    var objectSet = {};
    var code_Object = [];
    var packageSet = {};
    code_Object.push("\nexport type int64 = number | string\n");
    code_Object.push("export type int32 = number\n\n\n");
    less.walk(basePath, (v) => {
        for (let object of v.enums) {
            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                if (object.title) {
                    code_Object.push("/**\n");
                    code_Object.push(" * ");
                    code_Object.push(object.title);
                    code_Object.push("\n */\n");
                }
                code_Object.push("export enum ");
                code_Object.push(object.name);
                code_Object.push(" {\n");
                for (let fd of object.items) {
                    if (fd.title) {
                        code_Object.push("\t/**\n");
                        code_Object.push("\t * ");
                        code_Object.push(fd.title);
                        code_Object.push("\n\t */\n");
                    }
                    code_Object.push("\t");
                    code_Object.push(fd.name);
                    code_Object.push(" = ");
                    code_Object.push(JSON.stringify(fd.value));
                    code_Object.push(",\n");
                }
                code_Object.push("}\n\n");
            }
        }
        for (let object of v.objects) {
            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                if (object.title) {
                    code_Object.push("/**\n");
                    code_Object.push(" * ");
                    code_Object.push(object.title);
                    code_Object.push("\n */\n");
                }
                code_Object.push("export interface ");
                code_Object.push(object.name);
                code_Object.push(" {\n\n");
                for (let fd of object.fields) {
                    if (fd.title) {
                        code_Object.push("\t/**\n");
                        code_Object.push("\t * ");
                        code_Object.push(fd.title);
                        code_Object.push("\n\t */\n");
                    }
                    code_Object.push("\t");
                    code_Object.push(fd.name);
                    if (!fd.required) {
                        code_Object.push("?");
                    }
                    code_Object.push(": ");
                    code_Object.push(getType(fd, ""));
                    code_Object.push("\n\n");
                }
                code_Object.push("}\n\n");
            }
        }
        {
            let dir = path.dirname(v.name);
            let vs = packageSet[dir];
            if (vs === undefined) {
                packageSet[dir] = [v];
            }
            else {
                vs.push(v);
            }
        }
    });
    for (let p in packageSet) {
        let vs = packageSet[p];
        console.info(p);
        if (p == ".") {
            p = "index.ts";
        }
        else {
            p = p + ".ts";
        }
        let code = [];
        {
            let dir = path.dirname(p);
            let rdir = "./";
            let libdir = "../";
            if (dir != '.') {
                mkdirs(path.join(outDir, dir));
                rdir = "../".repeat(dir.split("/").length);
                libdir = "../".repeat(dir.split("/").length + 1);
            }
            code.push("import * as OS from ");
            code.push(JSON.stringify(rdir + "ObjectSet"));
            code.push("\n");
            code.push("import * as http from ");
            code.push(JSON.stringify(libdir + "http"));
            code.push("\n\n");
            code.push("export type int64 = number | string\n");
            code.push("export type int32 = number\n\n");
            code.push("export interface Error {\n\n");
            code.push("\t/**\n");
            code.push("\t * 错误码\n");
            code.push("\t */\n");
            code.push("\t");
            code.push("errno");
            code.push(": ");
            code.push("number");
            code.push("\n\n");
            code.push("\t/**\n");
            code.push("\t * 信息\n");
            code.push("\t */\n");
            code.push("\t");
            code.push("errmsg");
            code.push(": ");
            code.push("string");
            code.push("\n\n");
            code.push("}\n\n");
        }
        for (let v of vs) {
            lessCode(code, v, objectSet);
        }
        fs.writeFileSync(path.join(outDir, p), code.join(''));
    }
    fs.writeFileSync(path.join(outDir, "ObjectSet.ts"), code_Object.join(''));
}
exports.walk = walk;
function lessCode(vs, v, objectSet) {
    {
        let data = "any";
        for (let fd of v.response.fields) {
            if (fd.name == "data" && fd.typeSymbol !== undefined) {
                data = "OS." + fd.typeSymbol;
                break;
            }
        }
        vs.push("export function ");
        let n = path.basename(v.name);
        if (n == "in") {
            n = "In";
        }
        if (n == "for") {
            n = "For";
        }
        if (n == "if") {
            n = "If";
        }
        if (n == "is") {
            n = "Is";
        }
        vs.push(n);
        vs.push("(name: string,task: {\n\n");
        for (let fd of v.request.fields) {
            if (fd.title) {
                vs.push("\t/**\n");
                vs.push("\t * ");
                vs.push(fd.title);
                vs.push("\n\t */\n");
            }
            vs.push("\t");
            vs.push(fd.name);
            if (!fd.required) {
                vs.push("?");
            }
            vs.push(": ");
            vs.push(getType(fd, "OS."));
            vs.push("\n\n");
        }
        vs.push("}): ");
        vs.push(data);
        vs.push(" {\n");
        vs.push("\treturn ");
        if (v.request.method == 'GET') {
            vs.push("http.get(");
            vs.push("name, ");
            vs.push(JSON.stringify(v.name + ".json"));
            vs.push(", task");
            vs.push(")\n");
        }
        else {
            vs.push("http.post(");
            vs.push("name, ");
            vs.push(JSON.stringify(v.name + ".json"));
            vs.push(", task");
            vs.push(")\n");
        }
        vs.push("}\n\n");
    }
}
