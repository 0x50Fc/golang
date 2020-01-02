import * as less from "../Less"

let url = require("url");


function getType(fd: less.LessField): string {
    if (fd.isArray) {
        return "string"
    }
    switch (fd.type) {
        case less.FieldType.INT32:
        case less.FieldType.INT64:
            return "integer"
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            return "number"
        case less.FieldType.BOOLEAN:
            return "string"
        case less.FieldType.FILE:
            return "file"
    }
    return "string"
}

function getFormat(fd: less.LessField): string {
    if (fd.isArray) {
        return "string"
    }
    switch (fd.type) {
        case less.FieldType.INT32:
            return "int32";
        case less.FieldType.INT64:
            return "int64"
        case less.FieldType.FLOAT32:
        case less.FieldType.FLOAT64:
            return "number"
        case less.FieldType.BOOLEAN:
            return "Boolean"
        case less.FieldType.FILE:
            return "file"
    }
    return "string"
}

export function walk(basePath: string, vpc: string, pattern = /.*/i, schemes: string[] = ['https'], alias: string = '/', info: any = { title: 'API', version: '1.0' }, timeout: number = 30000): any {

    let paths: any = {};

    less.walk(basePath, (v: less.Less): void => {

        let name = "/" + v.name + ".json";

        if (!pattern.test(name)) {
            return;
        }

        let path: any = {};
        let parameters: any[] = [
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
        let consumes: string[] = [];

        if (v.request.method == "POST") {
            inType = "formData"
            consumes.push("application/x-www-form-urlencoded");
        }

        let handling = "MAPPING"

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
                handling = "PASSTHROUGH"
            }
        }

        if (handling == 'PASSTHROUGH') {
            parameters = [];
        }


        let id = v.name.replace(/[^a-zA-Z0-9]/g, '_') + '_json'

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
