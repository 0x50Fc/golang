
import * as less from "./Less"


export interface TableField {
    title: string
    name: string
    type: string
    length: number
    index: string
    unique: string
    default: string
}

export interface TableFieldSet {
    [name: string]: TableField
}

export interface Table {
    fields?: TableFieldSet
    autoIncrement?: number
    count?: number
    name?: string[]
}

export interface TableSet {
    [name: string]: Table
}

function escapeName(n: string): string {
    return "`" + n + "`"
}

function tableField(fd: less.LessField, getEnum: (name: string) => less.LessEnum | undefined, jsonType?: boolean): TableField {
    let n = fd.name.toLocaleLowerCase();
    let type = "VARCHAR";
    let defaultValue = "''";
    let length = fd.length === undefined ? 0 : fd.length;
    switch (fd.type) {
        case less.FieldType.BOOLEAN:
            type = "TINYINT";
            defaultValue = "0";
            break;
        case less.FieldType.INT32:
            type = "INT";
            defaultValue = "0";
            break;
        case less.FieldType.INT64:
            type = "BIGINT";
            defaultValue = "0";
            break;
        case less.FieldType.FLOAT32:
            type = "FLOAT";
            defaultValue = "0";
            break;
        case less.FieldType.FLOAT64:
            type = "DOUBLE";
            defaultValue = "0";
            break;
        case less.FieldType.OBJECT:
            if (jsonType) {
                type = 'JSON'
                length = 0
                defaultValue = 'NULL'
                break;
            }
            if (length == -1) {
                length = 0;
                type = "TEXT";
            } else if (length == -2) {
                length = 0;
                type = "BIGTEXT";
            } else {
                type = "TEXT"
            }
            defaultValue = '';
            break;
        case less.FieldType.STRING:
            if (length == -1) {
                length = 0;
                type = "TEXT";
                defaultValue = '';
            } else if (length == -2) {
                length = 0;
                type = "BIGTEXT";
                defaultValue = '';
            } else if (length == 0) {
                length = 64
            }
            break;
        case less.FieldType.ENUM:
            if (fd.typeSymbol === undefined) {
                type = "TINYINT";
                defaultValue = "0";
            } else {
                let e = getEnum(fd.typeSymbol);
                if (e == undefined) {
                    type = "TINYINT";
                    defaultValue = "0";
                } else {
                    let mi: number = 0;
                    let ml: number = 0;
                    for (let i of e.items) {
                        if (typeof i.value == 'string') {
                            if (i.value.length > ml) {
                                ml = i.value.length
                            }
                        } else {
                            if (Math.abs(i.value) > mi) {
                                mi = Math.abs(i.value);
                            }
                        }
                    }
                    if (ml == 0) {
                        if (mi < 0x0ff) {
                            type = "TINYINT";
                        } else if (mi < 0x0ffff) {
                            type = "SMALLINT";
                        } else if (mi < 0x0ffffffff) {
                            type = "INT";
                        } else {
                            type = "BIGINT";
                        }
                        if (e.items.length > 0) {
                            defaultValue = '' + e.items[0].value
                        } else {
                            defaultValue = "0";
                        }
                    } else {
                        type = "VARCHAR";
                        if (length < ml) {
                            length = ml;
                        }
                        if (e.items.length > 0) {
                            defaultValue = "'" + e.items[0].value + "'"
                        } else {
                            defaultValue = "''";
                        }
                    }
                }
            }
            break;
    }
    return {
        title: fd.title,
        name: n,
        type: type,
        default: defaultValue,
        length: length,
        index: fd.index === undefined ? '' : fd.index,
        unique: fd.unique === undefined ? '' : fd.unique,
    }
}

function tableFieldSQL(field: TableField): string {
    let vs: string[] = [];
    vs.push(escapeName(field.name));
    vs.push(' ');
    vs.push(field.type);
    if (field.length > 0) {
        vs.push('(');
        vs.push(field.length + '');
        vs.push(')');
    }
    if (field.default != '') {
        vs.push(' DEFAULT ');
        vs.push(field.default);
    }
    return vs.join('')
}

interface EnumSet {
    [name: string]: less.LessEnum
}

