

import * as less from "../Less"
import * as fs from "fs"
import { Document } from "./Document";

function getType(fd: less.LessField): string {
    switch (fd.type) {
        case less.FieldType.BOOLEAN:
            return ':boolean';
        case less.FieldType.ENUM:
        case less.FieldType.OBJECT:
            if (fd.typeSymbol !== undefined) {
                return ':' + fd.typeSymbol
            } else {
                return ':any'
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
    return ':string'
}


export function walk(basePath: string, outFile: string, name: string, title?: string): void {

    let objectSet: any = {};
    let doc = new Document();
    let vs: less.Less[] = [];

    doc.root.setValue('title', name);

    if (title) {
        doc.root.add().setValue('title', title);
    }

    less.walk(basePath, (v: less.Less): void => {

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
                    } else {
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

    vs = vs.sort((a: less.Less, b: less.Less): number => {
        let v = a.name.split('/').length - b.name.split('/').length;
        if (v == 0) {
            return a.name.localeCompare(b.name);
        }
        return v;
    });

    for (let v of vs) {

        let e = doc.root.add().setValue('title', v.name).setValue('@automatic', true)

        if (v.request.title) {
            e.add().setValue('title', v.request.title);
        }

        e.add().setValue('title', ':' + v.request.method);

        for (let fd of v.request.fields) {

            let v = e.add();

            if (fd.required) {
                v.setValue('title', fd.name);
            } else {
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
