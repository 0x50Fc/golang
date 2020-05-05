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
class Context {
    constructor(basePath) {
        this.basePath = basePath;
        this._enumSet = {};
        this._objectSet = {};
    }
    symbol(name) {
        let vs = name.split(/[_\/]/i);
        let n = [];
        for (let v of vs) {
            if (v.length > 0) {
                n.push(v.substr(0, 1).toLocaleUpperCase() + v.substr(1));
            }
        }
        return n.join('');
    }
    getType(fd, required) {
        let vs = [];
        if (fd.isArray) {
            vs.push("[]");
        }
        if (required === undefined) {
            required = false;
        }
        switch (fd.type) {
            case less.FieldType.INT32:
                if (!fd.required && !required) {
                    return "interface{}";
                }
                vs.push("int32");
                break;
            case less.FieldType.INT64:
                if (!fd.required && !required) {
                    return "interface{}";
                }
                vs.push("int64");
                break;
            case less.FieldType.FLOAT32:
                if (!fd.required && !required) {
                    return "interface{}";
                }
                vs.push("float32");
                break;
            case less.FieldType.FLOAT64:
                if (!fd.required && !required) {
                    return "interface{}";
                }
                vs.push("float64");
                break;
            case less.FieldType.BOOLEAN:
                if (!fd.required && !required) {
                    return "interface{}";
                }
                vs.push("bool");
                break;
            case less.FieldType.STRING:
                if (!fd.required && !required) {
                    return "interface{}";
                }
                vs.push("string");
                break;
            case less.FieldType.ENUM:
                if (!fd.required && !required) {
                    return "interface{}";
                }
                {
                    if (fd.typeSymbol !== undefined) {
                        let v = this._enumSet[fd.typeSymbol];
                        if (v !== undefined) {
                            if (v.type == less.LessEnumType.STRING) {
                                vs.push("string");
                            }
                            else {
                                vs.push("int32");
                            }
                            break;
                        }
                    }
                    vs.push("int32");
                }
                break;
            case less.FieldType.OBJECT:
                {
                    if (fd.typeSymbol !== undefined) {
                        let v = this._objectSet[fd.typeSymbol];
                        if (v !== undefined) {
                            vs.push("*" + this.symbol(v.name));
                            break;
                        }
                    }
                    vs.push("interface{}");
                }
                break;
            default:
                vs.push("interface{}");
                break;
        }
        return vs.join('');
    }
    getTypeInit(fd, required) {
        let vs = [];
        if (fd.isArray) {
            return "nil";
        }
        if (required === undefined) {
            required = false;
        }
        switch (fd.type) {
            case less.FieldType.INT32:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.INT64:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.FLOAT32:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.FLOAT64:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.BOOLEAN:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("false");
                break;
            case less.FieldType.STRING:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push('""');
                break;
            case less.FieldType.ENUM:
                if (!fd.required && !required) {
                    return "nil";
                }
                {
                    if (fd.typeSymbol !== undefined) {
                        let v = this._enumSet[fd.typeSymbol];
                        if (v !== undefined) {
                            vs.push(JSON.stringify(v.items[0].value));
                            break;
                        }
                    }
                    vs.push("0");
                }
                break;
            case less.FieldType.OBJECT:
                {
                    if (fd.typeSymbol !== undefined) {
                        let v = this._objectSet[fd.typeSymbol];
                        if (v !== undefined) {
                            vs.push("&" + this.symbol(v.name) + "{}");
                            break;
                        }
                    }
                    vs.push("nil");
                }
                break;
            default:
                vs.push("nil");
                break;
        }
        return vs.join('');
    }
    getTypeDefault(fd, required) {
        let vs = [];
        if (fd.isArray) {
            return "nil";
        }
        if (required === undefined) {
            required = false;
        }
        switch (fd.type) {
            case less.FieldType.INT32:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.INT64:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.FLOAT32:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.FLOAT64:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("0");
                break;
            case less.FieldType.BOOLEAN:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push("false");
                break;
            case less.FieldType.STRING:
                if (!fd.required && !required) {
                    return "nil";
                }
                vs.push('""');
                break;
            case less.FieldType.ENUM:
                if (!fd.required && !required) {
                    return "nil";
                }
                {
                    if (fd.typeSymbol !== undefined) {
                        let v = this._enumSet[fd.typeSymbol];
                        if (v !== undefined) {
                            vs.push(JSON.stringify(v.items[0].value));
                            break;
                        }
                    }
                    vs.push("0");
                }
                break;
            case less.FieldType.OBJECT:
                {
                    vs.push("nil");
                }
                break;
            default:
                vs.push("nil");
                break;
        }
        return vs.join('');
    }
    getDataType(v) {
        for (let fd of v.response.fields) {
            if (fd.name == "data") {
                return this.getType(fd, true);
            }
        }
        return "interface{}";
    }
    getDataTypeInit(v) {
        for (let fd of v.response.fields) {
            if (fd.name == "data") {
                return this.getTypeInit(fd, true);
            }
        }
        return "nil";
    }
    getDataTypeDefault(v) {
        for (let fd of v.response.fields) {
            if (fd.name == "data") {
                return this.getTypeDefault(fd, true);
            }
        }
        return "nil";
    }
    getFieldDecl(fd, jsonType = false) {
        let vs = [];
        vs.push(this.symbol(fd.name));
        vs.push("\t");
        vs.push(this.getType(fd));
        vs.push("\t");
        vs.push("`");
        if (fd.output === undefined || fd.output) {
            vs.push('json:');
            if (!fd.required || fd.type == less.FieldType.OBJECT || fd.isArray) {
                vs.push(JSON.stringify(fd.name + ",omitempty"));
            }
            else {
                vs.push(JSON.stringify(fd.name));
            }
        }
        else {
            vs.push('json:"-"');
        }
        vs.push(" name:");
        vs.push(JSON.stringify(fd.name.toLocaleLowerCase()));
        vs.push(' title:');
        vs.push(JSON.stringify(fd.title));
        if (fd.length !== undefined) {
            vs.push(' length:');
            vs.push(JSON.stringify(fd.length + ''));
        }
        if (fd.index !== undefined) {
            vs.push(' index:');
            vs.push(JSON.stringify(fd.index));
        }
        if (fd.unique !== undefined) {
            vs.push(' unique:');
            vs.push(JSON.stringify(fd.unique));
        }
        if (jsonType && fd.type == less.FieldType.OBJECT) {
            vs.push(' jsonType:"true"');
        }
        vs.push("`");
        return vs.join('');
    }
    checker(v) {
        for (let item of v.enums) {
            this._enumSet[item.name] = item;
        }
        for (let item of v.objects) {
            this._objectSet[item.name] = item;
        }
    }
    walk(fn) {
        return less.walk(this.basePath, (v) => {
            this.checker(v);
            fn(v);
        });
    }
}
exports.Context = Context;