function tableSQL(object: less.LessObject, getEnum: (name: string) => less.LessEnum | undefined, name: string, sql: (sql: string) => void, table?: Table, jsonType?: boolean): void {

    if (table === undefined || table.fields === undefined) {

        let autoIncrement: number | undefined;
        let vs: string[] = [];

        if (table && table.autoIncrement !== undefined) {
            autoIncrement = table.autoIncrement;
        }

        vs.push("CREATE TABLE IF NOT EXISTS ");
        vs.push(escapeName(name))
        vs.push(" (")

        vs.push("\r\n\tid BIGINT NOT NULL")

        if (autoIncrement !== undefined) {
            vs.push(" AUTO_INCREMENT")
        }

        vs.push("\t#ID\r\n")

        let indexs: TableField[] = []
        let uniques: TableField[] = []

        for (let fd of object.fields) {

            if (fd.name == "id") {
                continue
            }

            let field = tableField(fd, getEnum, jsonType)

            if (field.index != "") {
                indexs.push(field)
            }

            if (field.unique != "") {
                uniques.push(field)
            }

            vs.push("\t,")
            vs.push(tableFieldSQL(field));
            vs.push("\t#[字段] ")
            vs.push(field.title)
            vs.push("\r\n");

        }

        vs.push("\t, PRIMARY KEY(id) \r\n")

        for (let field of indexs) {
            vs.push("\t,INDEX ");
            vs.push(escapeName(field.name))
            vs.push(" (");
            vs.push(escapeName(field.name));
            vs.push(" ");
            vs.push(field.index);
            vs.push(")");

            vs.push("\t#[索引] ");
            vs.push(field.title);
            vs.push("\r\n");
        }

        for (let field of uniques) {
            vs.push("\t,UNIQUE INDEX ");
            vs.push(escapeName(field.name + '_u'))
            vs.push(" (");
            vs.push(escapeName(field.name));
            vs.push(" ");
            vs.push(field.index);
            vs.push(")");

            vs.push("\t#[唯一索引] ");
            vs.push(field.title);
            vs.push("\r\n");
        }

        if (autoIncrement === undefined) {
            vs.push(" ) ;\r\n")
        } else {
            vs.push(" ) AUTO_INCREMENT = ")
            vs.push(autoIncrement + '')
            vs.push(";\r\n")
        }

        sql(vs.join(''));

    } else {

        let vs: string[] = [];

        for (let fd of object.fields) {

            if (fd.name == "id") {
                continue
            }

            let field = tableField(fd, getEnum, jsonType)

            let v_fd = table.fields[field.name];

            if (v_fd === undefined) {
                vs.push("ALTER TABLE ");
                vs.push(escapeName(name))
                vs.push(" ADD COLUMN ");
                vs.push(tableFieldSQL(field));
                vs.push(";")
                vs.push("\t#[增加字段] ");
                vs.push(field.title);
                vs.push("\r\n");
            } else if (v_fd.type != field.type || v_fd.length != field.length) {
                vs.push("ALTER TABLE ");
                vs.push(escapeName(name));
                vs.push(" CHANGE ");
                vs.push(escapeName(field.name));
                vs.push(" ");
                vs.push(tableFieldSQL(field));

                vs.push(";")
                vs.push("\t#[修改字段] ");
                vs.push(field.title);
                vs.push("\r\n");
            }

            if (field.index != '' && (!v_fd || v_fd.index != field.index)) {
                vs.push("CREATE INDEX ");
                vs.push(escapeName(field.name));
                vs.push(" ON ");
                vs.push(escapeName(name));
                vs.push(" (");
                vs.push(escapeName(field.name));
                vs.push(" ");
                vs.push(field.index);
                vs.push(");")
                vs.push("\t#[创建索引] ");
                vs.push(field.title);
                vs.push("\r\n");
            }

            if (field.unique != '' && (!v_fd || v_fd.unique != field.unique)) {
                vs.push("CREATE UNIQUE INDEX ");
                vs.push(escapeName(field.name + '_'));
                vs.push(" ON ");
                vs.push(escapeName(name));
                vs.push(" (");
                vs.push(escapeName(field.name));
                vs.push(" ");
                vs.push(field.index);
                vs.push(");")
                vs.push("\t#[创建唯一索引] ");
                vs.push(field.title);
                vs.push("\r\n");
            }

        }

        sql(vs.join(''));
    }
}

export function walk(basePath: string, prefix: string, sql: (sql: string) => void, tableSet?: TableSet, jsonType?: boolean): TableSet {

    let r: TableSet = {}

    less.walk(basePath, (v: less.Less): void => {

        let enumSet: EnumSet = {};

        for (let i of v.enums) {
            enumSet[i.name] = i;
        }

        let getEnum = (name: string): less.LessEnum | undefined => {
            return enumSet[name]
        }

        for (let object of v.objects) {

            var isDBObject = false;

            for (let tag of object.tags) {
                if (tag.name == "type" && tag.text == "db") {
                    isDBObject = true;
                    break;
                }
            }

            if (!isDBObject) {
                continue;
            }

            var name = object.name.toLocaleLowerCase();

            if (r[name] !== undefined) {
                continue;
            }

            console.info(">>", name, object.title);

            var table: Table | undefined = tableSet === undefined ? undefined : tableSet[name]

            if (table && table.count !== undefined && table.count > 0) {
                for (var i = 1; i <= table.count; i++) {
                    tableSQL(object, getEnum, prefix + i + '_' + name, sql, table, jsonType)
                }
            } else if (table && table.name !== undefined && table.name.length > 0) {
                for (let n of table.name) {
                    tableSQL(object, getEnum, prefix + n + '_' + name, sql, table, jsonType)
                }
            } else {
                tableSQL(object, getEnum, prefix + name, sql, table, jsonType)
            }

            let fields: TableFieldSet = {};

            for (let fd of object.fields) {
                if (fd.name == "id") {
                    continue
                }
                let n = fd.name.toLocaleLowerCase();
                fields[n] = tableField(fd, getEnum, jsonType)
            }

            r[name] = {
                fields: fields,
                autoIncrement: table !== undefined ? table.autoIncrement : undefined,
                count: table !== undefined ? table.count : undefined,
            }
        }
    });

    return r;
}