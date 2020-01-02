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
    code_Object.push("export interface Error {\n\n");
    code_Object.push("\t/**\n");
    code_Object.push("\t * 错误码\n");
    code_Object.push("\t */\n");
    code_Object.push("\t");
    code_Object.push("errno");
    code_Object.push(": ");
    code_Object.push("number");
    code_Object.push("\n\n");
    code_Object.push("\t/**\n");
    code_Object.push("\t * 信息\n");
    code_Object.push("\t */\n");
    code_Object.push("\t");
    code_Object.push("errmsg");
    code_Object.push(": ");
    code_Object.push("string");
    code_Object.push("\n\n");
    code_Object.push("}\n\n");
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
        lessTaskCode(outDir, v, objectSet);
        lessCode(outDir, v, objectSet);
    });
    fs.writeFileSync(path.join(outDir, "ObjectSet.ts"), code_Object.join(''));
}
exports.walk = walk;
function lessTaskCode(outDir, v, objectSet) {
    let vs = [];
    let p = v.name + ".task.ts";
    console.info(">>", p);
    {
        let dir = path.dirname(p);
        let rdir = "./";
        if (dir) {
            mkdirs(path.join(outDir, dir));
            rdir = "../".repeat(dir.split("/").length);
        }
        vs.push("import * as OS from ");
        vs.push(JSON.stringify(rdir + "ObjectSet"));
        vs.push("\n\n");
        vs.push("export type int64 = OS.int64\n");
        vs.push("export type int32 = OS.int32\n");
        vs.push("export type Error = OS.Error\n");
        vs.push("\n");
    }
    if (v.request.title) {
        vs.push("/**\n");
        vs.push(" * ");
        vs.push(v.request.title);
        vs.push("\n */\n");
    }
    vs.push("export interface Task {\n");
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
    vs.push("}\n\n");
    fs.writeFileSync(path.join(outDir, p), vs.join(''));
}
function lessCode(outDir, v, objectSet) {
    let vs = [];
    let basename = path.basename(v.name);
    let p = v.name + ".less.ts";
    console.info(">>", p);
    if (fs.existsSync(path.join(outDir, p))) {
        return;
    }
    {
        let dir = path.dirname(p);
        let rdir = "./";
        if (dir) {
            mkdirs(path.join(outDir, dir));
            rdir = "../".repeat(dir.split("/").length);
        }
        vs.push("import * as OS from ");
        vs.push(JSON.stringify(rdir + "ObjectSet"));
        vs.push("\n");
        vs.push("import { Task } from ");
        vs.push(JSON.stringify("./" + basename + ".task"));
        vs.push("\n");
        vs.push("\n");
        vs.push("export type int64 = OS.int64\n");
        vs.push("export type int32 = OS.int32\n");
        vs.push("export type Error = OS.Error\n");
        vs.push("\n");
    }
    {
        let data = "any";
        for (let fd of v.response.fields) {
            if (fd.name == "data" && fd.typeSymbol !== undefined) {
                data = "OS." + fd.typeSymbol;
                break;
            }
        }
        if (v.request.title) {
            vs.push("/**\n");
            vs.push(" * ");
            vs.push(v.request.title);
            vs.push("\n */\n");
        }
        vs.push("export function handle");
        vs.push("(task: Task): ");
        vs.push(data);
        vs.push(" {\n");
        vs.push("}\n\n");
    }
    fs.writeFileSync(path.join(outDir, p), vs.join(''));
}
