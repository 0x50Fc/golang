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
const Document_1 = require("./Document");
function getType(fd) {
    switch (fd.type) {
        case less.FieldType.BOOLEAN:
            return ':boolean';
        case less.FieldType.ENUM:
        case less.FieldType.OBJECT:
            if (fd.typeSymbol !== undefined) {
                return ':' + fd.typeSymbol;
            }
            else {
                return ':any';
            }
        case less.FieldType.INT32:
            return ':int32';
        case less.FieldType.INT64:
            return ':int64';
        case less.FieldType.FLOAT32:
            return ':float32';
        case less.FieldType.FLOAT64:
            return ':float64';
    }
    return ':string';
}
function walk(basePath, outFile, name, title) {
    let objectSet = {};
    let doc = new Document_1.Document();
    let vs = [];
    doc.root.setValue('title', name);
    if (title) {
        doc.root.add().setValue('title', title);
    }
    less.walk(basePath, (v) => {
        vs.push(v);
        for (let object of v.enums) {
            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                let e = doc.root
                    .add()
                    .setValue('title', object.name)
                    .setValue('@automatic', true)
                    .setValue('@enum', (object.type == less.LessEnumType.INT32 ? 'int32' : 'string'));
                if (object.title) {
                    e.add().setValue('title', object.title);
                }
                for (let item of object.items) {
                    let v = e.add()
                        .setValue('title', item.name + ':' + item.value);
                    if (item.title) {
                        v.add().setValue('title', item.title);
                    }
                }
            }
        }
        for (let object of v.objects) {
            if (objectSet[object.name] === undefined) {
                objectSet[object.name] = object;
                let e = doc.root
                    .add()
                    .setValue('title', object.name)
                    .setValue('@automatic', true);
                let table = false;
                for (let tag of object.tags) {
                    if (tag.name == 'type' && tag.text == 'db') {
                        e.setValue('@table', object.name.toLocaleLowerCase());
                        table = true;
                        break;
                    }
                }
                for (let fd of object.fields) {
                    let v = e.add();
                    if (fd.required) {
                        v.setValue('title', fd.name);
                    }
                    else {
                        v.setValue('title', fd.name + '?');
                    }
                    if (fd.title) {
                        v.add().setValue('title', fd.title);
                    }
                    v.add().setValue('title', getType(fd));
                    if (table) {
                        if (fd.length !== undefined) {
                            v.add().setValue('title', ':length:' + fd.length);
                        }
                        if (fd.index !== undefined) {
                            v.add().setValue('title', ':index:' + fd.index);
                        }
                        if (fd.unique !== undefined) {
                            v.add().setValue('title', ':unique:' + fd.unique);
                        }
                    }
                }
            }
        }
    });
    vs = vs.sort((a, b) => {
        let v = a.name.split('/').length - b.name.split('/').length;
        if (v == 0) {
            return a.name.localeCompare(b.name);
        }
        return v;
    });
    for (let v of vs) {
        let e = doc.root.add().setValue('title', v.name).setValue('@automatic', true);
        if (v.request.title) {
            e.add().setValue('title', v.request.title);
        }
        e.add().setValue('title', ':' + v.request.method);
        for (let fd of v.request.fields) {
            let v = e.add();
            if (fd.required) {
                v.setValue('title', fd.name);
            }
            else {
                v.setValue('title', fd.name + '?');
            }
            if (fd.title) {
                v.add().setValue('title', fd.title);
            }
            v.add().setValue('title', getType(fd));
        }
        for (let fd of v.response.fields) {
            if (fd.name == 'data') {
                let v = e.add().setValue('title', ':data');
                if (fd.title) {
                    v.add().setValue('title', fd.title);
                }
                v.add().setValue('title', getType(fd));
                break;
            }
        }
    }
    fs.writeFileSync(outFile, JSON.stringify(doc.save(), undefined, 4));
}
exports.walk = walk;
