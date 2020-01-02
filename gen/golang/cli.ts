import * as less from "../Less"
import * as fs from "fs"
import * as path from "path"
import { Context } from "./context";


export function walk(basePath: string, outDir: string): void {

    console.info(basePath, ">>", outDir);

    if (!fs.existsSync(outDir)) {
        fs.mkdirSync(outDir);
    }

    let ctx = new Context(basePath);

    let inc_less = fs.readFileSync(path.join(__dirname, "inc/api.go_")).toString()
    let ns = path.basename(path.normalize(outDir))


    ctx.walk((v: less.Less): void => {

        {

            let n = ctx.symbol(v.name);
            let p = path.join(outDir, "API_" + n + ".go");

            fs.writeFileSync(p,

                inc_less
                    .replace(/#NS#/g, ns)
                    .replace(/#NAME#/g, n)
                    .replace(/#METHOD#/g, JSON.stringify(v.request.method || 'GET'))
                    .replace(/#TASK#/g, n + 'Task')
                    .replace(/#DATA#/g, ctx.getDataType(v))
                    .replace(/#DATA_DEFAULT#/g, ctx.getDataTypeDefault(v))
                    .replace(/#DATA_INIT#/g, ctx.getDataTypeInit(v))

                , { encoding: "utf-8" })
        }

        {

            let n = ctx.symbol(v.name) + "Task";
            let p = path.join(outDir, n + ".go");

            let vs: string[] = [
                "package ", ns, "\n\n",
                "type ", n, " struct {\n",
            ];

            for (let fd of v.request.fields) {
                vs.push("\t");
                vs.push(ctx.getFieldDecl(fd));
                vs.push("\n");
            }

            vs.push("}\n\n");

            vs.push("func (T *");
            vs.push(n);
            vs.push(") GetName() string {\n");
            vs.push("\t");
            vs.push("return ");
            vs.push(JSON.stringify(v.name + ".json"))
            vs.push("\n}\n\n");

            vs.push("func (T *");
            vs.push(n);
            vs.push(") GetTitle() string {\n");
            vs.push("\t");
            vs.push("return ");
            vs.push(JSON.stringify(v.request.title));
            vs.push("\n}\n\n");

            fs.writeFileSync(p,

                vs.join('')

                , { encoding: "utf-8" })

        }

        {
            for (let item of v.enums) {

                let n = ctx.symbol(item.name);
                let p = path.join(outDir, n + ".go");

                let vs: string[] = [
                    "package ", ns, "\n\n"
                ];

                for (let i of item.items) {
                    vs.push("const ");
                    vs.push(n);
                    vs.push("_");
                    vs.push(i.name);
                    vs.push(" = ");
                    vs.push(JSON.stringify(i.value));
                    vs.push("\n");
                }

                vs.push("\n");

                fs.writeFileSync(p,

                    vs.join('')

                    , { encoding: "utf-8" })
            }
        }

        {
            for (let item of v.objects) {
                let n = ctx.symbol(item.name);
                let p = path.join(outDir, n + ".go");
                let isDBObject = false;

                for (let tag of item.tags) {
                    if (tag.name == "type" && tag.text == "db") {
                        isDBObject = true;
                    }
                }

                let vs: string[] = [
                    "package ", ns, "\n\n"
                ];

                if (isDBObject) {
                    vs.push("import (\n");
                    vs.push('\t"git.sc.weibo.com/kk/microservice/golang/db"\n');
                    vs.push(")\n\n");
                }

                vs.push("type ");
                vs.push(n);
                vs.push(" struct {\n");

                if (isDBObject) {
                    vs.push("\tdb.Object\n");
                }

                for (let fd of item.fields) {
                    if (isDBObject && fd.name == "id") {
                        continue;
                    }
                    vs.push("\t");
                    vs.push(ctx.getFieldDecl(fd));
                    vs.push("\n");
                }

                vs.push("}\n\n");

                if (isDBObject) {
                    vs.push("func (O *");
                    vs.push(n);
                    vs.push(") GetName() string {\n");
                    vs.push("\treturn ");
                    vs.push(JSON.stringify(item.name.toLocaleLowerCase()));
                    vs.push("\n}\n\n");

                    vs.push("func (O *");
                    vs.push(n);
                    vs.push(") GetTitle() string {\n");
                    vs.push("\treturn ");
                    vs.push(JSON.stringify(item.title));
                    vs.push("\n}\n\n");

                }

                fs.writeFileSync(p,

                    vs.join('')

                    , { encoding: "utf-8" })
            }
        }


    });

}