
import * as less from "./Less"
import * as swagger from "./Swagger"
import * as aliSwagger from "./ali/swagger"
import * as postman from "./Postman"
import * as golangCli from "./golang/cli"
import * as golangServer from "./golang/server"
import * as process from "process"
import * as sql from "./sql"
import * as cli from "./cli"
import * as srv from "./srv"
import * as srvlib from "./srv-lib"
import * as mindmap from "./mindmap/less"
import * as fs from "fs"
import * as path from "path"
import * as phpCli from "./php/cli"

let commander = require('commander');

commander
    .command('less <inDir>')
    .description('quick generate less ')
    .usage("less <inDir>")
    .action(function (inDir: string) {

        less.walk(inDir, (v: less.Less): void => {

            process.stdout.write(JSON.stringify(v, undefined, 4));
            process.stdout.write("\n");

        });

    });

commander
    .command('swagger <inDir> <baseURL> [title] [version]')
    .description('quick generate swagger ')
    .usage("swagger <inDir> <baseURL> [title] [version]")
    .action(function (inDir: string, baseURL: string, title?: string, version?: string) {
        let v = swagger.walk(inDir, baseURL, title, version);
        process.stdout.write(JSON.stringify(v, undefined, 4));
        process.stdout.write("\n");

    });

commander
    .command('ali-swagger <inDir> <outFile> <vpc> [limit]')
    .description('quick generate ali-swagger ')
    .usage("ali-swagger <inDir> <outFile> <vpc> [limit]")
    .action(function (inDir: string, outFile: string, vpc: string, limit?: string) {
        let v: any = aliSwagger.walk(inDir, vpc);
        if (limit === undefined) {
            fs.writeFileSync(outFile + ".swagger", JSON.stringify(v, undefined, 4))
        } else {
            let paths: any = v.paths;
            let keys: string[] = [];
            for (let key in v.paths) {
                keys.push(key);
            }
            keys.sort();
            let l = parseInt(limit) || 10;
            let n = 0;
            let p = 1;
            let ps: any = {};

            for (let key of keys) {
                ps[key] = paths[key];
                n++;
                if (n >= l) {
                    v.paths = ps;
                    fs.writeFileSync(outFile + '-' + p + ".swagger", JSON.stringify(v, undefined, 4))
                    ps = {};
                    n = 0;
                    p++;
                }
            }

            if (n > 0) {
                v.paths = ps;
                fs.writeFileSync(outFile + '-' + p + ".swagger", JSON.stringify(v, undefined, 4))
                ps = {};
                n = 0;
                p++;
            }
        }

    });

commander
    .command('postman <inDir> <baseURL> [title] [version]')
    .description('quick generate postman ')
    .usage("postman <inDir> <baseURL> [title] [version]")
    .action(function (inDir: string, baseURL: string, title?: string, version?: string) {
        let v = postman.walk(inDir, baseURL, title, version);
        process.stdout.write(JSON.stringify(v, undefined, 4));
        process.stdout.write("\n");
    });

commander
    .command('golang <inDir> <outDir>')
    .description('quick generate golang ')
    .usage("golang <inDir> <outDir>")
    .action(function (inDir: string, outDir: string) {
        golangServer.walk(inDir, outDir);
    });

commander
    .command('golang-cli <inDir> <outDir>')
    .description('quick generate golang-cli ')
    .usage("golang-cli <inDir> <outDir>")
    .action(function (inDir: string, outDir: string) {
        golangCli.walk(inDir, outDir);
    });

commander
    .command('sql <inDir> <dest> <configFile>')
    .description('quick generate sql ')
    .usage("sql <inDir> <dest> <configFile>")
    .action(function (inDir: string, dest: string, configFile: string) {
        var config: any = {};
        if (fs.existsSync(configFile)) {
            let v = fs.readFileSync(configFile).toString();
            try {
                config = JSON.parse(v)
            }
            catch (e) {
                console.info("[JSON]", e);
            }
        }
        let vs: string[] = [];
        let tableSet = sql.walk(inDir, config.prefix || '', (v: string): void => {
            vs.push(v);
        }, config.tableSet);
        fs.writeFileSync(dest + ".sql", vs.join(''), { encoding: 'utf-8' });
        fs.writeFileSync(dest + '.json', JSON.stringify({
            prefix: config.prefix || '',
            autoIncrement: config.autoIncrement,
            tableSplitSet: config.tableSplitSet,
            tableSet: tableSet
        }, undefined, 4), { encoding: 'utf-8' });

    });

commander
    .command('cli <inDir> <outDir>')
    .description('quick generate cli ')
    .usage("cli <inDir> <outDir>")
    .action(function (inDir: string, outDir: string) {
        cli.walk(inDir, outDir);
    });

commander
    .command('srv <inDir> <outDir>')
    .description('quick generate srv ')
    .usage("srv <inDir> <outDir>")
    .action(function (inDir: string, outDir: string) {
        srv.walk(inDir, outDir);
    });

commander
    .command('srv-lib <inDir> <outDir>')
    .description('quick generate srv-lib ')
    .usage("srv-lib <inDir> <outDir>")
    .action(function (inDir: string, outDir: string) {
        srvlib.walk(inDir, outDir);
    });

commander
    .command('mindmap <inDir> <outFile> <name> [title]')
    .description('quick generate mindmap ')
    .usage("mindmap <inDir> <outFile> <name> [title]")
    .action(function (inDir: string, outFile: string, name: string, title?: string) {
        mindmap.walk(inDir, outFile, name, title);
    });

commander
    .command('php-cli <inDir> <outFile> <ns>')
    .description('quick generate php cli ')
    .usage("php-cli <inDir> <outFile> <namespace>")
    .action(function (inDir: string, outFile: string, ns: string) {
        phpCli.walk(inDir, outFile, ns);
    });


commander.parse(process.argv);