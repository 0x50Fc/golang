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
let url = require("url");
function getType(fd) {
    if (fd.isArray) {
        return "string";
    }
    switch (fd.type) {
        case less.FieldType.INT32:
        case less.FieldType.INT64:
            return "integer";
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            return "number";
        case less.FieldType.BOOLEAN:
            return "string";
        case less.FieldType.FILE:
            return "file";
    }
    return "string";
}
function getFormat(fd) {
    if (fd.isArray) {
        return "string";
    }
    switch (fd.type) {
        case less.FieldType.INT32:
            return "int32";
        case less.FieldType.INT64:
            return "int64";
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            return "number";
        case less.FieldType.BOOLEAN:
            return "Boolean";
        case less.FieldType.FILE:
            return "file";
    }
    return "string";
}
function walk(basePath, vpc, pattern = /.*/i, schemes = ['https'], alias = '/', info = { title: 'API', version: '1.0' }, timeout = 30000) {
    let paths = {};
    less.walk(basePath, (v) => {
        let name = "/" + v.name + ".json";
        if (!pattern.test(name)) {
            return;
        }
        let path = {};
        let parameters = [
            {
                name: 'Cookie',
                description: 'Cookie',
                in: 'header',
                pattern: '',
                required: false,
                type: 'string',
                format: 'string'
            }
        ];
        let inType = "query";
        let consumes = [];
        if (v.request.method == "POST") {
            inType = "formData";
            consumes.push("application/x-www-form-urlencoded");
        }
        let handling = "MAPPING";
        for (let fd of v.request.fields) {
            let type = getType(fd);
            parameters.push({
                name: fd.name,
                description: fd.title,
                in: inType,
                pattern: fd.pattern || '',
                required: fd.required,
                type: getType(fd),
                format: getFormat(fd)
            });
            if (type == 'file' || fd.name == '$body') {
                handling = "PASSTHROUGH";
            }
        }
        if (handling == 'PASSTHROUGH') {
            parameters = [];
        }
        let id = v.name.replace(/[^a-zA-Z0-9]/g, '_') + '_json';
        path[v.request.method.toLowerCase()] = {
            "x-aliyun-apigateway-paramater-handling": handling,
            "x-aliyun-apigateway-auth-type": "ANONYMOUS",
            "x-aliyun-apigateway-backend": {
                "type": "HTTP-VPC",
                "vpcAccessName": vpc,
                "path": name,
                "method": v.request.method.toLocaleLowerCase(),
                "timeout": timeout,
            },
            "x-aliyun-apigateway-system-parameters": [{
                    "systemName": "CaClientIp",
                    "backendName": "X-Real-IP",
                    "location": "header"
                }, {
                    "systemName": "CaClientUa",
                    "backendName": "X-User-Agent",
                    "location": "header"
                }, {
                    "systemName": "CaDomain",
                    "backendName": "X-Host",
                    "location": "header"
                }, {
                    "systemName": "CaHttpSchema",
                    "backendName": "X-Schema",
                    "location": "header"
                }, {
                    "systemName": "CaDomain",
                    "backendName": "X-Host",
                    "location": "header"
                }
            ],
            schemes: schemes,
            operationId: id,
            consumes: consumes,
            produces: ["application/json"],
            parameters: parameters,
            responses: {
                "200": {
                    description: "OK"
                }
            },
            summary: v.request.title
        };
        paths[name] = path;
    });
    return {
        basePath: alias,
        schemes: schemes,
        swagger: '2.0',
        paths: paths,
        info: info
    };
}
exports.walk = walk;
